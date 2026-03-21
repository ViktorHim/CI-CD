#!/bin/bash

echo "Building App..."
cd src
make clean
make
cd ..
echo "App was been builded."