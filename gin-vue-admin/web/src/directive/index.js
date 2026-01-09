/**
 * 指令集成文件
 * 统一导出所有自定义指令，便于在main.js中一次性注册
 */
import auth from './auth'
import clickOutSide from './clickOutSide'

/**
 * 安装所有指令
 * @param {App} app - Vue应用实例
 */
export const setupDirectives = (app) => {
  // 注册权限指令
  app.use(auth)
  // 注册外部点击指令
  app.use(clickOutSide)
}

export default {
  install: setupDirectives
}
