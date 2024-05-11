<template>
  <div class="container">
    <el-card>
      <h2>关于本项目</h2>
      <p>本项目基于 Vue 3.x + Gin 开发，主要用于个人毕业设计，旨在为用户提供一个简单、易用、实用的水果售卖管理系统。</p>
      <p>项目地址：https://github.com/cukor666/NourishOrchard</p>
      <p>作者：CukorZhong</p>
      <p>邮箱：cukorzhong@gamil.com</p>
      <h2>技术栈：</h2>
      <ul>
        <li>前端：
          <ul>
            <li>Vue 3.x</li>
            <li>Element-Plus</li>
            <li>Axios</li>
            <li>ECharts</li>
            <li>Vue-Router</li>
            <li>Pinia</li>
            <li>Sass</li>
          </ul>
        </li>
        <li>后端：
          <ul>
            <li>Gin</li>
            <li>Gorm</li>
            <li>MySQL</li>
            <li>Redis</li>
            <li>JWT</li>
            <li>Validator</li>
          </ul>
        </li>
      </ul>
      <div class="logo"></div>
    </el-card>
  </div>
</template>

<script setup>

</script>

<style scoped lang="scss">
.container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 80vh;
}

.el-card {
  width: 60%;
  border-radius: 15px;
  position: relative;
  z-index: 10;
  overflow: hidden;
  animation: changeBoxShadow 10s linear infinite;


  &::before {
    content: "";
    position: absolute;
    top: 50%;
    left: 50%;
    $r: 950px;
    width: $r;
    height: $r;
    // 居中
    transform: translate(-50%, -50%);
    // 锥形渐变背景
    background: conic-gradient(#fd2a2a, #ffd817, #eaff35, #fd2a2a);
    z-index: -2;
    animation: rotate 10s linear infinite;
  }

  @keyframes rotate {
    100% {
      transform: translate(-50%, -50%) rotate(360deg);
    }
  }

  &::after {
    content: "";
    position: absolute;
    top: 40%;
    left: 40%;
    $r: 200px;
    width: $r;
    height: $r;
    border-radius: 30%;
    // 居中
    //transform: translate(-5%, -5%);
    background: #fff;
    z-index: -1;
    animation: rotate 10s linear infinite;
    box-shadow: 100px 100px 10px #fff, -100px -80px 10px #fff;
  }
}


@keyframes changeBoxShadow {
  50% {
    box-shadow: 84px 80px 100px #F7E722, -77px -52px 100px red;
  }
}

h2 {
  margin-bottom: 5px;
}

p {
  margin-bottom: 5px;
}

ul {
  // 去除ul默认的列表符号
  list-style: none;
  margin-left: 20px;

  li {
    margin-bottom: 5px;
  }
}

.logo {
  position: absolute;
  $r: 300px;
  width: $r;
  height: $r;
  border-radius: 15px;
  transform: translate(274px, 100px);

  background-color: #ccc;
  background-image: url("https://cukor-resource-1318313222.cos.ap-guangzhou.myqcloud.com/img/cukorlogo.jpg");
  background-size: 450px;
  background-repeat: no-repeat;
  background-position: center;
  animation: movein 1s ease-in-out 1;
  animation-fill-mode: forwards;
}

@keyframes movein {
  80% {
    transform: translate(274px, -370px);
  }
  100% {
    transform: translate(274px, -358px);
  }

}
</style>