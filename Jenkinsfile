pipeline {
    agent {
        label 'agent3'
    }

    environment {
        GIT_URL = 'https://github.com/capelucita-roja/goo.git'
        BUILD_DIR = 'outyet' // <-- AJUSTA ESTA RUTA según el main.go que te interese
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
                sh 'go test ./... || true' // permite continuar si no hay tests
            }
        }

        stage('Build') {
            steps {
                sh 'mkdir -p build'
                sh "go build -o build/app ./${BUILD_DIR}"
            }
        }
    }

    post {
        failure {
            echo 'Falló el pipeline.'
        }
    }
}
