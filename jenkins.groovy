// 运行模块
env.MODULE="tudou-one-api-go"
// 端口必须改
env.ORIGINAL_PORT=3000

// 基础依赖模块，多个逗号隔开
env.MODULE_BASE=""
env.IS_MULT_MODULE=false
env.JAR_NAME="app.jar"
// 是否启用云原生部署架构，false使用“--network=host”模式，暴漏所有端口.只有在docker部署时有用
env.META_CLOUD=false

env.HOST_BASEDIR="/opt/compose/conf/app"

// Docker 参数
env.HARBOR_ADDR="harbor.e-tudou.com"
env.HARBOR_REPO="bpc-app-python"
env.HARBOR_VERSION="latest"
