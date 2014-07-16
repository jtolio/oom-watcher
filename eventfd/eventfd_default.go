// Copyright (C) 2014 JT Olds, see provided LICENSE file
// +build !linux

package eventfd

import (
	"errors"
)

type EventFD struct{}

func NewEventFD() (*EventFD, error) {
	return nil, errors.New("not implemented")
}

func (efd *EventFD) Close() error {
	return nil
}

// ReadEvents returns the count of events that have occurred since the last
// call. If no events have transpired, blocks until at least one does.
func (efd *EventFD) ReadEvents() (count int64, err error) {
	return 0, errors.New("not implemented")
}

func (efd *EventFD) Fd() int {
	return -1
}
