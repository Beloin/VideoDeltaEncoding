package video

import "beloin.com/dataencoding/pkg/video"

const file = "./examples/EletroFunk Deboxe Jiraya Uai As Meninas Se Envolve #shorts.mp4"

func GetVideoFrames(n int) {
  buf1 := video.GetFrameAsBmp(file, 0)
  buf2:= video.GetFrameAsBmp(file, 1)
}
