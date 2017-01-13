package main

import (
	"database/sql"
	"log"
	"time"
)

// Informer is an iterface allowing
// to talk to storage
type Informer interface {
	GetDevice(token string) (int32, error)
	InsertPoint(p *Point, tokenID int32) error
	GetPoints(token string) ([]*Point, error)
}

// PGInformer is implementation of Informer
// using postgres DB
type PGInformer struct {
	db *sql.DB
}

// NewPGInformer returns pointer to
// newly initialized PGInformer
func NewPGInformer(db *sql.DB) *PGInformer {
	i := &PGInformer{
		db: db,
	}
	return i
}

// GetDevice checks if device with token exists in database,
// updates it's last seen
// and returns id of the device.
// If not - it creates device and returns id
func (i *PGInformer) GetDevice(token string) (int32, error) {
	var id int32
	query := `INSERT INTO devices (token, created_at, last_seen_at)
  VALUES ($1, NOW(), NOW())
  ON CONFLICT (token) DO UPDATE SET
  last_seen_at = NOW()
  RETURNING id`
	err := i.db.QueryRow(query, token).Scan(&id)

	return id, err
}

// InsertPoint saves point in the database
func (i *PGInformer) InsertPoint(p *Point, tokenID int32) error {
	query := `INSERT INTO points
  (type, speed, direction, time_stamp, device_id, location)
  VALUES ($1, $2, $3, $4, $5,
  ST_GeomFromText($6, 4326))`
	_, err := i.db.Exec(query, p.Type, p.Speed, p.Direction, p.TimeStamp, tokenID,
		p.Location)
	return err
}

// GetPoints returns all the points of device identified by
// token ordered by time_stamp desc
func (i *PGInformer) GetPoints(token string) ([]*Point, error) {
	query := `SELECT points.speed, points.direction, points.time_stamp,
	ST_X(points.location::geometry), ST_Y(points.location::geometry)
	FROM points INNER JOIN devices
	ON points.device_id = devices.id
	WHERE devices.token = $1
	ORDER BY points.time_stamp`

	rows, err := i.db.Query(query, token)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]*Point, 0, 100)

	for rows.Next() {
		log.Printf("Got point!")
		point := &Point{}
		ts := time.Time{}
		var speed64, direction64, lat, lng float64

		if err := rows.Scan(&speed64, &direction64, &ts, &lng, &lat); err != nil {

			return nil, err
		}
		point.Speed = float32(speed64)
		point.Direction = float32(direction64)
		point.TimeStamp = ts
		point.Location = &PGLocation{
			Lat: lat,
			Lng: lng,
		}
		result = append(result, point)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
