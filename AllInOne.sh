#!/bin/bash

OUT_FILE="demo"

if [ $# -ne 0 ]
then
    # clean
    make distclean
    rm $OUT_FILE
else
    qmake
    make
    CC=gcc go build -o $OUT_FILE demo.go
    LD_LIBRARY_PATH=./ ./${OUT_FILE}
fi
