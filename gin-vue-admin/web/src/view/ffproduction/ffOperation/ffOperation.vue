<template>
    <div class="flex gap-4 p-4 h-[calc(100vh-120px)]">
        <!-- 左侧产品列表 -->
        <div class="w-350px bg-white rounded-lg shadow-sm p-4 flex flex-col">
            <div class="mb-4">
                <h3 class="text-lg font-semibold m-0">产品列表</h3>
            </div>

            <!-- 产品筛选 -->
            <div class="mb-4">
                <el-input v-model="productSearch"
                    placeholder="搜索产品名称或款号"
                    clearable
                    prefix-icon="Search"
                    @input="handleProductSearch" />
            </div>

            <!-- 产品列表 -->
            <div class="flex-1 overflow-hidden">
                <el-scrollbar height="calc(100vh - 240px)">
                    <div class="px-2 pt-1 pb-1">
                        <div v-for="product in filteredProducts"
                            :key="product.id"
                            class="product-item flex items-center justify-between p-3 mb-2 bg-gray-50 rounded-md cursor-pointer transition-all hover:bg-gray-100 hover:shadow-md hover:-translate-y-0.5"
                            :class="{
                'border-2 border-blue-500 shadow-lg bg-blue-50 -translate-y-0.5': selectedProduct?.id === product.id,
                'border-2 border-transparent': selectedProduct?.id !== product.id
              }"
                            @click="selectProduct(product)">
                            <div class="flex-1 min-w-0">
                                <div class="text-sm font-medium mb-1 truncate">
                                    {{ product.productName }}
                                </div>
                                <div class="text-xs text-gray-500 mb-1.5">款号: {{ product.productCode }}</div>
                                <div class="flex gap-1">
                                    <el-tag size="small">{{ product.productColor }}</el-tag>
                                    <el-tag :type="getProductFlagType(product.productFlag)"
                                        size="small">
                                        {{ product.productFlag }}
                                    </el-tag>
                                </div>
                            </div>
                            <el-icon class="text-base text-gray-400 transition-all"
                                :class="{ 'text-blue-500': selectedProduct?.id === product.id }">
                                <ArrowRight />
                            </el-icon>
                        </div>
                    </div>
                    <el-empty v-if="filteredProducts.length === 0"
                        description="暂无产品数据"
                        :image-size="80" />
                </el-scrollbar>
            </div>
        </div>

        <!-- 右侧工序管理 -->
        <div class="flex-1 bg-white rounded-lg shadow-sm p-4 overflow-hidden">
            <div v-if="!selectedProduct"
                class="h-full flex items-center justify-center">
                <el-empty description="请先选择一个产品"
                    :image-size="120" />
            </div>

            <div v-else
                class="h-full flex flex-col">
                <!-- 工序表头 -->
                <div class="flex justify-between items-center mb-4 pb-4 border-b border-gray-100">
                    <div>
                        <h3 class="text-lg font-semibold m-0 mb-2">工序管理</h3>
                        <div class="text-sm text-gray-500">
                            当前产品: {{ selectedProduct.productName }} ({{ selectedProduct.productCode }})
                        </div>
                    </div>
                    <el-button type="primary"
                        icon="Plus"
                        @click="openDrawer()">
                        新增工序
                    </el-button>
                </div>

                <!-- 工序表格 -->
                <div class="flex-1 overflow-hidden">
                    <el-table v-loading="operationLoading"
                        :data="operationData"
                        row-key="id"
                        border
                        stripe
                        height="calc(100vh - 240px)">
                        <el-table-column prop="id"
                            label="序号"
                            width="80"
                            align="center" />
                        <el-table-column prop="operationName"
                            label="工序名称"
                            min-width="200"
                            align="center">
                            <template #default="{ row }">
                                <el-tag type="primary"
                                    size="large">{{ row.operationName }}</el-tag>
                            </template>
                        </el-table-column>
                        <el-table-column prop="referencePrice"
                            label="参考单价"
                            min-width="150"
                            align="center">
                            <template #default="{ row }">
                                <span class="text-base font-semibold text-red-500">
                                    ¥{{ row.referencePrice?.toFixed(2) || '0.00' }}
                                </span>
                            </template>
                        </el-table-column>
                        <el-table-column label="操作"
                            :min-width="appStore.operateMinWith"
                            fixed="right"
                            align="center">
                            <template #default="{ row }">
                                <el-button type="primary"
                                    link
                                    icon="Edit"
                                    @click="openDrawer(row)">
                                    编辑
                                </el-button>
                                <el-button type="danger"
                                    link
                                    icon="Delete"
                                    @click="handleDelete(row)">
                                    删除
                                </el-button>
                            </template>
                        </el-table-column>
                    </el-table>

                    <el-empty v-if="!operationLoading && operationData.length === 0"
                        description="暂无工序数据，请添加工序"
                        :image-size="100"
                        class="mt-8" />
                </div>
            </div>
        </div>

        <!-- 新增/编辑工序抽屉 -->
        <OperationDrawer v-model="drawerVisible"
            :data="drawerData"
            :product="selectedProduct"
            @success="handleDrawerSuccess" />
    </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ArrowRight } from '@element-plus/icons-vue'
