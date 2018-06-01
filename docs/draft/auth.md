# 管理接口

* 管理员登陆

## 管理员登陆

路径：`/api/auth/login`

方法: `POST`

请求参数:

* username：用户名
* password：密码

请求示例：

```shell
curl -X POST \
  $API_BASE/api/auth/login \
  -H 'content-type: application/json' \
  -d '{
    "username": "xxx",
    "password": "xxx"
  }'
```

返回值：

```json
{
  "success": true,
  "code": 20000,
  "message": "login successfully",
  "payload": {
    "token": "xxx"
  }
}
```
