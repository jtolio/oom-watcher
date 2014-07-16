// Copyright (C) 2014 JT Olds, see provided LICENSE file
// +build !linux

package eventfd

type EventFD struct{}

func NewEventFD() (*EventFD, error) {
	return new(EventFD), nil
}

func (efd *EventFD) Close() error {
	return nil
}

// ReadEvents returns the count of events that have occurred since the last
// call. If no events have transpired, blocks until at least one does.
func (efd *EventFD) ReadEvents() (count int64, err error) {
	select {}
}

func (efd *EventFD) Fd() int {
	return -1
}
