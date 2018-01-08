/*
Should add variables instead of all these hardcoded paths

*/
pipeline {
    agent any
    stages {
	stage('checkout') {
	    steps {
		echo "checking out!"
		checkout([$class: 'GitSCM', branches: [[name: '*/master']], doGenerateSubmoduleConfigurations: false, extensions: [], submoduleCfg: [], userRemoteConfigs: [[credentialsId: 'fbb643e7-f9da-4d49-9828-8533a44065a7', url: 'git@github.com:DenLilleMand/blog.git']]])
	    }
	}
	stage('compile') {
	    steps {
		dir('./') {
		    echo "Compiling"
		}
	    }
	}
	stage('deployment') {
	    steps {
		sh "sudo lsof -i tcp:80 | awk 'NR!=1 {print $2}' | sudo xargs kill"
		sh 'sudo rm -rf /home/www/denlillemand.com'
		sh 'sudo mkdir -p /home/www/denlillemand.com/static'
		sh 'sudo mkdir -p /home/www/denlillemand.com/templates'
		sh 'sudo cp -rf ./dist /home/www/denlillemand.com/static'
		sh 'sudo cp -rf ./templates /home/www/denlillemand.com/templates'
		sh 'sudo cp blog /home/www/denlillemand/'
		sh 'sudo chmod +x /home/www/denlillemand/blog'
		sh 'sudo chmod +x deploy.sh'
		sh 'deploy.sh'
	    }
	}
    }
}

