<!-- 家庭页面 -->
<template>
  <el-card>
    <template #header>
      <div class="card-header" style="height: 10px;">
        <span>🏡 我的家庭</span>
        <el-button class="button" @click="initFamily" text v-if="!useFamilyStore().id">创建家庭</el-button>
      </div>
    </template>

    <div class="family-info">
      <el-avatar :size="120" :src="useFamilyStore().avatar"/>
      <div class="family-info-content">
        <div class="family-info-name">{{ useFamilyStore().name }}</div>
        <div class="family-info-desc">{{ useFamilyStore().desc }}</div>
      </div>
    </div>
  </el-card>

  <FamilyInitDialog :showDialog="showFamilyInitDialog" @closeDialog="closeFamilyInitDialog"/>
</template>

<script setup>

import {onMounted, reactive, ref} from "vue";
import {getFamily} from "@/api/family/family";
import FamilyInitDialog from "@/views/family/dialog/FamilyInitDialog.vue";
import {useFamilyStore} from "@/stores/family";

const showFamilyInitDialog = ref(false)

const initFamily = () => {
  showFamilyInitDialog.value = true
}

const closeFamilyInitDialog = (isShow) => {
  showFamilyInitDialog.value = isShow
}

onMounted(() => {
      getFamily().then((res) => {
        if (res.success) {
          useFamilyStore().id = res.data.id
          useFamilyStore().name = res.data.name
          useFamilyStore().desc = res.data.desc
          useFamilyStore().avatar = res.data.avatar
        }
      })
    }
)
</script>

<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-family: "Wawati SC", sans-serif;
  font-size: 15px;
}

.card-header .button {
  font-family: "Wawati SC", sans-serif;
  font-size: 15px;
  color: #0d84ff;
}

.family-info {
  display: flex;
  align-items: center;
  padding-bottom: 20px;
  border-bottom: 2px solid #ccc;
  margin-bottom: 20px;
}

.family-info-content {
  padding-left: 50px;
  flex: 1;
}

.family-info-content .family-info-name {
  font-size: 30px;
  font-family: "Wawati SC", sans-serif;
  color: #000000;
}

.family-info-content .family-info-desc {
  font-size: 12px;
  font-family: "Weibei SC", sans-serif;
  color: #b7b7ad;
}

.family-info-members {
  padding-top: 10px;
  display: flex;
  flex-wrap: wrap;
}

.member-avatar {
  overflow: hidden;
  text-overflow: ellipsis; /* 如果需要文字隐藏时显示省略号 */
  white-space: nowrap; /* 如果需要文字隐藏时显示省略号 */
}
</style>