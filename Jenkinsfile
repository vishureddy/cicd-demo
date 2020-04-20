pipeline{
        agent any
        tools {
               go 'default-go'
          } 
        environment{
                DOCKER_TAG = getDockerTag()
                DOCKER_REGISTRY_URL  = "localhost:5000"
                K8S_URL = "192.168.43.173:5000"
                IMAGE_URL_WITH_TAG = "${DOCKER_REGISTRY_URL}/go-app:${DOCKER_TAG}"
                K8S_IMAGE_URL_WITH_TAG = "${K8S_URL}/go-app:${DOCKER_TAG}"
        }
        stages{
                stage("SCM Checkout") {
                        steps {
                                git 'https://github.com/vishureddy/cicd-demo.git'
                        }
                }
                stage("Pre-Test---Getting Dependency packages") {
                        steps {
                                sh "go get -u golang.org/x/lint/golint"
                                sh "go get -u github.com/logrusorgru/aurora"
                        }
                }
                stage("Testing") {
                        steps {
                                sh "go fmt"
                                sh "go vet"
                                sh "go run -race ."
                                sh "go test"
                        }
                }
                stage('Buid Docker Image'){
                        steps{
                                sh "sudo docker build . -t ${IMAGE_URL_WITH_TAG}"
                        }

                }
                stage('Pushing Docker Image to local Registry'){
                        steps{
                                sh "sudo docker push ${IMAGE_URL_WITH_TAG}"
                        }
                }
                stage('Testing Docker image in the registry'){
                        steps{
                                sh "curl -X GET http://localhost:5000/v2/_catalog"
                        }
                }
                stage("Deploying the image to k8's cluster"){
                        steps{
                                sh "echo ${K8S_IMAGE_URL_WITH_TAG}"
                                sh "kubectl set image deployment/hello-go hello-go=${K8S_IMAGE_URL_WITH_TAG}"
                        }
                }

}
}

def getDockerTag(){
    def tag  = sh script: 'git rev-parse HEAD', returnStdout: true
    return tag
}
