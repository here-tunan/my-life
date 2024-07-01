<template>
  <!-- 日历 -->
  <el-calendar v-model="calendar">
    <template #date-cell="{ data }">
      <div class="day-calendar">
        <el-icon v-if="data.isSelected" class="add-item-icon" @click="addItem(data)">
          <Plus/>
        </el-icon>
        <p class="day-show">
          {{ data.day.split("-").slice(2).join("") }}
        </p>
      </div>
      <div class="day-items">
        <CalendarDayItems :day-items=exerciseItemsByEveryDay.get(data.day) />
      </div>
    </template>
  </el-calendar>

  <!-- 弹窗 -->
  <ExerciseAddDialog :show-dialog="dialogVisible" :date="formatStrDate(calendar)" @updateDialog="updateDialogVisible"/>
</template>

<script setup>
import {computed, onMounted, ref, watch} from "vue";
import {reactive} from 'vue'
import {Plus} from "@element-plus/icons-vue";
import CalendarDayItems from "@/components/calendar/DayItems.vue";
import {listExercise} from "@/api/excercise/exercise";
import ExerciseAddDialog from "@/components/exercise/ExerciseAddDialog.vue";

// 当前的日期对应的日历数据
const exerciseLogs = ref([])

const calendar = ref(new Date())

// 初始化
onMounted(() => {
  getExerciseLogs(calendar)
})

watch(calendar, (cur, old) => {
  if (cur.getMonth() !== old.getMonth()) {
    console.log("cur", cur)
    console.log("old", old)
    getExerciseLogs(cur)
  }
})

// 获取当前日期前后各40天的数据
const getExerciseLogs = () => {
  let startDate = new Date()
  startDate.setDate(calendar.value.getDate() - 40)
  let endDate = new Date()
  endDate.setDate(calendar.value.getDate() + 40)
  let param = {
    startDate: formatStrDate(startDate),
    endDate: formatStrDate(endDate),
  }
  console.log(param)
  listExercise(param).then(res => {
    if (res.success) {
      console.log(res.data)
      exerciseLogs.value = res.data
      // 刷新页面
      getExerciseLogs(calendar)
    }
  })
}


const formatStrDate = (date) => {
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  return `${year}-${month}-${day}`;
}

// 弹出框信息
const dialogVisible = ref(false)
const newItem = reactive({
  description: '',
  date: '',
  duration: 30,
})
const updateDialogVisible = isShow => {
  dialogVisible.value = isShow
}

// 点击新增item的按钮
const addItem = (selectDate) => {
  newItem.date = formatStrDate(new Date(selectDate.day));
  dialogVisible.value = true
}

// 获取所有数据的日期分组
const exerciseItemsByEveryDay = computed(() => {
  const groupedMap = new Map();
  if (exerciseLogs.value === null) {
    return new Map()
  }
  exerciseLogs.value.forEach(log => {
    const date = log.date;
    if (groupedMap.has(date)) {
      groupedMap.get(date).push(log);
    } else {
      groupedMap.set(date, [log]);
    }
  });
  return groupedMap;
});

</script>

<style scoped>

.day-calendar {
  display: flex;
}

.day-show {
  align-self: center;
  position: relative;
  /* 居中 */
  margin: 0 auto;
}

.add-item-icon {
  position: sticky;
  margin-left: 1px;
}
</style>
