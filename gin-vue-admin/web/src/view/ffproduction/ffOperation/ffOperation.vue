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
                            :key="product.ID"
                            class="product-item flex items-center justify-between p-3 mb-2 bg-gray-50 rounded-md cursor-pointer transition-all hover:bg-gray-100 hover:shadow-md hover:-translate-y-0.5"
                            :class="{
                'border-2 border-blue-500 shadow-lg bg-blue-50 -translate-y-0.5': selectedProduct?.ID === product.ID,
                'border-2 border-transparent': selectedProduct?.ID !== product.ID
              }"
                            @click="selectProduct(product)">
                            <div class="flex-1 min-w-0">
                                <div class="text-sm font-medium mb-1 truncate">
                                    {{ product.productName }}
                                </div>
                                <div class="text-xs text-gray-500 mb-1.5">款号: {{ product.productCode }}</div>
                                <div class="flex gap-1">
                                    <el-tag size="small">{{ product.productColor }}</el-tag>
                                    <FfDicTag dicType="productFlag" :dicValue="product.productFlag" />
                                </div>
                            </div>
                            <el-icon class="text-base text-gray-400 transition-all"
                                :class="{ 'text-blue-500': selectedProduct?.ID === product.ID }">
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
                        row-key="ID"
                        border
                        stripe
                        height="calc(100vh - 240px)">
                        <el-table-column prop="ID"
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
import FfDicTag from '@/components/ffDicTag/index.vue'
import { getProductList } from '@/api/ffProduct'
import { getOperationList, createOperation, updateOperation, deleteOperation } from '@/api/ffOperation'

defineOptions({
  name: 'FfOperation'
})

const appStore = useAppStore()

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

// 产品搜索
const handleProductSearch = () => {
  // 搜索是通过计算属性实现的，这里可以添加额外逻辑
}

// 选择产品
const selectProduct = (product) => {
  selectedProduct.value = product
  getOperationData()
}

// 获取产品列表（使用真实API）
const getProductData = async () => {
  productLoading.value = true
  try {
    const res = await getProductList({ productFlag: '1' ,page: 1, pageSize: 100 })
    if (res.code === 0) {
      productData.value = res.data.list
    } else {
      ElMessage.error('获取产品列表失败')
    }
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

// 获取工序数据（使用真实API）
const getOperationData = async () => {
  if (!selectedProduct.value) {
    operationData.value = []
    return
  }

  operationLoading.value = true
  try {
    const res = await getOperationList({
      productId: selectedProduct.value.ID,
      page: 1,
      pageSize: 100
    })
    if (res.code === 0) {
      operationData.value = res.data.list || []
    } else {
      ElMessage.error('获取工序列表失败')
      operationData.value = []
    }
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
    const data = {
      ...formData,
      productId: selectedProduct.value.ID
    }

    if (!isEdit) {
      delete data.ID
    }

    let res
    if (!isEdit) {
      res = await createOperation(data)
    } else {
      res = await updateOperation(data)
    }

    if (res.code === 0) {
      ElMessage.success(res.msg)
      drawerVisible.value = false
      getOperationData()
    } else {
      ElMessage.error(res.msg || '操作失败')
    }
  } catch (error) {
    ElMessage.error('操作失败')
  } finally {
    operationLoading.value = false
  }
}

// 删除工序（使用真实API）
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
        const res = await deleteOperation({ ID: row.ID })
        if (res.code === 0) {
          ElMessage.success(res.msg)
          getOperationData()
        } else {
          ElMessage.error(res.msg || '删除失败')
        }
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
