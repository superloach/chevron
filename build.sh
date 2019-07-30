#!/bin/bash
cd $(dirname $BASH_SOURCE)
rm -rf ./build
mkdir -p ./build
cython -3 ./chevron.py -o ./build/chevron.c --embed
gcc -Os -I/usr/include/python3.7m -L/usr/lib/x86_64-linux-gnu -o ./build/chevron ./build/chevron.c -lpython3.7m -lpthread -lm -lutil -ldl
rm ./build/chevron.c
./build/chevron
