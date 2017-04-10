#! /bin/bash
echo "ready to build the ss_golang"

# create objs folder
mkdir -p objs

go build -o ./objs/ss_golang ./src/

echo "build ss_golang successfully"
