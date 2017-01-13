package main

import (
	"database/sql/driver"
	"fmt"
)

// PGLocation is representation of POINT posgis field
type PGLocation struct {
	Lng float64 // Longitude in degrees
	Lat float64 // Latitude in degrees
}

// Value implements db.Driver interface so it could be used
// for the insert into postgres
func (l *PGLocation) Value() (driver.Value, error) {
	layout := fmt.Sprintf("POINT(%f %f)", l.Lng, l.Lat)
	return layout, nil
}

// Scan implements Scanner interface
func (l *PGLocation) Scan(val interface{}) error {
	// TODO implement scan interface to use it directly with DB

	return nil
}

// NewPGLocation returns newly initialized PGLocation
func NewPGLocation(lng, lat float64) *PGLocation {
	location := &PGLocation{
		Lng: lng,
		Lat: lat,
	}
	return location
}
