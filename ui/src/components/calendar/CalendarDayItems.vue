<template>
  <div class="day-items">
    <div>
      <div v-for="item in limitedItems">
        <p>{{ item.title }}</p>
      </div>
    </div>

    <div class="more-items-icon" v-if="showDropdownArrow" @click="toggleDropdown">
      <el-icon>
        <ArrowDown/>
      </el-icon>
    </div>
  </div>
</template>

<script setup>

// 父组件传来的参数
import {computed} from "vue";
import {ArrowDown} from "@element-plus/icons-vue";

const props = defineProps({
  //2024-04-12 格式
  day: String,
  dayItems: []
});

// 限制最多展示两个item
const limit = 2;

// 是否展示下拉
const showDropdownArrow = computed(() => {
      return props.dayItems !== undefined && props.dayItems.length > limit
    }
)

const limitedItems = computed(() => {
      return props.dayItems !== undefined ? props.dayItems.slice(0, limit) : [];
    }
)

const toggleDropdown = () => {
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


</style>