<template>
  <el-select v-model="selectedValue" v-bind="$attrs">
    <el-option v-for="item in dictList" :key="item.value" :label="item.label" :value="item.value" />
  </el-select>
</template>

<script setup>
import { ref, watch, computed } from 'vue'
import { getDict } from '@/utils/dictionary'

defineOptions({
  name: 'ffDicSelect'
})

const props = defineProps({
  dicType: {
    type: String,
    required: true
  },
  modelValue: {
    type: [String, Number],
    default: null
  }
})

const emit = defineEmits(['update:modelValue'])

const dictList = ref([])

const selectedValue = computed({
  get: () => {
    const item = dictList.value.find(item => item.value == props.modelValue)
    return item ? item.value : props.modelValue
  },
  set: (val) => {
    console.log('选中项目刷新: ', val);
    emit('update:modelValue', val)
  }
})

const loadDict = async () => {
  try {
    dictList.value = await getDict(props.dicType)
  } catch (error) {
    console.error('获取字典失败:', error)
  }
}

watch(() => props.dicType, () => {
  loadDict()
}, { immediate: true })
</script>
