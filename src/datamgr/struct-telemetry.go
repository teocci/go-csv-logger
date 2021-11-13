// Package datamgr
// Created by RTT.
// Author: teocci@yandex.com on 2021-Sep-28
package datamgr

import (
	"strconv"
	"time"

	"github.com/teocci/go-csv-logger/src/timemgr"
)

type RTT struct {
	Seq            int64     `json:"seq" csv:"seq"`
	DroneID        int64     `json:"drone_id" csv:"drone_id"`
	FlightID       int64     `json:"flight_id" csv:"flight_id"`
	TimeBootMs     uint32    `json:"time_boot_ms" csv:"time_boot_ms"`
	Lat            int32     `json:"lat" csv:"lat"`
	Lon            int32     `json:"lon" csv:"lon"`
	Alt            int32     `json:"alt" csv:"lat"`
	Roll           float32   `json:"roll" csv:"lat"`
	Pitch          float32   `json:"pitch" csv:"lat"`
	Yaw            float32   `json:"yaw" csv:"lat"`
	BatVoltage     float32   `json:"battery_voltage" csv:"battery_voltage"`
	BatCurrent     float32   `json:"battery_current" csv:"battery_current"`
	BatPercent     float32   `json:"battery_percentage" csv:"battery_percentage"`
	BatTemperature float32   `json:"battery_temperature" csv:"battery_temperature"`
	Temperature    float32   `json:"temperature" csv:"temperature"`
	LastUpdate     time.Time `json:"last_update" csv:"last_update"`
}

func ParseRTT(data []string) *RTT {
	droneID, _ := strconv.Atoi(data[1])
	sessionID, _ := strconv.Atoi(data[2])
	timeBootMs, _ := strconv.ParseFloat(data[0], 64)
	lat, _ := strconv.ParseFloat(data[3], 64)
	long, _ := strconv.ParseFloat(data[4], 64)
	alt, _ := strconv.ParseFloat(data[5], 64)
	roll, _ := strconv.ParseFloat(data[6], 64)
	pitch, _ := strconv.ParseFloat(data[7], 64)
	yaw, _ := strconv.ParseFloat(data[8], 64)
	temp, _ := strconv.ParseFloat(data[9], 64)
	batVol, _ := strconv.ParseFloat(data[10], 64)
	batCurr, _ := strconv.ParseFloat(data[11], 64)
	batPct, _ := strconv.ParseFloat(data[12], 64)
	batTemp, _ := strconv.ParseFloat(data[13], 64)
	gpsTime, _ := strconv.ParseFloat(data[14], 64)

	return &RTT{
		DroneID:        int64(droneID),
		FlightID:       int64(sessionID),
		TimeBootMs:     uint32(timeBootMs),
		Lat:            int32(lat),
		Lon:            int32(long),
		Alt:            int32(alt),
		Roll:           float32(roll),
		Pitch:          float32(pitch),
		Yaw:            float32(yaw),
		BatVoltage:     float32(batVol),
		BatCurrent:     float32(batCurr),
		BatPercent:     float32(batPct),
		BatTemperature: float32(batTemp),
		Temperature:    float32(temp),
		LastUpdate:     timemgr.UnixTime(float32(gpsTime)),
	}
}