import { useAppStore } from '@/pinia'
import OperationDrawer from './components/OperationDrawer.vue'

defineOptions({
  name: 'FfOperation'
})

const appStore = useAppStore()

// ==================== 模拟数据 ====================
// 模拟产品数据
const mockProducts = [
  {
    id: 1,
    productCode: 'P001',
    productName: '春季新款连衣裙',
    productColor: '蓝色',
    productImg:
      'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg',
    accountId: 1001,
    productOrderNum: 100,
    productFlag: '生产中'
  },
  {
    id: 2,
    productCode: 'P002',
    productName: '夏季短袖T恤',
    productColor: '白色',
    productImg:
      'https://fuss10.elemecdn.com/a/3f/3302e58f9a181d2509f3dc0fa68b0jpeg.jpeg',
    accountId: 1002,
    productOrderNum: 200,
    productFlag: '未生产'
  },
  {
    id: 3,
    productCode: 'P003',
    productName: '秋季卫衣',
    productColor: '灰色',
    productImg: '',
    accountId: 1001,
    productOrderNum: 150,
    productFlag: '已完成'
  },
  {
    id: 4,
    productCode: 'P004',
    productName: '冬季羽绒服',
    productColor: '黑色',
    productImg:
      'https://fuss10.elemecdn.com/1/34/19aa98b1fcb2781c4fba33d850549jpeg.jpeg',
    accountId: 1003,
    productOrderNum: 80,
    productFlag: '生产中'
  },
  {
    id: 5,
    productCode: 'P005',
    productName: '牛仔裤',
    productColor: '深蓝',
    productImg: '',
    accountId: 1002,
    productOrderNum: 120,
    productFlag: '未生产'
  }
]

// 模拟工序数据（按产品ID存储）
const mockOperations = {
  1: [
    // 产品ID为1的工序
    { id: 1, productId: 1, operationName: '裁剪', referencePrice: 5.5 },
    { id: 2, productId: 1, operationName: '缝纫', referencePrice: 8.0 },
    { id: 3, productId: 1, operationName: '熨烫', referencePrice: 3.0 },
    { id: 4, productId: 1, operationName: '质检', referencePrice: 2.5 }
  ],
  2: [
    { id: 5, productId: 2, operationName: '印花', referencePrice: 6.0 },
    { id: 6, productId: 2, operationName: '缝纫', referencePrice: 5.0 },
    { id: 7, productId: 2, operationName: '包装', referencePrice: 1.5 }
  ],
  3: [
    { id: 8, productId: 3, operationName: '裁剪', referencePrice: 7.0 },
    { id: 9, productId: 3, operationName: '缝纫', referencePrice: 10.0 },
    { id: 10, productId: 3, operationName: '绣花', referencePrice: 12.0 }
  ],
  4: [
    { id: 11, productId: 4, operationName: '充绒', referencePrice: 15.0 },
    { id: 12, productId: 4, operationName: '缝纫', referencePrice: 12.0 },
    { id: 13, productId: 4, operationName: '质检', referencePrice: 5.0 },
    { id: 14, productId: 4, operationName: '包装', referencePrice: 2.0 }
  ],
  5: [
    { id: 15, productId: 5, operationName: '裁剪', referencePrice: 4.0 },
    { id: 16, productId: 5, operationName: '缝纫', referencePrice: 6.0 }
  ]
}

// 用于生成新工序ID
let nextOperationId = 17

// ==================== 产品相关 ====================
const productLoading = ref(false)
const productData = ref([])
const productSearch = ref('')
const selectedProduct = ref(null)

