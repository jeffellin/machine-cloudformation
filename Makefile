default: build

version := "v0.1.0"
version_description := "Docker Machine Plugin for Amazon Cloud Formation"
human_name := "Cloud Formation Driver"

mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := "docker-machine-driver-amazoncf"
bin_suffix := ""

clean:
	rm bin/docker-machine*

compile:
	GOGC=off go build -o bin/$(current_dir)$(BIN_SUFFIX) docker-machine-driver-amazoncf/machine-driver-amazoncf.go

print-success:
	@echo
	@echo "Plugin built."
	@echo
	@echo "To use it, either run 'make install' or set your PATH environment variable correctly."

build: compile print-success

cross:
	for os in darwin windows linux; do \
		for arch in amd64 386; do \
			GOOS=$$os GOARCH=$$arch BIN_SUFFIX=_$$os-$$arch $(MAKE) compile & \
		done; \
	done; \
	wait

install:
	cp bin/$(current_dir) /usr/local/bin/$(current_dir)
