# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**gin-vue-admin** is a full-stack management system framework built with modern technologies:
- **Backend**: Go 1.24 + Gin 1.10 + GORM 1.25 + Casbin 2.103
- **Frontend**: Vue 3.5.7 + Vite 6.2.3 + Pinia 2.2.2 + Element Plus 2.10.2
- **Architecture**: Front-end and back-end separation with plugin-based extensibility
- **Databases**: MySQL, PostgreSQL, SQLite, SQL Server, MongoDB support
- **Cloud Storage**: Aliyun OSS, AWS S3, MinIO, Qiniu, Tencent COS, Huawei OBS, Cloudflare R2

## Directory Structure

```
gin-vue-admin/
├── server/                 # Go backend (open server/ in IDE, not root)
│   ├── api/v1/            # API controllers (system, example, ffproduction modules)
│   ├── config/            # Configuration structures
│   ├── core/              # Core initialization (zap, viper, server)
│   ├── initialize/        # Initialization modules (db, router, redis, validator)
│   ├── middleware/        # Gin middleware (auth, cors, jwt, logging, etc)
│   ├── model/             # Data models + request/response DTOs
│   ├── plugin/            # Plugin system (announcement, email, etc)
│   ├── router/            # Route definitions
│   ├── service/           # Business logic layer
│   ├── mcp/               # MCP (Model Context Protocol) integration
│   ├── utils/             # Utilities (upload, timer, ast, etc)
│   ├── config.yaml        # Main configuration file
│   ├── go.mod             # Go dependencies
│   └── main.go            # Entry point
├── web/                   # Vue frontend
│   ├── src/
│   │   ├── api/           # API service wrappers
│   │   ├── components/    # Global reusable components
│   │   ├── core/          # Core config and setup
│   │   ├── directive/     # Custom directives (v-auth)
│   │   ├── pinia/         # State management (user, router, dictionary)
│   │   ├── router/        # Route configuration
│   │   ├── style/         # Global styles (SCSS)
│   │   ├── utils/         # Utility functions
│   │   ├── view/          # Page components
│   │   ├── plugin/        # Frontend plugins
│   │   ├── App.vue        # Root component
│   │   └── main.js        # Entry point
│   ├── package.json       # Dependencies
│   ├── vite.config.js     # Vite configuration
│   ├── uno.config.js      # UnoCSS atomic CSS config
│   └── eslint.config.mjs  # ESLint configuration
└── deploy/                # Docker and deployment configs
```

## Backend Development

### Architecture Principles

**Strict Layered Architecture** (non-negotiable):
- **Router** → **API** → **Service** → **Model** (unidirectional dependency)
- Each layer has single responsibility; cross-layer calls are forbidden
- API layer never touches database directly; Service layer never handles `gin.Context`

**enter.go Pattern** (mandatory):
- All `api/`, `service/`, `router/` directories use `enter.go` to expose groups
- Global instances (`service.ServiceGroupApp`, `api.ApiGroupApp`, `router.RouterGroupApp`) are the only inter-module communication points
- Prevents circular dependencies

### Common Backend Commands

```bash
# From server/ directory
go generate                    # Download dependencies
go run .                       # Run development server (port 8888)
go build                       # Build binary
go test ./...                  # Run all tests
go test ./service/...          # Run specific package tests

# Swagger API documentation
swag init                      # Generate Swagger docs (creates docs/ folder)
# Then visit http://localhost:8888/swagger/index.html

# Code generation (if using auto-code feature)
go run . --auto-code           # Generate CRUD code
```

### Backend Code Organization

**Model Layer** (`model/`):
- Data models inherit `global.GVA_MODEL` (includes ID, CreatedAt, UpdatedAt)
- Request models in `model/request/` with `json` and `form` tags
- Response models in `model/response/`
- **Critical**: Maintain strict type consistency across data/request/response models

**Service Layer** (`service/`):
- Pure business logic, no HTTP/Gin code
- Functions return `(result, error)` pattern
- Access database through GORM
- Example: `service.ServiceGroupApp.SysUserService.CreateUser(user)`

**API Layer** (`api/`):
- HTTP request handlers with complete Swagger comments
- Calls Service layer via `service.ServiceGroupApp`
- Handles parameter validation and response formatting
- **Mandatory Swagger format**:
  ```go
  // CreateUser creates a new user
  // @Tags     User
  // @Summary  Create user
  // @Security ApiKeyAuth
  // @accept   application/json
  // @Produce  application/json
  // @Param    data body request.CreateUserRequest true "User data"
  // @Success  200  {object} response.Response{msg=string}
  // @Router   /user/createUser [post]
  func (a *UserApi) CreateUser(c *gin.Context) { ... }
  ```

**Router Layer** (`router/`):
- Maps HTTP paths to API handlers
- Configures middleware per route
- Uses route groups for organization
- Example: `router.RouterGroupApp.UserRouter.InitUserRouter(r)`

### Plugin Development

Reference: `server/plugin/announcement/` (canonical example)

Structure:
```
server/plugin/[name]/
├── api/enter.go              # API group definition
├── api/[module].go           # API handlers
├── config/config.go          # Plugin config struct
├── initialize/
│   ├── api.go               # Register APIs
│   ├── gorm.go              # DB migrations (AutoMigrate)
│   ├── menu.go              # Create menu items
│   ├── router.go            # Register routes
│   └── viper.go             # Load config
├── model/[model].go         # Data models
├── model/request/[req].go   # Request DTOs
├── router/enter.go          # Router group
├── router/[module].go       # Route definitions
├── service/enter.go         # Service group
├── service/[module].go      # Business logic
└── plugin.go                # Plugin entry (implements system.Plugin)
```

