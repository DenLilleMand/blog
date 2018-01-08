#!/bin/bash
set +e
#sudo lsof -i tcp:80 | awk 'NR!=1 {print $2}' | sudo xargs kill
forever stopall
sudo rm -rf /home/www/denlillemand.com
sudo mkdir -p /home/www/denlillemand.com/static
sudo mkdir -p /home/www/denlillemand.com/templates
sudo mkdir -p /home/www/denlillemand.com/migrations
sudo cp -rf ./static/* /home/www/denlillemand.com/static/
sudo cp -rf ./templates /home/www/denlillemand.com/templates
sudo cp -rf ./migrations /home/www/denlillemand.com/migrations
sudo cp blog /home/www/denlillemand.com/
sudo cp run.sh /home/www/denlillemand.com/
sudo cp goose /home/www/denlillemand.com/
sudo chmod +x /home/www/denlillemand.com/blog
sudo chmod +x /home/www/denlillemand.com/goose

/home/www/denlillemand.com/goose -dir="migrations/" postgres "user=denlillemand dbname=blog sslmode=disable" up

forever start run.sh
