bin-dir:=bin
cmd-dir:=cmd
bins:=$(addprefix $(bin-dir)/,$(subst $(cmd-dir)/,,$(shell find $(cmd-dir) -type d -mindepth 1 -maxdepth 1)))
srcs:=$(shell find . -type f -name "*.go")

build: $(bins)

$(bin-dir)/%: $(cmd-dir)/%/main.go $(srcs)
	go build -o $@ $<

test:
	@echo $(bin-dir)
	@echo $(cmd-dir)
	@echo $(bins)
