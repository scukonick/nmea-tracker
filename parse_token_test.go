package main

import "testing"

func TestParseToken(t *testing.T) {
	input := "Token: 12312314124"
	_, err := parseToken(input)
	if err == nil {
		t.Error("Expected error - invalid token, got: nil")
	}

	input = "Token: 1234567890123456789012345 12312"
	_, err = parseToken(input)
	if err == nil {
		t.Error("Expected error - invalid token, got: nil")
	}

	input = "Token: 1234567890123456789012345"
	expectedToken := "1234567890123456789012345"
	token, err := parseToken(input)
	if err != nil {
		t.Errorf("Expected nil error, got: %v", err)
	}
	if token != expectedToken {
		t.Errorf("Expected token %s, got: %s", expectedToken, token)
	}
}
