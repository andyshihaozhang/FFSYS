<template>
    <FfDrawer v-model="visible"
        :title="isEdit ? '编辑员工' : '新增员工'"
        @cancel="handleCancel"
        @confirm="handleConfirm">
        <el-form ref="formRef"
            :model="formData"
            :rules="formRules"
            label-width="80px">
            <el-form-item label="员工名"
                prop="employeeName">
                <el-input v-model="formData.employeeName"
                    placeholder="请输入员工名"
                    clearable />
            </el-form-item>

            <el-form-item label="手机号"
                prop="phone">
                <el-input v-model="formData.phone"
                    placeholder="请输入手机号"
                    clearable
                    maxlength="11" />
            </el-form-item>

            <el-form-item label="年龄"
                prop="age">
                <el-input-number v-model="formData.age"
                    :min="18"
                    :max="65"
                    controls-position="right"
                    style="width: 100%" />
            </el-form-item>

            <el-form-item label="性别"
                prop="sex">
                <el-radio-group v-model="formData.sex">
                    <el-radio label="男">男</el-radio>
                    <el-radio label="女">女</el-radio>
                </el-radio-group>
            </el-form-item>

            <el-form-item label="角色"
                prop="authorityIds">
                <el-select v-model="formData.authorityIds"
                    placeholder="请选择角色"
                    style="width: 100%">
                    <el-option label="组长"
                        value="组长" />
                    <el-option label="员工"
                        value="员工" />
                </el-select>
            </el-form-item>

            <el-form-item label="状态"
                prop="employeeStatus">
                <el-select v-model="formData.employeeStatus"
                    placeholder="请选择状态"
                    style="width: 100%">
                    <el-option label="在职"
                        value="在职" />
                    <el-option label="离职"
                        value="离职" />
                    <el-option label="休假"
                        value="休假" />
                </el-select>
            </el-form-item>
        </el-form>
    </FfDrawer>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import FfDrawer from '@/components/ffDrawer/index.vue'

defineOptions({
  name: 'EmployeeDrawer'
})

const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false
  },
  data: {
    type: Object,
    default: () => null
  }
})

const emit = defineEmits(['update:modelValue', 'success'])

// 默认表单数据
const getDefaultFormData = () => ({
  id: null,
  employeeName: '',
  phone: '',
  age: null,
  sex: '男',
  authorityIds: '员工',
  employeeStatus: '在职'
})

const formRef = ref(null)
const formData = ref(getDefaultFormData())

// 是否为编辑模式
const isEdit = computed(() => !!props.data?.id)

// 双向绑定 visible
const visible = computed({
  get: () => props.modelValue,
  set: (val) => emit('update:modelValue', val)
})

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
  employeeName: [
    { required: true, message: '请输入员工名', trigger: 'blur' },
    { min: 2, max: 20, message: '长度在 2 到 20 个字符', trigger: 'blur' }
  ],
  phone: [
    { required: true, message: '请输入手机号', trigger: 'blur' },
    {
      pattern: /^1[3-9]\d{9}$/,
      message: '请输入正确的手机号',
      trigger: 'blur'
    }
  ],
  age: [
    { required: true, message: '请输入年龄', trigger: 'blur' },
    {
      type: 'number',
      min: 18,
      max: 65,
      message: '年龄必须在18-65岁之间',
      trigger: 'blur'
    }
  ],
  sex: [{ required: true, message: '请选择性别', trigger: 'change' }],
  authorityIds: [{ required: true, message: '请选择角色', trigger: 'change' }],
  employeeStatus: [{ required: true, message: '请选择状态', trigger: 'change' }]
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
</script>
