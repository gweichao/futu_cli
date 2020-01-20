#!/bin/sh

# 根据情况修改如下变量，程序名称
	export PROGNAME=futu_cli

	OSName=`uname`
	OSName=${OSName:0:5}
	SPLIT=":"
	EXT=""

	export MAINP=`pwd`
# 将 /d/ 改成 d:/
	if [ ! "$OSName" = "Linux" ]; then
		echo "	change dir for win "
		export MAINP=${MAINP:1:1}":"${MAINP:2}
		export SPLIT=";"
		export EXT=".exe"
	fi
# ${string:position:length}
# =两边不加空格
#`pwd`  /d/gocode/....
#$(cd `dirname $0`; pwd) /d/gocode/...

	export GOPATH="$MAINP/gopath$SPLIT$MAINP$SPLIT$GOPATH"

	echo 
	echo "	该脚本编译输出到 ./bin/$PROGNAME$EXT "
	echo "	-- 提示，要运行 go test，请到 ./src/$PROGNAME，运行参考genv.sh "
	echo 
	
	echo "	当前gopath：$GOPATH"
#	echo "	MAINP：$MAINP"

	cd ./src/$PROGNAME
	go build -ldflags "-s -w" -o ../../bin/$PROGNAME$EXT
	cd ../../
	
	echo 
	echo 
	echo "	编译结果参见:"
	echo 
	dir ./bin/$PROGNAME* -l

	echo 
	echo 
	echo "	完成"
	echo 
