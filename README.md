# NoviceTask - Script

## API

- `POST /signup`

  注册用户

  ```
  {
  	"email": "xxxx",
  	"username": "xxxx",
  	"pwd": ""
  }
  ```

  成功则返回：

  Status Code `200`

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

  Status Code `401`

  ```
  {
  	"success": false,
  	"msg": "xxxxx",
  	"data": {}
  }
  ```

- `POST /login`

  登录

  ```
  {
  	"email":"xxxx",
  	"pwd":"xxxx"
  }
  ```

  - 成功则返回：

    Status Code `200`

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

    Status Code `401`

    ```
    {
    	"success": false,
    	"msg": "xxxxx",
    	"data": {}
    }
    ```

- `GET /user/id`

  获取用户信息（需要jwt验证）

  http header：

  ```
  Authorization: Bearer xxxxxxxxxx
  ```

  - Authorization无误则返回：

    Status Code `200`

    ```
    {
    	"success": true,
    	"msg": "",
    	"data":{
    		"id":xx
    		"email":"xxxx",
    		"username":"xxxx",
    		"createTime": "",
    		"role": false
    	}
    }
    ```

  - 若Authorization异常：

    Status Code `401 `

    ```
    {
    	"success": false,
    	"msg": "",
    	"data": {}
    }
    ```

- `POST /user/id`

  修改用户的信息（需要jwt验证）

  http header：

  ```
  Authorization: Bearer xxxxxxxxxx
  ```

  Request:

  ```
  {
  	"email": "xxxx",
  	"username": "xxxx",
  	"pwd": "xxxxx",
  }
  ```

  - 成功则返回：

    Status Code `200`

    ```
    {
    	"success": true,
    	"msg": "",
    	"data": {
    		"id":xx
    		"email":"xxxx",
    		"username":"xxxx",
    		"createTime": "",
    		"role": false
    	}
    }
    ```

    失败则返回：

    Status Code `401`

    ```
    {
    	"success": false,
    	"msg": "xxxxx",
    	"data": {}
    }
    ```

需要admin权限的操作：

如果没有权限则返回403

- `GET /user/all`

  查看所有用户信息

  Response:

  ```
  {
      "success": true,
      "msg": "data query succeeded",
      "data": [
          {
              "id": 1,
              "email": "366333@test",
              "username": "12333",
              "pwd": "xxxxxxx",
              "createTime": "xxxxxx",
              "role": false
          },
  		......
          {
              "id": 11,
              "email": "666666@test.com",
              "username": "3333",
              "pwd": "xxxxxxxxx",
              "createTime": "2022-04-02T23:19:05.739542+08:00",
              "role": false
          }
      ]
  }
  ```

- `DELETE /user/:id`

  删除用户

  Response:

  状态码 `204`

  ```
  {
  	"success": true,
  	"msg": "deleted successfully",
  	"data": null
  }
  ```

> ### 如果有邮箱验证的话：
>
> - `POST /signup`
>
>   注册用户
>
>   通过验证器后
>
>   redirect to `/user/verify?email=`
>
> - `GET /signup/verify?email=`
>
>   Server发送验证邮件
>
> - `POST /signup/verify?email=`
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
>   要去问一下）

## 数据库

`users`

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



## 密码加密

md5

还没搞

## JWT

还没搞
