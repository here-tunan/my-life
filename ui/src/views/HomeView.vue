<!-- 首页页面 -->
<template>
  <!-- 头部 -->
  <Header/>
  <!-- 侧边栏 -->
  <Sidebar/>

  <!-- 内容 -->
  <div class="content-box" :class="{ 'content-collapse': sidebar.collapse }">

    <!--  标签  -->
    <Tags/>

    <!-- 具体内容 -->
    <div class="content">
      <router-view v-slot="{ Component }">
        <!--        <transition name="move" mode="out-in">-->
        <!--        <keep-alive :include="tags.nameList">-->
        <keep-alive>
          <component :is="Component"></component>
        </keep-alive>
        <!--        </transition>-->
      </router-view>
    </div>
  </div>

</template>
<script setup>
import Header from "@/components/Header.vue";
import Sidebar from "@/components/Sidebar.vue";
import Tags from "@/components/Tags.vue";

import {useSidebarStore} from '@/stores/sidebar'
import {useTagsStore} from '@/stores/tags';

const sidebar = useSidebarStore()
const tags = useTagsStore()

</script>