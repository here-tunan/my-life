<template>
  <div class="container">
    <!-- 查询表头 -->
    <el-form label-width="50px">
      <el-row>
        <el-form-item label="时间">
          <el-date-picker
              v-model="form.dateSelect"
              type="daterange"
              unlink-panels
              range-separator="-"
              start-placeholder="Start date"
              end-placeholder="End date"
              value-format="YYYY-MM-DD"
          />
        </el-form-item>
        <el-form-item label="收/支" style="padding-left: 20px">
          <el-radio-group v-model="form.typeId">
            <el-radio size="default" border
                      v-for="item in transactionTypes"
                      :label="item.id"
            >{{ item.name }}
            </el-radio>
          </el-radio-group>
        </el-form-item>
      </el-row>

      <!-- action button -->
      <el-form-item>
        <el-button :icon="Search" type="primary" @click="onSearch">查一查</el-button>
        <el-button type="info" :icon="Remove" @click="clearForm">清空查询条件</el-button>
        <el-button type="success" :icon="Edit" @click="onBatchAdd">手工记账</el-button>
        <el-button type="success" :icon="UploadFilled" @click="onExcelAdd">Excel导入</el-button>
      </el-form-item>
    </el-form>
    <!-- 页面表单 -->
    <vxe-table
        :row-config="{height: 35}"
        round
        stripe
        height="600px"
        border
        show-overflow
        :data="tableData"
        :column-config="{resizable: true}"
        :edit-config="{trigger: 'dblclick', mode: 'cell'}"
        :tooltip-config="tooltipConfig"
        show-footer
        :footer-method="tableFooterMethod"
        :sort-config="{multiple: true, trigger: 'cell'}"
    >
      <vxe-column type="seq" fixed="left" width="50"></vxe-column>

      <vxe-column field="amount" title="金额" width="100px" fixed="left"
                  :edit-render="{autofocus: '.vxe-input--inner'}" sortable>
        <template #edit="{ row }">
          <vxe-input v-model="row.amount" type="float" step="1"></vxe-input>
        </template>
      </vxe-column>

      <vxe-column field="description" title="描述" width="300px" fixed="left" :edit-render="{}">
        <template #edit="{ row }">
          <vxe-input v-model="row.description" type="text" placeholder="请输入描述"></vxe-input>
        </template>
      </vxe-column>

      <vxe-column field="type" title="type" v-if="false"></vxe-column>
      <vxe-column field="typeName" width="140px" title="收/支" :edit-render="{}">
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
      <vxe-column field="categoryName" width="130px" title="类型" :edit-render="{}">
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

      <vxe-column title="操作" fixed="right" width="200">

        <template #default="{ row, column, rowIndex, columnIndex }">
          <vxe-button status="success" v-if="true" @click="tableSaveRow(row)">保存</vxe-button>
          <vxe-button status="warning" @click="tableDelRow(row, rowIndex)">删除</vxe-button>
        </template>
      </vxe-column>

    </vxe-table>

    <el-pagination
        v-model:current-page="curPageIndex"
        v-model:page-size="curPageSize"
        :page-sizes="[10, 20, 50, 100, 300]"
        :small="false"
        :background="true"
        layout="total, sizes, prev, pager, next, jumper"
        :total="dataTotal"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
    />

    <!--  excel 导入记账的弹窗  -->
    <ExcelTransactionRecordDialog :show-dialog="showExcelTransactionRecordDialog" @updateDialog="updateExcelDialog"/>

    <el-dialog class="transaction-dialog" width="1000px" v-model="showDialogBox" :model="dialogForm" title="记他个几笔！"
               draggable>
      <el-form label-width="60" ref="dialogFromRef" :model="dialogForm">
        <!--  表单第一行 -->
        <el-row>
          <el-col :span="7">
            <el-form-item label="收/支">
              <el-select v-model="dialogForm.selectType" value-key="id"
                         style="width:100px">
                <el-option
                    v-for="item in TRANSACTION_TYPES"
                    :key="item.id"
                    :label="item.name"
                    :value="item"
                />
              </el-select>
            </el-form-item>
          </el-col>

          <el-col :span="7">
            <el-form-item label="几笔">
              <el-input-number v-model="dialogForm.batchNum" autocomplete="off" placeholder="填个数字"
                               style="width: 100px"/>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="哪天">
              <el-date-picker
                  v-model="dialogForm.date"
                  type="date"
                  placeholder="Pick a date"
                  style="width: 140px"
                  :shortcuts="shortcuts"
                  value-format="YYYY-MM-DD"
              />
            </el-form-item>
          </el-col>
        </el-row>

        <!--  表单第二行 -->
        <el-row>
          <el-col :span="7">
            <el-form-item label="分类">
              <el-select v-model="dialogForm.selectCategory" value-key="id"
                         style="width:100px">
                <el-option
                    v-for="item in categories"
                    :key="item.id"
                    :label="item.name"
                    :value="item"
                />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="7">
            <el-form-item label="账户">
              <el-select v-model="dialogForm.selectAccount" value-key="id"
                         style="width:100px">
                <el-option
                    v-for="item in accounts"
                    :key="item.id"
                    :label="item.name"
                    :value="item"
                />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>

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

      <template #footer>
      <span class="dialog-footer">
        <el-button type="primary"
                   @click="dialogBoxBatchAdd">生成{{ dialogForm.batchNum > 0 ? dialogForm.batchNum : 0 }}条</el-button>
        <el-button type="primary" @click="dialogBoxSave">记上这几笔</el-button>
        <el-button @click="dialogBoxFormClear">清空选择</el-button>
        <el-button @click="dialogBoxCancel">算了退出</el-button>
      </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>

