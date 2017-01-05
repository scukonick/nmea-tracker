package main

import "log"

// ConversationHandler receives and processes conversations from devices
type ConversationHandler interface {
	Handle(c *Conversation)
}

// PGConversationHandler is an  implementation of ConversationHandler
// which processes conversations and stores them to Postgres DB
type PGConversationHandler struct {
	informer Informer
}

// NewPGConversationHandler returns pointer to newly initialized PGConversationHandler
func NewPGConversationHandler(informer Informer) *PGConversationHandler {
	h := &PGConversationHandler{
		informer: informer,
	}
	return h
}

// Handle reads conversation, checks if token already exists in the database,
// reads NMEA lines and stores them as points in the database
func (h *PGConversationHandler) Handle(c *Conversation) {
	deviceID, err := h.informer.GetDevice(c.Token)
	if err != nil {
		log.Printf("Error when getting device from DB: %v", err)
		c.Close()
		return
	}
	log.Printf("Found device: %d", deviceID)
	for p := range c.GetPoints() {
		log.Printf("Got point: %+v", p)
		err = h.informer.InsertPoint(p, deviceID)
		if err != nil {
			log.Printf("Error when inserting point: %v", err)
		}
	}
}
