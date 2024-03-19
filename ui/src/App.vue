<template>
  <div :class="globalMode.mode">
    <el-config-provider :locale="zhCn">
      <router-view />
    </el-config-provider>
  </div>
</template>
<script setup>
import zhCn from 'element-plus/es/locale/lang/zh-cn';
import 'dayjs/locale/zh-cn';
import './assets/css/main.css';
import './assets/css/color-dark.css';
import './assets/css/color-light.css';
import {useConfigStore} from "@/stores/config";

const globalMode = useConfigStore()


</script>
