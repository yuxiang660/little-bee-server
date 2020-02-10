<h1 align="center">
    <img alt="LittleBee" title="Lumen" src="https://github.com/yuxiang660/little-bee-server/blob/master/.github/logo.jpg" width="140"> </br>
</h1>

<h4 align="center">
  “蜂工厂”
</h4>
<h4 align="center">
  作者 -- “搬砖蜜蜂”
</h4>

# Demo
![Demo](https://github.com/yuxiang660/little-bee-server/blob/master/.github/demo.gif)

# Introduction

- A clean architecture `Go` HTTP server with the help of:
  - [`Gin`](https://gin-gonic.com/)
  - [`GORM`](https://gorm.io/)
  - [`JWT-GO`](https://github.com/dgrijalva/jwt-go)
  - [`GO-Redis`](https://github.com/go-redis/redis)
  - [`Go-Dig`](https://github.com/uber-go/dig)

- Refer to the [blog](https://yuxiang660.github.io/little-bee-client/posts/4/2020-02-10---Little-Bee-Server-Intro/) for details.

# Project Structure

```go
.
├── Makefile               // 用Makefile管理项目的编译
├── README.md              // 解释文档
├── cmd                    // 程序入口文件夹
|   └── server             // 主程序文件夹
|       └── main.go        // 主程序，每一个Go项目有且只有一个main入口
├── configs                // 用户配置文件夹
|   └── config.toml        // 存储对服务器的所有配置
├── docs                   // 存储服务器的Swagger API文档
├── export                 // 存储服务器输出文件，包括数据信息、log信息等
├── go.mod                 // Go项目包管理文件
├── go.sum                 // Go项目包管理文件
|── internal               // 项目内部源文件(其他项目无法直接调用)
|   └── app                // 内部源文件主目录
|       ├── app.go         // 内部源文件初始化入口
|       ├── auther         // 身份认证模块
|       ├── auther.go      // 身份认证模块初始化入口
|       ├── config         // 用户配置文件解析模块
|       ├── controller     // 控制器模块，用于处理业务逻辑
|       ├── controller.go  // 控制器模块初始化入口
|       ├── errors         // 服务器错误模块，定义错误信息
|       ├── ginhelper      // gin框架utilities
|       ├── logger         // 日志模块
|       ├── logger.go      // 日志模块初始化入口
|       ├── model          // Model模块，提供接口给控制器存储数据
|       ├── model.go       // Model模块初始化入口
|       ├── routers        // 路由模块
|       |   ├── api        // REST API
|       |   ├── middleware // 中间件
|       |   ├── routers.go // 初始化REST API和中间件
|       |   └── swagger.go // API 文档入口
|       ├── routers.go     // 路由初始化入口，并开启HTTP服务
|       ├── store          // 数据库模块，提供接口给Model模块和数据库打交道
|       └── store.go       // 数据模块初始化入口
```
