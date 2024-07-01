<template>
  <div class="day-item" @click="itemDetailClick">
    <p>{{ props.item.title }}</p>
  </div>
</template>

<script setup>
const props = defineProps({
  item: {}
});

const itemDetailClick = () => {
  // 打开item的弹窗
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