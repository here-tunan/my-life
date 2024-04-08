<template>
  <div class="login-page">
    <el-card shadow="always" class="card-container">
      <el-row>
        <el-col :span="10">
          <img src="@/assets/img/gaoda.jpg" alt="小麻最美丽" class="login-left-img"/>
        </el-col>
        <el-col :span="14">
          <div class="login-right-form">
            <div class="login-text">
              <h2>欢迎登录</h2>
            </div>

            <el-form>
              <el-form-item label="Username">
                <el-input v-model="username" placeholder="Enter your username"></el-input>
              </el-form-item>
              <el-form-item label="Password">
                <el-input v-model="password" type="password" placeholder="Enter your password" @keyup.enter.native="enterClick"></el-input>
              </el-form-item>
              <el-button type="primary" @click="loginClick">Login</el-button>
            </el-form>
          </div>
        </el-col>
      </el-row>


    </el-card>
  </div>
</template>

<script setup>
import {ref} from "vue";
import {ElMessage} from "element-plus";
import {login} from "@/api/auth/auth";
import router from "@/router";

const username = ref('')
const password = ref('')

const loginClick = () => {
  // 处理登录逻辑
  if (username.value === '' || password.value === '') {
    ElMessage.warning("账号密码均不能为空")
    return
  }

  login(username.value, password.value).then(res => {
    if (res.success) {
      ElMessage.success('成功登录！')
      // 保存token
      localStorage.setItem("token", res.data.token)
      let expiredSeconds = res.data.expiredTime
      let now = new Date()
      console.log("now time:", now)
      now.setTime(now.getTime() + expiredSeconds * 1000)
      localStorage.setItem("validTime", now.getTime())
      console.log("token validTime:", localStorage.getItem("validTime"))
      // 跳转到首页
      router.push('/')
    } else {
      ElMessage.error(res.error)
    }
  })
}

const enterClick = () => {
  loginClick()
}
</script>

<style scoped>
.login-page {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-image: url('@/assets/img/login_bg_1.jpg');
  background-size: cover;
  background-position: center;
}

.card-container {
  width: 600px;
  height: 270px;
  opacity: 0.9;
  animation: pulse 5s infinite;
  border-radius: 75px; /* 设置圆角 */
  display: flex;
  flex-direction: row; /* 横向居中 */
  align-items: center;
}

@keyframes pulse {
  0% {
    transform: scale(1);
  }
  50% {
    transform: scale(1.02);
  }
  100% {
    transform: scale(1);
  }
}

.login-text {
  padding-bottom: 18px;
}

.login-left-img {
  width: 80%;
  border-radius: 46px;
  margin-left: 40px;
}

.login-right-form {
  display: flex;
  flex-direction: column;
  align-items: center;
}


.el-button {
  width: 100%;
}
</style>
