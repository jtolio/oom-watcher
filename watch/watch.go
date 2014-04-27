// Copyright (C) 2014 JT Olds, see provided LICENSE file

package watch

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jtolds/oom-watcher/eventfd"
)

type Watcher struct {
	cgroup_path string
}

func NewWatcher(cgroup_path string) *Watcher {
	return &Watcher{cgroup_path: cgroup_path}
}

func (w *Watcher) Watch(oom_cb func(count int64)) error {
	events, err := eventfd.NewEventFD()
	if err != nil {
		return err
	}
	defer events.Close()

	oom_control, err := os.Open(filepath.Join(
		w.cgroup_path, "memory.oom_control"))
	if err != nil {
		return err
	}
	defer oom_control.Close()

	event_control, err := os.OpenFile(filepath.Join(
		w.cgroup_path, "cgroup.event_control"), os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(event_control,
		"%d %d\n", events.Fd(), oom_control.Fd())
	event_control.Close()
	if err != nil {
		return err
	}

	for {
		count, err := events.ReadEvents()
		if err != nil {
			return err
		}
		oom_cb(count)
	}
}
