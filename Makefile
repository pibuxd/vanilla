# never remove .vanilla directory!!!!!!!

build:
	sudo go build -o /bin/vanilla .
	mkdir -p ${HOME}/.vanilla
	mkdir -p ${HOME}/.vanilla/tmp
	cp -r -n ${HOME}/vanilla/data ${HOME}/.vanilla
	cp -r ${HOME}/vanilla/pkg/scripts ${HOME}/.vanilla

remove:
	sudo rm /bin/vanilla
	rm -rf ${HOME}/.vanilla/scripts
