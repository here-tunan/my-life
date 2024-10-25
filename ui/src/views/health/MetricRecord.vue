<template>
  <div>
    <h1 style="margin-bottom: 20px">体重记录</h1>
    <el-row>
      <el-form-item label="日期">
        <el-date-picker
            v-model="recordDate"
            type="date"
            placeholder="Pick a date"
            style="width: 140px"
            :shortcuts="shortcuts"
            value-format="YYYY-MM-DD"
        />
      </el-form-item>

      <el-form-item label="体重">
        <el-input-number v-model="weightNum" :precision="2" :step="1"/>
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="updateWeight">更新</el-button>
      </el-form-item>
    </el-row>

    <h1 style="margin-top: 20px; margin-bottom: 20px">饮食记录</h1>
    <el-row>
      <h4 style="margin-bottom: 20px">功能待实现中...</h4>
    </el-row>

  </div>
</template>

<script setup lang="ts">
import {ref, onMounted} from 'vue';
import {putWeight} from '@/api/health/weight'
import {ElMessage} from "element-plus";


// 今天的日期
const recordDate = ref('');

// 默认输入体重
const weightNum = ref(100)

const updateWeight = () => {
  console.log(recordDate.value)
  // 添加到数据库中
  putWeight({
    day: recordDate.value,
    weight: weightNum.value
  }).then(res => {
    if (!res.success) {
      ElMessage({
        type: 'info',
        message: '保存失败',
      })
      return
    }
    ElMessage({
      type: 'success',
      message: '保存成功',
    })
  })
};

onMounted(() => {

  // 时间初始化
  const now = new Date();
  recordDate.value = now.toISOString().slice(0, 10);

});

</script>

<style lang="scss" scoped>
.el-form-item {
  margin-right: 20px;
}
</style>