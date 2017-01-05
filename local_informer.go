package main

import "database/sql"

// Informer is an iterface allowing
// to talk to storage
type Informer interface {
	GetDevice(token string) (int32, error)
	InsertPoint(p *Point, tokenID int32) error
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
