<template>
  <RouterView />
</template>

<style lang="scss">
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  background-color: #ebe6e6;
}

/* 滚动条 */
::-webkit-scrollbar {
  width: 8px;
  background-color: #ebe6e6;
}

/* 滚动条上的滑块 */
::-webkit-scrollbar-thumb {
  background-color: #f4cccc;
  border-radius: 15px;
}

// 引入 CSS
@import url(./assets/css/element-variables.scss)
</style>