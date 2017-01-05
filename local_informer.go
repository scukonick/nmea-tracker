package main

import "database/sql"

// Informer is an iterface allowing
// to talk to storage
type Informer interface {
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
