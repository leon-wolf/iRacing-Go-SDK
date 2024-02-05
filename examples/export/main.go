package main

import "github.com/leon-wolf/iRacing-Go-SDK"

func main() {
	var sdk irsdk.IRSDK
	sdk = irsdk.Init(nil)
	defer sdk.Close()
	sdk.ExportIbtTo("data.ibt")
	sdk.ExportSessionTo("data.yml")
}
