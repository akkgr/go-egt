package main

import (
	"time"
)

// Heartbeat Process heartbeat type message from mobile device
func Heartbeat(message []byte) (r []byte, err error) {

	imei := string(message[1:15])
	sim := string(message[16:20])

	t := time.Now()
	s := t.Format("02012006150405")
	r = []byte(s)

	r = append(r, []byte(imei)...)
	r = append(r, []byte(sim)...)

	return r, nil
}
