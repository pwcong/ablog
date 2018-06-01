# 分类接口

* 获取分类列表
* 获取分类信息
* 删除分类（需管理员权限）
* 修改分类（需管理员权限）
* 新增分类（需管理员权限）

## 获取分类列表

路径：`/api/categories`

方法: `GET`

其他：

* 支持分页

请求示例：

```shell
curl -X GET $API_BASE/api/categories
```

返回值：

```json
{
  "success": true,
  "code": 20000,
  "message": "get categories successfully",
  "payload": {
    "page_no": -1,
    "page_size": -1,
    "current_size": 0,
    "total_size": 0,
    "data": []
  }
}
```

## 获取分类信息

路径：`/api/category/:id`

方法: `GET`

路径参数:

* id：分类 ID

请求示例：

```shell
curl -X GET $API_BASE/api/category/1
```

返回值：

```json
{
  "success": true,
  "code": 20000,
  "message": "get category successfully",
  "payload": {
    "id": 1,
    "created_at": "2018-05-24T22:24:18+08:00",
    "updated_at": "2018-05-24T22:24:18+08:00",
    "deleted_at": null,
    "name": "Golang",
    "articles": null
  }
}
```

## 删除分类（需管理员权限）

路径：`/api/category/:id`

方法: `DELETE`

路径参数:

* id：分类 ID

请求示例：

```shell
curl -X DELETE $API_BASE/api/category/1
```

返回值：

```json
{
  "success": true,
  "code": 20000,
  "message": "delete category successfully",
  "payload": null
}
```

## 修改分类（需管理员权限）

路径：`/api/category/:id`

方法: `POST`

路径参数：

* id：分类 ID

请求参数:

* name：分类名称

请求示例：

```shell
curl -X POST \
  $API_BASE/api/category/1 \
  -H 'content-type: application/json' \
  -d '{
    "name": "ABlog"
  }'
```

返回值：

```json
{
  "success": true,
  "code": 20000,
  "message": "update category successfully",
  "payload": {
    "id": 1,
    "created_at": "2018-05-24T22:24:18+08:00",
    "updated_at": "2018-05-24T22:31:21.3545376+08:00",
    "deleted_at": null,
    "name": "ABlog",
    "articles": null
  }
}
```

## 新增分类（需管理员权限）

路径：`/api/category`

方法: `POST`

请求参数:

* name：分类名称

请求示例：

```shell
curl -X POST \
  $API_BASE/api/category \
  -H 'content-type: application/json' \
  -d '{
    "name": "Golang"
  }'
```

返回值：

```json
{
  "success": true,
  "code": 20000,
  "message": "add category successfully",
  "payload": {
    "id": 1,
    "created_at": "2018-05-24T22:24:18+08:00",
    "updated_at": "2018-05-24T22:24:18+08:00",
    "deleted_at": null,
    "name": "Golang",
    "articles": null
  }
}
```
