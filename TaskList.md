## level 0

> - [x] rest api:get /user
> - [x] respond改掉
> - [x] 每个接口写一个结构体
> - [x] gitignore
> - [x] md5
> - [ ] validator
> - [ ] jwt
> - [ ] 加group
> - [x] update改掉

- [x] [Go语法规范](http://www.gonglin91.com/2018/03/30/go-code-review-comments/)

- [x] 配置使用[go mod](https://www.jianshu.com/p/760c97ff644c)进行包管理，设置[GOPROXY](https://blog.csdn.net/sinat_34241861/article/details/110232463)

- [x] Echo

  [Golang echo 快速入门](https://www.tizi365.com/archives/28.html)

  [Echo中文文档](http://echo.topgoer.com/)

- [x] MVC

  [MVC设计模式浅析](https://www.jianshu.com/p/254e2478a4ca)

- [x] REST API

  [RESTful API 最佳实践 - 阮一峰](http://www.ruanyifeng.com/blog/2018/10/restful-api-best-practices.html)

- [ ] [JWT](token认证.md) (以及echo中间件)

  [基于JWT标准的用户认证接口实现](https://www.cnblogs.com/xiaohuochai/p/8440335.html)

  [JWT安全验证理解与实践(Golang)](https://blog.wangjunfeng.com/post/golang-jwt/#3-%E7%AD%BE%E5%90%8D-signature)

  https://www.cnblogs.com/jianga/p/12487267.html

  https://www.cnblogs.com/tomtellyou/p/12895437.html

- [x] 《图解HTTP》    `看完一半了！好耶`

- [x] PostgreSQL

  [PostgreSQL - 菜鸟教程](https://www.runoob.com/postgresql/postgresql-tutorial.html)

- [x] Go-pg

  [PostgreSQL client and ORM for Go](https://pg.uptrace.dev/)

- [ ] [validator](https://blog.csdn.net/guyan0319/article/details/105918559)

- [ ] 把代理重新搞一下

- [ ] 搭环境

  - [ ] nginx
  - [ ] docker
  - [ ] ssh




>
>#### 实现一个用户管理系统
>- 只需要支持：注册、登录、验证身份、管理用户（CURD）
>
>- Request使用 GET-URL 参数、其他-JSON 的数据交互格式，Reponse使用统一的 JSON 格式
>
>-  密码不应该使用明文，你需要找一些方法加密
>
>
>- 使用JWT鉴权
>
>- 用户具有 default 和 admin 权限，admin 权限可以管理用户信息
>
>- 数据库表结构自己设计，要求使用Postgresql，使用[go-pg](https://github.com/go-pg/pg)（想要了解orm的使用可以去看看go的**xorm**和**gorm**，不做任何要求）
>
>- 学会写API文档（描述你的接口期望接收什么样的请求，期望回复什么样的响应，以及必要的错误状态码说明）
>
>- [参考](https://github.com/KSkun/Simple-Go-User-System)
>
>---
>
>- [x] Echo hello-world
>- [x] 架构设计
>- [x] PostgreSQL & Go-pg
>- [x] models接口
>- [x] controllers接口
>- [ ] md5加密
>- [ ] validator
>- [ ] JWT
>
>Functions:
>
>- [ ] 注册功能
>- [ ] 登录功能
>- [ ] 验证身份功能
>- [ ] 管理用户功能
>- [ ] Docs

## Further

- [ ] TCP/IP

- [ ] SQL反模式

- [ ] Go Roadmap
