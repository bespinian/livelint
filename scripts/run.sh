#!/bin/bash
    
buildDate=$(date +'%Y-%m-%d_%T')
go build -o bin/livelint -ldflags "-X main.githash=`git rev-parse --short HEAD` -X main.date=$buildDate -X main.gittag=`git tag --points-at HEAD`"
echo ${buildDate}