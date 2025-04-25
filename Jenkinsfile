
pipeline {
    agent { label 'agent3' }

    stages {
        stage('Clonar repositorio') {
            steps {
                git 'https://github.com/capelucita-roja/goo.git'
            }
        }

        stage('Instalar go-junit-report') {
            steps {
                sh 'go install github.com/jstemmer/go-junit-report@latest'
            }
        }

        stage('Build') {
            steps {
                sh '''
                    cd gotests
                    go build ./...
                '''
            }
        }

        stage('Ejecutar pruebas') {
            steps {
                sh '''
                    go version
                    cd gotests
                    go mod tidy
                    go test -v ./... | /home/jenkins/go/bin/go-junit-report > result.xml
                '''
            }
            post {
                always {
                    junit '**/result.xml'
                }
            }
        }
    }
}