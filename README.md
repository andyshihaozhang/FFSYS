# gin-vue-admin

一个现代化的全栈管理系统框架，基于 Go + Vue 3 构建，提供前后端分离、插件化、可扩展的解决方案。
本项目基于gin-vue-admin进行改造 https://www.gin-vue-admin.com

[English](./README_EN.md) | 中文

## 🌟 主要特性

- ✅ **完整的RBAC权限控制** - 基于Casbin的灵活权限管理体系
- ✅ **前后端完全分离** - 前后端独立开发、独立部署
- ✅ **插件化架构** - 模块化设计，支持自定义插件开发
- ✅ **代码自动生成** - 支持CRUD代码一键生成
- ✅ **多数据库支持** - MySQL、PostgreSQL、SQLite、SQL Server、MongoDB
- ✅ **云存储集成** - 阿里云OSS、AWS S3、MinIO、七牛云、腾讯云COS等
- ✅ **Swagger API文档** - 自动生成API文档，便于前后端协作
- ✅ **完善的中间件** - 支持认证、授权、日志、CORS等
- ✅ **实时日志** - 完整的操作日志和审计记录
- ✅ **表单生成器** - 可视化表单设计和生成

## 🛠️ 技术栈

### 后端技术栈

| 技术 | 版本 | 说明 |
|------|------|------|
| Go | 1.24.0 | 编程语言 |
| Gin | 1.10.0 | Web框架 |
| GORM | 1.25+ | ORM框架 |
| Casbin | 2.103.0 | 权限管理 |
| JWT | 5.2.2 | 身份验证 |
| Viper | 1.19.0 | 配置管理 |
| Zap | 1.27.0 | 日志系统 |
| Redis | 9.7.0 | 缓存系统 |

### 前端技术栈

| 技术 | 版本 | 说明 |
|------|------|------|
| Vue | 3.5.7 | 前端框架 |
| Vite | 6.2.3 | 构建工具 |
| Element Plus | 2.10.2 | UI组件库 |
| Pinia | 2.2.2 | 状态管理 |
| Vue Router | 4.4.3 | 路由管理 |
| Axios | 1.8.2 | HTTP客户端 |
| UnoCSS | 66.4.2 | 原子化CSS |
| ECharts | 5.5.1 | 数据可视化 |

## 📁 项目结构

```
gin-vue-admin/
├── server/                 # 后端服务 (Go)
│   ├── api/v1/            # API控制器层
│   │   ├── system/        # 系统模块
│   │   ├── example/       # 示例模块
│   │   └── ffproduction/  # 业务模块
│   ├── config/            # 配置结构体
│   ├── core/              # 核心启动
│   ├── docs/              # Swagger文档
│   ├── global/            # 全局变量
│   ├── initialize/        # 初始化模块
│   ├── middleware/        # 中间件
│   ├── model/             # 数据模型
│   │   ├── system/
│   │   ├── request/
│   │   └── response/
│   ├── plugin/            # 插件系统
│   │   ├── announcement/  # 公告插件
│   │   └── email/         # 邮件插件
│   ├── router/            # 路由定义
│   ├── service/           # 业务服务层
│   ├── source/            # 初始化数据
│   ├── utils/             # 工具函数
│   ├── config.yaml        # 配置文件
│   ├── go.mod             # 依赖管理
│   └── main.go            # 程序入口
│
├── web/                   # 前端应用 (Vue 3)
│   ├── public/            # 静态资源
│   ├── src/
│   │   ├── api/           # API接口封装
│   │   ├── components/    # 公共组件
│   │   ├── core/          # 核心配置
│   │   ├── directive/     # 自定义指令
│   │   ├── pinia/         # 状态管理
│   │   │   └── modules/   # 状态模块
│   │   ├── plugin/        # 前端插件
│   │   ├── router/        # 路由配置
│   │   ├── style/         # 全局样式
│   │   ├── utils/         # 工具函数
│   │   ├── view/          # 页面组件
│   │   ├── App.vue        # 根组件
│   │   └── main.js        # 程序入口
│   ├── package.json       # 依赖配置
│   ├── vite.config.js     # Vite配置
│   ├── uno.config.js      # UnoCSS配置
│   └── eslint.config.mjs  # ESLint配置
│
├── deploy/                # 部署配置
│   └── Dockerfile         # Docker配置
│
└── CLAUDE.md             # AI开发指南
```

## 🚀 快速开始

### 前置要求

- **后端**: Go 1.24+
- **前端**: Node.js 16+ (推荐 18+)
- **数据库**: MySQL 5.7+ / PostgreSQL 12+ / SQLite 3+
- **缓存**: Redis 5.0+ (可选)

### 后端开发

1. **进入后端目录**
```bash
cd server
```

2. **配置数据库** (`config.yaml`)
```yaml
mysql:
  path: 127.0.0.1:3306
  config: charset=utf8mb4&parseTime=True&loc=Local
  dbname: gva_db
  username: root
  password: 123456
```

3. **运行开发服务**
```bash
go run .
```

服务将在 `http://localhost:8888` 启动

4. **生成API文档**
```bash
swag init
# 访问 http://localhost:8888/swagger/index.html
```

5. **编译生产版本**
```bash
go build
```

### 前端开发

1. **进入前端目录**
```bash
cd web
```

2. **安装依赖**
```bash
npm install
# 或使用 yarn/pnpm
yarn install
```

