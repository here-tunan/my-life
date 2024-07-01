<template>
  <div class="day-item" @click="itemDetailClick">
    <p>{{ props.item.title }}</p>
  </div>

  <ExerciseFormDialog :show-dialog="dialogVisible" :date="dialogDate" item="props.item" @updateDialog="updateDialogVisible" />
</template>

<script setup>
import {ref} from "vue";
import ExerciseFormDialog from "@/components/health/ExerciseFormDialog.vue";

const props = defineProps({
  item: {}
});

// 弹出框信息
const dialogVisible = ref(false)
const dialogDate = ref('')
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