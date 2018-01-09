#!/bin/bash
set +e
rm -rf static/*
webpack
go build
cp ../../../../bin/goose ./
cp -rf dist/* static/
go test -cover | grep % > test_coverage.txt
node test_coverage.js
rm test_coverage.txt
