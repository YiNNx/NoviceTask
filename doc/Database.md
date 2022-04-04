## Database

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