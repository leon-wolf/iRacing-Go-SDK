package irsdk

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/hidez8891/shm"
	"github.com/leon-wolf/iRacing-Go-SDK/lib/winevents"
)

// IRSDK is the main SDK object clients must use
type IRSDK struct {
	r             reader
	h             *header
	session       Session
	s             []string
	tVars         *TelemetryVars
	lastValidData int64
	debug         bool
}

func (sdk *IRSDK) WaitForData(timeout time.Duration) bool {
	if !sdk.IsConnected() {
		initIRSDK(sdk)
	}
	if winevents.WaitForSingleObject(timeout) {
		return readVariableValues(sdk)
	}
	return false
}

func (sdk *IRSDK) GetVar(name VariableName) (Variable, error) {
	if !sessionStatusOK(sdk.h.status) {
		return Variable{}, fmt.Errorf("session is not active")
	}
	sdk.tVars.mux.Lock()
	if v, ok := sdk.tVars.vars[name]; ok {
		sdk.tVars.mux.Unlock()
		return v, nil
	}
	sdk.tVars.mux.Unlock()
	return Variable{}, fmt.Errorf("telemetry variable %q not found", name)
}

func (sdk *IRSDK) GetSession() Session {
	return sdk.session
}

func (sdk *IRSDK) GetLastVersion() int {
	if !sessionStatusOK(sdk.h.status) {
		return -1
	}
	sdk.tVars.mux.Lock()
	last := sdk.tVars.lastVersion
	sdk.tVars.mux.Unlock()
	return last
}

func (sdk *IRSDK) GetSessionData(path string) (string, error) {
	if !sessionStatusOK(sdk.h.status) {
		return "", fmt.Errorf("session not connected")
	}
	return getSessionDataPath(sdk.s, path)
}

func (sdk *IRSDK) IsConnected() bool {
	if sdk.h != nil {
		if sessionStatusOK(sdk.h.status) && (sdk.lastValidData+connTimeout > time.Now().Unix()) {
			return true
		}
	}

	return false
}

// ExportIbtTo ExportTo exports current memory data to a file
func (sdk *IRSDK) ExportIbtTo(fileName string) error {
	rbuf := make([]byte, fileMapSize)
	_, err := sdk.r.ReadAt(rbuf, 0)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(fileName, rbuf, 0644)
	return err
}

// ExportSessionTo ExportTo exports current session yaml data to a file
func (sdk *IRSDK) ExportSessionTo(fileName string) error {
	y := strings.Join(sdk.s, "\n")
	err := os.WriteFile(fileName, []byte(y), 0644)
	return err
}

func (sdk *IRSDK) BroadcastMsg(msg Msg) {
	if msg.P2 == nil {
		msg.P2 = 0
	}
	winevents.BroadcastMsg(broadcastMsgName, msg.Cmd, msg.P1, msg.P2, msg.P3)
}

// Close clean up sdk resources
func (sdk *IRSDK) Close() {
	err := sdk.r.Close()
	handleError(err)
}

// Init creates an SDK instance to operate with
func Init() IRSDK {

	debug := flag.Bool("debug", false, "debug mode")
	flag.Parse()
	r, err := shm.Open(fileMapName, fileMapSize)
	if err != nil {
		log.Fatal(err)
	}

	sdk := IRSDK{r: r, lastValidData: 0, debug: *debug}
	winevents.OpenEvent(dataValidEventName)
	initIRSDK(&sdk)
	return sdk
}

func initIRSDK(sdk *IRSDK) {
	h := readHeader(sdk.r)
	sdk.h = &h
	sdk.s = nil
	if sdk.tVars != nil {
		sdk.tVars.vars = nil
	}
	if sessionStatusOK(h.status) {
		if sdk.debug {
			dumpSessionDataToFile(sdk, "session.yaml")
		}
		sRaw := readSessionDataYaml(sdk.r, &h)
		err := yaml.Unmarshal([]byte(sRaw), &sdk.session)
		if err != nil {
			log.Fatal(err)
		}
		sdk.s = strings.Split(sRaw, "\n")
		sdk.tVars = readVariableHeaders(sdk.r, &h)
		readVariableValues(sdk)
	}
}

func sessionStatusOK(status int) bool {
	return (status & stConnected) > 0
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
