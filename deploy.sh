#!/bin/bash
set +e
sudo lsof -i tcp:80 | awk 'NR!=1 {print $2}' | sudo xargs kill
sudo rm -rf /home/www/denlillemand.com
sudo mkdir -p /home/www/denlillemand.com/static
sudo mkdir -p /home/www/denlillemand.com/templates
sudo cp -rf ./dist /home/www/denlillemand.com/static
sudo cp -rf ./templates /home/www/denlillemand.com/templates
sudo cp blog /home/www/denlillemand/
sudo chmod +x /home/www/denlillemand/blog
cd /home/www/denlillemand.com
./blog --port=80 --static='/home/www/denlillemand.com/static/' --templates='/home/www/denlillemand/templates/' & disown


