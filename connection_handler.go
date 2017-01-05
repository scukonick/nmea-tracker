package main

import "database/sql"

// ConversationHandler receives and processes conversations from devices
type ConversationHandler interface {
	Handle(c *Conversation)
}

// PGConversationHandler is an  implementation of ConversationHandler
// which processes conversations and stores them to Postgres DB
type PGConversationHandler struct {
	db *sql.DB
}

// NewPGConversationHandler returns pointer to newly initialized PGConversationHandler
func NewPGConversationHandler(db *sql.DB) *PGConversationHandler {
	h := &PGConversationHandler{
		db: db,
	}
	return h
}

// Handle reads conversation, checks if token already exists in the database,
// reads NMEA lines and stores them as points in the database
func (h *PGConversationHandler) Handle(c *Conversation) {

}
