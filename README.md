# simple-vote 投票选举系统API

## 完成情况说明
数据字典只保存了简单的极少字段，为区分候选人及投票活动，默认名称具有唯一性
主要的功能：如投票，创建候选人及创建投票活动，还有管理员控制的相关功能都有实现

## 待优化点
- 发送邮件: 目前发送邮件api有封装完成，但没实现具体的邮件发送逻辑，可用一个异步任务或消息队列实现具体发送流程
- 完善测试用例及代码覆盖率: 目前只实现了17个测试用例
- 优化及扩展设计: 缓存优化，用户及邮件部分单独实现公共服务等
- 代码及错误处理更加完善

## 目录结构
```
simple-vote
└── conf
    └── app.ini (配置文件)
└── docs
    └── sql
        └── vote.sql (sql文件)
    └── docs.go (swagger api文档相关的文件)
    └── swagger.json
    └── swagger.yaml
└── middleware (中间件)
    └── jwt
└── models (数据库模型)
└── pkg
    └── app
    └── constants
    └── e (错误代码及描述信息)
    └── file
    └── gredis
    └── logging
    └── setting
    └── util
└── routers
└── service
└── testing（测试用例）
    └── api_test.go
    └── utils.go
└── Dockerfile
└── go.mod
└── main.go
└── README.md
```

## Global 响应示例
```
{
  code: 0,
  message: 'success',
  data:{
    ...
  } 
}

```

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
[GIN-debug] PUT    /admin/v1/vote/theme/:id  --> github.com/afsxt/simple-vote/routers/admin/v1.UpdateTheme (3 handlers)
[GIN-debug] POST   /admin/v1/vote/candidates --> github.com/afsxt/simple-vote/routers/admin/v1.AddCandidates (3 handlers)
[GIN-debug] POST   /admin/v1/vote/theme/:id/state --> github.com/afsxt/simple-vote/routers/admin/v1.ChangeThemeState (3 handlers)
[GIN-debug] GET    /admin/v1/vote/theme/:themeID --> github.com/afsxt/simple-vote/routers/admin/v1.GetThemeResult (3 handlers)
[GIN-debug] GET    /admin/v1/vote/theme/:themeID/candidate/:candidateID/users --> github.com/afsxt/simple-vote/routers/admin/v1.GetCandidateUsers (3 handlers)
[GIN-debug] POST   /api/v1/vote/verify       --> github.com/afsxt/simple-vote/routers/api/v1.VerifyUser (3 handlers)
[GIN-debug] POST   /api/v1/vote              --> github.com/afsxt/simple-vote/routers/api/v1.Vote (3 handlers)
[GIN-debug] GET    /api/v1/vote/theme/:themeID --> github.com/afsxt/simple-vote/routers/api/v1.GetVoteDetails (3 handlers)
2023/02/19 10:28:30 [info] start http server listening :8000
```

Swagger Api文档

本地服务起来后可以访问如: http://localhost:8000/swagger/index.html

![image](https://github.com/afsxt/simple-vote/blob/main/docs/api.png)


错误code及描述信息
```
SUCCESS:        "ok",
ERROR:          "fail",
INVALID_PARAMS: "请求参数错误",

ERROR_CANDIDATE_EXIST:      "该候选人已经存在",
ERROR_CANDIDATE_EXIST_FAIL: "检查候选人是否存在失败",
ERROR_CANDIDATE_ADD_FAIL:   "新增候选人失败",

ERROR_THEME_EXIST:                    "该主题已经存在",
ERROR_THEME_EXIST_FAIL:               "检查主题是否存在失败",
ERROR_THEME_NOT_EXIST:                "该主题不存在",
ERROR_THEME_ADD_FAIL:                 "新增主题失败",
ERROR_THEME_GET_CANDIDATE_COUNT_FAIL: "获取选举主题候选人总数失败",
ERROR_THEME_COUNT_NOT_ENOUGH:         "该主题候选人不够",

ERROR_USER_ADD_FAIL:         "新增用户失败",
ERROR_USER_CHECK_VALID_FAIL: "较验用户是否合法失败",
ERROR_USER_INVALID:          "非法用户",
ERROR_USER_GET_VOTE_FAIL:    "获取候选人支持用户失败",

ERROR_VOTE_ADD_FAIL:        "新增投票失败",
ERROR_VOTE_AGAIN_FAILE:     "该用户已经对该主题投过票",
ERROR_VOTE_GET_DETAIL_FAIL: "用户获取投票详情失败",
```

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