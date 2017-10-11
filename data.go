package main

import (
	"encoding/binary"
	"fmt"
	"log"
)

var (
	dataStart = 18
	recordLen = 132
)

// Data Process heartbeat type message from mobile device
func Data(message []byte) (r []byte, err error) {

	imei := string(message[1:16])
	log.Printf("imei: %s", imei)

	dataLen := binary.BigEndian.Uint16(message[16:18])
	records := int(dataLen) / recordLen

	for i := 0; i < records; i++ {
		startpos := dataStart + i*recordLen
		endpos := startpos + recordLen
		row(message[startpos:endpos])
	}

	return r, nil
}

func row(message []byte) (r []byte, err error) {

	// Required
	log.Printf("SignalRecord: %x", message[:4])
	SignalRecord := binary.LittleEndian.Uint32(message[:4])
	log.Printf("SignalRecord: %v", SignalRecord)

	log.Printf("SignalDateTimeRead: %x", message[4:16])
	SignalDateTimeRead := string(message[4:16])
	log.Printf("SignalDateTimeRead: %s", SignalDateTimeRead)

	log.Printf("SignalEvent: %x", message[16:17])
	SignalEvent := fmt.Sprintf("%x", message[16:17])
	log.Printf("SignalEvent: %v", SignalEvent)

	log.Printf("SignalTagID: %x", message[17:25])
	SignalTagID := fmt.Sprintf("%x", message[17:25])
	log.Printf("SignalTagID: %s", SignalTagID)

	// Non required
	SignalUTCTime := string(message[25:31])
	log.Printf("SignalUTCTime: %s", SignalUTCTime)
	SignalPosStat := string(message[31:32])
	log.Printf("SignalPosStat: %s", SignalPosStat)
	SignalLatitude := string(message[32:39])
	log.Printf("SignalLatitude: %s", SignalLatitude)
	SignalLatitudeRef := string(message[39:40])
	log.Printf("SignalLatitudeRef: %s", SignalLatitudeRef)
	SignalLongtitude := string(message[40:48])
	log.Printf("SignalLongtitude: %s", SignalLongtitude)
	SignalLongtitudeRef := string(message[48:49])
	log.Printf("SignalLongtitudeRef: %s", SignalLongtitudeRef)
	SignalSpeed := string(message[49:54])
	log.Printf("SignalSpeed: %s", SignalSpeed)
	SignalHeading := string(message[54:59])
	log.Printf("SignalHeading: %s", SignalHeading)
	SignalUTCDate := string(message[59:65])
	log.Printf("SignalUTCDate: %s", SignalUTCDate)
	SignalMagneticVariation := string(message[65:70])
	log.Printf("SignalMagneticVariation: %s", SignalMagneticVariation)
	SignalMagRef := string(message[70:71])
	log.Printf("SignalMagRef: %s", SignalMagRef)
	SignalFixMode := string(message[71:72])
	log.Printf("SignalFixMode: %s", SignalFixMode)
	SignalSateliteUsed := fmt.Sprintf("%x", message[72:73])
	log.Printf("SignalSateliteUsed: %s", SignalSateliteUsed)
	SignalHDOP := string(message[73:78])
	log.Printf("SignalHDOP: %s", SignalHDOP)
	SignalMslAltitude := string(message[78:83])
	log.Printf("SignalMslAltitude: %s", SignalMslAltitude)
	SignalBattery := string(message[83:87])
	log.Printf("SignalBattery: %s", SignalBattery)
	SignalTemperature := string(message[87:90])
	log.Printf("SignalTemperature: %s", SignalTemperature)
	SignalLight := fmt.Sprintf("%x", message[90:92])
	log.Printf("SignalLight: %s", SignalLight)
	RSSI := string(message[92:94])
	log.Printf("RSSI: %s", RSSI)
	LAC := string(message[94:98])
	log.Printf("LAC: %s", LAC)
	Ci := string(message[98:102])
	log.Printf("Ci: %s", Ci)
	MCC := string(message[102:105])
	log.Printf("MCC: %s", MCC)
	MNC := string(message[105:107])
	log.Printf("MNC: %s", MNC)
	MSIN := string(message[107:117])
	log.Printf("MSIN: %s", MSIN)

	return r, nil
}
