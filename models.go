package main

import (
	"time"

	"github.com/jinzhu/gorm"
)

// LogRecord Record for logging
type LogRecord struct {
	SignalRecord            string
	SignalDateTimeRead      string
	SignalEvent             string
	SignalTagID             string
	SignalUTCTime           string
	SignalPosStat           string
	SignalLatitude          string
	SignalLatitudeRef       string
	SignalLongtitude        string
	SignalLongtitudeRef     string
	SignalSpeed             string
	SignalHeading           string
	SignalUTCDate           string
	SignalMagneticVariation string
	SignalMagRef            string
	SignalFixMode           string
	SignalSateliteUsed      string
	SignalHDOP              string
	SignalMslAltitude       string
	SignalBattery           string
	SignalTemperature       string
	SignalLight             string
	RSSI                    string
	LAC                     string
	Ci                      string
	MCC                     string
	MNC                     string
	MSIN                    string
}

// DataRecord Record for storing in database
type DataRecord struct {
	gorm.Model
	SignalIMEI              string
	SignalRecord            uint64
	SignalDateTimeRead      time.Time
	SignalTimeStampServer   time.Time
	SignalTagID             string
	SignalMPYID             int
	SignalEPYID             int
	SignalYpallilosID       int
	SignalSEL               int
	SignalEvent             int
	SignalBattery           string
	SignalLatitude          float32
	SignalLongtitude        string
	SignalLatitudeRef       string
	SignalLongtitudeRef     string
	SignalSpeed             string
	SignalHeading           string
	SignalMagneticVariation string
	SignalFixMode           string
	SignalSateliteUsed      string
	SignalHDOP              string
	SignalMslAltitude       string
	SignalTemperature       string
	SignalLight             string
	SignalChckedByTimer     bool
	Signaluniqueid          string
	SignalIsTsepy           bool
	SignalRealTimeInsertud  time.Time
	SignalUTCDateTime       time.Time
	SignalPosStat           bool
	SignalMagRef            bool
	RSSI                    string
	Lac                     string
	Ci                      string
	MCC                     string
	MNC                     string
	MSIN                    string
	IsTransfered            bool
	SignalMPYEpopti         int
	SignalEPYEpopti         int
	Sound                   string
	isAlarm                 bool
	SignalResponsible       int
	PersonID                int
	VehicleID               int
	IsMobile                bool
}
