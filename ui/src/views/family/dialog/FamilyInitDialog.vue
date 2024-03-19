<template>
  <el-dialog v-model="props.showDialog" @close="closeDialog" draggable title="创建新家庭" width="800px">
    <div class="family-init-dialog">
      <el-steps :active="active" align-center>
        <el-step title="步骤1" description="给你的家庭起一个名字"/>
        <el-step title="步骤2" description="描述一下这个家庭"/>
        <el-step title="步骤3" description="一个酷酷的头像"/>
      </el-steps>

      <div style="height: 90px;">
        <el-form label-width="50px">
          <div v-show="active === 0">
            <el-form-item>
              <el-input class="family-name" v-model="form.name" placeholder="一个好听点的名字🥳"/>
            </el-form-item>
          </div>
          <div v-show="active === 1">
            <el-form-item>
              <el-input class="family-desc" v-model="form.desc" :autosize="{ minRows: 4, maxRows: 6 }" type="textarea"
                        placeholder="或许可以使用一个家庭宣言🧐"/>
            </el-form-item>
          </div>
          <div v-show="active === 2">
            <el-form-item>
              <div class="family-avatar">
                <AvatarUpload :avatar-img="form.avatar" @newAvatar="updateNewAvatar"/>
              </div>
            </el-form-item>
          </div>
        </el-form>
      </div>
    </div>
    <template #footer>
      <span class="dialog-footer">
        <el-button type="primary" style="margin-top: 22px" @click="preClick"
                   :disabled="buttonDisable.pre">上一步</el-button>
        <el-button type="primary" style="margin-top: 22px" @click="nextClick"
                   :disabled="buttonDisable.next">下一步</el-button>
        <el-button type="success" style="margin-top: 22px" @click="confirmClick"
                   :disabled="buttonDisable.confirm">确认</el-button>
      </span>
    </template>
  </el-dialog>
</template>
<script setup>

import {reactive, ref, watch} from "vue";
import {ElMessage, genFileId} from "element-plus";
import AvatarUpload from "@/components/AvatarUpload.vue";
import {createFamily} from "@/api/family/family";
import {useFamilyStore} from "@/stores/family";

const active = ref(0)
const upload = ref()


const buttonDisable = reactive({
  pre: true,
  next: false,
  confirm: true
})

const form = reactive({
  name: '',
  desc: '',
  avatar: 'https://yaodao-images.oss-cn-hangzhou.aliyuncs.com/my-vue-lif/1700552942808.jpg',
  familyAvatarUrl: '',
})

const props = defineProps({
  showDialog: {
    type: Boolean,
    default: false
  }
})

const updateNewAvatar = (newAvatar) => {
  form.avatar = newAvatar
  console.log("新头像URL：" + newAvatar)
}

const emit = defineEmits(['closeDialog']);

const handleExceed = (files) => {
  if (upload.value) {
    upload.value.clearFiles();
  }
  const file = files[0]
  file.uid = genFileId()
  if (upload.value) {
    upload.value.handleStart(file);
  }
}

const preClick = () => {
  active.value--
  if (active.value < 0) {
    active.value = 0
  }
}

const nextClick = () => {
  if (active.value === 0) {
    // 校验家庭名称
    if (form.name === '') {
      ElMessage.warning("请先请一个好听的名字")
      return
    } else if (form.name.length > 30) {
      ElMessage.warning("这么长的名字更适合作为描述，请不要超过30个字")
      return;
    }
  } else if (active.value === 1) {
    // 校验家庭名称
    if (form.desc === '') {
      ElMessage.warning("一个家庭值得有一段描述")
      return;
    } else if (form.desc.length > 200) {
      ElMessage.warning("描述有点长，尽可能用200个以内的字来能描述这个小家～")
      return;
    }
  }
  // active 最大是2
  active.value++
  if (active.value > 2) {
    active.value = 2
  }
  console.log(active.value)
}

const confirmClick = () => {
  // 完成
  createFamily(form).then((res) => {
    if (res.success) {
      ElMessage.success("成功创建" + form.name)
      useFamilyStore().id = res.data.id
      useFamilyStore().name = res.data.name
      useFamilyStore().desc = res.data.desc
      useFamilyStore().avatar = res.data.avatar
      active.value = 3
      closeDialog()
    }
  })
}

const closeDialog = () => {
  emit('closeDialog', false);
}

watch(active, () => {
  buttonDisable.pre = false
  buttonDisable.next = false
  buttonDisable.confirm = false
  if (active.value === 0) {
    buttonDisable.pre = true
  }
  if (active.value >= 2) {
    buttonDisable.next = true
  }
  if (active.value < 2) {
    buttonDisable.confirm = true
  }
})

</script>

<style scoped>
.family-name {
  padding-top: 70px;
}
.family-desc {
  padding-top: 40px;
}
.family-avatar {
  padding-top: 40px;
  padding-left: 278px;
}
</style>