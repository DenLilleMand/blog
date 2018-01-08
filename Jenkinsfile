pipeline {
    agent none
    stages {
	stage('checkout') {
	    agent any
	    steps {
		echo "checking out!"
		checkout([$class: 'GitSCM', branches: [[name: '*/master']], doGenerateSubmoduleConfigurations: false, extensions: [], submoduleCfg: [], userRemoteConfigs: [[credentialsId: 'fbb643e7-f9da-4d49-9828-8533a44065a7', url: 'git@github.com:DenLilleMand/blog.git']]])
	    }
	}
	stage('build') {
	    agent any
	    steps {
		echo "Build step..."
		sh 'sudo chmod +x deploy.sh'
		sh './deploy.sh'
	    }
	}
	stage('deployment') {
	    agent {
		dockerfile {
		    args '-v /home/www/denlillemand.com/blog/:/root/blog/'
		}
	    }
	    steps {
		    sh 'chmod +x /root/blog/run.sh'
		    sh '/root/blog/run.sh'
	    }
	}
    }
}
