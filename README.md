# simple-vote

## 安装

```
$ go get github.com/afsxt/simple-vote
```

## 如何运行

### 必须

- Mysql
- Redis

### 准备

创建一个 `vote` 数据库，并且导入建表的 [SQL](https://github.com/afsxt/simple-vote/blob/master/docs/sql/vote.sql)

### 配置

你应该修改 `conf/app.ini` 配置文件

```
[database]
Type = mysql
User = root
Password = root
Host = 127.0.0.1:3306
Name = vote
TablePrefix = vote_

[redis]
Host = 127.0.0.1:6379
Password =
MaxIdle = 30
MaxActive = 30
IdleTimeout = 200
...
```


### 运行

```
$ cd $GOPATH/src/simple-vote

$ go run main.go 
```

项目的运行信息和已存在的 API's

```
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /swagger/*any             --> github.com/swaggo/gin-swagger.CustomWrapHandler.func1 (3 handlers)
[GIN-debug] GET    /ping                     --> github.com/afsxt/simple-vote/routers.InitRouter.func1 (3 handlers)
[GIN-debug] POST   /admin/v1/vote/theme      --> github.com/afsxt/simple-vote/routers/admin/v1.AddTheme (3 handlers)
[GIN-debug] POST   /admin/v1/vote/candidates --> github.com/afsxt/simple-vote/routers/admin/v1.AddCandidates (3 handlers)
[GIN-debug] POST   /admin/v1/vote/theme/:id/state --> github.com/afsxt/simple-vote/routers/admin/v1.ChangeThemeState (3 handlers)
[GIN-debug] POST   /api/v1/vote/verify       --> github.com/afsxt/simple-vote/routers/api/v1.VerifyUser (3 handlers)
[GIN-debug] POST   /api/v1/vote              --> github.com/afsxt/simple-vote/routers/api/v1.Vote (3 handlers)
[GIN-debug] GET    /api/v1/vote/:themeID     --> github.com/afsxt/simple-vote/routers/api/v1.GetVoteDetails (3 handlers)
2023/02/18 10:33:15 [info] start http server listening :8000
```

Swagger Api文档
可以访问: http://localhost:8000/swagger/index.html

## 特性

- RESTful API
- Gorm
- Swagger
- logging
- Jwt-go
- Gin
- App configurable
- Cron
- Redis