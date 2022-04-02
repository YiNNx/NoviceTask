# API

- POST /signup

  注册用户

  通过验证器存入数据库就欧克！

  ```
  {
  	"email": "xxxx"
  	"username": "xxxx"
  	"pwh": ""
  	"nickname": ""
  	"intro": ""
  }
  ```

  成功则返回：

  状态码 `201`

  ```
  {
  	"success": true,
  	"msg": "",
  	"data": {
  		"token":"xxxxxx"
  	}
  }
  ```

  失败则返回：

  状态码 `401`

  ```
  {
  	"success": false,
  	"msg": "xxxxx",
  	"data": {}
  }
  ```

- POST /login

  登录

  ```
  {
  	"email":"xxxx"
  	"pwdhash":"xxxx"
  }
  ```

  - 成功则返回：

    状态码 `201`

    ```
    {
    	"success": true,
    	"msg": "",
    	"data": {
    		"token":"xxxxxx"
    	}
    }
    ```

    失败则返回：

    状态码 `401`

    ```
    {
    	"success": false,
    	"msg": "xxxxx",
    	"data": {}
    }
    ```

- GET /user/id

  获取用户信息（需要jwt验证）

  http header：

  ```
  Authorization: Bearer xxxxxxxxxx
  ```

  - 成功则返回：

    ```
    {
    	"success": true,
    	"msg": "",
    	"data":{
    		"email":"xxxx"
    		"username":"xxxx"
    		"nickname": ""
    		"intro": ""
    	}
    }
    ```

  - 若用户未登录/token错误：

    状态码 `401 `

    ```
    {
    	"success": false,
    	"msg": "",
    	"data": {}
    }
    ```

- POST /user/id

  修改用户的信息（需要jwt验证）

  http header：

  ```
  Authorization: Bearer xxxxxxxxxx
  ```

  Request:

  ```
  {
  	"email": "xxxx"
  	"username": "xxxx"
  	"pwdhash": "xxxxx"
  	"nickname": ""
  	"intro": ""
  }
  ```
  
  - 成功则返回：
  
    状态码 `200`
  
    ```
    {
    	"success": true,
    	"msg": "",
    	"data": {
    		"email":"xxxx"
    		"username":"xxxx"
    		"nickname": ""
    		"intro": ""
    	}
    }
    ```
  
    失败则返回：
  
    状态码 `401`
  
    ```
    {
    	"success": false,
    	"msg": "xxxxx",
    	"data": {}
    }
    ```

需要admin权限的操作：

如果没有权限则返回403

- GET /user/all

  查看所有用户信息

  Response:

  ```
  {
  	"success": true,
  	"msg": "",
  	"data": {[
  		{
  			"id": xx
              "email": "xxxx"
              "username": "xxxx"
              "nickname": ""
              "intro": ""
  		},
  		{
  			"id": xx
              "email": "xxxx"
              "username": "xxxx"
              "nickname": ""
              "intro": ""
  		},
  		{
  			"id": xx
              "email": "xxxx"
              "username": "xxxx"
              "nickname": ""
              "intro": ""
  		}
      ], "total": 3
  	}
  }
  ```

- DELETE /user/:id

  删除用户

  Response:

  状态码 `204`

  ```
  {
  	"success": true,
  	"msg": "",
  	"data": {}
  }
  ```

> ### 如果有邮箱验证的话：
>
> - POST /signup
>
>   注册用户
>
>   通过验证器后
>
> - 重定向至/user/verify?email=
>
> - GET /signup/verify?email=
>
>   Server发送验证邮件
>
> - POST /signup/verify?email=
>
>   ```
>   {
>   	"token": ""
>   }
>   ```
>
>   Client发送token，Server确认无误则重定向至/user/verify注册成功 
>
>   （暂时想到的是把数据都存在token里传给server 验证成功就存到数据库
>
>   要去找ryao问一下）

## 状态码

- `200 OK`
- ` StatusCreated = 201` 
- `204 NO CONTENT` 用户删除数据成功。
- `http.StatusTemporaryRedirect = 307` 临时重定向
- `400 Bad Request`：服务器不理解客户端的请求
- `401 Unauthorized` ：表示用户没有权限（令牌、用户名、密码错误）。
- `403 Forbidden`：用户通过了身份验证，但是不具有访问资源所需的权限。
- `404 Not Found`：所请求的资源不存在，或不可用。
- `500 Internal Server Error`：客户端请求有效，服务器处理时发生了意外。

# 数据库

users

```
                                       数据表 "public.users"
    栏位     |           类型           | Collation | Nullable |              Default
-------------+--------------------------+-----------+----------+-------------------------------
 id          | bigint                   |           | not null | nextval('users_id_seq'::regclass)
 email       | text                     |           | not null |
 username    | text                     |           | not null |
 pwd_hash    | text                     |           |          |
 create_time | timestamp with time zone |           |          | now()
 role        | boolean                  |           | not null | false

```



# 密码加密

md5

# JWT



# valivator