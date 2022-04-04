## API

#### 无需权限的api

- `POST /user`

  注册用户

  ```
  {
  	"email": "xxxx",
  	"username": "xxxx",
  	"pwd": "xxxxxxxx"
  }
  ```

  （字段均不能为空，且邮箱须有效格式）

  返回：

  ```
  {
  	"success": true,
  	"msg": "",
  	"data": {
  		"token":"xxxxxxxxxxxx"
  	}
  }
  ```

- `GET /user/token?email=xxxxx&pwd=xxxxx`

  用户登录

  返回：

  ```
  {
  	"success": true,
  	"msg": "",
  	"data": {
  		"token":"xxxxxxxxxxx"
  	}
  }
  ```

#### 需登录权限的api

http header：

```
Authorization: Bearer xxxxxxxxxx
```

token异常返回401.

- `GET /user/id`

  获取用户信息

  返回：

  ```
  {
  	"success": true,
  	"msg": "",
  	"data":{
  		"id":xx
  		"email":"xxxx",
  		"username":"xxxx",
  	}
  }
  ```


- `PUT /user/id`

  修改用户信息

  ```
  {
  	"email": "xxxx",
  	"username": "xxxx",
  	"pwd": "xxxxx",
  }
  ```

  （字段均不能为空，且邮箱须有效格式）

  返回：

  ```
  {
  	"success": true,
  	"msg": "",
  	"data": {}
  }
  ```

#### 需要admin权限的操作：

http header：

```
Authorization: Bearer xxxxxxxxxx
```

token异常返回401，role不为admin返回403

- `GET /user/all`

  查看所有用户信息

  返回:

  ```
  {
      "success": true,
      "msg": "data query succeeded",
      "data": [
          {
              "id": 1,
              "email": "xxxxxx",
              "username": "xxxxxx",
              "createTime": "xxxxxx",
              "role": false
          },
  		......
          {
              "id": 11,
              "email": "xxxxxx",
              "username": "xxxxx",
              "createTime": "xxxxxx",
              "role": false
          }
      ]
  }
  ```

- `DELETE /user/:id`

  删除用户

  http header：

  ```
  Authorization: Bearer xxxxxxxxxx
  ```

  返回:

  ```
  {
  	"success": true,
  	"msg": "",
  	"data": null
  }
  ```

请求失败的统一返回格式为

```
{
	"success": false,
	"msg": "xxxxx",
	"data": null
}
```

## 