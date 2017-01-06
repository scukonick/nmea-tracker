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
func (p *PGLocation) Scan(val interface{}) error {
	/*b, err := hex.DecodeString(string(val.([]uint8)))
	if err != nil {
		return err
	}

	r := bytes.NewReader(b)
	var wkbByteOrder uint8
	if err := binary.Read(r, binary.LittleEndian, &wkbByteOrder); err != nil {
		return err
	}

	var byteOrder binary.ByteOrder
	switch wkbByteOrder {
	case 0:
		byteOrder = binary.BigEndian
	case 1:
		byteOrder = binary.LittleEndian
	default:
		return fmt.Errorf("invalid byte order %v", wkbByteOrder)
	}

	var wkbGeometryType uint64
	if err := binary.Read(r, byteOrder, &wkbGeometryType); err != nil {
		return err
	}

	if err := binary.Read(r, byteOrder, p); err != nil {
		return err
	}
	*/
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
