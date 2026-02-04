<template>
  <div class="flex gap-4 p-4 h-[calc(100vh-120px)]">
    <!-- 左侧产品列表 -->
    <div class="w-350px bg-white rounded-lg shadow-sm p-4 flex flex-col">
      <div class="mb-4">
        <h3 class="text-lg font-semibold m-0">产品列表</h3>
      </div>

      <!-- 产品筛选 -->
      <div class="mb-4">
        <el-input
          v-model="productSearch"
          placeholder="搜索产品名称或款号"
          clearable
          prefix-icon="Search"
          @input="handleProductSearch"
        />
      </div>

      <!-- 产品列表 -->
      <div class="flex-1 overflow-hidden">
        <el-scrollbar height="calc(100vh - 240px)">
          <div class="px-2 pt-1 pb-1">
            <div
              v-for="product in filteredProducts"
              :key="product.ID"
              class="product-item flex items-center justify-between p-3 mb-2 bg-gray-50 rounded-md cursor-pointer transition-all hover:bg-gray-100 hover:shadow-md hover:-translate-y-0.5"
              :class="{
                'border-2 border-blue-500 shadow-lg bg-blue-50 -translate-y-0.5': selectedProduct?.ID === product.ID,
                'border-2 border-transparent': selectedProduct?.ID !== product.ID
              }"
              @click="selectProduct(product)"
            >
            <div class="flex-1 min-w-0">
              <div class="text-sm font-medium mb-1 truncate">
                {{ product.productName }}
              </div>
              <div class="text-xs text-gray-500 mb-1.5">款号: {{ product.productCode }}</div>
              <div class="flex gap-1">
                <el-tag size="small">{{ product.productColor }}</el-tag>
                <FfDicTag dicType="productFlag" :dicValue="product.productFlag" />
              </div>
            </div>
            <el-icon
              class="text-base text-gray-400 transition-all"
              :class="{ 'text-blue-500': selectedProduct?.ID === product.ID }"
            >
              <ArrowRight />
            </el-icon>
            </div>
          </div>
          <el-empty
            v-if="filteredProducts.length === 0"
            description="暂无产品数据"
            :image-size="80"
          />
        </el-scrollbar>
      </div>
    </div>

    <!-- 右侧物料管理 -->
    <div class="flex-1 bg-white rounded-lg shadow-sm p-4 overflow-hidden">
      <div v-if="!selectedProduct" class="h-full flex items-center justify-center">
        <el-empty description="请先选择一个产品" :image-size="120" />
      </div>

      <div v-else class="h-full flex flex-col">
        <!-- 物料表头 -->
        <div class="flex justify-between items-center mb-4 pb-4 border-b border-gray-100">
          <div>
            <h3 class="text-lg font-semibold m-0 mb-2">物料管理</h3>
            <div class="text-sm text-gray-500">
              当前产品: {{ selectedProduct.productName }} ({{ selectedProduct.productCode }})
            </div>
          </div>
          <el-button type="primary" icon="Plus" @click="openDialog('add')">
            新增物料
          </el-button>
        </div>

        <!-- 物料表格 -->
        <div class="flex-1 overflow-hidden">
          <el-table
            v-loading="materialLoading"
            :data="materialData"
            row-key="ID"
            border
            stripe
            height="calc(100vh - 240px)"
          >
            <el-table-column prop="ID" label="ID" width="80" align="center" />
            <el-table-column
              prop="materialName"
              label="物料名称"
              min-width="200"
              align="center"
            />
            <el-table-column
              prop="materialColor"
              label="物料颜色"
              min-width="200"
              align="center"
            >
              <template #default="{ row }">
                <el-tag type="primary" size="large">{{ row.materialColor }}</el-tag>
              </template>
            </el-table-column>
            <el-table-column
              prop="materialUsage"
              label="物料用量"
              min-width="200"
              align="center"
            />
            <el-table-column
              prop="materialUom"
              label="物料单位"
              min-width="200"
              align="center"
            />
            <el-table-column
              prop="materialPrice"
              label="参考单价"
              min-width="150"
              align="center"
            >
              <template #default="{ row }">
                <span class="text-base font-semibold text-red-500">
                  ¥{{ row.materialPrice?.toFixed(2) || '0.00' }}
                </span>
              </template>
            </el-table-column>
            <el-table-column
              prop="materialspecification"
              label="物料规格"
              min-width="200"
              align="center"
            />
            <el-table-column
              prop="materialType"
              label="物料类型"
              min-width="200"
              align="center"
            />
            <el-table-column
              prop="supplierId"
              label="供应商ID"
              min-width="200"
              align="center"
            />
            <el-table-column
              label="操作"
              :min-width="appStore.operateMinWith"
              fixed="right"
              align="center"
            >
              <template #default="{ row }">
                <el-button
                  type="primary"
                  link
                  icon="Edit"
                  @click="openDialog('edit', row)"
                >
                  编辑
                </el-button>
                <el-button
                  type="danger"
                  link
                  icon="Delete"
                  @click="handleDelete(row)"
                >
                  删除
                </el-button>
              </template>
            </el-table-column>
          </el-table>

          <el-empty
            v-if="!materialLoading && materialData.length === 0"
            description="暂无物料数据，请添加物料"
            :image-size="100"
            class="mt-8"
          />
        </div>
      </div>
    </div>

    <!-- 新增/编辑物料对话框 -->
    <FfDrawer
      v-model="dialogVisible"
      :title="dialogType === 'add' ? '新增物料' : '编辑物料'"
      @cancel="closeDialog"
      @confirm="handleSubmit"
    >
      <el-form
        ref="formRef"
        :model="formData"
        :rules="formRules"
        label-width="100px"
      >
        <el-form-item label="所属产品">
          <el-input
            :value="selectedProduct?.productName + ' (' + selectedProduct?.productCode + ')'"
            disabled
          />
        </el-form-item>

        <el-form-item label="物料名称" prop="materialName">
          <el-input
            v-model="formData.materialName"
            placeholder="请输入物料名称"
            clearable
            maxlength="100"
            show-word-limit
          />
        </el-form-item>

        <el-form-item label="物料颜色" prop="materialColor">
          <el-input
            v-model="formData.materialColor"
            placeholder="请输入物料颜色"
            clearable
            maxlength="100"
            show-word-limit
          />
        </el-form-item>

        <el-form-item label="物料用量" prop="materialUsage">
          <el-input
            v-model="formData.materialUsage"
            placeholder="请输入物料用量"
            clearable
            maxlength="100"
            show-word-limit
          />
        </el-form-item>

        <el-form-item label="物料单位" prop="materialUom">
          <el-input
            v-model="formData.materialUom"
            placeholder="请输入物料单位"
            clearable
            maxlength="100"
            show-word-limit
          />
        </el-form-item>

        <el-form-item label="物料单价" prop="materialPrice">
          <el-input-number
            v-model="formData.materialPrice"
            :min="0"
            :precision="2"
            :step="0.1"
            controls-position="right"
            style="width: 100%"
            placeholder="请输入物料单价"
          />
          <div class="text-gray-400 text-sm mt-2">单位：元</div>
        </el-form-item>

        <el-form-item label="物料规格" prop="materialspecification">
          <el-input
            v-model="formData.materialspecification"
            placeholder="请输入物料规格"
            clearable
            maxlength="100"
            show-word-limit
          />
        </el-form-item>

        <el-form-item label="物料类型" prop="materialType">
          <el-input
            v-model="formData.materialType"
            placeholder="请输入物料类型"
            clearable
            maxlength="100"
            show-word-limit
          />
        </el-form-item>

        <el-form-item label="供应商" prop="supplierId">
          <el-input
            v-model="formData.supplierId"
            placeholder="请输入供应商"
            clearable
            maxlength="100"
            show-word-limit
          />
        </el-form-item>
      </el-form>
    </FfDrawer>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ArrowRight } from '@element-plus/icons-vue'
