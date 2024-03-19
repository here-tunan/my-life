<template>
  <el-dialog v-model="props.showDialog" @close="closeDialog" draggable title="Excel 记账" width="1000px">
    <div>
      <el-upload
          drag
          action=""
          ref="upload"
          class="upload-demo"
          accept=".csv"
          :limit="1"
          :on-exceed="handleExceed"
          :auto-upload="false"
          :on-change="handleChange"
          :on-remove="handleRemove"
      >
        <el-icon class="el-icon--upload">
          <upload-filled/>
        </el-icon>
        <div class="el-upload__text">
          将文件拖到此处，或<em>点击上传</em></div>
        <template #tip>
          <div class="el-upload__tip text-red">
            只能上传.csv文件，且编码为utf-8 (使用Txt等编辑器打开csv文件，发现乱码说明编码方式有问题，只需要使用 Excel
            应用打开该文件然后另存为——选择保存为.csv文件即可)
          </div>
        </template>
      </el-upload>
    </div>

    <!-- 弹出框中的表格   -->
    <vxe-table
        height="300"
        class="vxe-edit-table"
        border
        show-overflow
        :data="diaLogBoxTableData"
        :column-config="{resizable: true}"
        :edit-config="{trigger: 'click', mode: 'row'}"
        :tooltip-config="tooltipConfig"
    >
      <vxe-column type="seq" fixed="left" width="60"></vxe-column>
      <vxe-column field="description" title="描述" fixed="left" :edit-render="{}">
        <template #edit="{ row }">
          <vxe-input v-model="row.description" type="text" placeholder="请输入描述"></vxe-input>
        </template>
      </vxe-column>

      <vxe-column field="amount" title="金额" width="83px" :edit-render="{autofocus: '.vxe-input--inner'}">
        <template #edit="{ row }">
          <vxe-input v-model="row.amount" type="float"></vxe-input>
        </template>
      </vxe-column>

      <vxe-column field="type" title="type" v-if="false"></vxe-column>
      <vxe-column field="typeName" width="90px" title="收/支" :edit-render="{}">
        <template #default="{ row }">
          <span>{{ getTransactionTypeNameById(row.type) }}</span>
        </template>
        <template #edit="{ row }">
          <vxe-select v-model="row.type" transfer>
            <vxe-option v-for="item in TRANSACTION_TYPES" :key="item.id" :value="item.id"
                        :label="item.name"></vxe-option>
          </vxe-select>
        </template>
      </vxe-column>

      <vxe-column field="category" title="category" v-if="false"></vxe-column>
      <vxe-column field="categoryName" width="120px" title="类型" :edit-render="{}">
        <template #default="{ row }">
          <span>{{ getCategoryNameById(row.category, categories) }}</span>
        </template>
        <template #edit="{ row }">
          <vxe-select v-model="row.category" transfer>
            <vxe-option v-for="item in categories" :key="item.id" :value="item.id" :label="item.name"></vxe-option>
          </vxe-select>
        </template>
      </vxe-column>

      <vxe-column field="account" title="account" v-if="false"></vxe-column>
      <vxe-column field="accountName" width="120px" title="账户" :edit-render="{}">
        <template #default="{ row }">
          <span>{{ getAccountNameById(row.account, accounts) }}</span>
        </template>
        <template #edit="{ row }">
          <vxe-select v-model="row.account" transfer>
            <vxe-option v-for="item in accounts" :key="item.id" :value="item.id" :label="item.name"></vxe-option>
          </vxe-select>
        </template>
      </vxe-column>

      <vxe-column field="time" width="180px" title="时间" :edit-render="{}">
        <template #edit="{ row }">
          <vxe-input v-model="row.time" type="datetime" placeholder="请选择时间" transfer></vxe-input>
        </template>
      </vxe-column>

      <vxe-column title="操作" fixed="right" width="100">
        <template #default="{ row, column, rowIndex, columnIndex }">
          <vxe-button status="warning" @click="dialogTableDelRow(rowIndex)">删除</vxe-button>
        </template>
      </vxe-column>

    </vxe-table>

    <el-button type="success" @click="submitUpload">解析文件</el-button>
    <el-button type="success" @click="dialogBoxSave">保存数据</el-button>
    <el-button @click="closeDialog">关闭弹窗</el-button>
  </el-dialog>
