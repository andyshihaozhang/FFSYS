/**
 * 网站配置文件
 */
import packageInfo from '../../package.json'
import { appConfig } from './config.js'

// ANSI 颜色代码
const colors = {
  reset: '\x1b[0m',
  bright: '\x1b[1m',
  cyan: '\x1b[36m',
  green: '\x1b[32m',
  blue: '\x1b[34m',
  magenta: '\x1b[35m',
  yellow: '\x1b[33m'
}

const colorText = (text, color) => `${color}${text}${colors.reset}`
const greenText = (text) => colorText(text, colors.green)
const cyanText = (text) => colorText(text, colors.cyan)
const brightText = (text) => `${colors.bright}${text}${colors.reset}`

// ASCII Art Logo for FF-SYS
const logo = `
  ███████╗███████╗      ███████╗██╗   ██╗███████╗
  ██╔════╝██╔════╝      ██╔════╝╚██╗ ██╔╝██╔════╝
  █████╗  █████╗  █████╗███████╗ ╚████╔╝ ███████╗
  ██╔══╝  ██╔══╝  ╚════╝╚════██║  ╚██╔╝  ╚════██║
  ██║     ██║           ███████║   ██║   ███████║
  ╚═╝     ╚═╝           ╚══════╝   ╚═╝   ╚══════╝
`

export const viteLogo = (env) => {
  if (appConfig.showViteLogo) {
    console.log(cyanText(logo))
    console.log(brightText(`  ${appConfig.appName} Management System`))
    console.log(greenText(`  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━`))
    console.log(`  ${colorText('版本', colors.yellow)}    │ ${colorText(`v${packageInfo.version}`, colors.green)}`)
    console.log(`  ${colorText('API文档', colors.yellow)} │ ${colorText(`http://127.0.0.1:${env.VITE_SERVER_PORT}/swagger/index.html`, colors.blue)}`)
    console.log(`  ${colorText('前端地址', colors.yellow)} │ ${colorText(`http://127.0.0.1:${env.VITE_CLI_PORT}`, colors.blue)}`)
    console.log(greenText(`  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━`))
    console.log(colorText(`  🚀 ${appConfig.appName} 启动成功！`, colors.green))
    console.log('')
  }
}
