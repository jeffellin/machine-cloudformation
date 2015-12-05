default: build
github_user := "jeffellin"
project := "github.com/$(github_user)/$(current_dir)"
version := "v0.1.2"
version_description := "Docker Machine Plugin for Amazon Cloud Formation"
human_name := "Cloud Formation Driver"
export GO15VENDOREXPERIMENT = 1
mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := "docker-machine-driver-amazoncf"
repo := "machine-cloudformation"
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

cleanrelease:
	github-release delete \
		--user $(github_user) \
		--repo $(repo) \
		--tag $(version)
	git tag -d $(version)
	git push origin :refs/tags/$(version)

release:  
	git tag $(version)
	git push --tags
	github-release release \
		--user $(github_user) \
		--repo $(repo) \
		--tag $(version) \
		--name $(human_name) \
		--description $(version_description)
	for os in darwin windows linux; do \
		for arch in amd64 386; do \
			github-release upload \
				--user $(github_user) \
				--repo $(repo) \
				--tag $(version) \
				--name $(current_dir)_$$os-$$arch \
				--file bin/$(current_dir)_$$os-$$arch; \
		done; \
	done
