pipeline {
    agent any

    triggers {
        // Khoảng 2 phút hỏi git một lần, có commit mới thì tự build.
        pollSCM('H/2 * * * *')
    }

    environment {
        // Cùng tên với khi chạy "make up", để Jenkins cập nhật đúng container đó.
        COMPOSE_PROJECT_NAME = 'manager-life'
    }

    stages {
        stage('Test') {
            steps {
                sh 'go vet ./... && go test ./...'
            }
        }
        stage('Build') {
            steps {
                sh 'docker compose build'
            }
        }
        stage('Deploy') {
            steps {
                sh 'docker compose up -d'
            }
        }
        stage('Check') {
            steps {
                // Gọi thử API vừa deploy; thử lại vài lần vì service cần vài giây để lên.
                sh 'curl -fsS --retry 10 --retry-delay 3 --retry-all-errors http://host.docker.internal:8081/hello'
            }
        }
    }
}
