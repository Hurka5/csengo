package gui

import (
	"log"
	"os"
	"time"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var (
  initialized = false
)

func PlaySong(song string) {
  f, err := os.Open(song)
	if err != nil {
		log.Fatal(err)
	}
  defer f.Close()

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

  if !initialized {
	  speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	  initialized = true
  }
	

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
