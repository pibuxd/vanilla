# Do not manually remove ~/.vanilla directory

.PHONY: all install uninstall clean setup

all: setup
	go build -o "./target/vanilla" .

setup:
	mkdir -p "./target"
	mkdir -p "${HOME}/.vanilla/tmp"
	cp -nr "./data" "${HOME}/.vanilla"
	cp -r "./pkg/scripts" "${HOME}/.vanilla"

clean:
	if [ -f "./target/vanilla" ] ; then rm "./target/vanilla"; fi

install:
	if [ -f "./target/vanilla" ] ; then install -sm 755 "./target/vanilla" "/bin/vanilla"; fi

uninstall:
	if [ -f "/bin/vanilla" ] ; then rm "/bin/vanilla"; fi
