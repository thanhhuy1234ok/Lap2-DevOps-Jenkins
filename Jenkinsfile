pipeline {
    agent any

    environment {
        DOCKER_IMAGE = 'thanhhuy1234ok/lap2-devops-jenkins'
        DOCKER_TAG = 'main'
    }

    stages {
        stage('Clone Repository') {
            steps {
                git branch: 'main', url: 'https://github.com/thanhhuy1234ok/Lap2-DevOps-Jenkins.git'
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    // Không sử dụng sudo, đảm bảo user Jenkins có quyền chạy Docker
                    sh 'docker build -t ${DOCKER_IMAGE}:${DOCKER_TAG} .'
                }
            }
        }

        stage('Run Tests') {
            steps {
                echo 'Running tests...'
            }
        }

        stage('Push to Docker Hub') {
            steps {
                script {
                    // Đăng nhập và đẩy image lên Docker Hub
                    docker.withRegistry('https://index.docker.io/v1/', 'docker-hub-credentials') {
                        docker.image("${DOCKER_IMAGE}:${DOCKER_TAG}").push()
                    }
                }
            }
        }

        stage('Deploy Golang to DEV') {
            steps {
                echo 'Deploying to DEV...'
                // Kéo image từ Docker Hub
                sh 'docker image pull ${DOCKER_IMAGE}:${DOCKER_TAG}'

                // Dừng container cũ nếu tồn tại
                sh 'docker container stop golang-jenkins || echo "this container does not exist"'

                // Tạo network nếu chưa có
                sh 'docker network create dev || echo "this network exists"'

                // Xóa container không dùng
                sh 'echo y | docker container prune'

                // Chạy container mới
                sh 'docker container run -d --rm --name server-golang -p 4000:3000 --network dev ${DOCKER_IMAGE}:${DOCKER_TAG}'
            }
        } 
    }
}