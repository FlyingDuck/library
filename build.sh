#!/usr/bin/env bash
RUN_NAME="library"

mkdir -p output/bin output/conf output/assets/upload
cp bootstrap.sh output/
chmod +x output/bootstrap.sh

#echo "gofmt start"
gofmt -l -w -s .
#echo "gofmt done"

#echo "build main program"
go build -tags=jsoniter -o output/bin/${RUN_NAME}
echo "build completed: output/bin/${RUN_NAME}"