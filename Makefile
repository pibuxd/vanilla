build:
	sudo go build -o /usr/bin/vanilla .
	mkdir ${HOME}/vanilla
	cp -r scripts ${HOME}/vanilla

remove:
	sudo rm /usr/bin/vanilla
	rm ${HOME}/vanilla

run:
	go run .
