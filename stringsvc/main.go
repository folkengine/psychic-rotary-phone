package main

import (
	"context"
	"errors"
	"strings"
)

// StringService provides operations on strings.
type StringService interface {
	Uppercase(context.Context, string) (string, error)
	Count(context.Context, string) int
}

type stringService struct{}

func (stringService) Uppercase(_ context.Context, s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(_ context.Context, s string) int {
	return len(s);
}

// ErrEmpty is returned with input
var ErrEmpty = errors.New("Empty String")

type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V 	string `json:"v"`
	Err string `json:"err,omitempty"`
}