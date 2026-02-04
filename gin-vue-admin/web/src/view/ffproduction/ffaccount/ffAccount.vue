<template>
    <div class="ff-custom-panel">
        <!-- 搜索区域 -->
        <div class="gva-search-box">
            <el-form ref="searchFormRef"
                :inline="true"
                :model="searchInfo">
                <el-form-item label="客户名">
                    <el-input v-model="searchInfo.accountName"
                        placeholder="请输入客户名"
                        clearable />
                </el-form-item>
                <el-form-item label="联系方式">
                    <el-input v-model="searchInfo.phone"
                        placeholder="请输入联系方式"
                        clearable />
                </el-form-item>
                <el-form-item label="联系地址">
                    <el-input v-model="searchInfo.address"
                        placeholder="请输入联系地址"
                        clearable />
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
                    @click="openDrawer()">新增客户</el-button>
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
                <el-table-column prop="accountName"
                    label="客户名"
                    min-width="150"
                    align="center" />
                <el-table-column prop="phone"
                    label="联系方式"
                    min-width="130"
                    align="center" />
                <el-table-column prop="address"
                    label="联系地址"
                    min-width="200"
                    align="center"
                    show-overflow-tooltip />
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
        <AccountDrawer v-model="drawerVisible"
            :data="drawerData"
            @success="handleDrawerSuccess" />
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useAppStore } from '@/pinia'
import AccountDrawer from './components/AccountDrawer.vue'
import { getAccountList, createAccount, updateAccount, deleteAccount } from '@/api/ffAccount'

defineOptions({
  name: 'FfAccount'
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
  accountName: '',
  phone: '',
  address: ''
})

// 抽屉相关
const drawerVisible = ref(false)
const drawerData = ref(null)

// ==================== 业务逻辑 ====================
const getTableData = async () => {
  loading.value = true
  try {
    const params = {
      page: page.value,
      pageSize: pageSize.value,
      ...searchInfo.value
    }
    const res = await getAccountList(params)
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
    accountName: '',
    phone: '',
    address: ''
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
      res = await updateAccount(formData)
    } else {
      res = await createAccount(formData)
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
  ElMessageBox.confirm(`确定要删除客户 "${row.accountName}" 吗？`, '删除确认', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
    .then(async () => {
      try {
        loading.value = true
        const res = await deleteAccount({ ID: row.ID })
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
    .catch(() => {})
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
