pipeline {
  agent {
    node {
      // spin up a node.js slave pod to run this build on
      label 'go11'
    }
  }
  options {
    // set a timeout of 20 minutes for this pipeline
    timeout(time: 20, unit: 'MINUTES')
  }
  stages {
    stage('preamble') {
      steps {
        script {
          openshift.withCluster() {
            openshift.withProject() {
              echo "Using project: ${openshift.project()}"
              echo "Using project: ${env.GIT_COMMIT}"
              echo "Using project: ${env.GIT_BRANCH}"
            }
          }
        }
      }
    }
    stage('checkout') {
      steps {
        script {
          openshift.withCluster() {
            openshift.withCredentials() {
              openshift.withProject() {
                checkout([$class           : 'GitSCM',
                          branches         : [[name: "*/*"]],
                          userRemoteConfigs: [[url: "https://github.com/tsungming/alameda.git"]]
                ]);
                dir("${WORKSPACE}") {
                  def commit_id = sh(returnStdout: true, script: 'git rev-parse --short HEAD').trim()
                        echo "${commit_id}"
                }
              }
            }
          }
        }
      }
    }
    stage('build') {
      steps {
        setGitHubPullRequestStatus context: 'test', message: 'test', state: 'SUCCESS'
      }      
    }
  }
}