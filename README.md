# Gin Admin - 简易后台管理系统后端

一个使用 Go + Gin 框架开发的简易后台管理系统后端，支持用户注册/登录（JWT 鉴权）、文章管理（CRUD + 分页搜索 + Redis 缓存），代码结构清晰、分层合理，适合学习和作为简历项目。

## 技术栈

- **Web 框架**：Gin v1.9+
- **ORM**：GORM v2 + MySQL
- **缓存**：Redis (go-redis/v9)
- **认证**：JWT (golang-jwt/jwt/v5) + bcrypt 密码加密
- **响应格式**：统一 Response 结构体 + 全局 recovery
- **中间件**：鉴权、参数校验
- **分层结构**：internal (handler / service / model / middleware) + pkg (工具包)
- **其他**：docker-compose 支持一键启动（待完善）

## 功能亮点

- 用户注册 / 登录（JWT token + 密码哈希）
- 文章管理：创建、列表（分页 + 关键词搜索 + Redis 缓存 5 分钟）、更新、删除
- 权限控制：只能修改/删除自己的文章
- 统一响应格式 + 全局 panic 恢复
- Redis 缓存热点列表数据，减少数据库压力
- 代码规范：table-driven 测试准备、命名清晰、注释适中

