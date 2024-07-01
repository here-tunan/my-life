<template>
  <!-- 日历 -->
  <el-calendar>
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
  <ExerciseFormDialog :show-dialog="dialogVisible" :date="dialogDate" @updateDialog="updateDialogVisible"/>
</template>

<script setup>
import {computed, ref} from "vue";
import {reactive} from 'vue'
import {Plus} from "@element-plus/icons-vue";
import CalendarDayItems from "@/components/calendar/CalendarDayItems.vue";
import ExerciseFormDialog from "@/components/health/ExerciseFormDialog.vue";


const getCurrentDateFormatted = () => {
  const now = new Date();
  return formatStrDate(now);
}

const formatStrDate = (date) => {
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  return `${year}-${month}-${day}`;
}

// 弹出框信息
const dialogVisible = ref(false)
const dialogDate = ref('')
const updateDialogVisible = isShow => {
  dialogVisible.value = isShow
}

const form = reactive({
  date: getCurrentDateFormatted(),
  tag: '',
  duration: 30,
  content: ''
})

const exerciseLogs = ref([
  {title: '打打太极拳', date: '2024-07-01'},
  {title: '篮球训练三小时', date: '2024-07-13'},
  {title: '游泳训练三小时', date: '2024-07-18'},
  {title: '啦啦啦1小时', date: '2024-07-18'},
  {title: '爬山老和山', date: '2024-07-18'},
  {title: '今天游泳了', date: '2024-07-16'},
  {title: '羽毛球', date: '2024-07-16'},
  {title: '健身 练腿日', date: '2024-07-16'},
  {title: 'Happy Day', date: '2024-07-16'},
  {title: 'nothing', date: '2024-07-16'},
  {title: '特殊的日期展示', date: '2024-06-26'},
])

const exerciseItemsByEveryDay = computed(() => {
  const groupedMap = new Map();
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


// 点击新增item的按钮
const addItem = (selectDate) => {
  dialogVisible.value = true
  dialogDate.value = formatStrDate(new Date(selectDate.day));
}

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
