# Etcd-Studio

#### 介绍
Webui for Etcd(v3)

#### 软件架构
软件架构说明
cd webui
npm run build
cd ..

./go-bindata.exe -o src/bindata.go -fs -prefix "webui/dist" webui/dist/...

cd src
go build -o ../EtcdStudio.exe . 
cd ..

./upx.exe EtcdStudio.exe


#### 安装教程

1. 安装nodejs
sudo rm -fv /etc/yum.repos.d/nodesource*
curl -sL https://rpm.nodesource.com/setup_12.x | bash -
sudo yum clean all
yum install -y nodejs

npm install -g cnpm --registry=https://registry.npm.taobao.org
cnpm install vue -g
cnpm install vue-cli -g
cnpm install vue-cli-service -g

#### 使用说明

