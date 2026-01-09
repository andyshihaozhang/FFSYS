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
              :key="product.id"
              class="product-item flex items-center justify-between p-3 mb-2 bg-gray-50 rounded-md cursor-pointer transition-all hover:bg-gray-100 hover:shadow-md hover:-translate-y-0.5"
              :class="{
                'border-2 border-blue-500 shadow-lg bg-blue-50 -translate-y-0.5': selectedProduct?.id === product.id,
                'border-2 border-transparent': selectedProduct?.id !== product.id
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
                <el-tag
                  :type="getProductFlagType(product.productFlag)"
                  size="small"
                >
                  {{ product.productFlag }}
                </el-tag>
              </div>
            </div>
            <el-icon
              class="text-base text-gray-400 transition-all"
              :class="{ 'text-blue-500': selectedProduct?.id === product.id }"
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
            row-key="id"
            border
            stripe
            height="calc(100vh - 240px)"
          >
            <el-table-column prop="id" label="ID" width="80" align="center" />
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

defineOptions({
  name: 'FfMaterial'
})

const appStore = useAppStore()

// ==================== 模拟数据 ====================
// 模拟产品数据
const mockProducts = [
  {
    id: 1,
    productCode: 'P001',
    productName: '春季新款连衣裙',
    productColor: '蓝色',
    productImg: 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg',
    accountId: 1001,
    productOrderNum: 100,
    productFlag: '生产中'
  },
  {
    id: 2,
    productCode: 'P002',
    productName: '夏季短袖T恤',
    productColor: '白色',
    productImg: 'https://fuss10.elemecdn.com/a/3f/3302e58f9a181d2509f3dc0fa68b0jpeg.jpeg',
    accountId: 1002,
    productOrderNum: 200,
    productFlag: '未生产'
  },
  {
    id: 3,
    productCode: 'P003',
    productName: '秋季卫衣',
    productColor: '灰色',
    productImg: '',
    accountId: 1001,
    productOrderNum: 150,
    productFlag: '已完成'
  },
  {
    id: 4,
    productCode: 'P004',
    productName: '冬季羽绒服',
    productColor: '黑色',
    productImg: 'https://fuss10.elemecdn.com/1/34/19aa98b1fcb2781c4fba33d850549jpeg.jpeg',
    accountId: 1003,
    productOrderNum: 80,
    productFlag: '生产中'
  },
  {
    id: 5,
    productCode: 'P005',
    productName: '牛仔裤',
    productColor: '深蓝',
    productImg: '',
    accountId: 1002,
    productOrderNum: 120,
    productFlag: '未生产'
  }
]

// 模拟物料数据（按产品ID存储）
const mockMaterial = {
  1: [ // 产品ID为1的工序
    {
      id: 1,
      productId: 1,
      materialName: '高档真丝面料',
      materialColor: '蓝色',
      materialUsage: 2.5,
      materialUom: '米',
      materialPrice: 158.00,
      materialspecification: '114cm×100cm',
      materialType: '主料',
      supplierId: 'S001'
    },
    {
      id: 2,
      productId: 1,
      materialName: '拉链',
      materialColor: '银色',
      materialUsage: 1,
      materialUom: '条',
      materialPrice: 3.50,
      materialspecification: '50cm',
      materialType: '辅料',
      supplierId: 'S002'
    },
    {
      id: 3,
      productId: 1,
      materialName: '纽扣',
      materialColor: '白色',
      materialUsage: 5,
      materialUom: '个',
      materialPrice: 0.80,
      materialspecification: '直径1.2cm',
      materialType: '辅料',
      supplierId: 'S003'
    },
    {
      id: 4,
      productId: 1,
      materialName: '缝纫线',
      materialColor: '蓝色',
      materialUsage: 200,
      materialUom: '米',
      materialPrice: 0.05,
      materialspecification: '40/2',
      materialType: '辅料',
      supplierId: 'S002'
    }
  ],
  2: [
    {
      id: 5,
      productId: 2,
      materialName: '纯棉针织布',
      materialColor: '白色',
      materialUsage: 1.2,
      materialUom: '米',
      materialPrice: 28.00,
      materialspecification: '160cm宽',
      materialType: '主料',
      supplierId: 'S001'
    },
    {
      id: 6,
      productId: 2,
      materialName: '印花浆料',
      materialColor: '彩色',
      materialUsage: 0.1,
      materialUom: '千克',
      materialPrice: 45.00,
      materialspecification: '水性环保',
      materialType: '辅料',
      supplierId: 'S004'
    },
    {
      id: 7,
      productId: 2,
      materialName: '领标',
      materialColor: '白色',
      materialUsage: 1,
      materialUom: '张',
      materialPrice: 0.50,
      materialspecification: '5cm×2cm',
      materialType: '辅料',
      supplierId: 'S003'
    }
  ],
  3: [
    {
      id: 8,
      productId: 3,
      materialName: '卫衣面料',
      materialColor: '灰色',
      materialUsage: 1.8,
      materialUom: '米',
      materialPrice: 35.00,
      materialspecification: '180cm宽加厚',
      materialType: '主料',
      supplierId: 'S001'
    },
    {
      id: 9,
      productId: 3,
      materialName: '绣花线',
      materialColor: '金色',
      materialUsage: 50,
      materialUom: '米',
      materialPrice: 0.15,
      materialspecification: '120D/2',
      materialType: '辅料',
      supplierId: 'S002'
    },
    {
      id: 10,
      productId: 3,
      materialName: '帽绳',
      materialColor: '灰色',
      materialUsage: 1.5,
      materialUom: '米',
      materialPrice: 2.00,
      materialspecification: '0.8cm圆绳',
      materialType: '辅料',
      supplierId: 'S003'
    }
  ],
  4: [
    {
      id: 11,
      productId: 4,
      materialName: '防水尼龙布',
      materialColor: '黑色',
      materialUsage: 3.0,
      materialUom: '米',
      materialPrice: 55.00,
      materialspecification: '150cm宽防水',
      materialType: '主料',
      supplierId: 'S005'
    },
    {
      id: 12,
      productId: 4,
      materialName: '90%白鸭绒',
      materialColor: '白色',
      materialUsage: 0.25,
      materialUom: '千克',
      materialPrice: 280.00,
      materialspecification: '90绒10毛',
      materialType: '填充料',
      supplierId: 'S006'
    },
    {
      id: 13,
      productId: 4,
      materialName: '金属拉链',
      materialColor: '黑色',
      materialUsage: 1,
      materialUom: '条',
      materialPrice: 8.50,
      materialspecification: '70cm双向',
      materialType: '辅料',
      supplierId: 'S002'
    },
    {
      id: 14,
      productId: 4,
      materialName: '内衬布',
      materialColor: '黑色',
      materialUsage: 2.5,
      materialUom: '米',
      materialPrice: 12.00,
      materialspecification: '150cm宽',
      materialType: '里料',
      supplierId: 'S001'
    }
  ],
  5: [
    {
      id: 15,
      productId: 5,
      materialName: '牛仔布',
      materialColor: '深蓝',
      materialUsage: 1.5,
      materialUom: '米',
      materialPrice: 32.00,
      materialspecification: '148cm宽12盎司',
      materialType: '主料',
      supplierId: 'S001'
    },
    {
      id: 16,
      productId: 5,
      materialName: '铜质五爪扣',
      materialColor: '古铜色',
      materialUsage: 1,
      materialUom: '套',
      materialPrice: 1.20,
      materialspecification: '17mm',
      materialType: '辅料',
      supplierId: 'S003'
    },
    {
      id: 17,
      productId: 5,
      materialName: '口袋布',
      materialColor: '白色',
      materialUsage: 0.3,
      materialUom: '米',
      materialPrice: 8.00,
      materialspecification: '斜纹布',
      materialType: '辅料',
      supplierId: 'S001'
    }
  ]
}

