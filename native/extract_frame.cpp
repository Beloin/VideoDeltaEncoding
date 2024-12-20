#include <libavcodec/avcodec.h>
#include <libavformat/avformat.h>
#include <libswscale/swscale.h>
#include <stdio.h>
#include <stdlib.h>

void save_frame_as_bmp(AVFrame *pFrame, int width, int height, int frameIndex) {
  char filename[32];
  snprintf(filename, sizeof(filename), "frame%04d.bmp", frameIndex);
  FILE *file = fopen(filename, "wb");
  if (!file) {
    fprintf(stderr, "Could not open file for writing: %s\n", filename);
    return;
  }

  // BMP Header
  unsigned char bmpfileheader[14] = {'B', 'M', 0, 0,  0, 0, 0,
                                     0,   0,   0, 54, 0, 0, 0};
  unsigned char bmpinfoheader[40] = {40, 0, 0, 0, 0, 0, 0,  0,
                                     0,  0, 0, 0, 1, 0, 24, 0};

  int imageSize = width * height * 3;
  bmpfileheader[2] = (unsigned char)(imageSize + 54);
  bmpfileheader[3] = (unsigned char)((imageSize + 54) >> 8);
  bmpfileheader[4] = (unsigned char)((imageSize + 54) >> 16);
  bmpfileheader[5] = (unsigned char)((imageSize + 54) >> 24);
  bmpinfoheader[4] = (unsigned char)(width);
  bmpinfoheader[5] = (unsigned char)(width >> 8);
  bmpinfoheader[6] = (unsigned char)(width >> 16);
  bmpinfoheader[7] = (unsigned char)(width >> 24);
  bmpinfoheader[8] = (unsigned char)(height);
  bmpinfoheader[9] = (unsigned char)(height >> 8);
  bmpinfoheader[10] = (unsigned char)(height >> 16);
  bmpinfoheader[11] = (unsigned char)(height >> 24);

  fwrite(bmpfileheader, 1, 14, file);
  fwrite(bmpinfoheader, 1, 40, file);

  for (int y = height - 1; y >= 0; y--) {
    fwrite(pFrame->data[0] + y * pFrame->linesize[0], 1, width * 3, file);
  }

  fclose(file);
}

int main(int argc, char *argv[]) {
  if (argc < 2) {
    fprintf(stderr, "Usage: %s <input_file>\n", argv[0]);
    return -1;
  }

  const char *input_filename = argv[1];
  av_register_all();

  AVFormatContext *pFormatCtx = NULL;
  if (avformat_open_input(&pFormatCtx, input_filename, NULL, NULL) != 0) {
    fprintf(stderr, "Could not open file: %s\n", input_filename);
    return -1;
  }

  if (avformat_find_stream_info(pFormatCtx, NULL) < 0) {
    fprintf(stderr, "Could not retrieve stream info.\n");
    return -1;
  }

  int videoStream = -1;
  for (int i = 0; i < pFormatCtx->nb_streams; i++) {
    if (pFormatCtx->streams[i]->codecpar->codec_type == AVMEDIA_TYPE_VIDEO) {
      videoStream = i;
      break;
    }
  }

  if (videoStream == -1) {
    fprintf(stderr, "No video stream found.\n");
    return -1;
  }

  AVCodecParameters *pCodecParams = pFormatCtx->streams[videoStream]->codecpar;
  AVCodec *pCodec = avcodec_find_decoder(pCodecParams->codec_id);
  if (!pCodec) {
    fprintf(stderr, "Unsupported codec.\n");
    return -1;
  }

  AVCodecContext *pCodecCtx = avcodec_alloc_context3(pCodec);
  avcodec_parameters_to_context(pCodecCtx, pCodecParams);

  if (avcodec_open2(pCodecCtx, pCodec, NULL) < 0) {
    fprintf(stderr, "Could not open codec.\n");
    return -1;
  }

  AVFrame *pFrame = av_frame_alloc();
  AVFrame *pFrameRGB = av_frame_alloc();

  if (!pFrame || !pFrameRGB) {
    fprintf(stderr, "Could not allocate frame.\n");
    return -1;
  }

  int numBytes = av_image_get_buffer_size(AV_PIX_FMT_RGB24, pCodecCtx->width,
                                          pCodecCtx->height, 32);
  uint8_t *buffer = (uint8_t *)av_malloc(numBytes * sizeof(uint8_t));
  av_image_fill_arrays(pFrameRGB->data, pFrameRGB->linesize, buffer,
                       AV_PIX_FMT_RGB24, pCodecCtx->width, pCodecCtx->height,
                       1);

  struct SwsContext *sws_ctx = sws_getContext(
      pCodecCtx->width, pCodecCtx->height, pCodecCtx->pix_fmt, pCodecCtx->width,
      pCodecCtx->height, AV_PIX_FMT_RGB24, SWS_BILINEAR, NULL, NULL, NULL);

  int frameIndex = 0;
  AVPacket packet;
  while (av_read_frame(pFormatCtx, &packet) >= 0) {
    if (packet.stream_index == videoStream) {
      if (avcodec_send_packet(pCodecCtx, &packet) == 0) {
        while (avcodec_receive_frame(pCodecCtx, pFrame) == 0) {
          sws_scale(sws_ctx, (uint8_t const *const *)pFrame->data,
                    pFrame->linesize, 0, pCodecCtx->height, pFrameRGB->data,
                    pFrameRGB->linesize);
          save_frame_as_bmp(pFrameRGB, pCodecCtx->width, pCodecCtx->height,
                            frameIndex++);
        }
      }
    }
    av_packet_unref(&packet);
  }

  av_free(buffer);
  av_frame_free(&pFrame);
  av_frame_free(&pFrameRGB);
  avcodec_free_context(&pCodecCtx);
  avformat_close_input(&pFormatCtx);
  return 0;
}
