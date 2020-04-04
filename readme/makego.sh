#!/bin/sh

OSName=`uname`
OSName=${OSName:0:5}

debugflag=0

export CURDIR=`pwd`
export GOP1=${GOPATH}

if test $debugflag -ne 0; then
    echo CURDIR 0--${CURDIR:0:1}--
    echo CURDIR 1--${CURDIR:1:1}--
    echo CURDIR 2--${CURDIR:2:1}--
    echo GOPATH--${GOPATH}--
fi

# 将 /d/ 改成 d:/
	if [ ! "$OSName" = "Linux" ]; then
		export CURDIR=${CURDIR:1:1}":"${CURDIR:2}
		echo "change dir for win: CURDIR="${CURDIR}
        # CURDIR=${CURDIR/\//} # 去掉第一个 / 
        # CURDIR=${CURDIR/\//:/} # 第二个/改为 :/
        export GOP1=${GOPATH//\\/\/} # 所有 \ 转为 /
	fi

export PROTOC=${GOP1}/bin/protoc.exe
export PROTO_DIR=${CURDIR}/pbproto
export PBGO_DIR=${CURDIR}/github.com/futuopen/ftapi4go/pb
# pbgo
export PLUGIN=gogofaster

# 典型值
# PROTOC:E:/gopath/bin/protoc.exe
# CURDIR:/e/futu/api/v3.20
# PROTO_DIR:/e/futu/api/v3.20/pbproto
# PBGO_DIR:/e/futu/api/v3.20/github.com/futuopen/ftapi4go/pb

echo ;
echo PROTOC:${PROTOC}
echo CURDIR:${CURDIR}
echo PROTO_DIR:${PROTO_DIR}
echo PBGO_DIR:${PBGO_DIR}
echo ;

# exit 0 # debug

if [ ! -d "${PBGO_DIR}" ]; then
    if test $debugflag -ne 0; then
        echo 创建文件夹 "${PBGO_DIR}"
    fi
  mkdir -p "${PBGO_DIR}"
fi

for file in ${PBGO_DIR}/*; do
    if test $debugflag -ne 0; then
        echo 删除 $file
    fi
    rm -rf $file
done

# 含路径
# for file in ${PROTO_DIR}/*; do
#     echo $file
# done

inc1=0
for file in `ls ${PROTO_DIR}/*.proto`
do
    let inc1++
    
    if test $debugflag -ne 0; then
    # debug
        if test $inc1 -le 8; then
            continue
        fi
        echo seq=${inc1}
    fi
 
    name=`echo ${file##*/} | awk -F. '{print $1}'`
    if test $debugflag -ne 0; then
        echo "file=$file, name=$name"
    fi

    typeset -l subdir
    subdir=${name//_/}
    if test $debugflag -ne 0; then
        echo subdir=${subdir}; echo '';
    fi
    
    echo "--seq=${inc1}-- in $file"
    
    if [ ! -d "${PBGO_DIR}/${subdir}" ]; then
        if test $debugflag -ne 0; then
            echo 创建文件夹 "${PBGO_DIR}/${subdir}"
        fi
        mkdir -p "${PBGO_DIR}/${subdir}"
    fi
# e.g. E:/gopath/bin/protoc.exe --proto_path=e:/futu/api/v3.20/pbproto --gogofaster_out=paths=source_relative:e:/futu/api/v3.20/github.com/futuopen/ftapi4go/pb/qotgetbasicqot Qot_GetBasicQot.proto
    if test $debugflag -ne 0; then
        echo 命令 "${PROTOC}" --proto_path="${PROTO_DIR}" --${PLUGIN}_out=paths=source_relative:"${PBGO_DIR}/${subdir}" ${name}.proto
    fi
    
    "${PROTOC}" --proto_path="${PROTO_DIR}" --${PLUGIN}_out=paths=source_relative:"${PBGO_DIR}/${subdir}" ${name}.proto

    if test $debugflag -ne 0; then
    # debug
        if test $inc1 -ge 10; then
            break
        fi
    fi
done

echo "";
read -p "Press enter key to exit " var
echo "";
