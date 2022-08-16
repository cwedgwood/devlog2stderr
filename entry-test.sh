#!/bin/sh

set -eu

# 'ps'
echo "-"
ls -l /proc/*/exe

/usr/local/sbin/devlog2stderr &

# 'ps'
echo "-"
ls -l /proc/*/exe

[ -e /dev/log ] || sleep 0.1 # logged needs time to start
[ -e /dev/log ] || ( echo "/dev/log not present" 1>&2 ; exit 1)

echo "-"
# test priorities
logger --id=1234 -p0 p0
logger --id=1234 -p1 p1
logger --id=4567 -p6 p6
logger --id=4567 -p7 p7

echo "-"
# test line-ending weirdness
( /bin/echo -en "one\r\ntwo\nthree " ) | logger

exit 0

