package main

import (
	"log"
	"os"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/wav"
)

func PlayNotificationSound() {
	// Assignment is required otherwise no conversion is made in os.Open
	notificationSound := resourceNotificationWav.StaticName
	f, err := os.Open(notificationSound)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	streamer, format, err := wav.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	<-done
}
