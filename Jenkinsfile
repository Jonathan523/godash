pipeline {
    environment {
        VERSION = "v1.0.0"
        PROJECT_NAME = JOB_NAME.split('/')
        IMAGE_NAME = "flohoss/${PROJECT_NAME[0]}"
        IMAGE = ''
    }
    agent any
    stages {
        stage('Building') {
            steps {
                script {
                    IMAGE = docker.build IMAGE_NAME
                }
            }
        }
        stage('Deploying') {
            steps {
                script {
                    docker.withRegistry( 'https://ghcr.io', 'githubContainer' ) {
                        if (BRANCH_NAME == "main") {
                            IMAGE.push("${VERSION}")
                            IMAGE.push("latest")
                        } else {
                            IMAGE.push("${BRANCH_NAME}")
                        }
                    }
                }
            }
        }
    }
}
