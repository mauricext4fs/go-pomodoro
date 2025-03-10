package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/wav"
)

func createTmpWaveFile(wave []byte) {
	f, err := os.CreateTemp("", "pomodoro_sound")
	if err != nil {
		log.Fatalln("Could not create a temporary file: ", err)
	}

	fmt.Println("Temp file name: ", f.Name())

	defer os.Remove(f.Name())

	_, err = f.Write(wave)
	if err != nil {
		log.Fatalln("Could not write to temporary file: ", f.Name())
	}
}

func PlayNotificationSound() {
	wR := bytes.NewReader(resourceNotificationWav.Content())
	createTmpWaveFile(resourceNotificationWav.Content())
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
