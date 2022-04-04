# NoviceTask0

- 支持：注册、登录、验证身份、管理用户（CURD）

- Request使用 GET-URL 参数、其他-JSON 的数据交互格式，Reponse使用统一的 JSON 格式

- 密码使用bcrypt加密


- 使用JWT鉴权
- 用户具有 default 和 admin 权限，admin 权限可以管理用户信息
- 数据库使用Postgresql & [go-pg](https://github.com/go-pg/pg)
