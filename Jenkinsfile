pipeline {
    agent DEVOPS

    environment {
        IMAGE_NAME = 'nikhil346/my-app'
    }

    stages {
        stage('Clone Repository') {
            steps {
                git 'https://github.com/nikhilvarma836/jenkins-pipeline-cicd.git'
            }
        }

        stage('Build') {
            steps {
                echo "Building Docker image..."
                sh 'docker build -t $Dockerfile .'
            }
        }

        stage('Test') {
            steps {
                echo 'Running tests (placeholder)...'
                sh 'echo "No tests implemented."'
            }
        }

        stage('Push to DockerHub') {
            steps {
                withCredentials([string(credentialsId: 'dockerhub-token', variable: 'DOCKERHUB')]) {
                    sh """
                        echo "$DOCKERHUB" | docker login -u nikhil346 --password-stdin
                        docker push $Dockerfile
                    """
                }
            }
        }

        stage('Deploy') {
            steps {
                echo 'Deploying app locally using Docker...'
                sh 'docker rm -f my-app || true'
                sh 'docker run -d -p 8080:8080 --name my-app $Dockerfile'
            }
        }
    }
}

