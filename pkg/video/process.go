package video

import (
	"bytes"
	"fmt"
	"io"
	"os"

)

// TODO: Implement a full lenght video get

// TODO: Convert this to use raw libav or use raw ffmpeg
// Example with raw ffmpeg: ffmpeg -i ${fileName} -vf "select=eq(n\,0)" -vframes 1 -format image2 -vcodec bmp out%03d.bmp
func GetFrameAsBmp(inFileName string, frameNum int) io.Reader {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		// Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}). // TODO: Change to bitmap
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "bmp"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}

	return buf
}

// ffprobe -v error -select_streams v:0 -count_packets \
//      -show_entries stream=nb_read_packets -of csv=p=0 input.mp4
func GetVideoInformation(fileName string) {
}
