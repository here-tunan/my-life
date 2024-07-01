<template>
  <div class="day-item" @click="itemDetailClick">
    <p>{{ item.description }}</p>
  </div>

  <ExerciseFormDialog :show-dialog="dialogVisible" :date="dialogDate" :item="props.item" @updateDialog="updateDialogVisible" />
</template>

<script setup>
import {reactive, ref} from "vue";
import ExerciseFormDialog from "@/components/exercise/ExerciseEditDialog.vue";

const props = defineProps({
  item: {}
});

const item = reactive(props.item)

// console.log(props.item)

// 弹出框信息
const dialogVisible = ref(false)
const dialogDate = ref('')

// 子组件的事件
const updateDialogVisible = isShow => {
  dialogVisible.value = isShow
}

const itemDetailClick = () => {
  dialogVisible.value = true
}

</script>

<style scoped>
.day-item p {
  font-size: 11px;
  margin-top: 1px;
  margin-bottom: 1px;
}

.day-item:hover {
  background-color: lightgray;
}
</style>