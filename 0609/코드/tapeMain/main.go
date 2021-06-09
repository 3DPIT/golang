package main

import (
	"src/github.com/headfirstgo/gadget"
)

type Player interface { //인터페이스 타입을 정의
	Play(string)
	Stop()
}

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	record, ok := player.(gadget.TapeRecorder)
	if ok {
		record.Record()
	}
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	TryOut(gadget.TapeRecorder{})
	TryOut(gadget.TapePlayer{})
}
