#!/bin/bash
#Update mysl create use and database

#let mysql some time to create the databases
sleep 2
#grab the godep library
go get github.com/tools/godep
#install the binary
cd /go/src/evonck/todo
godep restore
rm -r ./Godeps 
godep save
godep go install
cd /go/bin
#launch or API
./todo 'root:root@tcp(['${MYSQL_PORT_3306_TCP_ADDR}']:3306)/todo?charset=utf8&parseTime=True'
