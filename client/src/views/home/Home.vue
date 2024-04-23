<template>
  <div class="container">
    <el-carousel indicator-position="outside" style="width: 100%">
      <el-carousel-item v-for="(item, index) in swiper" :key="index">
        <img :src="item.img" alt="无法显示"/>
      </el-carousel-item>
    </el-carousel>
    <div class="main-content">
      <div class="left">
        <e-charts :option="option1"/>
      </div>
      <div class="right">
        <e-charts :option="option2"/>
      </div>
    </div>
    <div class="footer">
      <e-charts :option="option3"/>
    </div>
  </div>
</template>

<script setup>
import {computed, onMounted, ref} from "vue";
import request from "@/axios/request.js";
import * as echarts from 'echarts';

const swiper = ref([
  {
    id: 1,
    img: "https://pic.quanjing.com/u0/p9/QJ6660181509.jpg?x-oss-process=style/794wsy"
  },
  {
    id: 2,
    img: "https://pic.quanjing.com/bi/0c/QJ6375588182.jpg?x-oss-process=style/794wsy"
  },
  {
    id: 3,
    img: "https://pic.quanjing.com/gg/us/QJ6520051711.jpg?x-oss-process=style/794wsy"
  }
])

const option1Data = ref([
  {value: 1000, name: '用户'},
  {value: 120, name: '员工'},
  {value: 14, name: '管理员'},
])

const option1 = computed(() => {
  return {
    title: {
      text: '人员统计'
    },
    xAxis: {
      type: 'category',
      data: option1Data.value.map(item => item.name)
    },
    yAxis: {
      type: 'value'
    },
    series: [
      {
        data: option1Data.value.map(item => item.value),
        type: 'bar'
      }
    ]
  }
})

const option2Data = ref([
  {value: 40, name: '广西'},
  {value: 38, name: '广东'},
  {value: 32, name: '海南'},
])

const option2 = computed(() => {
  return {
    title: {
      text: '水果来源分布'
    },
    legend: {
      top: 'top'
    },
    toolbox: {
      show: true,
      feature: {
        mark: {show: true},
        dataView: {show: true, readOnly: false},
        restore: {show: true},
        saveAsImage: {show: true}
      }
    },
    series: [
      {
        name: '南丁格尔图',
        type: 'pie',
        radius: [50, 160],
        center: ['50%', '50%'],
        roseType: 'area',
        itemStyle: {
          borderRadius: 8
        },
        data: option2Data.value
      }
    ]
  }
})

function getVirtualData(year) {
  const date = +echarts.time.parse(year + '-01-01');
  const end = +echarts.time.parse(+year + 1 + '-01-01');
  const dayTime = 3600 * 24 * 1000;
  const data = [];
  for (let time = date; time < end; time += dayTime) {
    data.push([
      echarts.time.format(time, '{yyyy}-{MM}-{dd}', false),
      Math.floor(Math.random() * 500)
    ]);
  }
  return data;
}

const option3 = computed(() => {
  return {
    title: {
      top: 30,
      left: 'center',
      text: '订单热力图'
    },
    tooltip: {},
    visualMap: {
      min: 0,
      max: 100,
      type: 'piecewise',
      orient: 'horizontal',
      left: 'center',
      top: 65
    },
    calendar: {
      top: 120,
      left: 30,
      right: 30,
      cellSize: ['auto', 25],
      range: '2024',
      itemStyle: {
        borderWidth: 0.5
      },
      yearLabel: { show: false }
    },
    series: {
      type: 'heatmap',
      coordinateSystem: 'calendar',
      data: getVirtualData('2024')
    }
  }
})

const getOption1Data = async () => {
  try {
    let res = await request.get('/account/cnt')
    if (res.code === 200) {
      option1Data.value = res.data
    } else {
      console.error('获取数据失败')
    }
  } catch (error) {
    console.error(error)
  }
}

const getOption2Data = async () => {
  try {
    let res = await request.get('/fruit/origin-pie')
    if (res.code === 200) {
      option2Data.value = res.data
    } else {
      console.error('获取数据失败')
    }
  } catch (error) {
    console.error(error)
  }
}

onMounted(() => {
  getOption1Data()
  getOption2Data()
})

</script>

<style lang="scss" scoped>
.container {
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.main-content {
  width: 100%;
  height: 400px;
  margin-top: 20px;
  padding: 10px;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #e6f3ec;
  box-shadow: 0 5px 5px rgba(0, 0, 0, 0.3);
  overflow: hidden;

  .left {
    width: 50%;
    height: 100%;
    display: flex;
    justify-self: center;
    align-items: center;
  }

  .right {
    width: 50%;
    height: 100%;
    display: flex;
    justify-self: center;
    align-items: center;
  }
}

.footer {
  width: 100%;
  height: 300px;
  background-color: #f6eded;
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 20px;
  padding: 10px;
  box-shadow: 0 5px 5px rgba(0, 0, 0, 0.3);
}

img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

</style>