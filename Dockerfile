#! /bin/bash
FROM ubuntu:latest

RUN apt update 

RUN mkdir -p /root/blog/
RUN chmod -R 777 /root/blog/
