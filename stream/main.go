package main

import (
	"os/exec"
)

func main() {
	instantConvert("")
}

func instantConvert(mp4path string) {
	exec.Command("ffmpeg", "-i", mp4path, "-c:v", "libvpx", "-b:v", "1M", "-c:a", "libvorbis", "output.webm").Start()
}
