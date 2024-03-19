<template>
  <div>
    <TransactionStatistic v-if="dataIsReady" :familyId="familyInfo.id"/>
  </div>
</template>

<script setup>

import TransactionStatistic from "@/components/transaction/TransactionStatistic.vue";
import {onMounted, reactive, ref} from "vue";
import {getFamily} from "@/api/family/family";


const familyInfo = reactive({
  id: 0,
  name: '',
  desc: '',
  avatar: 'https://yaodao-images.oss-cn-hangzhou.aliyuncs.com/blog/edOdOEi751c.jpg',
  members: []
})

const dataIsReady = ref(false)

onMounted(() => {
      getFamily().then((res) => {
        if (res.success) {
          familyInfo.id = res.data.id
          familyInfo.name = res.data.name
          familyInfo.desc = res.data.desc
          familyInfo.avatar = res.data.avatar
          familyInfo.members = res.data.members
        }
        // 父组件数据加载完成后才加载子组件
        dataIsReady.value = true
      })
    }
)

</script>