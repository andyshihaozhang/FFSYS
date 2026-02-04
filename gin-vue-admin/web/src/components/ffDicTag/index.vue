<template>
  <el-tag v-if="dictItem" :type="dictItem.extend" :size="size">
    {{ dictItem.label }}
  </el-tag>
</template>

<script setup>
import { ref, watch } from 'vue'
import { getDict } from '@/utils/dictionary'

defineOptions({
  name: 'ffDicTag'
})

const props = defineProps({
  dicType: {
    type: String,
    required: true
  },
  size: {
    type: String,
    default: 'large'
  },
  dicValue: {
    type: [String, Number],
    required: true
  }
})

const dictItem = ref(null)

const loadDict = async () => {
  try {
    const dictList = await getDict(props.dicType)
    dictItem.value = dictList.find(item => item.value == props.dicValue) || null
  } catch (error) {
    console.error('获取字典失败:', error)
  }
}

watch(() => [props.dicType, props.dicValue], () => {
  loadDict() 
}, { immediate: true })
</script>

