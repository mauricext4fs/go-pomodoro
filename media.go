package main

import (
	"bytes"
	"log"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/wav"
)

func PlayNotificationSound() {
	wR := bytes.NewReader(resourceNotificationWav.Content())
	streamer, format, err := wav.Decode(wR)
	if err != nil {
		log.Fatal(err)
	}

	defer streamer.Close()
	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		log.Fatalf("Failed to init speaker: %v", err)
	}

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	<-done
}