## Frontend Development

### Architecture Principles

**Modular Architecture**:
- Page → API Service → Backend (unidirectional)
- No cross-module direct calls
- All API calls through `src/api/` files

**State Management**:
- Use Pinia stores in `src/pinia/modules/`
- Never modify state directly; use actions
- Example stores: `user.js`, `router.js`, `dictionary.js`

### Common Frontend Commands

```bash
# From web/ directory
npm install                    # Install dependencies
npm run dev                    # Development server (http://localhost:5173)
npm run serve                  # Alternative dev command
npm run build                  # Production build
npm run preview               # Preview production build
npm run lint                  # Run ESLint
```

### Frontend Code Organization

**API Layer** (`src/api/`):
- Encapsulate all backend calls
- Use `@/utils/request.js` for HTTP
- Include JSDoc comments
- Example:
  ```javascript
  /**
   * Get user list
   * @param {Object} data - Query params
   * @returns {Promise} User list
   */
  export const getUserList = (data) => {
    return service({
      url: '/user/getUserList',
      method: 'post',
      data
    })
  }
  ```

**Components** (`src/components/`):
- Reusable UI elements
- Single responsibility
- Props and events clearly defined
- Use Composition API with `<script setup>`

**Pages** (`src/view/`):
- Business page implementations
- Use Composition API
- Handle loading/error states
- Organized by business module

**State Management** (`src/pinia/modules/`):
- Composition API style stores
- Return state, computed, and actions
- Example structure:
  ```javascript
  export const useUserStore = defineStore('user', () => {
    const userInfo = ref({})
    const token = useStorage('token', '')

    const setUserInfo = (val) => { userInfo.value = val }
    const login = async (form) => { /* ... */ }

    return { userInfo, token, setUserInfo, login }
  })
  ```

## Frontend-Backend Integration

### Data Type Consistency (Critical)

**Must maintain identical types across layers**:
- Backend Go struct field → Frontend JavaScript/TypeScript type
- Example: Go `Status int` must be `status: number` in frontend
- Pointer types (`*string`) serialize to base type or null in JSON
- Enum fields must use consistent type (not mixed string/number)

### API Response Format

**Standard response structure**:
```json
{
  "code": 0,
  "data": { /* response data */ },
  "msg": "success"
}
```

**Pagination format**:
```json
{
  "page": 1,
  "pageSize": 10,
  "total": 100,
  "list": [ /* items */ ]
}
```

### Time Format

- Use ISO 8601 standard (e.g., `2024-01-23T10:30:00Z`)
- Backend: GORM handles automatically
- Frontend: Use `@vueuse/core` utilities for parsing

## Development Workflow

### Backend Development

1. **Design models** in `model/` (data, request, response)
2. **Implement service** in `service/` (business logic)
3. **Create API handlers** in `api/` (with Swagger comments)
4. **Define routes** in `router/`
5. **Generate Swagger**: `swag init`
6. **Test**: `go test ./...`

### Frontend Development

1. **Create API wrapper** in `src/api/`
2. **Build components** in `src/components/`
3. **Implement page** in `src/view/`
4. **Add state** in `src/pinia/` if needed
5. **Test**: `npm run lint`

### Testing

- Backend: `go test ./service/...` (unit tests)
- Frontend: ESLint validation with `npm run lint`
- Integration: Manual testing with both servers running

## Configuration

### Backend Config (`server/config.yaml`)

Key sections:
- `system`: Server port, environment, multipoint login
- `mysql`/`postgres`/`sqlite`: Database connection
- `redis`: Cache configuration
- `jwt`: Token settings
- `cors`: CORS configuration
- `oss`: Cloud storage settings

### Frontend Config

- `.env.development`: Dev environment variables
- `.env.production`: Production environment variables
- `vite.config.js`: Build and dev server settings
- `uno.config.js`: UnoCSS atomic CSS presets

## Important Notes

### IDE Setup

- **Backend**: Open `server/` directory in IDE (not root)
- **Frontend**: Open `web/` directory or use VSCode workspace
- **Workspace**: Use `gin-vue-admin.code-workspace` for both in VSCode

### Git Workflow

- Branch naming: `feature/*`, `fix/*`, `docs/*`
- Commit format: `[filename]: description` (e.g., `[user.go]: add user validation`)
- PR requires 2 maintainer approvals

### Performance Considerations

- Frontend: Use lazy loading for routes, virtual scrolling for large lists
- Backend: Index frequently queried fields, use Redis for caching
- Both: Minimize API calls, batch operations when possible

### Security Best Practices

- Never commit `.env` files or credentials
- Use JWT for authentication (configured in `config.yaml`)
- Implement RBAC with Casbin (role-based access control)
- Validate all user input in API layer
- Use CORS middleware appropriately

## Useful Resources

- **Documentation**: https://www.gin-vue-admin.com
- **Online Demo**: http://demo.gin-vue-admin.com (admin/123456)
- **Swagger Docs**: http://localhost:8888/swagger/index.html (after `swag init`)
- **Plugin Market**: https://plugin.gin-vue-admin.com/
- **Community**: QQ Group 971857775

## MCP Integration

The project includes MCP (Model Context Protocol) support for AI-assisted development:
- Located in `server/mcp/` directory
- Provides tools for code generation and analysis
- Consult `.cursor/rules/project_rules.md` for detailed MCP guidelines
