#!/bin/sh

HTTPServerLog =`grep \"log\" | grep HTTPSever | awk -F'\"' '{print $4}'`".INFO"

case $1 in
http)
	echo "监视日志文件:" $HTTPServerLog
	tail -f $HTTPServerLog
;;
*)
	tail -f $HTTPServerLog
;;
esac