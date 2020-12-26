#!bin/sh

curl -X POST http://pibux.pl:2137/upload -F "upload[]=@$1" -F "upload[]=@$1.txt" -H "Content-Type: multipart/form-data"
