#!/bin/bash
set +e
rm -rf static/*
webpack
go build
cp ../../../../bin/goose ./
cp ../../../../bin/sup ./
cp -rf dist/* static/
