<template>
  <FfDrawer
    v-model="visible"
    :title="isEdit ? '编辑客户' : '新增客户'"
    @cancel="handleCancel"
    @confirm="handleConfirm"
  >
    <el-form
      ref="formRef"
      :model="formData"
      :rules="formRules"
      label-width="90px"
    >
      <el-form-item label="客户名" prop="accountName">
        <el-input
          v-model="formData.accountName"
          placeholder="请输入客户名"
          clearable
        />
      </el-form-item>

      <el-form-item label="联系方式" prop="phone">
        <el-input
          v-model="formData.phone"
          placeholder="请输入联系方式"
          clearable
          maxlength="11"
        />
      </el-form-item>

      <el-form-item label="联系地址" prop="address">
        <el-input
          v-model="formData.address"
          type="textarea"
          :rows="3"
          placeholder="请输入联系地址"
          clearable
          maxlength="200"
          show-word-limit
        />
      </el-form-item>
    </el-form>
  </FfDrawer>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import FfDrawer from '@/components/ffDrawer/index.vue'

defineOptions({
  name: 'AccountDrawer'
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
  accountName: '',
  phone: '',
  address: ''
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
  accountName: [
    { required: true, message: '请输入客户名', trigger: 'blur' },
    { min: 2, max: 50, message: '长度在 2 到 50 个字符', trigger: 'blur' }
  ],
  phone: [
    { required: true, message: '请输入联系方式', trigger: 'blur' },
    {
      pattern: /^1[3-9]\d{9}$/,
      message: '请输入正确的手机号',
      trigger: 'blur'
    }
  ],
  address: [
    { required: true, message: '请输入联系地址', trigger: 'blur' },
    { min: 5, max: 200, message: '长度在 5 到 200 个字符', trigger: 'blur' }
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
</script>
