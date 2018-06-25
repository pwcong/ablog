# 计数接口

- 获取计数
- 增加计数

## 获取计数

路径：`/api/counter/:key/:id`

方法: `GET`

路径参数:

- key
- id

请求示例：

```shell
curl -X GET $API_BASE/api/counter/article/1
```

返回值：

```json
{
  "success": true,
  "code": 20000,
  "message": "get counter successfully",
  "payload": {
    "id": 8,
    "created_at": "2018-06-23T15:07:00+08:00",
    "updated_at": "2018-06-25T14:37:02+08:00",
    "deleted_at": null,
    "key": "article",
    "target_id": 1,
    "count": 287
  }
}
```

## 增加计数

路径：`/api/counter`

方法: `POST`

请求参数:

- key
- id

请求示例：

```shell
curl -X POST $API_BASE/api/counter \
  -H 'Content-Type: application/json' \
  -d '{
    "key": "article",
    "id": 1
  }'
```

返回值：

```json
{
  "success": true,
  "code": 20000,
  "message": "plus counter successfully",
  "payload": {
    "id": 8,
    "created_at": "2018-06-23T15:07:00+08:00",
    "updated_at": "2018-06-25T15:50:01.2997818+08:00",
    "deleted_at": null,
    "key": "article",
    "target_id": 1,
    "count": 288
  }
}
```
