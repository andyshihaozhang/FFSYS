<template>
    <FfDrawer v-model="visible"
        :title="isEdit ? '编辑产品' : '新增产品'"
        @cancel="handleCancel"
        @confirm="handleConfirm">
        <el-form ref="formRef"
            :model="formData"
            :rules="formRules"
            label-width="100px">
            <el-form-item label="产品款号"
                prop="productCode">
                <el-input v-model="formData.productCode"
                    placeholder="请输入产品款号"
                    clearable />
            </el-form-item>

            <el-form-item label="产品名称"
                prop="productName">
                <el-input v-model="formData.productName"
                    placeholder="请输入产品名称"
                    clearable />
            </el-form-item>

            <el-form-item label="产品颜色"
                prop="productColor">
                <el-input v-model="formData.productColor"
                    placeholder="请输入产品颜色"
                    clearable />
            </el-form-item>

            <el-form-item label="产品图片"
                prop="productImg">
                <el-upload class="product-uploader"
                    :action="uploadUrl"
                    :headers="uploadHeaders"
                    :show-file-list="false"
                    :on-success="handleUploadSuccess"
                    :before-upload="beforeUpload"
                    accept="image/*">
                    <el-image v-if="formData.productImg"
                        :src="formData.productImg"
                        fit="cover"
                        class="product-image" />
                    <el-icon v-else
                        class="product-uploader-icon">
                        <Plus />
                    </el-icon>
                </el-upload>
                <div class="text-gray-400 text-sm mt-2">支持jpg、png格式，大小不超过2MB</div>
            </el-form-item>

            <el-form-item label="客户"
                prop="accountId">
                <el-select v-model="formData.accountId"
                    placeholder="请选择客户"
                    style="width: 100%">
                    <el-option v-for="item in props.accountOptions"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value" />
                </el-select>
            </el-form-item>

            <el-form-item label="下单数量"
                prop="productOrderNum">
                <el-input-number v-model="formData.productOrderNum"
                    :min="1"
                    controls-position="right"
                    style="width: 100%"
                    placeholder="请输入下单数量" />
            </el-form-item>

            <el-form-item label="生产标志"
                prop="productFlag">
                <FfDicSelect v-model="formData.productFlag"
                    dicType="productFlag"
                    placeholder="请选择生产标志"
                    style="width: 100%" />
            </el-form-item>
        </el-form>
    </FfDrawer>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { Plus } from '@element-plus/icons-vue'
import { useUserStore } from '@/pinia'
import FfDrawer from '@/components/ffDrawer/index.vue'
import FfDicSelect from '@/components/ffDicSelect/index.vue'

defineOptions({
  name: 'ProductDrawer'
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
  accountOptions: {
    type: Array,
    default: () => []
  }
})

const emit = defineEmits(['update:modelValue', 'success'])

const userStore = useUserStore()

// 默认表单数据
const getDefaultFormData = () => ({
  ID: null,
  productCode: '',
  productName: '',
  productColor: '',
  productImg: '',
  accountId: null,
  productOrderNum: null,
  productFlag: 1
})

const formRef = ref(null)
const formData = ref(getDefaultFormData())

// 是否为编辑模式
const isEdit = computed(() => !!props.data?.ID)

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

// 上传配置
const uploadUrl = computed(() => {
  return import.meta.env.VITE_BASE_API + '/fileUploadAndDownload/upload'
})

const uploadHeaders = computed(() => {
  return {
    'x-token': userStore.token
  }
})

// 表单验证规则
const formRules = {
  productCode: [
    { required: true, message: '请输入产品款号', trigger: 'blur' },
    { min: 2, max: 50, message: '长度在 2 到 50 个字符', trigger: 'blur' }
  ],
  productName: [
    { required: true, message: '请输入产品名称', trigger: 'blur' },
    { min: 2, max: 100, message: '长度在 2 到 100 个字符', trigger: 'blur' }
  ],
  productColor: [
    { required: true, message: '请输入产品颜色', trigger: 'blur' }
  ],
  accountId: [{ required: true, message: '请选择客户', trigger: 'change' }],
  productOrderNum: [
    { required: true, message: '请输入下单数量', trigger: 'blur' },
    { type: 'number', min: 1, message: '下单数量必须大于0', trigger: 'blur' }
  ],
  productFlag: [
    { required: true, message: '请选择生产标志', trigger: 'change' }
  ]
}

// 图片上传前校验
const beforeUpload = (file) => {
  const isImage = file.type.startsWith('image/')
  const isLt2M = file.size / 1024 / 1024 < 2

  if (!isImage) {
    ElMessage.error('只能上传图片文件!')
    return false
  }
  if (!isLt2M) {
    ElMessage.error('图片大小不能超过 2MB!')
    return false
  }
  return true
}

// 图片上传成功回调
const handleUploadSuccess = (response) => {
  if (response.code === 0) {
    formData.value.productImg = response.data.file.url
    ElMessage.success('图片上传成功')
  } else {
    ElMessage.error(response.msg || '图片上传失败')
  }
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

<style lang="scss" scoped>
.product-uploader {
  :deep(.el-upload) {
    border: 1px dashed var(--el-border-color);
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
    transition: var(--el-transition-duration-fast);

    &:hover {
      border-color: var(--el-color-primary);
    }
  }

  .product-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 120px;
    height: 120px;
    text-align: center;
    line-height: 120px;
  }

  .product-image {
    width: 120px;
    height: 120px;
    display: block;
  }
}
</style>
