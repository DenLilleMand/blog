#!/bin/bash
forever start -c /bin/bash /home/www/denlillemand.com/run.sh --port=80 --dbuser=denlillemand --dbname=blog --static-dir="/home/www/denlillemand.com/static/" --template-dir="/home/www/denlillemand.com/templates/"