// 用于生成新工序ID
let nextmaterialId = 18

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

// 获取生产标志标签类型
const getProductFlagType = (flag) => {
  const typeMap = {
    未生产: 'info',
    生产中: 'warning',
    已完成: 'success'
  }
  return typeMap[flag] || 'info'
}

// 产品搜索
const handleProductSearch = () => {
  // 搜索是通过计算属性实现的，这里可以添加额外逻辑
}

// 选择产品
const selectProduct = (product) => {
  selectedProduct.value = product
  getmaterialData()
}

// 获取产品列表（使用模拟数据）
const getProductData = async () => {
  productLoading.value = true
  try {
    // 模拟API延迟
    await new Promise(resolve => setTimeout(resolve, 300))
    productData.value = [...mockProducts]
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
    id: null,
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

// 获取物料数据（使用模拟数据）
const getmaterialData = async () => {
  if (!selectedProduct.value) {
    materialData.value = []
    return
  }

  materialLoading.value = true
  try {
    // 模拟API延迟
    await new Promise(resolve => setTimeout(resolve, 300))
    const productId = selectedProduct.value.id
    materialData.value = mockMaterial[productId] ? [...mockMaterial[productId]] : []
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
      id: null,
      productId: selectedProduct.value.id,
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
      id: row.id,
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

// 提交表单（使用模拟数据）
const handleSubmit = async () => {
  const valid = await formRef.value?.validate()
  if (!valid) return

  try {
    materialLoading.value = true
    // 模拟API延迟
    await new Promise(resolve => setTimeout(resolve, 300))

    const productId = selectedProduct.value.id

    if (dialogType.value === 'add') {
      // 新增物料
      const newmaterial = {
        id: nextmaterialId++,
        productId: productId,
        materialName: formData.value.materialName,
        materialColor: formData.value.materialColor,
        materialUsage: formData.value.materialUsage,
        materialUom: formData.value.materialUom,
        materialPrice: formData.value.materialPrice,
        materialspecification: formData.value.materialspecification,
        materialType: formData.value.materialType,
        supplierId: formData.value.supplierId
      }

      if (!mockMaterial[productId]) {
        mockMaterial[productId] = []
      }
      mockMaterial[productId].push(newmaterial)

      ElMessage.success('新增成功')
    } else {
      // 编辑物料
      const materialList = mockMaterial[productId]
      if (materialList) {
        const index = materialList.findIndex(op => op.id === formData.value.id)
        if (index !== -1) {
          materialList[index] = {
            ...materialList[index],
            materialName: formData.value.materialName,
            materialColor: formData.value.materialColor,
            materialUsage: formData.value.materialUsage,
            materialUom: formData.value.materialUom,
            materialPrice: formData.value.materialPrice,
            materialspecification: formData.value.materialspecification,
            materialType: formData.value.materialType,
            supplierId: formData.value.supplierId
          }
        }
      }

      ElMessage.success('更新成功')
    }

    closeDialog()
    getmaterialData()
  } catch (error) {
    ElMessage.error('操作失败')
  } finally {
    materialLoading.value = false
  }
}

// 删除物料（使用模拟数据）
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
        // 模拟API延迟
        await new Promise(resolve => setTimeout(resolve, 300))

        const productId = selectedProduct.value.id
        const materialList = mockMaterial[productId]

        if (materialList) {
          const index = materialList.findIndex(op => op.id === row.id)
          if (index !== -1) {
            materialList.splice(index, 1)
          }
        }

        ElMessage.success('删除成功')
        getmaterialData()
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

