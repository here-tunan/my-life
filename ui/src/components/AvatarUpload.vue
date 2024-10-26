<template>
  <div class="avatar">
    <div class="avatar-image" @click="avatarClick">
      <el-avatar :size="100" :src="props.avatarImg"/>
      <span class="avatar-edit"><el-icon><CameraFilled/></el-icon></span>
    </div>
  </div>

  <el-dialog title="上传头像" v-model="showDialog" width="600px" draggable>
    <vue-cropper
        ref="cropper"
        :src="imgSrc"
        :ready="cropImage"
        :zoom="cropImage"
        :cropmove="cropImage"
        style="width: 100%; height: 400px"
    ></vue-cropper>

    <template #footer>
				<span class="dialog-footer">
					<el-button class="crop-demo-btn" type="primary">选择图片
						<input class="crop-input" type="file" name="image" accept="image/*" @change="setImage"/>
					</el-button>
					<el-button type="primary" @click="saveAvatar">裁剪完成</el-button>
				</span>
    </template>
  </el-dialog>
</template>

<script setup>
import VueCropper from 'vue-cropperjs';
import 'cropperjs/dist/cropper.css';
import {ref} from "vue";
import {ElMessage} from "element-plus";
import {CameraFilled} from "@element-plus/icons-vue";
import {getOssInfo, upload} from "@/api/fundation/oss";


const imgSrc = ref('');
const cropImg = ref('');
const cropper = ref('');
const showDialog = ref(false)

const props = defineProps({
  avatarImg: {
    type: String,
    default: 'https://yaodao-images.oss-cn-hangzhou.aliyuncs.com/my-vue-lif/1700552942808.jpg'
  }
})

const emit = defineEmits(['newAvatar']);

const setImage = (e) => {
  const file = e.target.files[0];
  if (!file.type.includes('image/')) {
    return;
  }
  let fileLmt = file.size / 1024 / 1024;
  if (fileLmt > 2) {
    ElMessage.warning("请传入少于2M的图片")
    return;
  }
  const reader = new FileReader();
  reader.onload = (event) => {
    showDialog.value = true;
    imgSrc.value = event.target.result;
    cropper.value && cropper.value.replace(event.target.result);
  };
  reader.readAsDataURL(file);
}

// 裁剪后
const cropImage = () => {
  cropImg.value = cropper.value.getCroppedCanvas().toDataURL();
}

// 确认上传这个图片
const saveAvatar = () => {
  let filename = new Date().getTime() + ".jpg"
  // 查询oss信息
  getOssInfo().then((res) => {
    if (res.success) {
      // 官网文档https://github.com/fengyuanchen/cropperjs#methods
      cropper.value.getCroppedCanvas().toBlob((data) => {
        let fileFormData = new FormData()
        // key表示上传到Bucket内的Object的完整路径，例如exampledir/exampleobject.txt，完整路径中不能包含Bucket名称。
        // filename表示待上传的本地文件名称。
        fileFormData.append('key', "my-vue-lif/" + filename);
        fileFormData.append('policy',res.data.policy)
        fileFormData.append('OSSAccessKeyId', res.data.accessId)
        fileFormData.append('success_action_status', '200')
        fileFormData.append('signature', res.data.signature)
        fileFormData.append('file', data)
        upload(res.data.host, fileFormData).then((result) => {
          // console.log(result)
          if (result.status === 200) {
            let newAvatar = res.data.host + "/my-vue-lif/" + filename
            emit('newAvatar', newAvatar);
          }
          showDialog.value= false
        })
      });
      return
    }
    ElMessage.error("上传失败")
  })

  ElMessage.success("上传成功")
}

// 点击头像，打开弹出框
const avatarClick = () => {
  showDialog.value = true;
  imgSrc.value = props.avatarImg;
}

</script>

<style scoped>
.avatar {
  text-align: center;
}

.avatar-image {
  position: relative;
  margin: auto;
  width: 100px;
  height: 100px;
  background: #f8f8f8;
  border: 1px solid #eee;
  border-radius: 50px;
  overflow: hidden;
}

.avatar-edit {
  display: flex;
  justify-content: center;
  align-items: center;
  position: absolute;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  opacity: 0;
  transition: opacity 0.3s ease;
}

.avatar-image i {
  color: #eee;
  font-size: 25px;
}

.avatar-image:hover .avatar-edit {
  opacity: 1;
}

.crop-demo-btn {
  position: relative;
}

.crop-input {
  position: absolute;
  width: 100px;
  height: 40px;
  left: 0;
  top: 0;
  opacity: 0;
  cursor: pointer;
}
</style>