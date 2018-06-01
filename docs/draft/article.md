# 文章接口

* 新增文章（需管理员权限）
* 获取文章列表
* 获取文章详情
* 获取文章列表（筛选）
* 评论文章
* 修改文章（需管理员权限）
* 删除文章（需管理员权限）

## 新增文章（需管理员权限）

路径：`/api/article`

方法: `POST`

请求参数:

* title：标题
* content：内容
* banner：横幅 url
* category_id：分类 ID
* tag_ids：标签 ID 数组

请求示例：

```shell
curl -X POST \
  $API_BASE/api/article \
  -H 'content-type: application/json' \
  -d '{
    "title": "Test",
    "content": "Hello ABlog!",
    "banner": "http://ablog.png",
    "category_id": 1,
    "tag_ids": [1]
  }'
```

返回值：

```json
{
  "success": true,
  "code": 20000,
  "message": "add article successfully",
  "payload": {
    "id": 1,
    "created_at": "2018-05-24T22:42:49.3283467+08:00",
    "updated_at": "2018-05-24T22:42:49.3403381+08:00",
    "deleted_at": null,
    "title": "Test",
    "content": "Hello ABlog!",
    "banner": "http://ablog.png",
    "category": {
      "id": 1,
      "created_at": "2018-05-24T22:24:18+08:00",
      "updated_at": "2018-05-24T22:42:49.338339+08:00",
      "deleted_at": null,
      "name": "ABlog",
      "articles": null
    },
    "category_id": 1,
    "tags": [
      {
        "id": 1,
        "created_at": "2018-05-24T22:34:40+08:00",
        "updated_at": "2018-05-24T22:34:40+08:00",
        "deleted_at": null,
        "name": "HAHA",
        "articles": null
      }
    ],
    "evaluations": null
  }
}
```

## 获取文章列表

路径：`/api/articles`

方法: `GET`

其他：

* 支持分页

请求示例：

```shell
curl -X GET $API_BASE/api/articles
```

返回值：

```json
{
  "success": true,
  "code": 20000,
  "message": "get articles successfully",
  "payload": {
    "page_no": -1,
    "page_size": -1,
    "current_size": 1,
    "total_size": 1,
    "data": []
  }
}
```

## 获取文章详情

路径：`/api/article/:id`

方法: `GET`

路径参数:

* id：文章 ID

请求示例：

```shell
curl -X GET $API_BASE/api/article/1
```

返回值：

```json
{
  "success": true,
  "code": 20000,
  "message": "get article successfully",
  "payload": {
    "id": 1,
    "created_at": "2018-05-24T22:42:49+08:00",
    "updated_at": "2018-05-24T22:42:49+08:00",
    "deleted_at": null,
    "title": "Test",
    "content": "Hello ABlog!",
    "banner": "http://ablog.png",
    "category": {
      "id": 1,
      "created_at": "2018-05-24T22:24:18+08:00",
      "updated_at": "2018-05-24T22:42:49+08:00",
      "deleted_at": null,
      "name": "ABlog",
      "articles": null
    },
    "category_id": 1,
    "tags": [
      {
        "id": 1,
        "created_at": "2018-05-24T22:34:40+08:00",
        "updated_at": "2018-05-24T22:34:40+08:00",
        "deleted_at": null,
        "name": "HAHA",
        "articles": null
      }
    ],
    "evaluations": null
  }
}
```

## 获取文章列表（筛选）

路径：`/api/articles/:flag/:id`

路径参数：

* flag：筛选类型，可选值有：

  * category：按分类筛选
  * tag：按标签筛选

* id：筛选类型 ID

方法: `GET`

其他：

* 支持分页

请求示例：

```shell
curl -X GET $API_BASE/api/articles/category/1
```

返回值：

