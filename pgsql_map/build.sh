#!/bin/bash

SOURCE="$0"
while [ -h "$SOURCE"  ]; do # resolve $SOURCE until the file is no longer a symlink
    DIR="$( cd -P "$( dirname "$SOURCE"  )" && pwd  )"
    SOURCE="$(readlink "$SOURCE")"
    [[ $SOURCE != /*  ]] && SOURCE="$DIR/$SOURCE" # if $SOURCE was a relative symlink, we need to resolve it relative to the path where the symlink file was located
done
DIR="$( cd -P "$( dirname "$SOURCE"  )" && pwd  )"

out=`go install github.com/asyoume/postgres/pgsql_map  2>&1 >/dev/null`

if [ $? -eq 0 ];then
   echo  -e  "\033[32m程序编译成功,开始测试例子\033[0m"
   pgsql_map -in "$DIR/test.json"  -o "$DIR/_out/type.go"
else
    echo  -e  "\033[31m程序编译出错,请检查代码哦\033[0m"
    echo "$out"
fi