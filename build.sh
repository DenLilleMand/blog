#!/bin/bash
set +e
rm -rf static/*
webpack
go build
cp ../../../../bin/goose ./
cp -rf dist/* static/
go test -cover
node test_coverage.js
