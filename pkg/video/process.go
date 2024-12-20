package video

import (
	"bytes"
	"fmt"
	"io"
	"os"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func ExampleReadFrameAsBmp(inFileName string, frameNum int) io.Reader {
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
