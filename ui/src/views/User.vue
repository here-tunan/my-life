<template>
  <div class="container">
    <el-container>
      <!--  侧边栏  -->
      <el-aside width="200px">
        <el-menu :default-active="activeIndex" @select="handleSelect">
          <el-menu-item index="1">
            <span>个人资料</span>
          </el-menu-item>
          <el-menu-item index="2">
            <span>密码修改</span>
          </el-menu-item>
        </el-menu>
      </el-aside>

      <el-main>
        <!-- 个人资料  -->
        <div class="personal-info" v-if="activeIndex === '1'">
          <h3>个人资料</h3>
          <el-divider/>
          <el-form :model="form" :rules="rules" ref="ruleRef" label-width="80px">
            <el-row>
              <el-col :span="14">
                <el-form-item label="账户id">
                  <el-input v-model="form.account" disabled/>
                </el-form-item>
                <el-form-item label="昵称" prop="name">
                  <el-input v-model="form.name"/>
                </el-form-item>
                <el-form-item label="描述" prop="desc">
                  <el-input v-model="form.desc" type="textarea" :autosize="{ minRows: 4, maxRows: 5 }"/>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <div style="padding-top: 20px">
                  <avatar-upload :avatar-img="form.avatar" @newAvatar="getNewAvatar"/>
                </div>
              </el-col>
            </el-row>

            <el-form-item>
              <el-button type="primary" @click="submitForm(ruleRef)">保存修改</el-button>
            </el-form-item>
          </el-form>
        </div>

        <!-- 账号设置 -->
        <div class="account-setting" v-if="activeIndex === '2'">
          <h3>密码修改</h3>
          <el-divider/>
          <el-form :model="passwordForm" :rules="passwordRules" ref="passwordRuleRef" label-width="80px">
            <el-form-item label="新密码" prop="first">
              <el-col :span="12">
                <el-input v-model="passwordForm.first" type="password" show-password/>
              </el-col>
            </el-form-item>
            <el-form-item label="确认密码" prop="second">
              <el-col :span="12">
                <el-input v-model="passwordForm.second" type="password" show-password/>
              </el-col>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="submitPasswordForm(passwordRuleRef)">更新密码</el-button>
            </el-form-item>
          </el-form>
        </div>
      </el-main>
    </el-container>
  </div>
</template>
<script setup>

import {reactive, ref, watch} from "vue";
import {useUserStore} from "@/stores/user";
import AvatarUpload from "@/components/AvatarUpload.vue";
import {putUser} from "@/api/user/user";
import {ElMessage} from "element-plus";

const activeIndex = ref("1")

// 如果直接让 form.avatar = avatar 会导致表单数据改变时，store中的值也改变，因为这是指向了同一个响应对象
// const { avatar, username, desc } = toRefs(useUserStore());

//  --- 个人资料表单 ---
const form = reactive({
  account: useUserStore().account,
  avatar: useUserStore().avatar,
  name: useUserStore().username,
  desc: useUserStore().desc
})

const rules = reactive({
  name: [
    {required: true, message: '请输入一个用户名', trigger: 'blur'},
    {min: 0, max: 20, message: '请不要超出20个字符', trigger: 'blur'},
  ],
  desc: [
    {min: 0, max: 300, message: '请不要超出300个字符', trigger: 'blur'},
  ],
})

const ruleRef = ref()

// --- 密码修改表单 ---
const passwordForm = reactive({
  first: '',
  second: ''
})

// 检查两个密码输入框是否一致(这个方法要声明在passwordRules之前)
const checkSamePassword = (rule, value, callback) => {
  if (passwordForm.first !== value) {
    return callback(new Error('请确认您两次输入的新密码保持一致'))
  } else {
    callback()
  }
}

const passwordRules = reactive({
  first: [
    {required: true, message: '密码不能为空', trigger: 'blur'},
  ],
  second: [
    {required: true, message: '密码不能为空', trigger: 'blur'},
    {validator: checkSamePassword, trigger: 'blur'},
  ],
})

const passwordRuleRef = ref()

// 监听userStore的变化
watch(useUserStore(), () => {
  form.account = useUserStore().account;
  form.avatar = useUserStore().avatar;
  form.name = useUserStore().username;
  form.desc = useUserStore().desc;
})

// 菜单选中
const handleSelect = (key) => {
  activeIndex.value = key
}

// 新头像上传的回调
const getNewAvatar = (newAvatar) => {
  form.avatar = newAvatar
}

// 个人资料修改确定
const submitForm = async (formEl) => {
  if (!formEl) return
  await formEl.validate((valid, fields) => {
    if (valid) {
      console.log('submit!')
      // 执行保存逻辑
      let param = {
        avatar: form.avatar,
        name: form.name,
        desc: form.desc
      }
      putUser(param).then((res) => {
        if (res.success) {
          ElMessage.success('成功更新个人资料！')
          // 更新store的值
          useUserStore().avatar = res.data.avatar;
          useUserStore().username = res.data.name;
          useUserStore().desc = res.data.desc;
        }
      })
    } else {
      console.log('error submit!', fields)
    }
  })

}

// 密码修改
const submitPasswordForm = async (formEl) => {
  if (!formEl) return
  await formEl.validate((valid, fields) => {
    if (valid) {
      console.log('submit!')
      // 执行保存逻辑
      let param = {
        password: passwordForm.first,
      }
      putUser(param).then((res) => {
        if (res.success) {
          ElMessage.success('成功更新密码！')
        }
      })
    } else {
      console.log('error submit!', fields)
    }
  })
}
</script>

<style scoped>
</style>