import { useAppStore } from '@/pinia'
import FfDrawer from '@/components/ffDrawer/index.vue'
import FfDicTag from '@/components/ffDicTag/index.vue'
import { getProductList } from '@/api/ffProduct'
import { getMaterialList, createMaterial, updateMaterial, deleteMaterial } from '@/api/ffMaterial'

defineOptions({
  name: 'FfMaterial'
})

const appStore = useAppStore()

// ==================== 产品相关 ====================
const productLoading = ref(false)
const productData = ref([])
const productSearch = ref('')
const selectedProduct = ref(null)

// 筛选后的产品列表
const filteredProducts = computed(() => {
  if (!productSearch.value) {
    return productData.value
  }
  const searchLower = productSearch.value.toLowerCase()
  return productData.value.filter(
    (item) =>
      item.productName.toLowerCase().includes(searchLower) ||
      item.productCode.toLowerCase().includes(searchLower)
  )
})

// 产品搜索
const handleProductSearch = () => {
  // 搜索是通过计算属性实现的，这里可以添加额外逻辑
}

// 选择产品
const selectProduct = (product) => {
  selectedProduct.value = product
  getmaterialData()
}

// 获取产品列表（使用真实API）
const getProductData = async () => {
  productLoading.value = true
  try {
    const res = await getProductList({ page: 1, pageSize: 100 })
    if (res.code === 0) {
      productData.value = res.data.list
    } else {
      ElMessage.error('获取产品列表失败')
    }
  } catch (error) {
    ElMessage.error('获取产品列表失败')
  } finally {
    productLoading.value = false
  }
}

