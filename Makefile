# Do not remove /etc/vanilla directory

.PHONY: all install remove clean setup

all: setup
	go build -o "./target/vanilla" .

setup:
	mkdir -p "/etc/vanilla"
	mkdir -p "./target"
	mkdir -p "/etc/vanilla/tmp"
	mkdir -p "/etc/vanilla/data"
	cp -r "./pkg/scripts" "/etc/vanilla"
	if [ ! -f "/etc/vanilla/data/installed-packages.json" ] ; then echo "[]" >> /etc/vanilla/data/installed-packages.json
	if [ ! -f "/etc/vanilla/config.json" ] ; then echo "{}" >> "/etc/vanilla/config.json"

clean:
	if [ -f "./target/vanilla" ] ; then rm "./target/vanilla"; fi

install:
	if [ -f "./target/vanilla" ] ; then install -sm 755 "./target/vanilla" "/bin/vanilla"; fi

remove:
	if [ -f "/bin/vanilla" ] ; then rm "/bin/vanilla"; fi
