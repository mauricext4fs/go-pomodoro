package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/wav"
)

func createTmpWaveFile(wave []byte) *os.File {
	f, err := os.CreateTemp("", "pomodoro_sound")
	if err != nil {
		log.Fatalln("Could not create a temporary file: ", err)
	}

	fmt.Println("Temp file name: ", f.Name())

	numBytes := 0
	numBytes, err = f.Write(wave)

	if err != nil {
		log.Fatalln("Could not write to temporary file: ", f.Name())
	}

	if numBytes <= 1 {
		log.Fatalln("Oupss... nothing written to the temp wave file")
	}

	return f
}

func PlayNotificationSound() {
	f := createTmpWaveFile(resourceNotificationWav.Content())
	file, err := os.Open(f.Name())
	if err != nil {
		log.Fatalln("Temp wave file could not be open! ", err)
	}

	defer f.Close()

	streamer, format, err := wav.Decode(file)
	if err != nil {
		log.Fatal("Wav decode error: ", err)
	}
	defer streamer.Close()

	log.Println("Start playing Notification audio")
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	<-done
	log.Println("Finish playing Notification audio")
}
