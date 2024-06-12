<template>
  <el-row>
    <el-col :span="24">
      <el-calendar>
        <template #date-cell="{ data }">
          <div class="exercise-day-calendar">
            <el-icon v-if="data.isSelected" class="add-item-icon" @click="addItem(data)"><Plus/></el-icon>
            <p class="day-show">
              {{ data.day.split("-").slice(2).join("") }}
            </p>
          </div>
        </template>
      </el-calendar>
    </el-col>
  </el-row>
  <el-dialog v-model="dialogVisible" title="新增训练内容" width="40%">
    <el-form :model="form" label-width="auto" style="max-width: 600px">
      <el-form-item label="训练内容">
        <el-input v-model="form.name"/>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="dialogVisible = false">Create</el-button>
        <el-button @click="dialogVisible = false">Cancel</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<script setup>
import {ref} from "vue";
import {reactive} from 'vue'
import {Plus} from "@element-plus/icons-vue";


const form = reactive({
  name: '',
})
const dialogVisible = ref(false)

const addItem = (data) => {
  console.log(data)
  dialogVisible.value = true

  console.log("is today?" + isToday(data))
}

const isToday = (calendarData) => {
  const today = new Date();
  const formattedDate = today.toISOString().slice(0, 10).replace(/-/g, '-');
  console.log(formattedDate);
  return formattedDate === calendarData.day
}

</script>

<style scoped>

.exercise-day-calendar {
  display: flex;
}

.day-show {
  align-self: center;
  position: relative;
  /* 居中 */
  margin:0 auto;
}

.add-item-icon {
  position: absolute;
  margin-left: 1px;
}
</style>