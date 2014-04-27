// Copyright (C) 2014 JT Olds, see provided LICENSE file
// +build linux

package eventfd

/*
#include <sys/eventfd.h>

int64_t hostEndian64(void *data) {
  return *((int64_t *)data);
}
*/
import "C"

import (
	"fmt"
	"runtime"
	"syscall"
	"unsafe"
)

type EventFD struct {
	fd     int
	closed bool
}

func NewEventFD() (*EventFD, error) {
	fd, err := C.eventfd(0, C.EFD_CLOEXEC)
	if err != nil {
		return nil, err
	}
	if fd == -1 {
		return nil, fmt.Errorf("unknown eventfd error")
	}
	efd := &EventFD{fd: int(fd)}
	runtime.SetFinalizer(efd, func(efd *EventFD) {
		efd.Close()
	})
	return efd, nil
}

func (efd *EventFD) Close() error {
	if efd.closed {
		return nil
	}
	efd.closed = true
	runtime.SetFinalizer(efd, nil)
	return syscall.Close(efd.fd)
}

// ReadEvents returns the count of events that have occurred since the last
// call. If no events have transpired, blocks until at least one does.
func (efd *EventFD) ReadEvents() (count int64, err error) {
	var buf [8]byte
	n, err := syscall.Read(efd.fd, buf[:])
	if err != nil {
		return 0, err
	}
	if n != 8 {
		return 0, fmt.Errorf("eventfd returned less than 8 bytes: %d", n)
	}
	count = int64(C.hostEndian64(unsafe.Pointer(&buf[0])))
	return count, nil
}

func (efd *EventFD) Fd() int {
	return efd.fd
}
