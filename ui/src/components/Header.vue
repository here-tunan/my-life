<!--首页的Header控件-->
<template>
  <div class="header">
    <!--  折叠展开按钮  -->
    <div class="collapse-btn" @click="collapseChange">
      <el-icon v-if="sidebar.collapse">
        <Expand/>
      </el-icon>
      <el-icon v-else>
        <Fold/>
      </el-icon>
    </div>

    <div class="logo">This is my life!</div>

    <div class="header-right">
      <div class="header-user-con">
        <!-- 消息中心 小铃铛图标 -->
        <div class="btn-bell">
          <el-tooltip
              effect="dark"
              :content="message ? `有${message}条未读消息` : `消息中心`"
              placement="bottom"
          >
            <el-icon class="el-icon-notice"><Bell /></el-icon>
          </el-tooltip>
          <span class="btn-bell-badge" v-if="message"></span>
        </div>

        <div class="header-mode-toggle">
          <!-- 颜色模式切换 -->
          <el-switch
              v-model="darkOpen"
              :active-action-icon="Moon"
              :inactive-action-icon="Sunny"
              @change="modeToggle"
          />
        </div>

        <!-- 用户头像 -->
        <el-avatar class="user-avatar" :size="30" :src="avatarImg"/>

        <!-- 用户名下拉菜单 -->
        <el-dropdown class="user-name" trigger="click" @command="handleCommand">
					<span class="el-dropdown-link">
						{{ username }}
						<el-icon class="el-icon--right">
							<arrow-down/>
						</el-icon>
					</span>
          <template #dropdown>
            <el-dropdown-menu>
              <a href="https://gitee.com/yaodao666/vue-my-life" target="_blank">
                <el-dropdown-item>项目仓库</el-dropdown-item>
              </a>
              <el-dropdown-item command="user">个人中心</el-dropdown-item>
              <el-dropdown-item divided command="loginOut">退出登录</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
    </div>
  </div>
</template>

<script setup>
import {useSidebarStore} from "@/stores/sidebar";
import {useConfigStore} from "@/stores/config";
import {ArrowDown, Bell, Expand, Fold, Moon, Sunny} from "@element-plus/icons-vue";
import {computed, onMounted, ref} from "vue";
import {useRouter} from "vue-router";
import {getLoginUserInfo} from "@/api/user/user";
import {ElMessage} from "element-plus";
import {useUserStore} from "@/stores/user";

const sidebar = useSidebarStore()
const config = useConfigStore()

const avatarImg = computed(() => {
  return useUserStore().avatar
})

const username = computed(() => {
  return useUserStore().username
})
// const username = ref("yaodao")
const message = ref(2);

// 用户名下拉菜单选择事件
const router = useRouter();

// 默认关闭黑暗模式
const darkOpen = ref(true)

onMounted(() => {
  if (document.body.clientWidth < 1000) {
    collapseChange();
  }
});

onMounted(() => {
  getLoginUserInfo().then(
      (res) => {
        if (res.success) {
          useUserStore().account = res.data.account
          useUserStore().username = res.data.name
          useUserStore().desc = res.data.desc
          useUserStore().avatar = res.data.avatar
          localStorage.setItem("avatar", res.data.avatar)
        } else {
          ElMessage.error("获取用户信息失败")
        }
      }
  )
})

const modeToggle = () => {
  console.log(config.mode)
  config.toggle()
}

const handleCommand = (command) => {
  if (command === 'loginOut') {
    localStorage.removeItem('token');
    ElMessage.success('成功退出，请重新登录')
    router.push('/login');
  } else if (command === 'user') {
    router.push('/user');
  }
};

const collapseChange = () => {
  sidebar.handleCollapse();
};
</script>

<style scoped>
.header {
  position: relative;
  box-sizing: border-box;
  width: 100%;
  height: 70px;
  font-size: 22px;
  color: #fff;
}

.collapse-btn {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  float: left;
  padding: 0 21px;
  cursor: pointer;
}

.header .logo {
  float: left;
  width: 250px;
  line-height: 70px;
}

.header-right {
  float: right;
  padding-right: 50px;
}

.header-user-con {
  display: flex;
  height: 70px;
  align-items: center;
}

.header-mode-toggle {
  display: flex;
  height: 70px;
  align-items: center;
  margin-right: 10px;
}

.btn-fullscreen {
  transform: rotate(45deg);
  margin-right: 5px;
  font-size: 24px;
}

.btn-bell,
.btn-fullscreen {
  right: 10px;
  position: relative;
  width: 30px;
  height: 30px;
  text-align: center;
  border-radius: 18px;
  cursor: pointer;
  display: flex;
  align-items: center;
}

.btn-bell-badge {
  position: absolute;
  right: 4px;
  top: 0;
  width: 8px;
  height: 8px;
  border-radius: 4px;
  background: #f56c6c;
  color: #fff;
}

.btn-bell .el-icon-notice {
  color: #fff;
}

.user-name {
  margin-left: 10px;
}

.user-avator {
  margin-left: 20px;
}

.el-dropdown-link {
  color: #fff;
  cursor: pointer;
  display: flex;
  align-items: center;
}

.el-dropdown-menu__item {
  text-align: center;
}
</style>