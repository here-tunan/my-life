<template>

  <el-dialog v-model="props.showDialog" @close="closeDialog" title="修改训练内容" width="650px">
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
        <el-button type="primary" @click="save" style="margin-left: 60px">确认修改</el-button>
        <el-button @click="closeDialog">取消</el-button>
        <el-button type="danger" @click="deleteItem" style="margin-left: 280px">删除</el-button>
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
  item: {
    type: {},
    default: {
      id: '',
      title: '',
      date: '',
      tag: '',
      duration: '10',
    }
  }
})

// 接收父组件信息
const form = reactive({
  id: props.item.id,
  content: props.item.description,
  date: props.item.date,
  tag: props.item.tag,
  duration: props.item.duration,
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
  // form 中的内容恢复
  form.id = props.item.id;
  form.content = props.item.description;
  form.date = props.item.date;
  form.tag = props.item.tag;
  form.duration = props.item.duration;
}

const deleteItem = () => {
  ElMessage.warning("暂不支持")
}

const save = () => {
  if (form.tag.trim() !== '') {
    console.log('Form Data:', form);

    const param = {
      tag: form.tag,
      description: form.content,
      duration: form.duration,
      date: form.date,
    }

    createExercise(param).then(res => {
      if (res.success) {
        ElMessage.success('添加成功！请刷新页面')
        closeDialog()
      } else {
        console.error('请求过程中发生错误:', res);
        ElMessage.error('添加失败');
      }
    })
  } else {
    ElMessage.warning("请选择一项训练标签");
  }
}

</script>

<style scoped>
.form-item-left-label {
  text-align: left !important;
}
</style>