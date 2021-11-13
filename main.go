// Package go_mavlink_parser
// Created by RTT.
// Author: teocci@yandex.com on 2021-Sep-27
package main

import (
	"bytes"
	"io"
	"log"
	"math/rand"
	"time"

	"github.com/jszwec/csvutil"
	"github.com/teocci/go-csv-logger/src/csvmgr"
	"github.com/teocci/go-csv-logger/src/datamgr"
)

const (
	defaultLon = 1274446074
	defaultLat = 358242073
)

var (
	initConf datamgr.InitConf
	rtt      *datamgr.RTT
	csvl     *csvmgr.CSVLogger

	records []*datamgr.RTT

	headerSent = false
)

func main() {
	rand.Seed(time.Now().UnixNano())
	base := time.Now()

	initConf = datamgr.InitConf{
		Host:      "localhost",
		Port:      5562,
		ConnID:    0,
		ModuleTag: "tl",
		CompanyID: 1,
		DroneID:   4,
		FlightID:  124,
	}

	// init csvlogger
	csvl = csvmgr.NewCSVLogger(initConf)

	for i := 0; i < 100; i++ {
		current := time.Now()
		record := &datamgr.RTT{
			Seq:            int64(i),
			DroneID:        initConf.DroneID,
			FlightID:       initConf.FlightID,
			TimeBootMs:     uint32(current.Sub(base).Milliseconds()),
			Lat:            int32(defaultLat),
			Lon:            int32(defaultLon),
			Alt:            int32(genRandom(800, 1200)),
			Roll:           genRandomFloat(),
			Pitch:          genRandomFloat(),
			Yaw:            genRandomFloat(),
			BatVoltage:     30 + genRandomFloat(),
			BatCurrent:     1 + genRandomFloat(),
			BatPercent:     genRandomFloat() * 100,
			BatTemperature: 40 + genRandomFloat(),
			Temperature:    32 + genRandomFloat(),
			LastUpdate:     current,
		}
		records = append(records, record)
	}

	for _, record := range records {
		process(record)
		//fmt.Printf("%#v\n", record)
	}
}



func process(rtt *datamgr.RTT) {
	appendRecord(rtt)
}

func appendRecord(rtt *datamgr.RTT) {
	rtts := []datamgr.RTT{*rtt}
	b, err := csvutil.Marshal(rtts)
	if err != nil {
		log.Println("error:", err)
	}

	buf := bytes.NewBuffer(b)

	header, err := buf.ReadBytes('\n')
	if err != nil && err != io.EOF {
		log.Println("error:", err)
	}

	line, err := buf.ReadBytes('\n')
	if err != nil && err != io.EOF {
		log.Println("error:", err)
	}

	if !headerSent {
		csvl.Append <- header
		headerSent = true
	}
	csvl.Append <- line
}

func genRandom(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

func genRandomFloat() float32 {
	return rand.Float32()
}