3. **启动开发服务器**
```bash
npm run dev
# 或
npm run serve
```

应用将在 `http://localhost:5173` 启动

4. **生产构建**
```bash
npm run build
```

5. **预览生产版本**
```bash
npm run preview
```

## 📚 开发指南

### 后端开发指南

#### 严格的分层架构

```
路由层 (Router) → API层 (Api) → 服务层 (Service) → 模型层 (Model)
```

- **Model层**: 定义数据模型和请求/响应结构体
- **Service层**: 实现业务逻辑，纯业务代码不涉及HTTP
- **API层**: 处理HTTP请求和响应，调用Service层
- **Router层**: 定义路由和中间件

#### 常用命令

```bash
# 从 server/ 目录执行

# 下载依赖
go mod tidy

# 运行开发服务器
go run .

# 生成Swagger文档
swag init

# 运行测试
go test ./...

# 构建二进制文件
go build -o gin-vue-admin

# 代码生成 (如果配置了自动生成)
go run . --auto-code
```

#### 项目配置 (`server/config.yaml`)

**主要配置项**:
- `system.addr`: 服务地址和端口
- `mysql/postgres/sqlite`: 数据库连接
- `redis`: Redis缓存配置
- `jwt`: JWT令牌配置
- `cors`: CORS跨域配置
- `oss`: 云存储配置

### 前端开发指南

#### 组件开发规范

```vue
<script setup>
// 使用 Composition API
import { ref, computed } from 'vue'
import { getUserList } from '@/api/user'

// 状态定义
const users = ref([])
const loading = ref(false)

// 计算属性
const total = computed(() => users.value.length)

// 方法定义
const fetchUsers = async () => {
  loading.value = true
  try {
    const res = await getUserList()
    users.value = res.data
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="users-container">
    <el-table :data="users" v-loading="loading">
      <!-- 表格内容 -->
    </el-table>
  </div>
</template>
```

#### API接口编写

```javascript
import service from '@/utils/request'

/**
 * 获取用户列表
 * @param {Object} params - 查询参数
 * @param {number} params.page - 页码
 * @param {number} params.pageSize - 每页数量
 * @returns {Promise} 用户列表
 */
export const getUserList = (params) => {
  return service({
    url: '/user/getUserList',
    method: 'post',
    data: params
  })
}
```

#### 常用命令

```bash
# 从 web/ 目录执行

# 安装依赖
npm install

# 启动开发服务器
npm run dev

# 生产构建
npm run build

# 代码检查
npm run lint

# 预览生产构建
npm run preview
```

## 🔐 权限管理

项目使用Casbin进行权限管理，支持多种权限模型（RBAC、ABAC等）。

### 创建新的权限

在 `menu` 表中添加菜单项，自动生成对应权限。API权限通过 `cas_enforcer` 表管理。

### 权限验证

```go
// 后端权限检查
ok, err := global.GVA_ENFORCER.Enforce(userId, resource, action)
```

```javascript
// 前端权限指令
<button v-auth="['admin', 'system:user:create']">创建用户</button>
```

## 🔌 插件开发

项目支持插件化开发，可轻松扩展功能。

### 插件结构

```
server/plugin/[plugin-name]/
├── api/
├── initialize/
├── model/
├── router/
├── service/
└── plugin.go
```

### 创建新插件

参考 `server/plugin/announcement/` 目录了解完整插件结构。

### 注册插件

在 `main.go` 中注册插件：

```go
// 初始化插件
pluginManager.Register(pluginInstance)
```

## 💾 数据库支持

项目支持多种数据库：

- **MySQL** - 推荐，性能最优
- **PostgreSQL** - 企业级选择
- **SQLite** - 轻量级部署
- **SQL Server** - Windows环境
- **MongoDB** - 文档数据库

在 `config.yaml` 中配置：

```yaml
# MySQL
mysql:
  path: 127.0.0.1:3306
  dbname: gva_db
  username: root
  password: password

# PostgreSQL
postgres:
  path: 127.0.0.1:5432
  dbname: gva_db
  username: postgres
  password: password

# SQLite
sqlite:
  path: ./data.db
```

## ☁️ 云存储集成

支持多家云存储服务商：

- **阿里云 OSS**
- **AWS S3**
- **MinIO**
- **七牛云**
- **腾讯云 COS**
- **华为云 OBS**
- **Cloudflare R2**

在 `config.yaml` 中配置相应的云存储凭证。

## 📖 API文档

启动服务后，访问以下地址查看API文档：

- **Swagger UI**: http://localhost:8888/swagger/index.html

### 常见API响应格式

**成功响应**:
```json
{
  "code": 0,
  "msg": "success",
  "data": {
    // 响应数据
  }
}
```

**列表响应**:
```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "list": [ /* 数据列表 */ ],
    "total": 100,
    "page": 1,
    "pageSize": 10
  }
}
```

**错误响应**:
```json
{
  "code": 500,
  "msg": "错误信息",
  "data": null
}
```

## 🐳 Docker部署

### 构建Docker镜像

**后端镜像**:
```bash
cd server
docker build -t gin-vue-admin-server .
docker run -d -p 8888:8888 gin-vue-admin-server
```

**前端镜像**:
```bash
cd web
docker build -t gin-vue-admin-web .
docker run -d -p 80:80 gin-vue-admin-web
```

### Docker Compose部署

```bash
docker-compose up -d
```

参考 `deploy/docker-compose.yml` 了解详细配置。