# eiblog

在 [eiblog1.3.0](https://github.com/eiblog/eiblog) 基础上做一些修改：
 
 - 可以自定义配置mongodb和elasticsearch地址。
 - 删除界面上的评论和一些入口图标，使得界面更简洁。
 
 <br>
 
 ## 在docker部署eiblog服务
 
 ### (1) 编译
 
编译为linux二进制文件
 
 ```bash
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build 
```
 
 <br>
 
 ### (2) 构建eiblog镜像

点击查看[Dockerfile](./Dockerfile)文件，构建镜像：

 > docker build -t zhuyasen/eiblog:1.0 .
 
 <br>
 
### (3) 部署

部署的目录结构如下：
 
 ```
.
├── docker-compose.yml
├── eiblog
├── elasticsearch
└── mongodb
```

注：eiblog服务默认没有conf和static两个目录映射，如果需要自定义配置启动，则在docker-compose.yml中的eiblog下添加conf和static两个目录映射，并且从项目中复制conf和static两个目录过来，例如配置mongodb和elasticsearch地址(conf/app.yml末尾)，例如在容器中使用nginx反向代理，也需要把conf和static两个目录映射给nginx使用(tls证书、静态文件)。
 
点击查看[docker-compose.yml](./docker-compose.yml)文件，构建镜像：
 
> docker-compose up -d

<br>

后台管理 http://localhost:9000/admin/login 默认账号和密码都是admin，登陆后可以修改密码。

首页 http://localhost:9000 ，界面如下：

![home](home.png)

 