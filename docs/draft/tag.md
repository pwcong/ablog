# 标签接口

* 获取标签列表
* 获取标签信息
* 删除标签（需管理员权限）
* 修改标签（需管理员权限）
* 新增标签（需管理员权限）

## 获取标签列表

路径：`/api/tags`

方法: `GET`

其他：

* 支持分页

请求示例：

```shell
curl -X GET $API_BASE/api/tags
```

返回值：

```json
{
  "success": true,
  "code": 20000,
  "message": "get tags successfully",
  "payload": {
    "page_no": -1,
    "page_size": -1,
    "current_size": 0,
    "total_size": 0,
    "data": []
  }
}
```

## 获取标签信息

路径：`/api/tag/:id`

方法: `GET`

路径参数:

* id：标签 ID

请求示例：

```shell
curl -X GET $API_BASE/api/tag/1
```

返回值：

```json
{
  "success": true,
  "code": 20000,
  "message": "get tag successfully",
  "payload": {
    "id": 1,
    "created_at": "2018-05-24T22:34:40+08:00",
    "updated_at": "2018-05-24T22:34:40+08:00",
    "deleted_at": null,
    "name": "Test",
    "articles": null
  }
}
```

## 删除标签（需管理员权限）

路径：`/api/tag/:id`

方法: `DELETE`

路径参数:

* id：标签 ID

请求示例：

```shell
curl -X DELETE $API_BASE/api/tag/1
```

返回值：

```json
{
  "success": true,
  "code": 20000,
  "message": "delete tag successfully",
  "payload": null
}
```

## 修改标签（需管理员权限）

路径：`/api/tag/:id`

方法: `POST`

路径参数：

* id：标签 ID

请求参数:

* name：标签名称

请求示例：

```shell
curl -X POST \
  $API_BASE/api/tag/1 \
  -H 'content-type: application/json' \
  -d '{
    "name": "HAHA"
  }'
```

返回值：

```json
{
  "success": true,
  "code": 20000,
  "message": "update tag successfully",
  "payload": {
    "id": 1,
    "created_at": "2018-05-24T22:34:40+08:00",
    "updated_at": "2018-05-24T22:34:40+08:00",
    "deleted_at": null,
    "name": "HAHA",
    "articles": null
  }
}
```

## 新增标签（需管理员权限）

路径：`/api/tag`

方法: `POST`

请求参数:

* name：标签名称

请求示例：

```shell
curl -X POST \
  $API_BASE/api/tag \
  -H 'content-type: application/json' \
  -d '{
    "name": "Test"
  }'
```

返回值：

```json
{
  "success": true,
  "code": 20000,
  "message": "add tag successfully",
  "payload": {
    "id": 1,
    "created_at": "2018-05-24T22:34:39.5358777+08:00",
    "updated_at": "2018-05-24T22:34:39.5358777+08:00",
    "deleted_at": null,
    "name": "Test",
    "articles": null
  }
}
```
