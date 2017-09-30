package main

import "encoding/binary"
import "log"

var (
	dataStart = 18
	recordLen = 132
)

// Data Process heartbeat type message from mobile device
func Data(message []byte) (r []byte, err error) {

	imei := string(message[1:15])
	log.Printf("imei: %s", imei)

	sim := string(message[16:20])
	log.Printf("sim: %s", sim)

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
	SignalRecord := string(message[:4])
	log.Printf("imei: %s", SignalRecord)
	SignalDateTimeRead := string(message[4:16])
	log.Printf("imei: %s", SignalDateTimeRead)
	SignalEvent := string(message[16:17])
	log.Printf("imei: %s", SignalEvent)
	SignalTagID := string(message[17:25])
	log.Printf("imei: %s", SignalTagID)
	// Non required
	// SignalUTCTime := string(message[25:31])
	// SignalPosStat := string(message[31:32])
	// SignalLatitude := string(message[32:39])
	// SignalLatitudeRef := string(message[39:40])
	// SignalLongtitude := string(message[40:48])
	// SignalLongtitudeRef := string(message[48:49])
	// SignalSpeed := string(message[49:54])
	// SignalHeading := string(message[54:59])
	// SignalUTCDate := string(message[59:65])
	// SignalMagneticVariation := string(message[65:70])
	// SignalMagRef := string(message[70:71])
	// SignalFixMode := string(message[71:72])
	// SignalSateliteUsed := string(message[72:73])
	// SignalHDOP := string(message[73:78])
	// SignalMslAltitude := string(message[78:83])
	// SignalBattery := string(message[83:87])
	// SignalTemperature := string(message[87:90])
	// SignalLight := string(message[90:92])
	// RSSI := string(message[92:94])
	// LAC := string(message[94:98])
	// Ci := string(message[98:102])
	// MCC := string(message[102:105])
	// MNC := string(message[105:107])
	// MSIN := string(message[107:117])

	return r, nil
}
