package main

import (
	"database/sql/driver"
	"fmt"
)

// PGLocation is representation of POINT posgis field
type PGLocation struct {
	lng float64 // Longitude in degrees
	lat float64 // Latitude in degrees
}

// Value implements db.Driver interface so it could be used
// for the insert into postgres
func (l *PGLocation) Value() (driver.Value, error) {
	layout := fmt.Sprintf("POINT(%f %f)", l.lng, l.lat)
	return layout, nil
}

// NewPGLocation returns newly initialized PGLocation
func NewPGLocation(lng, lat float64) *PGLocation {
	location := &PGLocation{
		lng: lng,
		lat: lat,
	}
	return location
}
