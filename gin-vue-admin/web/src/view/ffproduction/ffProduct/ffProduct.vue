<template>
    <div>
        <!-- 搜索区域 -->
        <div class="gva-search-box">
            <el-form ref="searchFormRef"
                :inline="true"
                :model="searchInfo">
                <el-form-item label="产品款号">
                    <el-input v-model="searchInfo.productCode"
                        placeholder="请输入产品款号"
                        clearable />
                </el-form-item>
                <el-form-item label="产品名称">
                    <el-input v-model="searchInfo.productName"
                        placeholder="请输入产品名称"
                        clearable />
                </el-form-item>
                <el-form-item label="产品颜色">
                    <el-input v-model="searchInfo.productColor"
                        placeholder="请输入产品颜色"
                        clearable />
                </el-form-item>
                <el-form-item label="生产标志">
                    <el-select v-model="searchInfo.productFlag"
                        placeholder="请选择生产标志"
                        clearable>
                        <el-option label="未生产"
                            value="未生产" />
                        <el-option label="生产中"
                            value="生产中" />
                        <el-option label="已完成"
                            value="已完成" />
                    </el-select>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary"
                        icon="Search"
                        @click="onSubmit">查询</el-button>
                    <el-button icon="Refresh"
                        @click="onReset">重置</el-button>
                </el-form-item>
            </el-form>
        </div>

        <!-- 表格区域 -->
        <div class="gva-table-box">
            <div class="gva-btn-list">
                <el-button type="primary"
                    icon="Plus"
                    @click="openDrawer()">新增产品</el-button>
            </div>

            <el-table v-loading="loading"
                :data="tableData"
                row-key="id"
                border
                stripe>
                <el-table-column prop="id"
                    label="序号"
                    width="80"
                    align="center" />
                <el-table-column prop="productCode"
                    label="产品款号"
                    min-width="120"
                    align="center" />
                <el-table-column prop="productName"
                    label="产品名称"
                    min-width="150"
                    align="center" />
                <el-table-column prop="productColor"
                    label="产品颜色"
                    width="80"
                    align="center">
                    <template #default="{ row }">
                        <el-tag size="small">{{ row.productColor }}</el-tag>
                    </template>
                </el-table-column>
                <el-table-column prop="productImg"
                    label="产品图片"
                    min-width="100"
                    align="center">
                    <template #default="{ row }">
                        <el-image v-if="row.productImg"
                            :src="row.productImg"
                            :preview-src-list="[row.productImg]"
                            fit="cover"
                            style="width: 50px; height: 50px; cursor: pointer"
                            preview-teleported />
                        <span v-else
                            class="text-gray-400">暂无图片</span>
                    </template>
                </el-table-column>
                <el-table-column prop="accountId"
                    label="客户ID"
                    min-width="100"
                    align="center" />
                <el-table-column prop="productOrderNum"
                    label="下单数量"
                    width="100"
                    align="center" />
                <el-table-column prop="productFlag"
                    label="生产标志"
                    width="100"
                    align="center">
                    <template #default="{ row }">
                        <el-tag :type="getProductFlagType(row.productFlag)"
                            size="small">
                            {{ row.productFlag }}
                        </el-tag>
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

            <!-- 分页 -->
            <div class="gva-pagination">
                <el-pagination v-model:current-page="page"
                    v-model:page-size="pageSize"
                    :page-sizes="[10, 20, 50, 100]"
                    :total="total"
                    layout="total, sizes, prev, pager, next, jumper"
                    @size-change="handleSizeChange"
                    @current-change="handleCurrentChange" />
            </div>
        </div>

        <!-- 新增/编辑抽屉 -->
        <ProductDrawer v-model="drawerVisible"
            :data="drawerData"
            @success="handleDrawerSuccess" />
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useAppStore } from '@/pinia'
import ProductDrawer from './components/ProductDrawer.vue'

defineOptions({
  name: 'FfProduct'
})

const appStore = useAppStore()

// ==================== 数据定义 ====================
const loading = ref(false)
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const tableData = ref([])

