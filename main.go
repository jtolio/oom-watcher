// Copyright (C) 2014 JT Olds, see provided LICENSE file

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spacemonkeygo/flagfile"
	"github.com/jtolds/oom-watcher/watch"
)

var (
	cgroupPath = flag.String("cgroup", "", "Path to the cgroup to monitor, "+
		"e.g.: /sys/fs/cgroup/memory/mycgroup")
)

func main() {
	flagfile.Load()
	if *cgroupPath == "" || len(flag.Args()) == 0 {
		fmt.Printf("usage: %s --cgroup <path> -- <subprocess> [args...]\n",
			os.Args[0])
		os.Exit(1)
	}

	panic(watch.NewWatcher(*cgroupPath).Watch(func(count int64) {
		log.Printf("%d ooms detected, running %v %d times",
			count, flag.Args(), count)

		for i := int64(0); i < count; i++ {
			cmd := exec.Command(flag.Arg(0), flag.Args()[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				log.Printf("failed running %v: %s", flag.Args(), err)
			}
		}
	}))
}
