<template>
    <el-drawer v-model="visible"
        :title="title"
        :size="drawerSize"
        :show-close="false"
        :close-on-press-escape="false"
        :close-on-click-modal="false"
        :destroy-on-close="destroyOnClose">
        <template #header>
            <div class="flex justify-between items-center">
                <span class="text-lg">{{ title }}</span>
                <div>
                    <el-button @click="handleCancel">{{ cancelText }}</el-button>
                    <el-button type="primary"
                        :loading="loading"
                        @click="handleConfirm">
                        {{ confirmText }}
                    </el-button>
                </div>
            </div>
        </template>

        <slot />
    </el-drawer>
</template>

<script setup>
import { computed } from 'vue'
import { useAppStore } from '@/pinia'

defineOptions({
  name: 'FfDrawer'
})

const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false
  },
  title: {
    type: String,
    default: ''
  },
  size: {
    type: String,
    default: ''
  },
  loading: {
    type: Boolean,
    default: false
  },
  cancelText: {
    type: String,
    default: '取消'
  },
  confirmText: {
    type: String,
    default: '确定'
  },
  destroyOnClose: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['update:modelValue', 'cancel', 'confirm'])

const appStore = useAppStore()

const visible = computed({
  get: () => props.modelValue,
  set: (val) => emit('update:modelValue', val)
})

const drawerSize = computed(() => props.size || appStore.drawerSize)

const handleCancel = () => {
  emit('cancel')
  visible.value = false
}

const handleConfirm = () => {
  emit('confirm')
}
</script>