```json
{
  "success": true,
  "code": 20000,
  "message": "get articles successfully",
  "payload": {
    "page_no": -1,
    "page_size": -1,
    "current_size": 1,
    "total_size": 1,
    "data": [
      {
        "id": 1,
        "created_at": "2018-05-24T22:42:49+08:00",
        "updated_at": "2018-05-24T22:42:49+08:00",
        "deleted_at": null,
        "title": "Test",
        "content": "Hello ABlog!",
        "banner": "http://ablog.png",
        "category": {
          "id": 1,
          "created_at": "2018-05-24T22:24:18+08:00",
          "updated_at": "2018-05-24T22:46:53+08:00",
          "deleted_at": null,
          "name": "ABlog",
          "articles": null
        },
        "category_id": 1,
        "tags": [
          {
            "id": 1,
            "created_at": "2018-05-24T22:34:40+08:00",
            "updated_at": "2018-05-24T22:34:40+08:00",
            "deleted_at": null,
            "name": "HAHA",
            "articles": null
          }
        ],
        "evaluations": null
      }
    ]
  }
}
```

## 评论文章

路径：`/api/article/:id/evaluate`

方法: `POST`

路径参数：

* id：文章 ID

请求参数:

* score：评分
* content：评论内容

请求示例：

```shell
curl -X POST \
  $API_BASE/api/article/1/evaluate \
  -H 'content-type: application/json' \
  -d '{
    "content": "真棒~",
    "score": 5
  }'
```

返回值：

```json
{
  "success": true,
  "code": 20000,
  "message": "add evaluation successfully",
  "payload": {
    "id": 1,
    "created_at": "2018-05-24T22:57:35.2424649+08:00",
    "updated_at": "2018-05-24T22:57:35.246462+08:00",
    "deleted_at": null,
    "ip": "::1",
    "score": 5,
    "content": "真棒~",
    "article": {
      "id": 1,
      "created_at": "2018-05-24T22:42:49+08:00",
      "updated_at": "2018-05-24T22:57:35.2444652+08:00",
      "deleted_at": null,
      "title": "Test",
      "content": "Hello ABlog!",
      "banner": "http://ablog.png",
      "category": {
        "id": 0,
        "created_at": "0001-01-01T00:00:00Z",
        "updated_at": "0001-01-01T00:00:00Z",
        "deleted_at": null,
        "name": "",
        "articles": null
      },
      "category_id": 1,
      "tags": null,
      "evaluations": null
    },
    "article_id": 1
  }
}
```

## 修改文章（需管理员权限）

路径：`/api/article/:id`

方法: `POST`

路径参数

* id：文章 ID

请求参数:

* title：标题
* content：内容
* banner：横幅 url
* category_id：分类 ID
* tag_ids：标签 ID 数组

请求示例：

```shell
curl -X POST \
  $API_BASE/api/article/1 \
  -H 'content-type: application/json' \
  -d '{
    "title": "Test",
    "content": "Hello ABlog!",
    "banner": "http://ablog.png",
    "category_id": 1,
    "tag_ids": [1]
  }'
```

返回值：

```json
{
  "success": true,
  "code": 20000,
  "message": "update article successfully",
  "payload": {
    "id": 1,
    "created_at": "2018-05-24T22:42:49+08:00",
    "updated_at": "2018-05-24T23:04:00+08:00",
    "deleted_at": null,
    "title": "Test",
    "content": "Hello ABlog!",
    "banner": "http://ablog.png",
    "category": {
      "id": 1,
      "created_at": "2018-05-24T22:24:18+08:00",
      "updated_at": "2018-05-24T23:04:00.1127498+08:00",
      "deleted_at": null,
      "name": "ABlog",
      "articles": null
    },
    "category_id": 1,
    "tags": [
      {
        "id": 1,
        "created_at": "2018-05-24T22:34:40+08:00",
        "updated_at": "2018-05-24T22:34:40+08:00",
        "deleted_at": null,
        "name": "HAHA",
        "articles": null
      }
    ],
    "evaluations": null
  }
}
```

## 删除文章（需管理员权限）

路径：`/api/article/:id`

方法: `DELETE`

路径参数:

* id：分类 ID

请求示例：

```shell
curl -X DELETE $API_BASE/api/article/1
```

返回值：

```json
{
  "success": true,
  "code": 20000,
  "message": "delete article successfully",
  "payload": null
}
```
