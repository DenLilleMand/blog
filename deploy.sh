#!/bin/bash
set +e
sudo lsof -i tcp:80 | awk 'NR!=1 {print $2}' | sudo xargs kill
sudo rm -rf /home/www/denlillemand.com
sudo mkdir -p /home/www/denlillemand.com/static
sudo mkdir -p /home/www/denlillemand.com/templates
sudo mkdir -p /home/www/denlillemand.com/migrations
sudo cp -rf ./static/* /home/www/denlillemand.com/static/
sudo cp -rf ./templates /home/www/denlillemand.com/templates
sudo cp -rf ./migrations /home/www/denlillemand.com/migrations
sudo cp blog /home/www/denlillemand.com/
sudo cp goose /home/www/denlillemand.com/
sudo chmod +x /home/www/denlillemand.com/blog
sudo chmod +x /home/www/denlillemand.com/goose

/home/www/denlillemand.com/goose -dir="migrations/" postgres "user=denlillemand dbname=blog sslmode=disable" up

HAS_SESSION=$?
if [ $HAS_SESSION -eq 0 ]
then
    tmux attach-session -t $USER
else #$HAS_SESSION = 1
    tmux -2 new-session -s $SESSION
    tmux rename-window -t $SESSION:0 'Deployed code'

tmux send-keys "/home/denlillemand.com/blog --port=80 --dbuser=denlillemand --dbname=blog --static-dir='/home/www/denlillemand.com/static/' --template-dir='/home/www/denlillemand.com/templates/' 
"  Enter



