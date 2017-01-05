package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
)

// NmeaServer is implementation of tcp server
// accepting connections and sending them to handler
type NmeaServer struct {
	Address string
}

// NewNmeaServer returns newly initialized pointer to NmeaServer
func NewNmeaServer(address string) *NmeaServer {
	return &NmeaServer{Address: address}
}

// Serve listen on the Address and sends new conversations to handler
// to handle them in the new goroutine. implementation of ConversationHandler
// should be thread-safe.
func (s *NmeaServer) Serve(handler ConversationHandler) {
	l, err := net.Listen("tcp", s.Address)
	if err != nil {
		log.Fatal(err)
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		tcpConn := conn.(*net.TCPConn)
		log.Printf("Incoming connection from: %s", conn.RemoteAddr().String())
		conversation, err := getConversation(tcpConn)
		if err != nil {
			log.Printf("WARN: not correct input stream")
			conn.Close()
		}

		go handler.Handle(conversation)
	}
}

// getConversation reads from new tcp connection,
// parses it and returns *Conversation
func getConversation(input *net.TCPConn) (*Conversation, error) {
	scanner := bufio.NewScanner(input)
	if !scanner.Scan() {
		return nil, errors.New("No data from input")
	}

	token, err := parseToken(scanner.Text())
	if err != nil {
		return nil, fmt.Errorf("Invalid token: %v", err)
	}

	conversation := &Conversation{
		Token:      token,
		Body:       input,
		RemoteAddr: input.RemoteAddr().String(),
	}
	return conversation, nil
}

// parseToken accepts line with token,
// parses it and returns token. If there is some issue with
// line parsing, it returns non-nil error
func parseToken(line string) (string, error) {
	var token string
	elems := strings.Split(line, " ")
	if len(elems) != 2 {
		return token, fmt.Errorf("Wrong number of elements: %d", len(elems))
	}

	token = elems[1]
	if len(token) != 25 {
		return token, fmt.Errorf("Wrong length of token: %d", len(token))
	}

	return token, nil
}
