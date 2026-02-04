# 蛋蛋开发文档

## 服务模块设计
整体说明：
- 应用服务采用 golang，数据库用 mysql，缓存用 redis
- 推荐算法服务采用 python
- 采用微服务架构设计，加密服务、认证服务、用户服务、问答服务、评论服务、笔记服务、聊天服务、地球村服务、定时器服务、智能体员工调度服务

安全模块：
- 采用端到端解密，用户数据在传输过程中加密，确保数据安全
- 登录状态通过 cookie 校验
- 数据库中用户保存用户敏感数据的哈希值，登录时对比哈希值校验密码
- 日志打印避免打印敏感信息，如密码、手机号等
- 前端避免直接展示敏感信息，如密码、手机号等
- 前端渲染避免xss攻击，对用户输入进行转义

### 注册登录模块
- 功能：用户注册、登录、退出
- 接口：
    - 注册：POST /register
    - 登录：POST /login
    - 退出：POST /logout
    - 检查登录状态：GET /check-login

## 用户模块
- 功能：用户个人信息管理
- 接口：
    - 获取用户信息：GET /user
    - 更新用户信息：PUT /user

## 问答模块
- 功能：用户发布问题、回答问题
- 接口：
    - 发布问题：POST /question
    - 回答问题：POST /answer
    - 获取问题列表：GET /questions
    - 获取问题详情：GET /question/:id
    - 点赞问题：POST /question/:id/like
    - 取消点赞问题：POST /question/:id/unlike

## 评论模块
- 功能：用户对问题或回答进行评论
- 接口：
    - 发布评论：POST /comment
    - 获取评论列表：GET /comments
    - 获取评论详情：GET /comment/:id
    - 点赞评论：POST /comment/:id/like
    - 取消点赞评论：POST /comment/:id/unlike
    - 删除评论：DELETE /comment/:id
    - 回复评论：POST /comment/:id/reply

## 笔记模块
- 功能：用户发布笔记、查看笔记
- 接口：
    - 发布笔记：POST /note
    - 获取笔记列表：GET /notes
    - 获取笔记详情：GET /note/:id
    - 笔记分类列表：GET /note/categories
    - 获取分类下的笔记列表：GET /note/category/:id

## 聊天模块
- 功能：用户与智能体、员工进行聊天
- 接口：
    - 发送消息：POST /chat
    - 获取聊天记录：GET /chat/:id

## 地球村模块
- 功能：用户在地球村进行互动、交流
- 接口：
    - 加入村落：POST /earth-village/join
    - 退出村落：POST /earth-village/leave
    - 获取村落列表：GET /earth-villages
    - 获取村落详情：GET /earth-village/:id
    - 发送帖子到村落：POST /earth-village/:id/post
    - 获取帖子列表：GET /earth-village/:id/posts
    - 点赞帖子：POST /earth-village/:id/post/:postId/like
    - 取消点赞帖子：POST /earth-village/:id/post/:postId/unlike
    - 删除帖子：DELETE /earth-village/:id/post/:postId
    - 评论帖子：POST /earth-village/:id/post/:postId/reply
    - 获取评论列表：GET /earth-village/:id/post/:postId/replies

## 用户端设计

H5 网页，Vue3 框架，自己实现 UI 交互，使用动画库实现页面交互效果

- 支持中文、英文、日语、法语、阿拉伯语、韩语、德语、西班牙语
- 响应式布局，适配不同屏幕尺寸
- 优化加载速度，避免白屏时间过长
- 实现用户注册、登录、退出功能
- 实现问题发布、回答、查看、点赞、取消点赞、评论、回复、查看评论功能
- 实现笔记发布、查看、分类、查看分类下笔记功能
- 实现聊天功能，与智能体、员工进行互动
- 实现加入村落、退出村落、查看村落列表、查看村落详情、发送帖子到村落、查看帖子列表、点赞帖子、取消点赞帖子、删除帖子、评论帖子、查看评论功能

