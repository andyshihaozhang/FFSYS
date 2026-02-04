<template>
    <div class="ff-custom-panel">
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
                    <FfDicSelect v-model="searchInfo.productFlag"
                        dicType="productFlag"
                        placeholder="请选择生产标志"
                        clearable />
                </el-form-item>
                <el-form-item label="客户">
                    <el-select v-model="searchInfo.accountId"
                        placeholder="请选择客户"
                        clearable>
                        <el-option v-for="item in accountOptions"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value" />
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
                row-key="ID"
                border
                stripe>
                <el-table-column prop="ID"
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
                    width="100"
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
                        <CustomPic v-if="row.productImg" pic-type="file" :pic-src="row.productImg" preview />
                        <span v-else
                            class="text-gray-400">暂无图片</span>
                    </template>
                </el-table-column>
                <el-table-column prop="accountName"
                    label="客户名称"
                    min-width="120"
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
                        <FfDicTag dicType="productFlag" :dicValue="row.productFlag" />
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
                <el-pagination :current-page="page"
                    :page-size="pageSize"
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
            :accountOptions="accountOptions"
            @success="handleDrawerSuccess" />
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useAppStore } from '@/pinia'
import ProductDrawer from './components/ProductDrawer.vue'
import FfDicTag from '@/components/ffDicTag/index.vue'
import FfDicSelect from '@/components/ffDicSelect/index.vue'
import { getProductList, createProduct, updateProduct, deleteProduct } from '@/api/ffProduct'
import { getAccountOptions } from '@/api/ffAccount'

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
  productFlag: null,
  accountId: null
})

// 抽屉相关
const drawerVisible = ref(false)
const drawerData = ref(null)

// 客户选项
const accountOptions = ref([])

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
    } else {
      ElMessage.error(res.msg || '获取数据失败')
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
    productFlag: null,
    accountId: null
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
    } else {
      ElMessage.error(res.msg || '操作失败')
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
        const res = await deleteProduct({ ID: row.ID })
        if (res.code === 0) {
          ElMessage.success(res.msg)
          getTableData()
        } else {
          ElMessage.error(res.msg || '删除失败')
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
  loadAccountOptions()
})

const loadAccountOptions = async () => {
  try {
    const res = await getAccountOptions()
    if (res.code === 0) {
      accountOptions.value = res.data
    }
  } catch (error) {
    console.error('获取客户选项失败', error)
  }
}
</script>

<style lang="scss" scoped>
.gva-search-box {
  :deep(.el-form-item) {
    margin-bottom: 12px;
  }
}
</style>
