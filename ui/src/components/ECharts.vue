<template>
    <div ref="chartContainer" :style="{ width: width, height: height }"></div>
  </template>
  
  <script>
  import * as echarts from 'echarts';
  
  export default {
    name: 'ECharts',
    props: {
      width: {
        type: [String, Number],
        default: '100%'
      },
      height: {
        type: [String, Number],
        default: '400px'
      },
      options: {
        type: Object,
        required: true
      }
    },
    data() {
      return {
        chart: null
      };
    },
    mounted() {
      this.initChart();
    },
    beforeUnmount() {
      this.disposeChart();
    },
    methods: {
      initChart() {
        this.chart = echarts.init(this.$refs.chartContainer);
        this.chart.setOption(this.options);
      },
      disposeChart() {
        if (this.chart) {
          this.chart.dispose();
          this.chart = null;
        }
      }
    },
    watch: {
      options: {
        deep: true,
        handler(newOptions) {
          if (this.chart) {
            this.chart.setOption(newOptions);
          }
        }
      }
    }
  };
  </script>
