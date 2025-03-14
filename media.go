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

	//defer os.Remove(f.Name())

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
	//wR := bytes.NewReader(resourceNotificationWav.Content())
	file := createTmpWaveFile(resourceNotificationWav.Content())
	file.Seek(0, 0)

	streamer, format, err := wav.Decode(file)
	if err != nil {
		log.Fatal("37: ", err)
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
