#!bin/sh

name=$1

curl -X POST http://pibux.pl:2137/upload -F "file=@$1" -H "Content-Type: multipart/form-data"
