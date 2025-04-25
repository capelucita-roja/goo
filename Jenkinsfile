pipeline {
    agent { label 'agent3' }

    environment {
        GO111MODULE = 'on'
    }

    stages {
        stage('Clonar proyecto') {
            steps {
                git url: 'https://github.com/capelucita-roja/goo.git', branch: 'main'
            }
        }

        stage('Preparar entorno') {
            steps {
                sh 'go mod tidy'
            }
        }

        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }

        stage('Build') {
            steps {
                sh 'go build -o build/app'
            }
        }
    }

    post {
        failure {
            echo "Falló el pipeline."
        }
        success {
            echo "Pipeline ejecutado correctamente en agent3."
        }
    }
}
