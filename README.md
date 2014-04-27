oom-watcher
===========

This command can be run to watch a cgroup for OOM events and run some
subcommand.

Example usage:

	go build .
	./oom-watcher --cgroup /sys/fs/cgroup/memory/mycgroup -- \
		mail -s oom oom-alert@yourco.com
