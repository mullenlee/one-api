properties([
    parameters([
        gitParameter(
            branch: '',
            branchFilter: 'origin/(.*)',
            defaultValue: 'master',
            description: '选择分支',
            name: 'BRANCH',
            quickFilterEnabled: true,
            filterLength: 1,
            filterable: true,
            selectedValue: 'NONE',
            sortMode: 'NONE',
            tagFilter: '*',
            type: 'PT_BRANCH_TAG'
        ),
        [$class: 'BooleanParameterDefinition',
            name: 'CODE_CHECK',
            defaultValue: false,
            description: '代码检查'
        ],
        [
            $class: 'ChoiceParameter',
            choiceType: 'PT_SINGLE_SELECT',
            description: '请选择部署方式',
            filterLength: 1,
            filterable: false,
            name: 'DEPLOY_MODE',
            script: [
                $class: 'GroovyScript',
                script: [
                    classpath: [],
                    sandbox: true,
                    script:'''
                       return ['DKCompose', 'Docker','K8s','SpringBoot', 'buildTar']
                    '''
                ]
            ]
        ],
        [
            $class: 'ChoiceParameter',
            choiceType: 'PT_SINGLE_SELECT',
            description: '请选择部署环境',
            filterLength: 1,
            filterable: false,
            name: 'DEPLOY_PROFILE',
            script: [
                $class: 'GroovyScript',
                fallbackScript: [
                    classpath: [],
                    sandbox: false,
                    script:
                        "return ['Could not get The environemnts']"
                ],
                script: [
                    classpath: [],
                    sandbox: true,
                    script:'''
                       return ['dev', 'beta', 'prod']
                    '''
                ]
            ]
        ],
        [
            $class: 'CascadeChoiceParameter',
            choiceType: 'PT_SINGLE_SELECT',
            description: '请选择部署服务器',
            name: 'REMOTES',
            referencedParameters: 'DEPLOY_PROFILE',
            script:
                [
                $class: 'GroovyScript',
                fallbackScript: [
                    classpath: [],
                    sandbox: false,
                    script: "return ['Could not get Environment from Env Param']"
                    ],
                script: [
                    classpath: [],
                    sandbox: true,
                    script: '''
                    if (DEPLOY_PROFILE.startsWith("dev")){
                      return ["请选择","10.1.3.113","10.1.3.112","10.1.3.125"]
                    }
                    else if(DEPLOY_PROFILE.startsWith("beta")){
                      return ["请选择","10.1.3.123","10.1.3.122"]
                    }
                    else if(DEPLOY_PROFILE.startsWith("prod")){
                      return ["请选择","10.0.34.221"]
                    }
                    '''
                ]
            ]
        ],
        [
            $class: 'DynamicReferenceParameter',
            choiceType: 'ET_FORMATTED_HTML',
            name: 'BRANCH_RELEASE',
            description: '提测首次构建请设置RELEASE分支名称(命名规范: release/x_x_x)',
            omitValueField: true,
            referencedParameters: 'DEPLOY_PROFILE,BRANCH',
            script: [
                $class: 'GroovyScript',
                script: [
                    classpath: [],
                    sandbox: true,
                    script: """
                       if(DEPLOY_PROFILE.equals("beta") && BRANCH.equals("master")){
                            return '<input name="value" value="release/" type="text">'
                       }else {
                            return '<p>不需要</p>'
                       }
                    """
                ]
            ]
        ],
        [
            $class: 'DynamicReferenceParameter',
            choiceType: 'ET_FORMATTED_HTML',
            name: 'BRANCH_TAG',
            description: '版本首次发布请设置发布TAG分支名称(命名规范: -alpha|beta|rc|r|stable)',
            omitValueField: true,
            referencedParameters: 'DEPLOY_PROFILE,BRANCH',
            script: [
                $class: 'GroovyScript',
                script: [
                    classpath: [],
                    sandbox: true,
                    script: """
                       if(DEPLOY_PROFILE.equals("prod") && BRANCH.startsWith("release")){
                            return '<input name="value" value="-r" type="text">'
                       }else {
                            return '<p>不需要</p>'
                       }
                    """
                ]
            ]
        ],
        [
            $class: 'DynamicReferenceParameter',
            choiceType: 'ET_FORMATTED_HTML',
            name: 'RELEASE_PERIOD',
            description: '删除指定时间周期没有提交记录的Releas分支',
            omitValueField: true,
            referencedParameters: 'DEPLOY_PROFILE',
            script: [
                $class: 'GroovyScript',
                script: [
                    classpath: [],
                    sandbox: true,
                    script: """
                       if(DEPLOY_PROFILE.equals("prod")){
                            return '<select name="value" value="none"><option value="none">none</option><option value="7days">7天</option><option value="14days">14天</option><option value="30days">30天</option><option value="once">保留最近一次</option></select>'
                       }else {
                            return '<p>不需要</p>'
                       }
                    """
                ]
            ]
        ]
    ])
])

