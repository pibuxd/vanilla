#!bin/sh

wget --no-check-certificate -q --show-progress -r -np -nH --cut-dirs=3 -R index.html -P $HOME/.vanilla/packages/ "http://pibux.pl/repos/$1.tar.gz"

tar -xzf $HOME/.vanilla/packages/$1.tar.gz -C $HOME/.vanilla/packages/
cd $HOME/.vanilla/packages/$1 && sudo make build
