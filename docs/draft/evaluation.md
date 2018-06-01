# 分类接口

* 获取评论列表
* 获取评论详情
* 删除评论（需管理员权限）

## 获取评论列表

路径：`/api/evaluations/:id`

方法: `GET`

路径参数：

* id：文章 ID

其他：

* 支持分页

请求示例：

```shell
curl -X GET $API_BASE/api/evaluations/1
```

返回值：

```json
{
  "success": true,
  "code": 20000,
  "message": "get evaluations successfully",
  "payload": {
    "page_no": -1,
    "page_size": -1,
    "current_size": 1,
    "total_size": 1,
    "data": []
  }
}
```

## 获取评论详情

路径：`/api/evaluation/:id`

方法: `GET`

路径参数:

* id：分类 ID

请求示例：

```shell
curl -X GET $API_BASE/api/evaluation/1
```

返回值：

```json
{
  "success": true,
  "code": 20000,
  "message": "get evaluation successfully",
  "payload": {
    "id": 1,
    "created_at": "2018-05-24T22:57:35+08:00",
    "updated_at": "2018-05-24T22:57:35+08:00",
    "deleted_at": null,
    "ip": "::1",
    "score": 5,
    "content": "真棒~",
    "article": {
      "id": 0,
      "created_at": "0001-01-01T00:00:00Z",
      "updated_at": "0001-01-01T00:00:00Z",
      "deleted_at": null,
      "title": "",
      "content": "",
      "banner": "",
      "category": {
        "id": 0,
        "created_at": "0001-01-01T00:00:00Z",
        "updated_at": "0001-01-01T00:00:00Z",
        "deleted_at": null,
        "name": "",
        "articles": null
      },
      "category_id": 0,
      "tags": null,
      "evaluations": null
    },
    "article_id": 1
  }
}
```

## 删除评论（需管理员权限）

路径：`/api/evaluation/:id`

方法: `DELETE`

路径参数:

* id：分类 ID

请求示例：

```shell
curl -X DELETE $API_BASE/api/evaluation/1
```

返回值：

```json
{
  "success": true,
  "code": 20000,
  "message": "delete evaluation successfully",
  "payload": null
}
```
