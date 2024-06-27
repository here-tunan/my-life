<template>
  <div class="day-items">
    <div>
      <div v-for="item in showItems">
        <p>{{ item.title }}</p>
      </div>
    </div>

    <div class="more-items-icon" v-if="showDropdownArrow" @click="toggleDropdown">
      {{ moreItemsNum }} more items
    </div>
  </div>
</template>

<script setup>

// 父组件传来的参数
import {computed, ref} from "vue";

const props = defineProps({
  //2024-04-12 格式
  day: String,
  dayItems: []
});

// 限制最多展示三个item
const maxShowNum = 3;

const itemsNum = computed(() => {
  return props.dayItems !== undefined ? props.dayItems.length : 0
})

// 是否展示下拉，大于三个时展示
const showDropdownArrow = computed(() => {
      return itemsNum > maxShowNum
    }
)

// 需要展示的数量
const showItemsNum = computed(() => {
  if (itemsNum > 0) {
    if (itemsNum <= maxShowNum) {
      return itemsNum
    } else {
      return itemsNum - maxShowNum - 1
    }
  }
  return 0;
})

const moreItemsNum = computed(() => {
      return itemsNum - maxShowNum
    }
)

const showItems = computed(() => {
  return showItemsNum > 0 ? props.dayItems.slice(0, showItemsNum) : []
})


const toggleDropdown = () => {
  // 弹出多个事项
  console.log("展示下拉")
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