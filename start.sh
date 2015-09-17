#!/bin/bash
#Update mysl create use and database

#launch go
cd /srv/app/todo/src/evonck/todo/test
#go test
(cd /srv/app/todo/src/evonck/todo; go get)

(cd /srv/app/todo/src/evonck/todo; go install)
cd /srv/app/todo/src/evonck/todo
./todo 'root:@tcp(['${MYSQL_PORT_3306_TCP_ADDR}']:3306)/todo?charset=utf8&parseTime=True'



