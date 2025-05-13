pipeline {
    agent any

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
                sh 'docker build -t nikhil346/my-app .'
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
                echo 'Pushing Docker image to DockerHub...'
                withCredentials([usernamePassword(credentialsId: "$DOCKER_HUB_CREDENTIALS", usernameVariable: 'nikhil346', passwordVariable: 'DOCKERHUB_TOKEN')]) {
                    sh '''
                        echo "$DOCKERHUB_TOKEN" | docker login -u "nikhil346" --password-stdin
                        docker push nikhil346/my-app
                    '''

                }
            }
        }

        stage('Deploy') {
            steps {
                echo 'Deploying app locally using Docker...'
                sh 'docker rm -f my-app || true'
                sh 'docker run -d -p 8080:8080 --name my-app nikhil346'
            }
        }
    }
}
