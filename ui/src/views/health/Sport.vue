<template>
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
  <el-dialog v-model="dialogVisible" title="新增训练内容" width="40%">
    <el-form :model="form" label-width="auto" style="max-width: 600px">
      <el-form-item label="日期" class="form-item-left-label">
        <el-input v-model="form.date" placeholder="日期" disabled></el-input>
      </el-form-item>

      <el-form-item label="标签" class="form-item-left-label">
        <el-select
            v-model="form.tag"
            placeholder="请选择标签"
        >
          <el-option
              v-for="tag in tags"
              :key="tag.value"
              :label="tag.label"
              :value="tag.value">
          </el-option>
        </el-select>
      </el-form-item>

      <el-form-item label="时长（分钟）" class="form-item-left-label">
        <el-input-number v-model="form.duration" :min="1" :max="1440" :step="10"></el-input-number>
      </el-form-item>

      <el-form-item label="内容" class="form-item-left-label">
        <el-input v-model="form.content" type="textarea" placeholder="请输入训练内容描述"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="confirmAdd">Create</el-button>
        <el-button @click="dialogVisible = false">Cancel</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<script setup>
import {computed, ref} from "vue";
import {reactive} from 'vue'
import {Plus} from "@element-plus/icons-vue";
import CalendarDayItems from "@/components/calendar/CalendarDayItems.vue";
import {createExercise} from "@/api/excercise/exercise";
import {ElMessage} from "element-plus";

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

const dialogVisible = ref(false)
const tags = [
  {value: 'run', label: '跑步'},
  {value: 'swim', label: '游泳'},
  {value: 'chest', label: '胸部'},
  {value: 'leg', label: '腿部'},
  {value: 'back', label: '背部'},
  {value: 'shoulder', label: '肩部'},
]

const confirmAdd = () => {
  if (form.tag.trim() !== '') {
    console.log('Form Data:', form);
    dialogVisible.value = false;

    const param = {
      type: form.tag,
      description: form.content,
      duration: form.duration,
      date: form.date,
    }

    createExercise(param).then(res =>{
      if (res.success){
        console.log('锻炼记录创建成功');
        alert('添加成功！')
        form.tag = '';
        form.duration = 30;
        form.content = '';
      }else {
        console.error('请求过程中发生错误:', error);
        alert('添加失败');
      }
    })
  } else {
    alert("请选择一项训练标签");
  }
}

// 点击新增item的按钮
const addItem = (selectDate) => {
  dialogVisible.value = true
  form.date = formatStrDate(new Date(selectDate.day));
}

</script>

<style scoped>
.form-item-left-label {
  text-align: left !important;
}

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
  position: absolute;
  margin-left: 1px;
}
</style>
