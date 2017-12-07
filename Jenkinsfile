pipeline {
    agent any
    stages {
	stage('checkout') {
	    steps {
		git branch: 'master', credentialsId: '976f0493-2701-4900-89e0-4c472c49b1b7', url: 'git@github.com:DenLilleMand/blog.git'    
	    }
	}
	stage('compile') {
	    steps {
		dir('./') {
		    echo "Hello world!"
		}
	    }
	}
    }
}