pipeline {
  agent any
  tools {
    git 'Default'
    maven 'maven3.6.0'
    jdk 'jdk1.8'
  }
  options {
  	buildDiscarder logRotator(artifactDaysToKeepStr: '', artifactNumToKeepStr: '', daysToKeepStr: '10', numToKeepStr: '3')
  }
  environment {
    JAR_SHORT_NAME = 'app'
  }
  stages {
  	stage('编译清单') {
      steps {
        
        echo "--------------------------------外部参数 start-----------------------------------"
        loadProjectConfig()
        echo "--------------------------------外部参数 end  -----------------------------------"

        echo "代码检查 CODE_CHECK: ${params.CODE_CHECK}"
        echo "Release保留周期 RELEASE_PERIOD: ${RELEASE_PERIOD}"
        echo "提测首次构建 BRANCH_RELEASE: ${params.BRANCH_RELEASE}"
        echo "发布首次构建 BRANCH_TAG: ${params.BRANCH_TAG}"
      	echo "--------------------------------环境清单 start-----------------------------------"
      	sh "git --version"
      	sh "java -version"
      	sh "mvn -v"
      	echo "--------------------------------环境清单 end  -----------------------------------"

      	echo "--------------------------------软件清单 start-----------------------------------"
      	echo "部署方式：${params.DEPLOY_MODE}"
        
      	echo "编译环境：${params.DEPLOY_PROFILE}"
      	echo "项目空间(弃用)：${params.NAMESPACE}"

        echo "部署模块：${env.MODULE}"
        echo "运行端口：${params.ORIGINAL_PORT}"
        echo "部署分支：${params.BRANCH}"
         echo "部署服务器：${params.REMOTES}"
        echo "--------------------------------软件清单 end  -----------------------------------"      
      }
    }
    stage('检出代码') {
      steps {
        sh 'git checkout ${BRANCH}'
        withCredentials([usernamePassword(credentialsId: 'ops-jenkins', passwordVariable: 'GIT_PASSWORD', usernameVariable: 'GIT_USERNAME')]) {
            sh '''
            git config --local credential.helper "!p() { echo username=\\$GIT_USERNAME; echo password=\\$GIT_PASSWORD; }; p"
            git pull origin ${BRANCH}
            '''
        }
      }
    }
    stage('创建测试RELEASE分支') {
      when {
        expression { params.BRANCH_RELEASE != '' }
      }
      environment {
        MVN_VERSION = getReleaseVersion("${params.BRANCH_RELEASE}")
      }
      steps {
        echo "RELEASE名称: ${params.BRANCH_RELEASE}"
        echo "MVN RELEASE名称: ${env.MVN_VERSION}"
        sh """
          git checkout -b ${params.BRANCH_RELEASE}
          mvn versions:set -DnewVersion="${env.MVN_VERSION}" -s /data/requried/maven-3.6.0/conf/tudou_repo_settings.xml
          git add .
          git commit -m "提测第一次构建"
          git push -u origin ${params.BRANCH_RELEASE}
          mvn clean deploy -Dmaven.test.skip=true -P${DEPLOY_PROFILE} -s /data/requried/maven-3.6.0/conf/tudou_repo_settings.xml
        """
      }
    }
    stage('创建发布tag') {
      when {
        expression { params.BRANCH_TAG != '' }
      }
      steps {
        echo "TAG名称: ${params.BRANCH_TAG}"
        sh "git tag -a ${params.BRANCH_TAG} -m 'tag by ${params.BRANCH}'"
        sh "git push -u origin ${params.BRANCH_TAG}"
      }
    }
    stage('代码检查') {
      when {
        expression { params.CODE_CHECK == true}
      }
      steps {
        echo '开始代码检查'
        script {
          def scannerHome = tool 'SQScanner'
          echo "SQScanner path: ${scannerHome}"
          withSonarQubeEnv('SonarQubeServer') {
        	  sh "${scannerHome}/bin/sonar-scanner"
          }
        }
        
      }
    }
    stage('获取版本号') {
      when {
         expression { env.MODULE!="" }
      }
      steps {
      	script {
            echo "非多模块项目结构"

            //dockerImage = docker.build("${env.MODULE}:latest", "-f Dockerfile .")

            // 读取 pyproject.toml 文件内容
            def tomlContent = readFile('pyproject.toml')
            // 使用正则表达式提取 [tool.poetry] 下的 version 字段
            def versionMatch = tomlContent =~ /\[tool\.poetry\][^\[]+?version = "(.*?)"/
            if (versionMatch) {
                // 将版本保存到环境变量
                POETRY_VERSION = versionMatch[0][1]
                echo "Poetry version: ${POETRY_VERSION}"
                 env.HARBOR_VERSION=POETRY_VERSION
            } else {
                error "Unable to extract Poetry version from pyproject.toml"
            }
        }
      }
    }
    stage('编译指定模块') {
      when {
        expression { env.MODULE!="" }
      }
      steps {
        script {
            echo "发布版本 ${env.HARBOR_VERSION}"
            dockerImage = docker.build("${env.HARBOR_ADDR}/${env.HARBOR_REPO}/${env.MODULE}:${env.HARBOR_VERSION}", "-f Dockerfile .")
            dockerImage.push()
        }
      }
    }
    stage('Docker/DKCompose方式部署') {
      when {
        expression { params.DEPLOY_MODE == 'Docker' || params.DEPLOY_MODE == 'DKCompose'}
        expression { params.REMOTES != '' && params.REMOTES != '请选择'}
      }
      steps {
        echo "项目版本：${env.HARBOR_VERSION}"
        script {
          echo "发布 ${env.HARBOR_ADDR}/${env.HARBOR_REPO}/${env.MODULE}:${env.HARBOR_VERSION}"
          for(REMOTE in params.REMOTES.tokenize(',')){
              echo "部署服务器：${REMOTE} ${DEPLOY_MODE}"
              sh """
                  ssh -o StrictHostKeyChecking=no -l root ${REMOTE} 'bash ${env.HOST_BASEDIR}/deploy.sh ${DEPLOY_MODE} ${env.HARBOR_ADDR} ${env.HARBOR_REPO} ${env.MODULE} ${env.HARBOR_VERSION} ${REMOTE} ${env.ORIGINAL_PORT} ${DEPLOY_PROFILE} ${env.META_CLOUD} ${params.NAMESPACE} 2>&1'
              """
              echo "10S时间，喝杯水等等"
              sleep 10
              echo "部署服务器健康检查：${REMOTE}"
              sh """
                  ssh -o StrictHostKeyChecking=no -l root ${REMOTE} 'bash ${env.HOST_BASEDIR}/health.sh ${env.MODULE} 2>&1'
              """
            }
      	}
      }
    }
    stage('K8s方式部署') {
      when {
        expression { params.DEPLOY_MODE == 'K8s' }
      }
      environment {
        JAR_SHORT_NAME = getShEchoResult("basename '${env.JAR_NAME}' .jar")
      }
      steps {
        echo '开始K8s方式部署'
        echo "发布 ${env.HARBOR_ADDR}/${env.HARBOR_REPO}/${JAR_SHORT_NAME}:${env.HARBOR_VERSION}"
        script {
          for(REMOTE in params.REMOTES.tokenize(',')){
            echo "部署服务器：${REMOTE}"
            sh """
		        sh -x deploy.sh ${DEPLOY_MODE} ${env.HARBOR_ADDR} ${env.HARBOR_REPO} ${JAR_SHORT_NAME} ${env.HARBOR_VERSION} ${REMOTE} ${env.META_CLOUD}
		        """
            echo "部署服务器健康检查：${REMOTE}"
            sh """
                cat /opt/www/tools/cicd/devops-scripts/health_check.sh
		        """
          }
        }
      }
    }
    stage('清理RELEASE分支') {
      when {
        expression { params.RELEASE_PERIOD != null && params.RELEASE_PERIOD != '' && params.RELEASE_PERIOD != 'none' }
      }
      steps {
        sh "bash /opt/www/tools/cicd/devops-scripts/clean_release.sh ${params.RELEASE_PERIOD}"
      }
    }
    stage('清理工作区缓存') {
      steps {
        cleanWs(
            cleanWhenAborted: true,
            cleanWhenFailure: true,
            cleanWhenNotBuilt: true,
            cleanWhenSuccess: true,
            cleanWhenUnstable: true,
            disableDeferredWipeout: true,
            deleteDirs: true
        )
      }
    }
  }
  post {
    always { 
      echo 'I will always say Hello again!'
    }
  }
}

def getReleaseVersion(nversion){
  // 将字符串分割成一个列表
    def parts = nversion.split('/')

    // 将第一个元素转换为大写并添加到结果中
    def result = "${parts[1]}-${parts[0].toUpperCase()}"

    // 将剩余元素添加到结果中，并用下划线（_）进行连接
    for (int i = 2; i < parts.size(); i++) {
        result += "_${parts[i]}"
    }
    return result
}

// 获取 shell 命令输出内容
def getShEchoResult(cmd) {
    def getShEchoResultCmd = "ECHO_RESULT=`${cmd}`\necho \${ECHO_RESULT}"
    // 使用 returnStdout 时，返回的字符串末尾会追加一个空格。可以使用 .trim() 将其移除
    return sh (
        script: getShEchoResultCmd,
        returnStdout: true
    ).trim()
}

//加载项目里面的配置文件
def loadProjectConfig(){
    def jenkinsConfigFile="./jenkins.groovy"
    if (fileExists("${jenkinsConfigFile}")) {
        load "${jenkinsConfigFile}"
        echo "找到打包参数文件${jenkinsConfigFile}，加载成功"
    } else {
        echo "${jenkinsConfigFile}不存在,请在项目${jenkinsConfigFile}里面配置打包参数"
        sh "exit 1"
    }
}