</template>

<script setup>
import {onMounted, reactive, ref} from "vue";
import {UploadFilled} from "@element-plus/icons-vue";
import {ElMessage, ElMessageBox} from "element-plus";
import {postTransactionExcel} from "@/api/money/transaction/transactionExcel";
import {getTransactionTypeNameById, TRANSACTION_TYPES} from "@/enums/transactionType";
import {getAccountNameById, getCategoryNameById} from "@/api/money/money";
import {allTransactionCategory} from "@/api/money/transaction/transactionCategory";
import {allTransactionAccount} from "@/api/money/transaction/transactionAccount";
import {batchPutTransaction} from "@/api/money/transaction/transaction";

const props = defineProps({
  showDialog: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['updateDialog']);

const upload = ref()
const curFile = ref()

// --- 弹出框表数据 START --- //

const diaLogBoxTableData = ref([])
// 光标移动展示内容
const tooltipConfig = reactive({
  showAll: true
})
// 下拉账户数据
const accounts = ref([])
// 下拉分类数据
const categories = ref([])

// --- 弹出框表数据 END --- //

// 初始化数据
onMounted(() => {
      allTransactionCategory().then(res => {
        categories.value = res.data
      });

      allTransactionAccount().then(res => {
        accounts.value = res.data
      })
    }
)


const closeDialog = (file) => {
  // props.showDialog = false;// 这样是没有用的，得用下面的方式传递值
  emit('updateDialog', false);
  diaLogBoxTableData.value = []
}

const handleChange = (file) => {
  if (file == null) {
    curFile.value = null
    return
  }
  let fileTemp = file.raw
  console.log("文件类型为：", fileTemp.type)
  curFile.value = fileTemp
}

const handleRemove = () => {
  curFile.value = null
}

const submitUpload = () => {
  if (curFile.value == null) {
    ElMessage.warning("请先选择一个文件上传")
    return
  }
  // 清空表格
  diaLogBoxTableData.value = []
  console.log("开始处理上传文件, 名称： " + curFile.value.name, curFile.value)
  let fileFormData = new FormData()
  fileFormData.append('file', curFile.value)
  // 基于 multipart/form-data 指定传输数据为二进制类型进行http请求
  postTransactionExcel(fileFormData).then(
      res => {
        if (res.success) {
          console.log(res.data)
          ElMessage.success("上传解析成功！")
          diaLogBoxTableData.value = res.data
        }
      }).catch(
      error => {
        ElMessage.error("上传解析失败" + error)
      })
}

const handleExceed = () => {
  ElMessage.warning('最多只能上传一个文件！');
}

// 表格内的删除
const dialogTableDelRow = (rowIndex) => {
  diaLogBoxTableData.value.splice(rowIndex, 1)
}

// 保存按钮
const dialogBoxSave = () => {

  let params = diaLogBoxTableData.value
  if (params.length === 0) {
    ElMessage({
      type: 'Warning',
      message: '闹呢，啥都没填呢',
    })
    return
  }

  let rowNum = 0
  for (let param of params) {
    rowNum ++
    if (!param.account || !param.type || !param.category || !param.type || !param.userId || !param.time) {
      ElMessage({
        type: 'info',
        message: '先把内容填全喽啊，第' + rowNum + '行',
      })
      return;
    }
    param.amount = parseFloat(param.amount)
  }

  ElMessageBox.confirm(
      '确定要保存，不再改改了么？保存完了可就退出了。',
      '考虑一哈儿',
      {
        confirmButtonText: '昂嫩',
        cancelButtonText: '再改改',
        type: 'warning',
      }
  ).then(async () => {
    // 批量保存
    batchPutTransaction(params).then(
        res => {
          if (!res.success) {
            ElMessage({
              type: 'info',
              message: '保存失败，别问我为啥',
            })
            return
          }
          ElMessage({
            type: 'success',
            message: '保存成功，数量: ' + res.data + ',快去查询页面看看吧！',
          })
          closeDialog()
        }
    )
  })
}

</script>

<style scoped>

.upload-demo {
  width: 360px;
}

</style>