package main

import (
	"log"
	"time"

	irsdk "github.com/leon-wolf/iRacing-Go-SDK"
)

func main() {
	sdk := irsdk.Init()
	defer sdk.Close()

	// Wait for iRacing to be connected
	for !sdk.IsConnected() {
		log.Println("Waiting for iRacing connection...")
		time.Sleep(5 * time.Second)
	}

	log.Println("iRacing connected! Exporting data...")
	sdk.ExportIbtTo("data.ibt")
	sdk.ExportSessionTo("data.yml")
	log.Println("Data exported successfully!")
}
