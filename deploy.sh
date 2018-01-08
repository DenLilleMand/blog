#!/bin/bash
set +e
sudo lsof -i tcp:80 | awk 'NR!=1 {print $2}' | sudo xargs kill
sudo rm -rf /home/www/denlillemand.com
sudo mkdir -p /home/www/denlillemand.com/static
sudo mkdir -p /home/www/denlillemand.com/templates
sudo cp -rf ./dist /home/www/denlillemand.com/static
sudo cp -rf ./templates /home/www/denlillemand.com/templates
sudo cp blog /home/www/denlillemand.com/
sudo chmod +x /home/www/denlillemand.com/blog
cd /home/www/denlillemand.com
./blog --port=80 --dbuser=denlillemand --dbname=blog --static-dir='/home/www/denlillemand.com/static/' --template-dir='/home/www/denlillemand.com/templates/' & disown


