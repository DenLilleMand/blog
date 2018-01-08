#!/bin/bash
set +e
cd /home/www/denlillemand.com
./blog --port=80 --static='/home/www/denlillemand.com/static/' --templates='/home/www/denlillemand/templates/' & disown


