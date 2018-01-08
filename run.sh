#!/bin/bash
chmod +x /root/blog/blog
/root/blog/blog --port=80 --dbuser=denlillemand --dbname=blog --static-dir="/home/www/denlillemand.com/static/" --template-dir="/home/www/denlillemand.com/templates/"