import {computed, onMounted, reactive, ref} from "vue";
import {
  TRANSACTION_TYPE_NAME,
  TRANSACTION_TYPE_ID,
  TRANSACTION_TYPES,
  getTransactionTypeNameById
} from '@/enums/transactionType'
import {getShortcuts} from "@/utils/time";
import {getAccountNameById, getCategoryNameById} from "@/api/money/money";
import {Edit, Remove, Search, UploadFilled} from "@element-plus/icons-vue";
import {ElMessage, ElMessageBox} from "element-plus";
import {allTransactionCategory} from "@/api/money/transaction/transactionCategory";
import {allTransactionAccount} from "@/api/money/transaction/transactionAccount";
import {
  batchPutTransaction,
  deleteTransaction,
  putTransaction,
  queryTransaction
} from "@/api/money/transaction/transaction";
import ExcelTransactionRecordDialog from "@/views/money/ExcelTransactionRecordDialog.vue";

// --------------------- 页面数据 --------------------
const tooltipConfig = reactive({
  showAll: true
})
// 时间选择快捷键
const shortcuts = getShortcuts()
// 收入/支出
const transactionTypes = ref([])
// 下拉账户数据
const accounts = ref([])
// 下拉分类数据
const categories = ref([])
// 表格数据
const tableData = ref([])
// 表单数据
const form = reactive({
  userId: '',
  typeId: '',
  dateSelect: '',
})
// 分页信息
const curPageSize = ref(20)
const curPageIndex = ref(1)
const dataTotal = ref(0)
// 页脚计算统计信息
const tableFooterData = computed(() => {
  let sum1 = 0;
  let sum2 = 0;
  if (tableData.value != null && tableData.value.length !== 0) {
    for (let data of tableData.value) {
      if (data.amount != null && data.amount !== '' && data.amount !== 0) {
        if (data.type === TRANSACTION_TYPE_ID.EXPEND) {
          sum1 += parseFloat(data.amount) * 100
        } else {
          sum2 += parseFloat(data.amount) * 100
        }
      }
    }
    sum1 = sum1 / 100
    sum2 = sum2 / 100
  }
  return [['支出', sum1], ['收入', sum2]]
})
// ------------------- 页面数据 -----------------------


// -----------------  弹出框数据 -----------------------
const showExcelTransactionRecordDialog = ref(false)
const updateExcelDialog = isShow => {
  showExcelTransactionRecordDialog.value = isShow
}

// 弹出框是否弹出
const showDialogBox = ref(false)
const dialogFromRef = ref()
// 弹出框表单数据
const dialogForm = reactive({
  selectUser: '',
  selectType: {
    id: TRANSACTION_TYPE_ID.EXPEND,
    name: TRANSACTION_TYPE_NAME[TRANSACTION_TYPE_ID.EXPEND]
  },
  selectCategory: {
    id: '',
    name: '',
  },
  selectAccount: {
    id: '',
    name: '',
  },
  date: '',
  batchNum: 1
})
// 弹出框中的表数据
const diaLogBoxTableData = ref([])
// ------------------- END ----------------------------


