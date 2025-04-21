pipeline{
  agent any
  stages {
  stage('maven install') {
    steps {
         withMaven( maven: 'MyMaven'){
              bat 'mvn clean install'
          }
    }
  }

 }
}