// 搜索表单
const searchFormRef = ref(null)
const searchInfo = ref({
  productCode: '',
  productName: '',
  productColor: '',
  productFlag: ''
})

// 抽屉相关
const drawerVisible = ref(false)
const drawerData = ref(null)

// ==================== 工具函数 ====================
const getProductFlagType = (flag) => {
  const typeMap = {
    未生产: 'info',
    生产中: 'warning',
    已完成: 'success'
  }
  return typeMap[flag] || 'info'
}

// ==================== API 调用 (模拟数据) ====================
const getProductList = async (params) => {
  await new Promise((resolve) => setTimeout(resolve, 300))

  const mockData = [
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

  let filteredData = mockData.filter((item) => {
    if (params.productCode && !item.productCode.includes(params.productCode))
      return false
    if (params.productName && !item.productName.includes(params.productName))
      return false
    if (params.productColor && !item.productColor.includes(params.productColor))
      return false
    if (params.productFlag && item.productFlag !== params.productFlag)
      return false
    return true
  })

  const start = (params.page - 1) * params.pageSize
  const end = start + params.pageSize

  return {
    code: 0,
    data: {
      list: filteredData.slice(start, end),
      total: filteredData.length,
      page: params.page,
      pageSize: params.pageSize
    },
    msg: '查询成功'
  }
}

const createProduct = async (_data) => {
  await new Promise((resolve) => setTimeout(resolve, 300))
  return { code: 0, msg: '创建成功' }
}

const updateProduct = async (_data) => {
  await new Promise((resolve) => setTimeout(resolve, 300))
  return { code: 0, msg: '更新成功' }
}

const deleteProduct = async (_id) => {
  await new Promise((resolve) => setTimeout(resolve, 300))
  return { code: 0, msg: '删除成功' }
}

// ==================== 业务逻辑 ====================
const getTableData = async () => {
  loading.value = true
  try {
    const params = {
      page: page.value,
      pageSize: pageSize.value,
      ...searchInfo.value
    }
    const res = await getProductList(params)
    if (res.code === 0) {
      tableData.value = res.data.list
      total.value = res.data.total
      page.value = res.data.page
      pageSize.value = res.data.pageSize
    }
  } catch (error) {
    ElMessage.error('获取数据失败')
  } finally {
    loading.value = false
  }
}

const onSubmit = () => {
  page.value = 1
  getTableData()
}

const onReset = () => {
  searchFormRef.value?.resetFields()
  searchInfo.value = {
    productCode: '',
    productName: '',
    productColor: '',
    productFlag: ''
  }
  page.value = 1
  getTableData()
}

const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 打开抽屉
const openDrawer = (row = null) => {
  drawerData.value = row ? { ...row } : null
  drawerVisible.value = true
}

// 抽屉提交成功
const handleDrawerSuccess = async (formData, isEdit) => {
  try {
    loading.value = true
    let res
    if (isEdit) {
      res = await updateProduct(formData)
    } else {
      res = await createProduct(formData)
    }

    if (res.code === 0) {
      ElMessage.success(res.msg)
      drawerVisible.value = false
      getTableData()
    }
  } catch (error) {
    ElMessage.error('操作失败')
  } finally {
    loading.value = false
  }
}

// 删除
const handleDelete = (row) => {
  ElMessageBox.confirm(`确定要删除产品 "${row.productName}" 吗？`, '删除确认', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
    .then(async () => {
      try {
        loading.value = true
        const res = await deleteProduct(row.id)
        if (res.code === 0) {
          ElMessage.success(res.msg)
          getTableData()
        }
      } catch (error) {
        ElMessage.error('删除失败')
      } finally {
        loading.value = false
      }
    })
    .catch(() => {
      // 用户取消删除
    })
}

// ==================== 生命周期 ====================
onMounted(() => {
  getTableData()
})
</script>

<style lang="scss" scoped>
.gva-search-box {
  :deep(.el-form-item) {
    margin-bottom: 12px;
  }
}
</style>
