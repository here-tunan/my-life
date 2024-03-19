<template>
  <el-card>
    <template #header>
      <div class="card-header" style="height: 10px;">
        <span>🌱 我的信息</span>
        <el-button class="button" @click="toUserView">设置</el-button>
      </div>
    </template>

    <!-- 用户信息 -->
    <div class="user-info">
      <el-avatar :size="120" :src="avatarImg"></el-avatar>
      <div class="user-info-cont">
        <div class="user-info-name">{{ username }}</div>
        <div class="user-info-motto">{{ userDesc }}</div>
      </div>
    </div>
  </el-card>
</template>

<script setup>

import {computed} from "vue";
import {useUserStore} from "@/stores/user";
import {useFamilyStore} from "@/stores/family";
import router from "@/router";

const username = computed(() => {
  return useUserStore().username
})

const userDesc = computed(() => {
  return useUserStore().desc
})

const avatarImg = computed(() => {
  return useUserStore().avatar
})

const toUserView = () => {
  router.push('/user')
}
</script>


<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-family: "Wawati SC", sans-serif;
  font-size: 15px;
}

.user-info {
  display: flex;
  align-items: center;
  padding-bottom: 20px;
  border-bottom: 2px solid #ccc;
  margin-bottom: 20px;
}

.user-info-cont {
  padding-left: 50px;
  flex: 1;
}

.user-info-cont .user-info-name {
  font-size: 30px;
  font-family: "Wawati SC", sans-serif;
  color: #000000;
}

.user-info-cont .user-info-motto {
  font-size: 12px;
  font-family: "Weibei SC", sans-serif;
  color: #b7b7ad;
}

.card-header .button {
  font-family: "Wawati SC", sans-serif;
  font-size: 15px;
  color: #0d84ff;
}
</style>