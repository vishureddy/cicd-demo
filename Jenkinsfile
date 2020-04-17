pipeline{
        agent any
        environment{
                DOCKER_TAG = getDockerTag()
                DOCKER_REGISTRY_URL  = "localhost:5000"
                IMAGE_URL_WITH_TAG = "${DOCKER_REGISTRY_URL}/go-app:${DOCKER_TAG}"
        }
        stages{
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
                stage("Deploy to k8's"){
                        steps{
                                sh "chmod +x changeTag.sh"
                                sh "./changeTag.sh ${DOCKER_TAG}"
                                sh "scp -o StrictHostKeyChecking=no services.yaml node-app-pod.yaml visweswara@:192.168.43.173/home/visweswara/jjj/"
                                script{
	                         try{
                                 sh "ssh visweswara@192.168.43.173 kubectl apply -f ."
                                }catch(error){
                                        sh "ssh visweswara@192.168.43.173 kubectl create -f ."

                                  }
                        }

                }
        }
}
}

def getDockerTag(){
    def tag  = sh script: 'git rev-parse HEAD', returnStdout: true
    return tag
}
