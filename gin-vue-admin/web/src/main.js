import './style/element_visiable.scss'
import 'element-plus/theme-chalk/dark/css-vars.css'
import 'element-plus/dist/index.css'
import 'uno.css'
import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import start from '@/core'
import directives from '@/directive'
import pinia from '@/pinia'
import router from '@/router'
import App from './App.vue'
import '@/core/error-handel'

const app = createApp(App)

app.config.productionTip = false

app
  .use(start) // 应用初始化
  .use(ElementPlus) // el-plus组件初始化
  .use(pinia) // pina初始化
  .use(directives) // 指令初始化
  .use(router) // 路由初始化
  .mount('#app')
export default app
