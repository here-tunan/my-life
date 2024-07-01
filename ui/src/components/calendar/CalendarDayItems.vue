<template>
  <div class="day-items">
    <div>
      <div v-for="item in showItems">
        <DayItem :item="item"/>
      </div>
    </div>

    <div class="more-items-icon" v-if="showDropdownArrow" @click="clickMoreItems">
      {{ moreItemsNum }} more items
    </div>

    <el-dialog
        class="more-items-dialog"
        v-model="dialogVisible"
        width="300"
    >
      <div v-for="item in props.dayItems">
        <DayItem :item="item"/>
      </div>
    </el-dialog>
  </div>
</template>

<script setup>

// 父组件传来的参数
import {computed, ref} from "vue";
import DayItem from "@/components/calendar/DayItem.vue";

const props = defineProps({
  //2024-04-12 格式
  day: String,
  dayItems: []
});

const dialogVisible = ref(false)

// 限制最多展示三个item
const maxShowNum = 3;

const itemsNum = computed(() => {
  return props.dayItems !== undefined ? props.dayItems.length : 0
})

// 是否展示下拉，大于三个时展示
const showDropdownArrow = computed(() => {
      return itemsNum.value > maxShowNum
    }
)

// 需要展示的数量
const showItemsNum = computed(() => {
  if (itemsNum.value > 0) {
    if (itemsNum.value <= maxShowNum) {
      return itemsNum.value
    } else {
      return maxShowNum - 1
    }
  }
  return 0;
})

const moreItemsNum = computed(() => {
      return itemsNum.value - showItemsNum.value
    }
)

const showItems = computed(() => {
  return showItemsNum.value > 0 ? props.dayItems.slice(0, showItemsNum.value) : []
})


const clickMoreItems = () => {
  // 弹出多个事项
  console.log("展示下拉")
  dialogVisible.value = true

}

</script>

<style scoped>
.day-items {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.day-items {
  /*让div内部文字居中*/
  text-align: center;
}

.day-items p {
  font-size: 11px;
}

.more-items-icon {
  font-size: 11px;
  color: #918f8f;
}

</style>

<style>
.more-items-dialog {
  border-radius: 20px;
  background-color: rgb(228, 228, 248);
}
</style>