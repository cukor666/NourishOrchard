<template>
  <RouterView />
</template>

<style>
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
  background-color: #ADCAC3;
}

/* 滚动条上的滑块 */
::-webkit-scrollbar-thumb {
  background-color: #BFD2DB;
  border-radius: 5px;
}
</style>