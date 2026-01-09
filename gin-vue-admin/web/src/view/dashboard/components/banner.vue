<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/8
    @description: Banner轮播图组件 - 使用随机风景图
!-->

<template>
  <el-carousel class="-mt-2" :interval="5000" arrow="hover">
    <el-carousel-item
      class="cursor-pointer lg:h-40"
      v-for="(item, index) in banners"
      :key="index"
      @click="openLink(item.link)"
    >
      <el-image
        class="h-full w-full"
        :src="item.img"
        fit="cover"
        loading="lazy"
      >
        <template #error>
          <div class="h-full w-full flex items-center justify-center bg-gray-100">
            <el-icon :size="40" color="#ccc">
              <Picture />
            </el-icon>
          </div>
        </template>
      </el-image>
    </el-carousel-item>
  </el-carousel>
</template>

<script setup>
import { Picture } from '@element-plus/icons-vue'

/**
 * 打开外部链接
 * @param {string} link - 链接地址
 */
const openLink = (link) => {
  if (link) {
    window.open(link, '_blank')
  }
}

// 图床API基础地址
const IMAGE_API = 'https://image.ai6.me/random?type=img&dir=/1/landscape/'

/**
 * Banner配置
 * 使用随机图床API，每次刷新显示不同的风景图
 * 通过添加随机参数确保每张图片都不同
 */
const banners = [
  {
    img: `${IMAGE_API}&t=${Date.now()}`,
    link: 'https://image.ai6.me/browse/landscape'
  },
  {
    img: `${IMAGE_API}&t=${Date.now() + 1}`,
    link: 'https://image.ai6.me/browse/landscape'
  },
  {
    img: `${IMAGE_API}&t=${Date.now() + 2}`,
    link: 'https://image.ai6.me/browse/landscape'
  },
  {
    img: `${IMAGE_API}&t=${Date.now() + 3}`,
    link: 'https://image.ai6.me/browse/landscape'
  }
]
</script>

<style scoped lang="scss">
:deep(.el-carousel__item) {
  overflow: hidden;
}

:deep(.el-image) {
  transition: transform 0.3s ease;
}

:deep(.el-carousel__item:hover .el-image) {
  transform: scale(1.05);
}
</style>
