# 后端架构

## Model

### userModel 字段

```go
type User struct {
    gorm.Model//gorm框架的统一model字段，自动添加ID和created等
    Name     string `gorm:"varchar(20);not null"`
    Email    string `gorm:"varchar(50);not null;unique"`
    Password string `gorm:"size:255;not null"`
}
```

### result 字段

```go
type Result struct {
    Code int         `json:"code"`
    Data interface{} `json:"data"`
    Msg  string      `json:"msg"`
}
```

## UserController

控制层操控 service 访问数据库，需要三个函数：登录，注册，登出

一个 login API，该函数接收一个账号，一个密码

## UserService

login 函数，调用数据库

## route

路由负责接受前端的交互请求

```
前端 Axios POST /user/login
        ↓
Gin 路由匹配到 LoginHandler
        ↓
ShouldBindJSON 绑定用户名密码到 UserCredentials 结构体
        ↓
调用 user_service.Login() 验证用户信息
        ↓
访问数据库返回结果
        ↓
返回 JSON 结果给前端
```