// 筛选后的产品列表
const filteredProducts = computed(() => {
  if (!productSearch.value) {
    return productData.value
  }
  const searchLower = productSearch.value.toLowerCase()
  return productData.value.filter(
    (item) =>
      item.productName.toLowerCase().includes(searchLower) ||
      item.productCode.toLowerCase().includes(searchLower)
  )
})

// 获取生产标志标签类型
const getProductFlagType = (flag) => {
  const typeMap = {
    未生产: 'info',
    生产中: 'warning',
    已完成: 'success'
  }
  return typeMap[flag] || 'info'
}

// 产品搜索
const handleProductSearch = () => {
  // 搜索是通过计算属性实现的，这里可以添加额外逻辑
}

// 选择产品
const selectProduct = (product) => {
  selectedProduct.value = product
  getOperationData()
}

// 获取产品列表（使用模拟数据）
const getProductData = async () => {
  productLoading.value = true
  try {
    // 模拟API延迟
    await new Promise((resolve) => setTimeout(resolve, 300))
    productData.value = [...mockProducts]
  } catch (error) {
    ElMessage.error('获取产品列表失败')
  } finally {
    productLoading.value = false
  }
}

// ==================== 工序相关 ====================
const operationLoading = ref(false)
const operationData = ref([])

// 抽屉相关
const drawerVisible = ref(false)
const drawerData = ref(null)

// 获取工序数据（使用模拟数据）
const getOperationData = async () => {
  if (!selectedProduct.value) {
    operationData.value = []
    return
  }

  operationLoading.value = true
  try {
    // 模拟API延迟
    await new Promise((resolve) => setTimeout(resolve, 300))
    const productId = selectedProduct.value.id
    operationData.value = mockOperations[productId]
      ? [...mockOperations[productId]]
      : []
  } catch (error) {
    ElMessage.error('获取工序列表失败')
    operationData.value = []
  } finally {
    operationLoading.value = false
  }
}

// 打开抽屉
const openDrawer = (row = null) => {
  if (!selectedProduct.value) {
    ElMessage.warning('请先选择产品')
    return
  }
  drawerData.value = row ? { ...row } : null
  drawerVisible.value = true
}

// 抽屉提交成功
const handleDrawerSuccess = async (formData, isEdit) => {
  try {
    operationLoading.value = true
    // 模拟API延迟
    await new Promise((resolve) => setTimeout(resolve, 300))

    const productId = selectedProduct.value.id

    if (!isEdit) {
      // 新增工序
      const newOperation = {
        id: nextOperationId++,
        productId: productId,
        operationName: formData.operationName,
        referencePrice: formData.referencePrice
      }

      if (!mockOperations[productId]) {
        mockOperations[productId] = []
      }
      mockOperations[productId].push(newOperation)

      ElMessage.success('新增成功')
    } else {
      // 编辑工序
      const operationList = mockOperations[productId]
      if (operationList) {
        const index = operationList.findIndex((op) => op.id === formData.id)
        if (index !== -1) {
          operationList[index] = {
            ...operationList[index],
            operationName: formData.operationName,
            referencePrice: formData.referencePrice
          }
        }
      }

      ElMessage.success('更新成功')
    }

    drawerVisible.value = false
    getOperationData()
  } catch (error) {
    ElMessage.error('操作失败')
  } finally {
    operationLoading.value = false
  }
}

// 删除工序（使用模拟数据）
const handleDelete = (row) => {
  ElMessageBox.confirm(
    `确定要删除工序 "${row.operationName}" 吗？`,
    '删除确认',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }
  )
    .then(async () => {
      try {
        operationLoading.value = true
        // 模拟API延迟
        await new Promise((resolve) => setTimeout(resolve, 300))

        const productId = selectedProduct.value.id
        const operationList = mockOperations[productId]

        if (operationList) {
          const index = operationList.findIndex((op) => op.id === row.id)
          if (index !== -1) {
            operationList.splice(index, 1)
          }
        }

        ElMessage.success('删除成功')
        getOperationData()
      } catch (error) {
        ElMessage.error('删除失败')
      } finally {
        operationLoading.value = false
      }
    })
    .catch(() => {
      // 用户取消删除
    })
}

// ==================== 生命周期 ====================
onMounted(() => {
  getProductData()
})
</script>
