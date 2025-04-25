pipeline {
    agent {
        label 'agent3'
    }

    environment {
        GIT_URL = 'https://github.com/capelucita-roja/goo.git'
    }

    stages {
        stage('Clonar proyecto') {
            steps {
                git url: "${GIT_URL}", branch: 'master'
            }
        }

        stage('Verificar archivos') {
            steps {
                sh 'ls -R'
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
                // Ajusta el path si tu main.go está en otro lugar
                sh 'go build -o build/app ./cmd/doc'
            }
        }
    }

    post {
        failure {
            echo 'Falló el pipeline.'
        }
    }
}
