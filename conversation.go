package main

import (
	"bufio"
	"io"
	"log"
)

// Conversation is representation of incoming data from device
type Conversation struct {
	Token      string
	RemoteAddr string
	Body       io.ReadCloser
}

// GetPoints runs a goroutine which parses NMEA lines
// and writes to output channel *Points
func (c *Conversation) GetPoints() chan *Point {
	output := make(chan *Point)
	go func() {
		defer close(output)
		defer c.Body.Close()

		scanner := bufio.NewScanner(c.Body)
		for scanner.Scan() {
			log.Printf("Received line: %v", scanner.Text())
			p, err := NewPointFromNMEA(scanner.Text())
			if err != nil {
				log.Printf("Failed to parse nmea: %v", err)
				continue
			}
			output <- p
		}

		if scanner.Err() != nil {
			log.Printf("Oops, error: %v", scanner.Err())
		}
		log.Printf("No more points")
	}()
	return output
}

// Close closes connection of conversation.
// This function should be used when we by some reason
// don't want to read from socket anymore
func (c *Conversation) Close() error {
	return c.Body.Close()
}