// ==================== 物料相关 ====================
const materialLoading = ref(false)
const materialData = ref([])

// 对话框
const dialogVisible = ref(false)
const dialogType = ref('add') // 'add' | 'edit'
const formRef = ref(null)
const formData = ref({
    ID: null,
    productId: null,
    materialName: '',
    materialColor: '',
    materialUsage: null,
    materialUom: '',
    materialPrice: null,
    materialspecification: '',
    materialType: '',
    supplierId: ''
})

// 表单验证规则
const formRules = reactive({
  materialName: [
    { required: true, message: '请输入物料名称', trigger: 'blur' },
    { min: 2, max: 100, message: '长度在 2 到 100 个字符', trigger: 'blur' }
  ],
  materialPrice: [
    { required: true, message: '请输入参考单价', trigger: 'blur' },
    { type: 'number', min: 0, message: '参考单价不能小于0', trigger: 'blur' }
  ]
})

// 获取物料数据（使用真实API）
const getmaterialData = async () => {
  if (!selectedProduct.value) {
    materialData.value = []
    return
  }

  materialLoading.value = true
  try {
    const res = await getMaterialList({
      productId: selectedProduct.value.ID,
      page: 1,
      pageSize: 100
    })
    if (res.code === 0) {
      materialData.value = res.data.list || []
    } else {
      ElMessage.error('获取物料列表失败')
      materialData.value = []
    }
  } catch (error) {
    ElMessage.error('获取物料列表失败')
    materialData.value = []
  } finally {
    materialLoading.value = false
  }
}

// 打开对话框
const openDialog = (type, row = null) => {
  if (!selectedProduct.value) {
    ElMessage.warning('请先选择产品')
    return
  }

  dialogType.value = type
  if (type === 'add') {
    formData.value = {
      ID: null,
      productId: selectedProduct.value.ID,
      materialName: '',
      materialColor: '',
      materialUsage: null,
      materialUom: '',
      materialPrice: null,
      materialspecification: '',
      materialType: '',
      supplierId: ''
    }
  } else {
    formData.value = {
      ID: row.ID,
      productId: row.productId,
      materialName: row.materialName,
      materialColor: row.materialColor,
      materialUsage: row.materialUsage,
      materialUom: row.materialUom,
      materialPrice: row.materialPrice,
      materialspecification: row.materialspecification,
      materialType: row.materialType,
      supplierId: row.supplierId
    }
  }
  dialogVisible.value = true
}

// 关闭对话框
const closeDialog = () => {
  formRef.value?.resetFields()
  dialogVisible.value = false
}

// 提交表单（使用真实API）
const handleSubmit = async () => {
  const valid = await formRef.value?.validate()
  if (!valid) return

  try {
    materialLoading.value = true
    const data = {
      ...formData.value,
      productId: selectedProduct.value.ID
    }

    if (dialogType.value === 'add') {
      delete data.ID
    }

    let res
    if (dialogType.value === 'add') {
      res = await createMaterial(data)
    } else {
      res = await updateMaterial(data)
    }

    if (res.code === 0) {
      ElMessage.success(res.msg)
      closeDialog()
      getmaterialData()
    } else {
      ElMessage.error(res.msg || '操作失败')
    }
  } catch (error) {
    ElMessage.error('操作失败')
  } finally {
    materialLoading.value = false
  }
}

// 删除物料（使用真实API）
const handleDelete = (row) => {
  ElMessageBox.confirm(
    `确定要删除物料 "${row.materialName}" 吗？`,
    '删除确认',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }
  )
    .then(async () => {
      try {
        materialLoading.value = true
        const res = await deleteMaterial({ ID: row.ID })
        if (res.code === 0) {
          ElMessage.success(res.msg)
          getmaterialData()
        } else {
          ElMessage.error(res.msg || '删除失败')
        }
      } catch (error) {
        ElMessage.error('删除失败')
      } finally {
        materialLoading.value = false
      }
    })
    .catch(() => {
      // 用户取消删除
    })
}

// ==================== 生命周期 ====================
onMounted(() => {
  getProductData()
})
</script>

