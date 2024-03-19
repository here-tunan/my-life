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
      <el-table :data="categoryData" stripe>
        <el-table-column prop="id" label="Id" width="180" v-if="false"/>
        <el-table-column prop="name" label="名称" width="180"/>
        <el-table-column prop="type" label="类型" width="180" v-if="false"/>
        <el-table-column prop="typeName" label="类型" width="180"/>
        <el-table-column prop="desc" label="描述" width="380"/>

        <el-table-column fixed="right" label="Operations" width="120">
          <template #default="scope">
            <el-button link type="primary" @click="editClick(scope.row)">Edit</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-row>

    <!-- dialog box -->
    <el-dialog v-model="dialogFormVisible" title="Transaction category">
      <el-form :model="categoryInfo">
        <el-form-item label="name" :label-width="formLabelWidth">
          <el-input v-model="categoryInfo.name" autocomplete="off" placeholder="Please input a name"/>
        </el-form-item>
        <el-form-item label="desc" :label-width="formLabelWidth">
          <el-input v-model="categoryInfo.desc" autocomplete="off" placeholder="Please input a description"/>
        </el-form-item>
        <el-form-item label="type" :label-width="formLabelWidth">
          <el-select v-model="categoryTypeValue" value-key="id" placeholder="Please select a type">
            <el-option
                v-for="item in categoryTypes"
                :key="item.id"
                :label="item.name"
                :value="item"
            />
          </el-select>
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

const categoryData = ref([])
const categoryTypeValue = ref({
  id: '',
  name: ''
})

const formLabelWidth = '140px'


// open the dialog box
const dialogFormVisible = ref(false)

const categoryInfo = reactive({
  id: '',
  name: '',
  desc: '',
})

const categoryTypes = [
  {id: 1, name: "支出"},
  {id: 2, name: "收入"}
]

const search = () => {
  doSearch()
}

const doSearch = async () => {
  const res = await axios.get("http://127.0.0.1:3000/api/money/transactionCategory");

  if (res.status === 200 && res.data.success) {
    let resList = res.data.data
    for (let one of resList) {
      if (one.type === 1) {
        one.typeName = "支出"
      }
      if (one.type === 2) {
        one.typeName = "收入"
      }
    }
    categoryData.value = resList
  } else {
    categoryData.value = []
  }
}

const editClick = (row) => {
  categoryInfo.id = row.id
  categoryInfo.name = row.name
  categoryInfo.desc = row.desc
  categoryTypeValue.value.id = row.type

  dialogFormVisible.value = true
}

const dialogBoxCancel = () => {
  closeDialogBox()
}

const dialogBoxConfirm = async () => {
  let param = {
    id: categoryInfo.id,
    name: categoryInfo.name,
    desc: categoryInfo.desc,
    type: categoryTypeValue.value.id
  }

  if (param.id === '') {
    param.id = null
  }

  // 保存内容
  let res = await axios.put("http://127.0.0.1:3000/api/money/transactionCategory", param, {
    headers: {
      'Content-Type':
          'application/json',
    },
  })
  if (res.status === 200 && res.data.success) {
    closeDialogBox()
    await doSearch()
  }
}

const addClick = () => {
  categoryInfo.id = ''
  categoryInfo.name = ''
  categoryInfo.desc = ''
  categoryTypeValue.value.id = ''
  dialogFormVisible.value = true
}

const closeDialogBox = () => {
  categoryInfo.id = ''
  categoryInfo.name = ''
  categoryInfo.desc = ''
  categoryTypeValue.value.id = ''
  dialogFormVisible.value = false
}


</script>

<style scoped>


.dialog-footer button:first-child {
  margin-right: 10px;
}
</style>