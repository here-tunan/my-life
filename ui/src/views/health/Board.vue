<template>
  <div>

    <h1 style="margin-top: 20px; margin-bottom: 20px">体重看板</h1>
    <el-row>
      <div id="weightChart" style="width: 800px; height: 400px;"></div>
    </el-row>

  </div>
</template>

<script setup lang="ts">
import {ref, onMounted} from 'vue';
import {getWeightList} from '@/api/health/weight'

import * as echarts from 'echarts';

// 图表
const myChart = ref(null);

const xData = ref([]);
const yData = ref([]);

const setWeightChartData = (x, y) => {
  myChart.value.setOption({
    title: {
      text: '体重变化情况'
    },
    tooltip: {},
    xAxis: {
      data: x
    },
    yAxis: {
      min: function (value) {
        return value.min - 20;
      },
      scale: true,
    },
    series: [
      {
        name: '体重kg',
        type: 'line',
        data: y,
        colorBy: 'data',
        markLine: {
          silent: true,
          lineStyle: {
            color: '#333'
          },
          data: [
            {
              yAxis: 85
            },
            {
              yAxis: 100
            }
          ]
        }
      }
    ]
  });
};

onMounted(() => {
  // 图表初始化
  myChart.value = echarts.init(document.getElementById('weightChart'));

  // 查询最近三十次的数据记录
  getWeightList({
    page: 1,
    size: 30
  }).then(res => {
    console.log(res)
    if (res.success) {
      console.log(res)
      let data = res.data
      // 请按倒序显示
      for (let item of data) {
        xData.value.push(item.day)
        yData.value.push(item.weight)
      }
      xData.value.reverse()
      yData.value.reverse()
      setWeightChartData(xData.value, yData.value)
    }
  })

});

</script>

<style lang="scss" scoped>
.el-form-item {
  margin-right: 20px;
}
</style>