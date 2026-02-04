<template>
  <fc-designer ref="designer" :config="config"/>
</template>

<script setup>
  import { ref } from 'vue'
  import FcDesigner from '@form-create/designer'

  defineOptions({
    name: 'FormGenerator'
  })

  const designer = ref(null)

  const config = {
    fieldReadonly: false
  }
</script>
