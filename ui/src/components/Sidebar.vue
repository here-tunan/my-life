<template>
  <div class="sidebar">
    <!--  菜单  -->
    <el-menu
        class="sidebar-el-menu"
        :default-active="onRoutes"
        :collapse="sidebar.collapse"
        background-color="#324157"
        text-color="#bfcbd9"
        active-text-color="#20a0ff"
        router
    >
      <template v-for="item in items">
        <!--  有子菜单  -->
        <el-sub-menu v-if="item.subs" :index="item.index" :key="item.index">
          <template #title>
            <el-icon>
              <component :is="item.icon"></component>
            </el-icon>
            <span>{{ item.title }}</span>
          </template>
          <!-- 二级目录 -->
          <template v-for="subItem in item.subs">
            <el-sub-menu v-if="subItem.subs" :index="subItem.index" :key="subItem.index">
              <template #title>{{ subItem.title }}</template>
              <el-menu-item v-for="(threeItem, i) in subItem.subs" :key="i" :index="threeItem.index">
                {{ threeItem.title }}
              </el-menu-item>
            </el-sub-menu>
            <el-menu-item v-else :index="subItem.index">
              <el-icon v-if="subItem.icon">
                <component :is="subItem.icon"></component>
              </el-icon>
              {{ subItem.title }}
            </el-menu-item>
          </template>
        </el-sub-menu>

        <!-- 无子菜单 -->
        <el-menu-item v-else :index="item.index" :key="item.index">
          <el-icon>
            <component :is="item.icon"></component>
          </el-icon>
          <template #title>{{ item.title }}</template>
        </el-menu-item>
      </template>
    </el-menu>
  </div>
</template>

<script setup>
import {computed} from 'vue';
import {useSidebarStore} from '@/stores/sidebar';
import {useRoute} from 'vue-router';

const items = [
  {
    icon: 'Help',
    index: '/dashboard',
    title: '系统首页',
  },
  {
    icon: 'Notebook',
    index: '1',
    title: '我的账本',
    subs: [
      {
        index: '/transaction-record',
        icon: 'EditPen',
        title: '我要记账',
      },
      {
        index: '/my-bill',
        icon: 'Odometer',
        title: '账单统计',
      }
    ],
  },
  {
    icon: 'Basketball',
    title: '健康生活',
    index: '2',
    subs: [
      {
        index: '/balance-diet',
        icon: 'Apple',
        title: '均衡膳食',
      },
      {
        index: '/sport',
        icon: 'Football',
        title: '运动健康',
      },
      {
        index: '/uuuu',
        icon: 'MoonNight',
        title: '睡眠健康',
      },
      {
        index: '/weight',
        icon: 'Chicken',
        title: '体重管理',
      },
    ]
  },
  {
    icon: 'CoffeeCup',
    index: '3',
    title: '我的家庭',
    subs: [
      {
        index: '/family',
        icon: 'HomeFilled',
        title: '温馨小家',
      },
      {
        index: '/family-bill',
        icon: 'Odometer',
        title: '家庭账单',
      }
    ]
  },
];

const route = useRoute();

// 返回路由
const onRoutes = computed(() => {
  return route.path;
});

const sidebar = useSidebarStore();
</script>

<style scoped>
.sidebar {
  display: block;
  position: absolute;
  left: 0;
  top: 70px;
  bottom: 0;
  overflow-y: scroll;
}

.sidebar::-webkit-scrollbar {
  width: 0;
}

.sidebar-el-menu:not(.el-menu--collapse) {
  width: 250px;
}

.sidebar > ul {
  height: 100%;
}
</style>
