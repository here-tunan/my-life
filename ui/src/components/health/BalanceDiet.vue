<template>
  <div class="container">
    <h1>均衡膳食</h1>
    <el-calendar v-model="value" @mouseenter="handleDateMouseEnter">

      <template #date-cell>
        <el-icon style="border: 1px solid #ddd;border-radius: 5px;">
          <Plus/>
        </el-icon>
        <div style="align-items: center;text-align: center">每日饮食</div>
      </template>
    </el-calendar>
  </div>
</template>

<script setup>
import {ref} from "vue";
import {ElMessage} from "element-plus";
import {Plus} from "@element-plus/icons-vue";

const value = ref(new Date())

const handleDateMouseEnter = (date) => {
  ElMessage.success('当前时间段：' + date.toString())
}
</script>
