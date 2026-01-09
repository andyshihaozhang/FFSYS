<template>
  <FfDrawer
    v-model="visible"
    :title="isEdit ? '编辑工序' : '新增工序'"
    @cancel="handleCancel"
    @confirm="handleConfirm"
  >
    <el-form
      ref="formRef"
      :model="formData"
      :rules="formRules"
      label-width="100px"
    >
      <el-form-item label="所属产品">
        <el-input
          :value="productInfo"
          disabled
        />
      </el-form-item>

      <el-form-item label="工序名称" prop="operationName">
        <el-input
          v-model="formData.operationName"
          placeholder="请输入工序名称"
          clearable
          maxlength="100"
          show-word-limit
        />
      </el-form-item>

      <el-form-item label="参考单价" prop="referencePrice">
        <el-input-number
          v-model="formData.referencePrice"
          :min="0"
          :precision="2"
          :step="0.1"
          controls-position="right"
          style="width: 100%"
          placeholder="请输入参考单价"
        />
        <div class="text-gray-400 text-sm mt-2">单位：元</div>
      </el-form-item>
    </el-form>
  </FfDrawer>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import FfDrawer from '@/components/ffDrawer/index.vue'

defineOptions({
  name: 'OperationDrawer'
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
  product: {
    type: Object,
    default: () => null
  }
})

const emit = defineEmits(['update:modelValue', 'success'])

// 默认表单数据
const getDefaultFormData = () => ({
  id: null,
  productId: null,
  operationName: '',
  referencePrice: null
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

// 产品信息展示
const productInfo = computed(() => {
  if (props.product) {
    return `${props.product.productName} (${props.product.productCode})`
  }
  return ''
})

// 监听 data 变化，初始化表单数据
watch(
  () => props.data,
  (val) => {
    if (val) {
      formData.value = { ...val }
    } else {
      formData.value = getDefaultFormData()
      // 新增时设置产品ID
      if (props.product) {
        formData.value.productId = props.product.id
      }
    }
  },
  { immediate: true }
)

// 监听 product 变化，更新 productId
watch(
  () => props.product,
  (val) => {
    if (val && !props.data) {
      formData.value.productId = val.id
    }
  }
)

// 表单验证规则
const formRules = {
  operationName: [
    { required: true, message: '请输入工序名称', trigger: 'blur' },
    { min: 2, max: 100, message: '长度在 2 到 100 个字符', trigger: 'blur' }
  ],
  referencePrice: [
    { required: true, message: '请输入参考单价', trigger: 'blur' },
    { type: 'number', min: 0, message: '参考单价不能小于0', trigger: 'blur' }
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
