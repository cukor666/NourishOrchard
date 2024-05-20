# NourishOrchard

## 关于

开发者：CukorZhong.

项目中文名：滋养果园。
这个是一个基于Gin+vue3的水果售卖后台管理系统，主要是偏向后台管理系统。

项目站点：

+ 用户端：[http://cukor.cn](http://cukor.cn)
+ 管理员端：[http://cukor.cn:5173](http://cukor.cn:5173)

使用的技术栈如下：

+ 数据库：
  + 本地：MySQL8+Redis5
  + 线上：MySQL8+Redis7
  + 软件：Navicat

+ 前端：
  + 主体框架：Vue3+Js
  + 衍生框架：
    + UI框架：ElementPlus
    + 前后端通信：Axois
    + 状态管理：Pinia
    + CSS预处理器：Sass
  + 构建工具：vite+yarn
  + 开发工具：WebStorm+VSCode
  
+ 后端：
  + 主体框架：Gin+Golang
  + 衍生框架：
    + ORM框架：GORM
    + 缓冲数据库：go-redis
    + 鉴权与加密：JWT+bcrypt（带盐的MD5）
    + Excel文件框架：Excelize
    + 配置文件库：yaml
  + 开发工具：Goland+VSCode
  
+ 版本控制：
  + Git+Github Desktop

+ 服务器：
  + 商家：腾讯云
  + 类型：Ubuntu Linux
  + 项目部署：Docker+Nginx
  
+ 其他工具：
  + 接口测试：ApiFox
  + 远程连接服务：XShell+XFtp


## 架构

后端1.0：Gin目录，已弃用，使用server目录作为后端服务器

后端2.0：server，主要管理后端业务

前端1.0：vue-project目录，已弃用，并将原来的代码放到vue-project.zip中，使用vue目录作为前端部分。

前端2.0：vue，主要管理前端界面展示部分。使用yarn作为包管理器

内部项目：internal，主要用于管理员和员工信息的数据库注入，和一些系统的配置项。

------

项目最新架构：MVC+DDD

项目的最外层是传统MVC架构，MVC各层内部采用了DDD的架构思想，DDD是“领域驱动设计”的英文简称，其核心思想是将同一类业务代码放到同一个文件夹（包、模块）下进行管理。以便产品与开发之间的相互协作。

## 构建与部署

### 前端

使用vite打包工具，包管理工具使用yarn.

1. 修改在axios文件夹下的request.js文件

```js
const local = "..."		// 本地ip+port
const server = "..."	// 线上ip+port

const request = axios.create({
  baseURL: server,	// 选择合适的ip
  timeout: 5000,
});
```

2. 执行打包命令

```shell
yarn build
```

3. 生成dist文件，后续服务器部署需要将其放到服务器的指定目录下。结合nginx完成项目前端部署。

### 数据库

本地使用的MySQL8版本，线上版本使用Docker容器化的MySQL8版本。Redis在本地使用Redis5版本，线上使用Docker容器化部署的Redis7版本。

### 后端

1. 修改config文件夹下的config.yaml

```yaml
systemConfig:
  host: localhost
  port: 
  secret: 
  issuer: 

mysql:
  host: localhost
  port: 
  user: 
  password: 
  dbname: nourish_orchard2
  param: charset=utf8mb4&parseTime=True&loc=Local	# mysql8

redis:
  host: localhost
  port: 
  password: 
  db: 0

jwt:
  secretKey: 
  issuer: 
```

2. 解决跨域问题

修改router文件夹下的variable.go文件

```go
package router

const (
	ClientDomain       = "..." // 本地前端的ip端口
	ClientOnlineDomain = "..."
)

var domains = [...]string{ClientDomain, ClientOnlineDomain}

```

目前为了方便项目的线上演示，也已经采用允许所有主机能直接访问都后端服务器的方式进行跨域问题的解决。后续考虑采用Nginx的反向代理方式解决跨域请求问题。

3. 执行构建命令

```shell
set GOARCH=amd64
go env -w GOARCH=amd64
set GOOS=linux
go env -w GOOS=linux
go build
go env -w GOARCH=amd64
go env -w GOOS=windows
```

4. 生成linux下可执行文件，后续发送到服务器上直接运行，完成项目的后端部署。

## 功能

本部分仅描述各个功能的开发需求，详细进度情况apifox文档。前后端交互使用的数据以JSON格式为准。大体内容如下：

```json
{
    "code": 200,
    "msg": "响应成功",
    "data": 0
}
```

+ code：状态码
+ msg：提示消息
+ data：具体数据

### 注册功能

前端发送对应的信息表单转换为JSON格式发送给后端，后端接收后进行参数校验，然后存入数据库中，在存入数据库时，一共需要操作两张表，分别是`account`和`user`表（因为管理员和员工的信息都是通过内部进行数据库注入，所以只有用户需要完成注册功能）。在操作这两张表时使用了事务管理，保证数据的一致性和原子性。

### 登录功能

前端发送对应的信息表单转换为JSON格式发送给后端，后端接收到数据后进行参数校验，然后比对账号、密码以及权限三者是否完全一致，如果不一致返回给前端登录失败提示消息，前端做出对应的友好提示的界面给用户展示；如果一致返回给前端登录成功的提示信息，并将后端生成的token数据放到`data`中返回给前端，然后前端接收到token之后存放到`localstorage`中。

### 内部信息注入（内部功能）

主要是管理员和员工的信息注入到数据库中。

管理员和员工信息都应该以excel表的形式存储于特定的位置，然后通过excelize库进行到读取然后通过GORM导入到数据库中。

导包：

```shell
go get -u github.com/xuri/excelize/v2
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

### 登录用户信息的个人管理

登录的人包括：

+ 用户
+ 管理员
+ 员工

该功能需要做权限验证，前端应该把登录的token传给后端，后端接收到token后，对token进行解析，然后查询数据库中对应的账户，将结果以JSON格式返回给前端，前端以美观、友好的界面展示给用户。

查看个人信息和更新个人信息应该放在同一个前端页面，所以只需要一个PUT请求。先将数据查询处理展示，前端给用户提供一个“修改信息”的按钮，当点击了这个按钮之后就可以对“用户”的个人信息进行更新，然后再点击提交按钮完成更新操作，为了防止用户误操作，所以再更新之前应该要弹出一个对话框，询问用户是否要更新自己的个人信息。如果用户确定要更新个人信息，则后端做出对应的修改操作。后端修改完毕之后把修改之后的数据发送给前端，前端以美观、友好的界面展示给用户。

### 其他

其他请看apifox具体文档。

