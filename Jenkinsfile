

pipeline {
    agent { label 'agent3' }

    stages {
        stage('Checkout') {
            steps {
                git  'https://github.com/capelucita-roja/goo.git', branch: 'master'
            }
        }
        stage('Test') {
            steps {
                sh 'go test ./... -v -json > result.json'
            }
        }
    }


}