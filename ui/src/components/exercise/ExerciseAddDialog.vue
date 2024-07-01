<template>

  <el-dialog v-model="props.showDialog" @close="closeDialog" title="新增训练内容" width="650px">
    <el-form :model="form" label-width="auto" style="max-width: 600px">
      <el-form-item label="日期" class="form-item-left-label">
        <el-input v-model="props.date" placeholder="日期" disabled></el-input>
      </el-form-item>

      <el-form-item label="标签" class="form-item-left-label">
        <el-select
            v-model="form.tag"
            placeholder="请选择标签"
        >
          <el-option
              v-for="tag in tags"
              :key="tag.value"
              :label="tag.value"
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
        <el-button type="primary" @click="confirmAdd">创建</el-button>
        <el-button @click="closeDialog">取消</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<script setup>
import {reactive} from "vue";
import {createExercise} from "@/api/excercise/exercise";
import {ElMessage} from "element-plus";

// 传来的参数
const props = defineProps({
  showDialog: {
    type: Boolean,
    default: false
  },
  date: ''
})

// 接收父组件信息
const form = reactive({
  content: '',
  date: '',
  tag: '',
  duration: 30,
})

const tags = [
  {value: '跑步', label: '跑步'},
  {value: '游泳', label: '游泳'},
  {value: '胸部', label: '胸部'},
  {value: '腿部', label: '腿部'},
  {value: '背部', label: '背部'},
  {value: '肩部', label: '肩部'},
]

const emit = defineEmits(['updateDialog']);

const closeDialog = () => {
  emit('updateDialog', false);
}

const confirmAdd = () => {
  if (form.tag.trim() === '') {
    ElMessage.warning("请选择一项训练标签");
    return
  }

  if (form.content.trim() === '') {
    ElMessage.warning("内容不能为空");
    return
  }

  console.log('Form Data:', form);

  const param = {
    type: form.tag,
    description: form.content,
    duration: form.duration,
    date: props.date,
  }

  createExercise(param).then(res => {
    if (res.success) {
      ElMessage.success('添加成功！')
      closeDialog()
    } else {
      console.error('请求过程中发生错误:', res);
      ElMessage.error('添加失败');
    }
  })
}

</script>

<style scoped>
.form-item-left-label {
  text-align: left !important;
}
</style>