# NoviceTask - Script

## API

请求失败的统一返回格式为

```
{
	"success": false,
	"msg": "xxxxx",
	"data": {}
}
```

#### 无需权限的api

- `POST /user`

  注册用户

  ```
  {
  	"email": "xxxx",
  	"username": "xxxx",
  	"pwd": "xxx"
  }
  ```

  成功：

  ```
  {
  	"success": true,
  	"msg": "",
  	"data": {
  		"token":"xxxxxx"
  	}
  }
  ```

- `GET /token?email=xxxxx&pwd=xxxxx`

  成功：

  ```
  {
  	"success": true,
  	"msg": "",
  	"data": {
  		"token":"xxxxxx"
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

  成功：

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


- `POST /user/id`

  修改用户的信息

  ```
  {
  	"email": "xxxx",
  	"username": "xxxx",
  	"pwd": "xxxxx",
  }
  ```
  
  成功：

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

token无误 且 role为admin. token异常返回401，role不为admin返回403

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
              "email": "xxxxxx",
              "username": "xxxxxx",
              "pwd": "xxxxxxx",
              "createTime": "xxxxxx",
              "role": false
          },
  		......
          {
              "id": 11,
              "email": "xxxxxx",
              "username": "xxxxx",
              "pwd": "xxxxxxxxx",
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

  Response:

  ```
  {
  	"success": true,
  	"msg": "",
  	"data": null
  }
  ```

> ### 如果有邮箱验证的话：
>
> - `POST /user`
>
>   注册用户
>
>   redirect to `/user/verify?email=`
>
> - `GET /user/verify?email=`
>
>   Server发送验证邮件
>
> - `POST /user/verify?email=`
>
>   ```
>  {
>   	"token": ""
>   }
>   ```
> 
>   Client发送token，Server确认无误则重定向并注册成功 
>
>   redis暂存 时效与验证码相同

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

## JWT

快了快了
