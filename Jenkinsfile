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
	    steps {
		sh 'sudo chmod +x run.sh'
		sh 'forever stopall'
		sh 'forever start -c /bin/bash run.sh'
	    }
	}
    }
}
