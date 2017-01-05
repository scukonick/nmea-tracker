package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/adrianmo/go-nmea"
)

const (
	gprmc       = "$GPRMC"
	gnrmc       = "$GNRMC"
	timeLayout1 = "020106150405"
	timeLayout2 = "02010615040500"
	timeLayout3 = "020106150405000"
)

// Point represents one track point
type Point struct {
	Type      string    // Type of source NMEA message
	Lat       float64   // Latitude in degrees
	Lng       float64   // Longtitude in degrees
	Speed     float32   // Speed in m/s
	Direction float32   // Direction in degrees
	TimeStamp time.Time // Time stamp of the point
}

// NewPointFromNMEA parses input NMEA string
// and returns pointer to newly initialized point.
// If it encounters some error it'll return non-nil error.
// For now it accepts only GPRMC, GNRMC messages.
func NewPointFromNMEA(input string) (*Point, error) {
	m, err := nmea.Parse(input)
	if err != nil {
		return nil, err
	}
	if m.GetSentence().Type != nmea.PrefixGPRMC {
		return nil, fmt.Errorf("Wrong message type: %s", m.GetSentence().Type)
	}

	gprmc := m.(nmea.GPRMC)
	if gprmc.Validity != "A" {
		return nil, fmt.Errorf("Validity is wrong: %s", gprmc.Validity)
	}
	var lat, lng interface{}
	lat = gprmc.Latitude
	lng = gprmc.Longitude

	latFloat := reflect.ValueOf(lat).Float()
	lngFloat := reflect.ValueOf(lng).Float()

	speed := float32(gprmc.Speed * 0.514444)

	timeStamp, err := parseTime(gprmc.Date + gprmc.Time)
	if err != nil {
		return nil, err
	}

	p := &Point{
		Type:      gprmc.Type,
		Lat:       latFloat,
		Lng:       lngFloat,
		Speed:     speed,
		Direction: float32(gprmc.Course),
		TimeStamp: timeStamp,
		// TimeStamp: gprmc.Time,
	}
	log.Printf("Point: %+v", p)
	return p, nil
}

func parseTime(input string) (time.Time, error) {
	for _, layout := range []string{timeLayout1, timeLayout2, timeLayout3} {
		t, err := time.Parse(layout, input)
		if err == nil {
			return t, nil
		}

	}
	return time.Time{}, errors.New("Failed to parse time")
}
