<template>
  <div class="container">
    <el-row>
      <el-form label-width="200px">
        <el-form-item>
          <el-button type="primary" @click="search">查询</el-button>
          <el-button type="primary" @click="addClick">新增</el-button>
        </el-form-item>
      </el-form>
    </el-row>
    <el-row>
      <el-table :data="accountData" stripe>
        <el-table-column prop="id" label="Id" width="180" v-if="false"/>
        <el-table-column prop="name" label="名称" width="180"/>
        <el-table-column prop="desc" label="描述" width="180"/>

        <el-table-column fixed="right" label="Operations" width="120">
          <template #default="scope">
            <el-button link type="primary" @click="editClick(scope.row)">Edit</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-row>
    <!-- dialog box -->
    <el-dialog v-model="dialogFormVisible" title="Transaction account">
      <el-form :model="accountInfo">
        <el-form-item label="name" :label-width="formLabelWidth">
          <el-input v-model="accountInfo.name" autocomplete="off" placeholder="Please input a name"/>
        </el-form-item>
        <el-form-item label="desc" :label-width="formLabelWidth">
          <el-input v-model="accountInfo.desc" autocomplete="off" placeholder="Please input a description"/>
        </el-form-item>
      </el-form>
      <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogBoxCancel">Cancel</el-button>
        <el-button type="primary" @click="dialogBoxConfirm">
          Confirm
        </el-button>
      </span>
      </template>
    </el-dialog>
  </div>
</template>


<script setup>
import {reactive, ref} from "vue";
import axios from "axios";

const accountData = ref([])


const formLabelWidth = '140px'


// open the dialog box
const dialogFormVisible = ref(false)

const accountInfo = reactive({
  id: '',
  name: '',
  desc: '',
})

const search = () => {
  doSearch()
}

const doSearch = async () => {
  const res = await axios.get("http://127.0.0.1:3000/api/money/transactionAccount");

  if (res.status === 200 && res.data.success) {
    let resList = res.data.data
    for (let one of resList) {
      if (one.type === 1) {
        one.typeName = "支出"
      }
    }
    accountData.value = resList
  } else {
    accountData.value = []
  }
}

const editClick = (row) => {
  accountInfo.id = row.id
  accountInfo.name = row.name
  accountInfo.desc = row.desc

  dialogFormVisible.value = true
}

const dialogBoxCancel = () => {
  closeDialogBox()
}

const dialogBoxConfirm = async () => {
  let param = {
    id: accountInfo.id,
    name: accountInfo.name,
    desc: accountInfo.desc,
  }
  if (param.id === '') {
    param.id = null
  }
  // 保存内容
  let res = await axios.put("http://127.0.0.1:3000/api/money/transactionAccount", param, {
    headers: {
      'Content-Type': 'application/json',
    },
  })
  if (res.status === 200 && res.data.success) {
    closeDialogBox()
    await doSearch()
  }
}

const addClick = () => {
  accountInfo.id = ''
  accountInfo.name = ''
  accountInfo.desc = ''
  dialogFormVisible.value = true
}

const closeDialogBox = () => {
  accountInfo.id = ''
  accountInfo.name = ''
  accountInfo.desc = ''
  dialogFormVisible.value = false
}


</script>

<style scoped>
.dialog-footer button:first-child {
  margin-right: 10px;
}
</style>