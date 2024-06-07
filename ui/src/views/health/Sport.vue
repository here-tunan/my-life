<template>
  <el-row :gutter="20">
    <el-col :span="10">
      <div class="date-picker">
        <div class="block">
          <span class="demonstration">健身记录:</span>
          <el-date-picker v-model="value2" type="daterange" unlink-panels range-separator="To"
            start-placeholder="Start date" end-placeholder="End date" :shortcuts="shortcuts" :size="size" />
        </div>
        <div>
          <ECharts :width="width" :height="height" :options="chartOptions" />
        </div>
      </div>
    </el-col>
    <el-col :span="14">
      <div class="collapse">
        <span class="collapseSpan">相关文章：</span>
        <el-collapse v-model="activeName" accordion>
          <el-collapse-item title="腿部训练" name="1">
            <div class="train-a">
              <a><router-link to="/leg-train"> Consistent with real life: in line with the process and logic of
              real life, and comply with languages and habits that the users are
              used to;</router-link></a>
            </div>
          </el-collapse-item>
          <el-collapse-item title="肩部训练" name="2">
            <div class="train-a">
              <a>肩部训练</a>
            </div>
          </el-collapse-item>
          <el-collapse-item title="胸部训练" name="3">
            <div class="train-a">
              <a><router-link to="/chest-train">练胸啊，胸弟</router-link></a>
            </div>
          </el-collapse-item>
          <el-collapse-item title="背部训练" name="4">
            <div class="train-a">
              <a>背部训练</a>
            </div>
          </el-collapse-item>
          <el-collapse-item title="腰部训练" name="5">
            <div class="train-a">
             <a>腰部训练</a>
            </div>
          </el-collapse-item>
        </el-collapse>
      </div>
    </el-col>
  </el-row>
  <el-row>
    <el-col :span="24">
      <div class="trainPlan">
        <span class="stration">训练计划:</span>
        <el-calendar>
          <template #date-cell="{ data }">
            <p :class="data.isSelected ? 'is-selected' : ''">
              {{ data.day.split("-").slice(1).join("-") }}
              {{ data.isSelected ? "✔️" : "" }}
            </p>
            <el-icon style="border: 1px solid #ddd;border-radius: 5px;" id="icon-button" @click="dialogVisible = true">
              <Plus />
            </el-icon>
          </template>
        </el-calendar>
      </div>
    </el-col>
  </el-row>
  <el-dialog v-model="dialogVisible" title="新增训练内容" width="40%">
    <el-form :model="form" label-width="auto" style="max-width: 600px">
      <el-form-item label="训练内容">
        <el-input v-model="form.name" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="dialogVisible = false">Create</el-button>
        <el-button @click="dialogVisible = false">Cancel</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<script setup>
import { ref } from "vue";
import { reactive } from 'vue'
import ECharts from '@/components/ECharts.vue'

const width = ref('400px');
const height = ref('200px');
const chartOptions = reactive({
  tooltip: {
    trigger: 'item'
  },
  series: [
    {
      type: 'pie',
      data: [
        {
          value: 335,
          name: '腿部训练'
        },
        {
          value: 234,
          name: '肩部训练'
        },
        {
          value: 548,
          name: '背部训练'
        },
        {
          value: 156,
          name: '腰部训练'
        },
        {
          value: 183,
          name: '胸部训练'
        }
      ]
    }
  ]
});

const form = reactive({
  name: '',
})
const dialogVisible = ref(false)

const onSubmit = () => {
  console.log('submit!')
}

const value2 = ref("");

const shortcuts = [
  {
    text: "Last week",
    value: () => {
      const end = new Date();
      const start = new Date();
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 7);
      return [start, end];
    },
  },
  {
    text: "Last month",
    value: () => {
      const end = new Date();
      const start = new Date();
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 30);
      return [start, end];
    },
  },
  {
    text: "Last 3 months",
    value: () => {
      const end = new Date();
      const start = new Date();
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 90);
      return [start, end];
    },
  },
];

//  const activeName = ref("1")
</script>

<style scoped>

a {
  color: rgb(93, 16, 238);
}
a:hover {
  color: #1989fa;
}

.is-selected {
  color: #1989fa;
}

.el-row {
  margin-bottom: 20px;
}

.trainPlan .stration {
  display: block;
  padding: 0 0 20px 0;
  text-align: left;
  font-size: 20px;
  flex: 1;
}

.collapse .collapseSpan {
  display: block;
  padding: 0 0 20px 0;
  text-align: left;
  font-size: 20px;
  flex: 1;
}

.date-picker {
  padding: 0;
  flex-wrap: wrap;
}

.date-picker .block {
  padding: 0 0 20px 0;
  text-align: left;
  border-right: solid 1px var(--el-border-color);
  flex: 1;
}

.date-picker .block:last-child {
  border-right: none;
}

.date-picker .demonstration {
  display: block;
  color: black;
  font-size: 20px;
  margin-bottom: 20px;
}
.train-a {
  display: flex;
  padding-left: 8px;
}
</style>
<style>
.el-collapse-item__header {
  padding-left: 5px;
  font-size: 16px;
}
</style>