# Namecheap setup of dns
for jenkins.denlillemand.com i just did a URL redirect record, which amazingly works okay with the jenkins github plugin.

Then i have the std A Record for host @ with my ip 
and the CNAME record for www pointing to denlillemand.com

# Jenkins
How to handle the fact that it kills everything persistent:
https://wiki.jenkins.io/display/JENKINS/ProcessTreeKiller

#Setting up Jenkins & Github
Hmm ... i used the jenkins package on ubuntu, 
which pretty much gives it to you for free on port 8080.
Just followed the instructions. 

The integration with github is pretty terrible as allways. but i do think when you get over the fact that 
jenkins as an environment where you can generate standard snippets e.g. cloning from github, 
and that the github plugin can generate that id thingy needed its okay. 
I found something out about credentials in jenkins aswell, aparently even though you have created some credentials, if the e.g. github plugin is looking for a speicific kind e.g.  ssh, username/pass or secret text it will not show any of the other credentials you have created. 


# Jenkins and docker
Nice enough jenkins has built in support for running docker containers, i fought a long time trying to get forever.js execute a binary Go file but eventually gave up, i think docker is better as a future time investment anyway.

Remember to run: 

sudo usermod -a -G docker $USER
sudo usermod -a -G docker jenkins

And remember to reboot the VM for jenkins to get the changes and/or restart the SSH connection if you just want to 
run the commands your self


Mainly using: https//jenkins.io/doc/book/pipeline/docker/ as a reference

https://go.cloudbees.com/docs/cloudbees-documentation/cje-user-guide/index.html#docker-workflow

# Installing docker on ubuntu
https://docs.docker.com/engine/installation/linux/docker-ce/ubuntu/#install-docker-ce-1

# Setting up database

1.CREATE DATABASE blog;
2.CREATE USER denlillemand;
3.GRANT ALL PRIVILEGES ON DATABASE blog TO denlillemand;

Remember to log into the 'blog' database before running the following:
4.CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


# Creating migrations
$ goose -dir="migrations" postgres "user=<dbusrname> dbname=<dbname> sslmode=disable"  create <migrationname> <sql|go>

# Setting up SSL when just hosting a golang executable directly
TODO

ref: https://www.kaihag.com/https-and-go/
