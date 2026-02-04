<template>
  <FfDrawer
    v-model="visible"
    :title="isEdit ? '编辑分配' : '新增分配'"
    @cancel="handleCancel"
    @confirm="handleConfirm"
  >
    <el-form
      ref="formRef"
      :model="formData"
      :rules="formRules"
      label-width="100px"
    >
      <el-form-item label="员工" prop="employeeId">
        <el-select
          v-model="formData.employeeId"
          placeholder="请选择员工"
          clearable
          filterable
        >
          <el-option
            v-for="employee in employeeList"
            :key="employee.ID"
            :label="employee.employeeName"
            :value="employee.ID"
          />
        </el-select>
      </el-form-item>

      <el-form-item label="单价" prop="unitPrice">
        <el-input-number
          v-model="formData.unitPrice"
          :min="0"
          :precision="2"
          :step="0.1"
          controls-position="right"
          style="width: 100%"
          placeholder="请输入单价"
        />
        <div class="text-gray-400 text-sm mt-2">单位：元</div>
      </el-form-item>

      <el-form-item label="分配数量" prop="assignmentQuantity">
        <el-input-number
          v-model="formData.assignmentQuantity"
          :min="1"
          :step="1"
          controls-position="right"
          style="width: 100%"
          placeholder="请输入分配数量"
        />
      </el-form-item>
    </el-form>
  </FfDrawer>
</template>

<script setup>
import { ref, computed, watch, onMounted } from 'vue'
import FfDrawer from '@/components/ffDrawer/index.vue'
import { getEmployeeList } from '@/api/ffEmployee'

defineOptions({
  name: 'AssignmentDrawer'
})

const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false
  },
  data: {
    type: Object,
    default: () => null
  },
  operation: {
    type: Object,
    default: () => null
  }
})

const emit = defineEmits(['update:modelValue', 'success'])

// 默认表单数据
const getDefaultFormData = () => ({
  ID: null,
  employeeId: null,
  unitPrice: null,
  assignmentQuantity: 1
})

const formRef = ref(null)
const formData = ref(getDefaultFormData())
const employeeList = ref([])

// 是否为编辑模式
const isEdit = computed(() => !!props.data?.ID)

// 双向绑定 visible
const visible = computed({
  get: () => props.modelValue,
  set: (val) => emit('update:modelValue', val)
})

// 获取员工列表
const getEmployees = async () => {
  try {
    const res = await getEmployeeList({ page: 1, pageSize: 100 })
    if (res.code === 0) {
      employeeList.value = res.data.list || []
    }
  } catch (error) {
    console.error('获取员工列表失败', error)
  }
}

// 监听 data 变化，初始化表单数据
watch(
  () => props.data,
  (val) => {
    if (val) {
      formData.value = { ...val }
    } else {
      formData.value = getDefaultFormData()
    }
  },
  { immediate: true }
)

// 表单验证规则
const formRules = {
  employeeId: [
    { required: true, message: '请选择员工', trigger: 'change' }
  ],
  unitPrice: [
    { required: true, message: '请输入单价', trigger: 'blur' },
    { type: 'number', min: 0, message: '单价不能小于0', trigger: 'blur' }
  ],
  assignmentQuantity: [
    { required: true, message: '请输入分配数量', trigger: 'blur' },
    { type: 'number', min: 1, message: '分配数量不能小于1', trigger: 'blur' }
  ]
}

// 取消
const handleCancel = () => {
  formRef.value?.resetFields()
  formData.value = getDefaultFormData()
}

// 确认提交
const handleConfirm = async () => {
  const valid = await formRef.value?.validate()
  if (!valid) return

  emit('success', { ...formData.value }, isEdit.value)
}

// 关闭后重置表单
watch(visible, (val) => {
  if (!val) {
    formRef.value?.resetFields()
    formData.value = getDefaultFormData()
  }
})

// 初始化
onMounted(() => {
  getEmployees()
})
</script>
