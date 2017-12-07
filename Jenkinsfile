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
		sh 'sudo cp -rf ./ /home/www/denlillemand.com/'
		sh 'sudo cd /home/www/denlillemand.com/'
		sh 'sudo ./blog --port=80 & disown'
	    }
	}
    }
}

