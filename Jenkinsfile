pipeline {
    agent any
    triggers {
        pollSCM '* * * * *'
    }
    stages {
        stage('Build') {
            steps {
                echo "Building.."
                sh '''
                go build  superapp/main.go  -o superapp
                
                '''
            }
        }
        stage('Test') {
            steps {
                echo "Testing.."
                sh ''' echo 'some testing'
                '''
            }
        }
        stage('Deliver') {
            steps {
                echo 'Deliver....'
                sh '''
                echo "doing delivery stuff.."
                '''
            }
        }
    }
}