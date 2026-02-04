<template>
    <div class="ff-custom-panel">
        <!-- 搜索区域 -->
        <div class="gva-search-box">
            <el-form ref="searchFormRef"
                :inline="true"
                :model="searchInfo">
                <el-form-item label="员工名">
                    <el-input v-model="searchInfo.employeeName"
                        placeholder="请输入员工名"
                        clearable />
                </el-form-item>
                <el-form-item label="手机号">
                    <el-input v-model="searchInfo.phone"
                        placeholder="请输入手机号"
                        clearable />
                </el-form-item>
                <el-form-item label="性别">
                    <FfDicSelect v-model="searchInfo.sex"
                        dicType="gender"
                        placeholder="请选择性别"
                        clearable />
                </el-form-item>
                <el-form-item label="角色">
                    <FfDicSelect v-model="searchInfo.authorityIds"
                        dicType="employeeAuthority"
                        placeholder="请选择角色"
                        clearable />
                </el-form-item>
                <el-form-item label="状态">
                    <FfDicSelect v-model="searchInfo.employeeStatus"
                        dicType="employeeStatus"
                        placeholder="请选择状态"
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
                    @click="openDrawer()">新增员工</el-button>
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
                <el-table-column prop="employeeName"
                    label="员工名"
                    min-width="120"
                    align="center" />
                <el-table-column prop="phone"
                    label="手机号"
                    min-width="130"
                    align="center" />
                <el-table-column prop="age"
                    label="年龄"
                    width="80"
                    align="center" />
                <el-table-column prop="sex"
                    label="性别"
                    width="80"
                    align="center">
                    <template #default="{ row }">
                        <FfDicTag dicType="gender" :dicValue="row.sex" />
                    </template>
                </el-table-column>
                <el-table-column prop="authorityIds"
                    label="角色"
                    width="80"
                    align="center">
                    <template #default="{ row }">
                        <FfDicTag dicType="employeeAuthority" :dicValue="row.authorityIds" />
                    </template>
                </el-table-column>
                <el-table-column prop="employeeStatus"
                    label="状态"
                    width="80"
                    align="center">
                    <template #default="{ row }">
                        <FfDicTag dicType="employeeStatus" :dicValue="row.employeeStatus" />
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
        <EmployeeDrawer v-model="drawerVisible"
            :data="drawerData"
            @success="handleDrawerSuccess" />
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useAppStore } from '@/pinia'
import EmployeeDrawer from './components/EmployeeDrawer.vue'
import FfDicTag from '@/components/ffDicTag/index.vue'
import FfDicSelect from '@/components/ffDicSelect/index.vue'
import { getEmployeeList, createEmployee, updateEmployee, deleteEmployee } from '@/api/ffEmployee'

defineOptions({
  name: 'FfEmployee'
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
  employeeName: '',
  phone: '',
  sex: null,
  authorityIds: null,
  employeeStatus: null
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
    const res = await getEmployeeList(params)
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
    employeeName: '',
    phone: '',
    sex: null,
    authorityIds: null,
    employeeStatus: null
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
      res = await updateEmployee(formData)
    } else {
      res = await createEmployee(formData)
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
  ElMessageBox.confirm(
    `确定要删除员工 "${row.employeeName}" 吗？`,
    '删除确认',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }
  )
    .then(async () => {
      try {
        loading.value = true
        const res = await deleteEmployee({ ID: row.ID })
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
})
</script>

<style lang="scss" scoped>
.gva-search-box {
  :deep(.el-form-item) {
    margin-bottom: 12px;
  }
}
</style>
