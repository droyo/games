#!/bin/sh
# acmewatch ./run.sh

pidfile=/tmp/mazewalker.pid
binfile=/tmp/mazewalker.bin
oldpid=$(cat $pidfile)
[ -n "$oldpid" ] && kill $oldpid
echo $$ > $pidfile
go build -o $binfile && exec $binfile
