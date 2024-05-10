<template>
  <div class="container">
    <h1>支付页面</h1>
    <el-card>
      <div class="qrcode"></div>
      <el-button type="primary" @click="pay">模拟支付</el-button>
    </el-card>
  </div>
</template>


<script setup>

const pay = () => {
  console.log('模拟支付')
}

</script>



<style scoped lang="scss">
.container {
  width: 100%;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
}
.el-card {
  border-radius: 15px;
  margin-top: 50px;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  .el-button {
    transform: translate(154px, 10px);
  }
}
.qrcode {
  width: 400px;
  height: 300px;
  background-color: #f2f2f2;
  background-image: url("@/assets/收付款.png");
  background-size: 400px;
  background-position: 0px -85px;
  background-repeat: no-repeat;
}
</style>