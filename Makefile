# never remove .vanilla directory!!!!!!!

build:
	sudo go build -o /bin/vanilla .
	mkdir ${HOME}/.vanilla
	cp -r ${HOME}/vanilla/data ${HOME}/.vanilla

remove:
	sudo rm /bin/vanilla

run:
	go run .