// 初始化数据
onMounted(() => {
      transactionTypes.value = TRANSACTION_TYPES

      allTransactionCategory().then(res => {
        categories.value = res.data
      });

      allTransactionAccount().then(res => {
        accounts.value = res.data
      })
    }
)


// ---------------------- 页面按钮 --------------------
// 查询按钮
const onSearch = () => {
  // process transaction detail query
  queryTransactionDetail()
}

// 我要记账按钮
const onBatchAdd = () => {
  openDialogBox()
}

// Excel 导入
const onExcelAdd = () => {
  showExcelTransactionRecordDialog.value = true
}

// 清空查询条件按钮
const clearForm = () => {
  form.userId = '';
  form.dateSelect = '';
}

// 单个保存
const tableSaveRow = async (row) => {
  let param = row
  param.amount = parseFloat(row.amount)
  // 单个保存
  putTransaction(param).then(
      res => {
        if (!res.success) {
          ElMessage({
            type: 'info',
            message: '保存失败',
          })
          return
        }
        ElMessage({
          type: 'success',
          message: '保存成功',
        })
      }
  )
}

// 删除单行
const tableDelRow = (row, rowIndex) => {
  ElMessageBox.confirm(
      '确定要删除么？',
      '就怕你是误点了啊',
      {
        confirmButtonText: '我很确定',
        cancelButtonText: '误点误点～',
        type: 'warning',
      }
  ).then(async () => {
    // 删除的逻辑
    console.log(row.id)
    deleteTransaction(row.id).then(
        res => {
          if (!res.success) {
            ElMessage({
              type: 'info',
              message: '删除失败',
            })
            return
          }
          // 表格中删除
          tableData.value.splice(rowIndex, 1)
          ElMessage({
            type: 'success',
            message: '删除成功',
          })
        }
    )
  })
}

// 查询账单信息
const queryTransactionDetail = async () => {
  // collect params
  let param = {}
  if (form.dateSelect !== '') {
    param.startTime = form.dateSelect[0]
    param.endTime = form.dateSelect[1]
  }

  param.pageSize = curPageSize.value
  param.pageIndex = curPageIndex.value
  if (form.typeId !== '') {
    param.type = parseInt(form.typeId)
  }

  queryTransaction(param).then(
      res => {
        if (res.success) {
          tableData.value = res.data
          dataTotal.value = res.total
        } else {
          tableData.value = []
        }
      }
  )
}

// 页脚方法
const tableFooterMethod = () => {
  return tableFooterData.value
}

// 分页按钮
const handleSizeChange = () => {
  queryTransactionDetail()
}
const handleCurrentChange = () => {
  queryTransactionDetail()
}

// ---------------------- END --------------------


// ---------------------- 弹窗按钮 ----------------
// 先生成这几条的按钮
const dialogBoxBatchAdd = () => {
  if (dialogForm.batchNum > 0) {
    for (let i = 0; i < dialogForm.batchNum; i++) {
      // console.log(dialogForm)
      diaLogBoxTableData.value.push({
        amount: 0,
        // 选的类型
        type: dialogForm.selectType ? dialogForm.selectType.id : '',
        // 选的分类
        category: dialogForm.selectCategory ? dialogForm.selectCategory.id : '',
        // 选的账户
        account: dialogForm.selectAccount ? dialogForm.selectAccount.id : '',
        // 选的时间
        time: dialogForm.date ? dialogForm.date + " 00:00:00" : '',
      })
    }
  }
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

  for (let param of params) {
    if (!param.account || !param.type || !param.category || !param.type || !param.time) {
      ElMessage({
        type: 'info',
        message: '先把内容填全喽啊',
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
          dialogBoxFormClear()
          closeDialogBox()
        }
    )
  })
}

const dialogBoxFormClear = () => {
  dialogForm.selectUser = null;
  dialogForm.selectCategory = null;
  dialogForm.selectAccount = null;
  dialogForm.date = '';
  dialogForm.batchNum = 1;
}

const dialogBoxCancel = () => {
  closeDialogBox()
}

const openDialogBox = () => {
  showDialogBox.value = true;
}

const closeDialogBox = () => {
  showDialogBox.value = false;
  diaLogBoxTableData.value = []
}
// ---------------------- END ----------------


</script>

<style scoped>

</style>
