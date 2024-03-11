<template>
  <div class="container">
    <el-container>
      <el-aside :style="asideStyle()">
        <aside-component></aside-component>
      </el-aside>
      <el-container class="header-main">
        <el-header>
          <header-component></header-component>
        </el-header>
        <el-main>
          <router-view></router-view>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script setup>
import AsideComponent from '@/components/home/AsideComponent.vue';
import HeaderComponent from '@/components/home/HeaderComponent.vue';
import {useAsideStore} from '@/stores/aside'

const asideStore = useAsideStore()

const asideStyle = () => {
  return {
    width: asideStore.asideWidth + 'px',
    minHeight: 100 + 'vh',
    backgroundColor: '#200A3E',
    transition: 'all 1s'
  }
}

</script>

<style lang="scss" scoped>
.el-aside {
  box-shadow: 5px 0px 5px #000;
  animation: changeAside 5s infinite;
}

.el-header {
  padding: 0;
}

.header-main {
  margin-left: 10px;
}

.el-main {
  padding-left: 5px;
  padding-right: 0;
}

@keyframes changeAside {
  0% {
    border-radius: 0px;
  }

  25% {
    border-radius: 0px 15px 15px 0px;
  }

  50% {
    border-radius: 0px 30px 30px 0px;
  }

  75% {
    border-radius: 0px 15px 15px 0px;
  }

  100% {
    border-radius: 0px;
  }
}
</style>