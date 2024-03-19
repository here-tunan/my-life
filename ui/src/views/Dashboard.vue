<template>
  <div>
    <el-row :gutter="20">
      <el-col :span="12">
        <UserCard/>
      </el-col>
      <el-col :span="12">
        <FamilyCard/>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import FamilyCard from "@/components/family/FamilyCard.vue";
import UserCard from "@/components/user/UserCard.vue";
import {ref} from "vue";

const hasFamily = ref(false)


</script>

