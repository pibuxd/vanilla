#!bin/sh

name=$1

curl -X POST http://pibux.pl:2137/upload -F "upload[]=@$1" -F "upload[]=@$1.json" -H "Content-Type: multipart/form-data"
