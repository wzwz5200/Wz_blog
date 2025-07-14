以下是基于 Go Fiber 的个人博客项目 API 设计中，每个接口的请求体（Request Body）和返回值（Response Body）定义，均为 JSON 格式，遵循 RESTful 风格。

---

## 一、用户模块

### 1. 用户注册
- `POST /api/v1/auth/register`
#### 请求体：
```json
{
  "name": "wz",
  "email": "wz@example.com",
  "password": "123456"
}
```
#### 返回值：
```json
{
  "message": "注册成功",
  "user": {
    "id": 1,
    "name": "wz",
    "email": "wz@example.com"
  }
}
```

### 2. 用户登录
- `POST /api/v1/auth/login`
#### 请求体：
```json
{
  "email": "wz@example.com",
  "password": "123456"
}
```
#### 返回值：
```json
{
  "message": "登录成功",
  "token": "JWT_TOKEN",
  "user": {
    "id": 1,
    "name": "wz"
  }
}
```

### 3. 获取用户信息
- `GET /api/v1/users/:id`
#### 返回值：
```json
{
  "id": 1,
  "name": "wz",
  "email": "wz@example.com"
}
```

---

## 二、文章模块

### 1. 获取文章列表
- `GET /api/v1/posts`
#### 返回值：
```json
{
  "total": 2,
  "posts": [
    {
      "id": 1,
      "title": "第一篇博客",
      "summary": "摘要",
      "author": "wz",
      "created_at": "2025-07-14T12:00:00Z"
    },
    {
      "id": 2,
      "title": "第二篇博客",
      "summary": "摘要",
      "author": "wz",
      "created_at": "2025-07-14T12:10:00Z"
    }
  ]
}
```

### 2. 获取文章详情
- `GET /api/v1/posts/:id`
#### 返回值：
```json
{
  "id": 1,
  "title": "第一篇博客",
  "content": "博客内容",
  "tags": ["Go", "Fiber"],
  "category": "编程",
  "author": "wz",
  "created_at": "2025-07-14T12:00:00Z"
}
```

### 3. 创建文章
- `POST /api/v1/posts`
#### 请求体：
```json
{
  "title": "第一篇博客",
  "content": "博客内容",
  "category_id": 1,
  "tags": ["Go", "Fiber"]
}
```
#### 返回值：
```json
{
  "message": "文章创建成功",
  "post_id": 1
}
```

### 4. 更新文章
- `PUT /api/v1/posts/:id`
#### 请求体：
```json
{
  "title": "更新标题",
  "content": "更新内容",
  "category_id": 2,
  "tags": ["Golang"]
}
```
#### 返回值：
```json
{
  "message": "文章更新成功"
}
```

### 5. 删除文章
- `DELETE /api/v1/posts/:id`
#### 返回值：
```json
{
  "message": "文章已删除"
}
```

---

## 三、评论模块

### 1. 获取文章评论
- `GET /api/v1/posts/:postId/comments`
#### 返回值：
```json
{
  "comments": [
    {
      "id": 1,
      "user": "访客A",
      "content": "写得不错！",
      "created_at": "2025-07-14T12:20:00Z"
    }
  ]
}
```

### 2. 发表评论
- `POST /api/v1/posts/:postId/comments`
#### 请求体：
```json
{
  "content": "写得不错！"
}
```
#### 返回值：
```json
{
  "message": "评论成功",
  "comment_id": 1
}
```

### 3. 回复评论
- `POST /api/v1/comments/:commentId/reply`
#### 请求体：
```json
{
  "content": "谢谢回复"
}
```
#### 返回值：
```json
{
  "message": "回复成功",
  "reply_id": 5
}
```

---

## 四、分类与标签模块

### 获取分类 / 标签
- `GET /api/v1/categories`
- `GET /api/v1/tags`
#### 返回值：
```json
{
  "list": [
    { "id": 1, "name": "编程" },
    { "id": 2, "name": "Go" }
  ]
}
```

### 管理员新增 / 修改 / 删除分类或标签（请求体通用）
```json
{
  "name": "新分类名称"
}
```

#### 返回值：
```json
{
  "message": "操作成功",
  "id": 1
}
```

---

## 五、管理员后台模块

### 获取用户列表
- `GET /api/v1/admin/users`
#### 返回值：
```json
{
  "users": [
    { "id": 1, "name": "wz", "email": "wz@example.com" }
  ]
}
```

### 删除用户 / 评论 / 文章等（通用）
- `DELETE /api/v1/admin/resource/:id`
#### 返回值：
```json
{
  "message": "删除成功"
}
```

如需更多字段或功能（例如分页参数、排序、状态码等），可在此基础上扩展。

