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
	stage('build') {
	    steps {
		echo "Build step..."
		sh 'sudo chmod +x deploy.sh'
		sh './deploy.sh'
	    }
	}
	stage('deployment') {
	    agent {
		docker( 
		    image 'ubuntu:latest'
		    args ' -p 80:80 -v /home/www/denlillemand.com/blog:/root/'
		)

	    }
	    stages {
		steps {
		    sh 'sudo chmod +x /root/run.sh'
		    sh '/root/run.sh'
		}
	    }
	}
    }
}

