<template>
    <div>
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
                    <el-select v-model="searchInfo.sex"
                        placeholder="请选择性别"
                        clearable>
                        <el-option label="男"
                            value="男" />
                        <el-option label="女"
                            value="女" />
                    </el-select>
                </el-form-item>
                <el-form-item label="角色">
                    <el-select v-model="searchInfo.authorityIds"
                        placeholder="请选择角色"
                        clearable>
                        <el-option label="组长"
                            value="组长" />
                        <el-option label="员工"
                            value="员工" />
                    </el-select>
                </el-form-item>
                <el-form-item label="状态">
                    <el-select v-model="searchInfo.employeeStatus"
                        placeholder="请选择状态"
                        clearable>
                        <el-option label="在职"
                            value="在职" />
                        <el-option label="离职"
                            value="离职" />
                        <el-option label="休假"
                            value="休假" />
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
                    @click="openDrawer()">新增员工</el-button>
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
                        <el-tag :type="row.sex === '男' ? 'primary' : 'danger'"
                            size="small">
                            {{ row.sex }}
                        </el-tag>
                    </template>
                </el-table-column>
                <el-table-column prop="authorityIds"
                    label="角色"
                    width="80"
                    align="center">
                    <template #default="{ row }">
                        <el-tag :type="row.authorityIds === '组长' ? 'warning' : 'info'"
                            size="small">
                            {{ row.authorityIds }}
                        </el-tag>
                    </template>
                </el-table-column>
                <el-table-column prop="employeeStatus"
                    label="状态"
                    width="80"
                    align="center">
                    <template #default="{ row }">
                        <el-tag :type="getStatusType(row.employeeStatus)"
                            size="small">
                            {{ row.employeeStatus }}
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
  sex: '',
  authorityIds: '',
  employeeStatus: ''
})

// 抽屉相关
const drawerVisible = ref(false)
const drawerData = ref(null)

// ==================== 工具函数 ====================
const getStatusType = (status) => {
  const typeMap = {
    在职: 'success',
    离职: 'info',
    休假: 'warning'
  }
  return typeMap[status] || 'info'
}

// ==================== API 调用 (模拟数据) ====================
const getEmployeeList = async (params) => {
  await new Promise((resolve) => setTimeout(resolve, 300))

  const mockData = [
    {
      id: 1,
      employeeName: '张三',
      phone: '13800138001',
      age: 28,
      sex: '男',
      authorityIds: '组长',
      employeeStatus: '在职'
    },
    {
      id: 2,
      employeeName: '李四',
      phone: '13800138002',
      age: 25,
      sex: '男',
      authorityIds: '员工',
      employeeStatus: '在职'
    },
    {
      id: 3,
      employeeName: '王五',
      phone: '13800138003',
      age: 30,
      sex: '女',
      authorityIds: '员工',
      employeeStatus: '休假'
    },
    {
      id: 4,
      employeeName: '赵六',
      phone: '13800138004',
      age: 32,
      sex: '男',
      authorityIds: '组长',
      employeeStatus: '在职'
    },
    {
      id: 5,
      employeeName: '孙七',
      phone: '13800138005',
      age: 27,
      sex: '女',
      authorityIds: '员工',
      employeeStatus: '离职'
    }
  ]

  let filteredData = mockData.filter((item) => {
    if (params.employeeName && !item.employeeName.includes(params.employeeName))
      return false
    if (params.phone && !item.phone.includes(params.phone)) return false
    if (params.sex && item.sex !== params.sex) return false
    if (params.authorityIds && item.authorityIds !== params.authorityIds)
      return false
    if (params.employeeStatus && item.employeeStatus !== params.employeeStatus)
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

const createEmployee = async (_data) => {
  await new Promise((resolve) => setTimeout(resolve, 300))
  return { code: 0, msg: '创建成功' }
}

const updateEmployee = async (_data) => {
  await new Promise((resolve) => setTimeout(resolve, 300))
  return { code: 0, msg: '更新成功' }
}

const deleteEmployee = async (_id) => {
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
    sex: '',
    authorityIds: '',
    employeeStatus: ''
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
        const res = await deleteEmployee(row.id)
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
