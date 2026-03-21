#!/bin/bash
set -e
echo "Building App"
chmod +x usr/bin/latin_checker
make deb
echo "Done!"