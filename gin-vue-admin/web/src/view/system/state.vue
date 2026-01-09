<template>
  <div class="gva-container2">
    <!-- 第一行：系统监控标题栏 -->
    <div class="bg-white dark:bg-slate-900 text-gray-800 dark:text-gray-400 rounded shadow p-3 mb-4 flex justify-between items-center">
      <div class="flex items-center gap-3">
        <el-icon class="text-xl text-active"><Monitor /></el-icon>
        <span class="text-base font-bold">系统监控</span>
        <el-tag v-if="lastUpdateTime" type="info" size="small">
          更新: {{ lastUpdateTime }}
        </el-tag>
      </div>
      <el-button
        type="primary"
        :icon="loading ? 'Loading' : 'Refresh'"
        :loading="loading"
        circle
        @click="reload"
      />
    </div>

    <!-- 初始加载骨架屏 -->
    <div v-if="!state.os && loading" v-loading="loading" class="p-20">
      <el-skeleton animated :rows="8" />
    </div>

    <!-- 主要内容区域 -->
    <div v-if="state.os" v-loading="loading" class="grid grid-cols-1 gap-4">
      <!-- 第二行：CPU监控 -->
      <div class="bg-white dark:bg-slate-900 text-gray-800 dark:text-gray-400 rounded shadow p-4">
        <div class="flex justify-between items-center mb-3">
          <div class="flex items-center gap-2">
            <el-icon class="text-lg text-warning"><Cpu /></el-icon>
            <span class="text-base font-bold">CPU使用率监控</span>
          </div>
          <el-tag type="warning" size="small">{{ state.cpu.cores }} 物理核心</el-tag>
        </div>
        <div class="w-full overflow-x-auto">
          <div
            ref="cpuChartRef"
            class="w-full h-64"
            :style="{ minWidth: state.cpu && state.cpu.cpus ? `${Math.max(600, state.cpu.cpus.length * 60)}px` : '600px' }"
          />
        </div>
      </div>

      <!-- 第三行：运行环境、内存使用、磁盘使用 -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
        <!-- 运行环境 -->
        <div class="bg-white dark:bg-slate-900 text-gray-800 dark:text-gray-400 rounded shadow p-4">
          <div class="flex items-center gap-2 mb-3">
            <el-icon class="text-lg text-primary"><Setting /></el-icon>
            <span class="text-base font-bold">运行环境</span>
          </div>
          <div class="space-y-2">
            <div class="flex justify-between items-center p-2 bg-gray-50 dark:bg-slate-800 rounded hover:bg-gray-100 dark:hover:bg-slate-700 transition-all">
              <span class="text-sm text-gray-600 dark:text-gray-400">操作系统</span>
              <span class="text-sm font-semibold">{{ state.os.goos }}</span>
            </div>
            <div class="flex justify-between items-center p-2 bg-gray-50 dark:bg-slate-800 rounded hover:bg-gray-100 dark:hover:bg-slate-700 transition-all">
              <span class="text-sm text-gray-600 dark:text-gray-400">CPU核心</span>
              <span class="text-sm font-semibold">{{ state.os.numCpu }} 核</span>
            </div>
            <div class="flex justify-between items-center p-2 bg-gray-50 dark:bg-slate-800 rounded hover:bg-gray-100 dark:hover:bg-slate-700 transition-all">
              <span class="text-sm text-gray-600 dark:text-gray-400">编译器</span>
              <span class="text-sm font-semibold">{{ state.os.compiler }}</span>
            </div>
            <div class="flex justify-between items-center p-2 bg-gray-50 dark:bg-slate-800 rounded hover:bg-gray-100 dark:hover:bg-slate-700 transition-all">
              <span class="text-sm text-gray-600 dark:text-gray-400">Go版本</span>
              <span class="text-sm font-semibold">{{ state.os.goVersion }}</span>
            </div>
            <div class="flex justify-between items-center p-2 bg-gray-50 dark:bg-slate-800 rounded hover:bg-gray-100 dark:hover:bg-slate-700 transition-all">
              <span class="text-sm text-gray-600 dark:text-gray-400">协程数</span>
              <el-tag type="success" size="small">{{ state.os.numGoroutine }}</el-tag>
            </div>
          </div>
        </div>

        <!-- 内存使用 -->
        <div v-if="state.ram" class="bg-white dark:bg-slate-900 text-gray-800 dark:text-gray-400 rounded shadow p-4">
          <div class="flex items-center gap-2 mb-3">
            <el-icon class="text-lg text-danger"><Odometer /></el-icon>
            <span class="text-base font-bold">内存使用</span>
          </div>
          <div class="flex flex-col items-center gap-3">
            <el-progress
              type="dashboard"
              :percentage="state.ram.usedPercent"
              :width="140"
              :color="getProgressColor(state.ram.usedPercent)"
            >
              <template #default="{ percentage }">
                <div class="text-center">
                  <div class="text-2xl font-bold text-gray-800 dark:text-gray-200">{{ percentage }}%</div>
                  <div class="text-xs text-gray-500 dark:text-gray-400 mt-1">已使用</div>
                </div>
              </template>
            </el-progress>
            <div class="w-full space-y-2">
              <div class="flex justify-between items-center p-2 bg-gray-50 dark:bg-slate-800 rounded">
                <span class="text-sm text-gray-600 dark:text-gray-400">总容量</span>
                <span class="text-sm font-semibold">{{ (state.ram.totalMb / 1024).toFixed(2) }} GB</span>
              </div>
              <div class="flex justify-between items-center p-2 bg-gray-50 dark:bg-slate-800 rounded">
                <span class="text-sm text-gray-600 dark:text-gray-400">已使用</span>
                <span class="text-sm font-semibold text-warning">{{ (state.ram.usedMb / 1024).toFixed(2) }} GB</span>
              </div>
              <div class="flex justify-between items-center p-2 bg-gray-50 dark:bg-slate-800 rounded">
                <span class="text-sm text-gray-600 dark:text-gray-400">可用</span>
                <span class="text-sm font-semibold text-success">{{ ((state.ram.totalMb - state.ram.usedMb) / 1024).toFixed(2) }} GB</span>
              </div>
            </div>
          </div>
        </div>

        <!-- 磁盘使用 -->
        <div v-if="state.disk" class="bg-white dark:bg-slate-900 text-gray-800 dark:text-gray-400 rounded shadow p-4">
          <div class="flex justify-between items-center mb-3">
            <div class="flex items-center gap-2">
              <el-icon class="text-lg text-success"><Document /></el-icon>
              <span class="text-base font-bold">磁盘使用</span>
            </div>
            <el-tag type="info" size="small">{{ state.disk.length }} 个分区</el-tag>
          </div>
          <div class="space-y-3 max-h-64 overflow-y-auto">
            <div
              v-for="(item, index) in state.disk.slice(0, 3)"
              :key="index"
              class="p-3 bg-gray-50 dark:bg-slate-800 rounded hover:bg-gray-100 dark:hover:bg-slate-700 transition-all"
            >
              <div class="flex justify-between items-center mb-2">
                <el-tag type="primary" size="small">{{ item.mountPoint }}</el-tag>
                <span class="text-xs font-semibold text-gray-600 dark:text-gray-400">{{ item.usedGb }} / {{ item.totalGb }} GB</span>
              </div>
              <el-progress
                :percentage="item.usedPercent"
                :color="getProgressColor(item.usedPercent)"
                :stroke-width="8"
              >
                <template #default="{ percentage }">
                  <span class="text-xs font-semibold ml-2">{{ percentage.toFixed(1) }}%</span>
                </template>
              </el-progress>
            </div>
            <div v-if="state.disk.length > 3" class="text-center text-sm text-gray-500 dark:text-gray-400 py-2">
              还有 {{ state.disk.length - 3 }} 个分区...
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup>
import { getSystemState } from '@/api/system'
import { onUnmounted, onMounted, ref, nextTick, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { Cpu, Document, Odometer, Monitor, Setting } from '@element-plus/icons-vue'
import * as echarts from 'echarts'

defineOptions({
  name: 'State'
})

// 响应式状态
const timer = ref(null)
const state = ref({})
const loading = ref(false)
const lastUpdateTime = ref('')
const cpuChartRef = ref(null)
let cpuChart = null

/**
 * 根据使用率百分比获取进度条颜色
 * @param {number} percentage 使用率百分比
 * @returns {string} 颜色值
 */
const getProgressColor = (percentage) => {
  if (percentage < 50) {
    return '#67C23A' // 绿色
  } else if (percentage < 75) {
    return '#E6A23C' // 橙色
  } else if (percentage < 90) {
    return '#F56C6C' // 红色
  } else {
    return '#F56C6C' // 深红色
  }
}

/**
 * 格式化时间
 */
const formatTime = () => {
  const now = new Date()
  const hours = String(now.getHours()).padStart(2, '0')
  const minutes = String(now.getMinutes()).padStart(2, '0')
  const seconds = String(now.getSeconds()).padStart(2, '0')
  return `${hours}:${minutes}:${seconds}`
}

/**
 * 初始化CPU图表
 */
const initCpuChart = () => {
  if (!cpuChartRef.value || !state.value.cpu) return

  if (cpuChart) {
    cpuChart.dispose()
  }

  cpuChart = echarts.init(cpuChartRef.value)

  const cpuData = state.value.cpu.cpus || []
  const cores = cpuData.map((_, index) => `核心${index}`)

  const option = {
    grid: {
      left: '50px',
      right: '30px',
      top: '30px',
      bottom: '30px',
      containLabel: true
    },
    tooltip: {
      trigger: 'axis',
      axisPointer: {
        type: 'shadow'
      },
      formatter: (params) => {
        const data = params[0]
        return `${data.name}<br/>${data.marker}使用率: ${data.value.toFixed(1)}%`
      }
    },
    xAxis: {
      type: 'category',
      data: cores,
      axisLabel: {
        interval: 0,
        rotate: cores.length > 8 ? 45 : 0,
        fontSize: 11
      }
    },
    yAxis: {
      type: 'value',
      max: 100,
      axisLabel: {
        formatter: '{value}%'
      }
    },
    series: [
      {
        name: 'CPU使用率',
        type: 'bar',
        data: cpuData.map(val => val.toFixed(1)),
        itemStyle: {
          color: (params) => {
            const value = params.value
            if (value < 50) return '#67C23A'
            if (value < 75) return '#E6A23C'
            if (value < 90) return '#F56C6C'
            return '#F56C6C'
          }
        },
        label: {
          show: true,
          position: 'top',
          formatter: '{c}%',
          fontSize: 10
        },
        barMaxWidth: 40
      }
    ]
  }

  cpuChart.setOption(option)

  // 响应式调整
  window.addEventListener('resize', () => {
    cpuChart?.resize()
  })
}

/**
 * 更新CPU图表
 */
const updateCpuChart = () => {
  if (!cpuChart || !state.value.cpu) return

  const cpuData = state.value.cpu.cpus || []

  cpuChart.setOption({
    series: [
      {
        data: cpuData.map(val => val.toFixed(1))
      }
    ]
  })
}

/**
 * 重新加载系统状态
 */
const reload = async () => {
  try {
    loading.value = true
    const { data } = await getSystemState()
    state.value = data.server
    lastUpdateTime.value = formatTime()

    // 初始化或更新图表
    if (state.value.cpu) {
      await nextTick()
      if (!cpuChart) {
        initCpuChart()
      } else {
        updateCpuChart()
      }
    }
  } catch (error) {
    ElMessage.error('获取系统状态失败')
    console.error('Failed to get system state:', error)
  } finally {
    loading.value = false
  }
}

// 监听state变化，初始化图表
watch(() => state.value.cpu, async (newVal) => {
  if (newVal && cpuChartRef.value) {
    await nextTick()
    initCpuChart()
  }
}, { deep: true })

// 组件挂载
onMounted(async () => {
  await reload()
})

// 初始加载
reload()

// 设置定时器，每10秒自动刷新
timer.value = setInterval(() => {
  reload()
}, 1000 * 60)

// 组件卸载时清除定时器和图表
onUnmounted(() => {
  clearInterval(timer.value)
  timer.value = null

  if (cpuChart) {
    cpuChart.dispose()
    cpuChart = null
  }

  window.removeEventListener('resize', () => {
    cpuChart?.resize()
  })
})
</script>

<style scoped lang="scss">
// 移除滚动条样式，使用项目全局配置
</style>
