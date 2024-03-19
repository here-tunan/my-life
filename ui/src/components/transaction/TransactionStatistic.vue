<template>
  <div class="container">
    <div class="param-select">
      <el-row :gutter="20">
        <el-col :span="10">
          <div class="time-pane">
            <el-tabs type="border-card" v-model="activeTab" @tab-change="tabChange">
              <el-tab-pane label="月账单" name="month">
                <el-date-picker
                    v-model="monthSelect"
                    value-format="YYYY-MM"
                    type="month"
                    placeholder="选择月份"
                    @change="monthSelectClick"
                />
              </el-tab-pane>

              <el-tab-pane label="年账单" name="year">
                <el-date-picker
                    v-model="yearSelect"
                    type="year"
                    value-format="YYYY"
                    placeholder="选择年份"
                    @change="yearSelectClick"
                />
              </el-tab-pane>

              <el-tab-pane label="自定义时间" name="custom">
                <el-date-picker
                    v-model="customDateSelect"
                    type="daterange"
                    unlink-panels
                    range-separator="到"
                    start-placeholder="Start date"
                    end-placeholder="End date"
                    value-format="YYYY-MM-DD"
                    @change="customSelectClick"
                />
              </el-tab-pane>
            </el-tabs>
          </div>
        </el-col>

        <el-col :span="10">
          <el-row>
            <!-- 收入支出选择 -->
            <div>
              <el-radio-group v-model="typeId">
                <el-radio-button
                    @change="typeChange"
                    size="default" border
                    v-for="item in transactionTypes"
                    :label="item.id"
                >{{ item.name }}
                </el-radio-button>
              </el-radio-group>
            </div>
          </el-row>

          <el-row>
            <div class="the-diff">{{theDiff}}</div>
          </el-row>

        </el-col>

      </el-row>
    </div>

    <!--  展示用的图表  -->
    <div class="show-chart">
      <div v-show="!noData">
        <div class="transaction-analysis">
          <el-row type="flex" justify="center" align="middle">
            <h1 class="analysis-text">{{ getTransactionTypeNameById(typeId) }}概览</h1>
            <el-divider/>
          </el-row>
          <el-row>
            <el-col :span="analysisDesc === '' ? 24 : 16">
              <div class="month-chart" id="month-chart"
                   :style="{ '--height': analysisDesc === '' ? '400px' : '200px' }"/>
            </el-col>
            <el-col :span="8" v-show="analysisDesc !== ''">
              <h5 class="analysis-desc">{{ analysisDesc }}</h5>
            </el-col>
          </el-row>
        </div>
        <div class="transaction-rank">
          <h1 class="rank-text">{{ getTransactionTypeNameById(typeId) }}排行榜</h1>
          <el-divider/>
          <div class="rank-table">
            <el-table :data="rankTableData" style="width: 100%" max-height="600" stripe>
              <el-table-column fixed prop="description" label="描述" width="250"/>
              <el-table-column prop="amount" label="金额" width="150"/>
              <el-table-column prop="time" label="时间" width="220"/>
              <el-table-column prop="userId" label="人" width="120" v-if="false"/>
              <el-table-column prop="userName" label="出账人" width="120" v-if="props.familyId !== 0">
                <template #default="scope">
                  <span>{{ getUserNameById(scope.row.userId, users) }}</span>
                </template>
              </el-table-column>

              <el-table-column prop="category" label="分类" width="120" v-if="false"/>
              <el-table-column prop="categoryName" label="分类" width="120">
                <template #default="scope">
                  <span>{{ getCategoryNameById(scope.row.category, categories) }}</span>
                </template>
              </el-table-column>

              <el-table-column prop="type" label="收/支" width="120" v-if="false"/>
              <el-table-column prop="typeName" label="收/支" width="120" v-if="false">
                <template #default="scope">
                  <span>{{ getTransactionTypeNameById(scope.row.type) }}</span>
                </template>
              </el-table-column>

              <el-table-column prop="account" label="账户" width="100" v-if="false"/>
              <el-table-column prop="accountName" label="账户" width="120">
                <template #default="scope">
                  <span>{{ getAccountNameById(scope.row.account, accounts) }}</span>
                </template>
              </el-table-column>

            </el-table>
          </div>

          <el-pagination
              v-model:current-page="currentPage"
              :page-sizes="[10, 50, 100, 300, 500]"
              v-model:page-size="pageSize"
              background layout="total, prev, pager, next, sizes"
              :total="tableDataTotal"
              style="padding-top: 10px"
              @size-change="handleSizeChange"
              @current-change="handleCurrentChange"
          />
        </div>
      </div>
      <div :style="{ display: noData ? 'block' : 'none'}">
        <el-empty description="本阶段无任何支出/收入记录"/>
      </div>
    </div>
  </div>
</template>

<script setup>
import {nextTick, onMounted, reactive, ref, shallowRef, watch} from "vue";
import * as echarts from "echarts";
import {queryTransaction, transactionAnalysis} from "@/api/money/transaction/transaction";
import {getTransactionTypeNameById, TRANSACTION_TYPE_ID, TRANSACTION_TYPES} from "@/enums/transactionType";
import {allTransactionCategory} from "@/api/money/transaction/transactionCategory";
import {allTransactionAccount} from "@/api/money/transaction/transactionAccount";
import {getAccountNameById, getCategoryNameById} from "@/api/money/money";
import {getLoginUserInfo, getUserNameById} from "@/api/user/user";
import {getFamily} from "@/api/family/family";


// 父组件传来的参数
const props = defineProps({
  familyId: {
    type: BigInt,
    default: 0,
  }
});

// 收入/支出
const transactionTypes = ref([])
const theDiff = ref('')
// 下拉账户数据
const accounts = ref([])
// 下拉分类数据
const categories = ref([])
// 用户信息(家庭下的所有成员)
const users = ref([])

const typeId = ref(TRANSACTION_TYPE_ID.EXPEND)

const activeTab = ref('month')
const monthSelect = ref('')
const yearSelect = ref('')
const customDateSelect = ref('')

const isInitChart = ref(false)

// 柱状图右边的描述
const analysisDesc = ref('')

// 表格相关
const rankTableData = ref()
const pageSize = ref(10)
const currentPage = ref(1)
const tableDataTotal = ref(110)

const showChart = shallowRef()

const showChartOption = reactive({
  series: [
    {
      type: 'pie',
      data: [],
      label: {
        normal: {
          formatter: '{b}:{d}% ( {c} )'
        }
      }
    }
  ]
})

const noData = ref(true)

const monthSelectClick = () => {
  initTableParam()
  let param = buildQueryParams();
  transactionsAnalysis(param)
  transactionsDetail(param)
}
const yearSelectClick = () => {
  initTableParam()
  let param = buildQueryParams();
  transactionsAnalysis(param)
  transactionsDetail(param)
}
const customSelectClick = () => {
  initTableParam()
  let param = buildQueryParams();
  transactionsAnalysis(param)
  transactionsDetail(param)
}
const tabChange = () => {
  initTableParam()
  let param = buildQueryParams();
  transactionsAnalysis(param)
  transactionsDetail(param)
}

const typeChange = () => {
  // ElMessage.success("收入/支出改变：" + typeId.value)
  let param = buildQueryParams();
  transactionsAnalysis(param)
  transactionsDetail(param)
}

const initTableParam = () => {
  currentPage.value = 1
  pageSize.value = 10
  rankTableData.value = []
  tableDataTotal.value = 0
  theDiff.value = ''
}

onMounted(() => {
  // 时间初始化
  const now = new Date();
  yearSelect.value = now.getFullYear().toString()
  monthSelect.value = now.getFullYear() + "-" + (now.getMonth() + 1)

  // 初始化
  transactionTypes.value = TRANSACTION_TYPES

  allTransactionCategory().then(res => {
    categories.value = res.data
  });

  allTransactionAccount().then(res => {
    accounts.value = res.data
  })

  // 相关用户信息
  if (props.familyId === 0) {
    // 用户信息
    getLoginUserInfo().then((res) => {
      if (res.success) {
        let userList = []
        userList.push(res.data)
        users.value = userList
      }
      let param = buildQueryParams();
      // 查询聚合信息
      transactionsAnalysis(param)
      // 查询排行具体信息
      transactionsDetail(param)
    })
  } else {
    // 获取家庭用户信息
    getFamily().then((res) => {
      if (res.success) {
        let userList = []
        for (let member of res.data.members) {
          userList.push({
            id: member.userId,
            name: member.name,
          })
        }
        users.value = userList
        let param = buildQueryParams();
        // 查询聚合信息
        transactionsAnalysis(param)
        // 查询排行具体信息
        transactionsDetail(param)
      }
    })
  }
})

// 分页按钮
const handleSizeChange = () => {
  // ElMessage.success("当前 size：" + pageSize.value)
  transactionsDetail(buildQueryParams())
}
const handleCurrentChange = () => {
  // ElMessage.success("当前 page：" + currentPage.value)
  transactionsDetail(buildQueryParams())
}

// 查询聚合信息
const transactionsAnalysis = (param) => {
  if (param.startTime === '' || param.endTime === '') {
    return
  }

  transactionAnalysis(param).then((res) => {
        let isNoData
        if (res.success) {
          let data = []
          if (res.data.total > 0) {
            buildTheDiff(res.data)
            buildAnalysisDesc(res.data)
            for (let categoryResult of res.data.categoryAggregations) {
              data.push({
                name: categoryResult.categoryName,
                value: categoryResult.amount
              })
            }
            data.sort((a, b) => b.value - a.value);
            isNoData = false
            showChartOption.series[0].data = data
          } else {
            isNoData = true
          }
        } else {
          isNoData = true
          // ElMessage.error("数据查询失败！")
        }
        noData.value = isNoData
      }
  )
}

// 构建收入和支出的差值
const buildTheDiff = (data) => {
  // 先乘100 再除100 避免精度的问题
  let diff = parseFloat(data.income) * 100 - parseFloat(data.expenditure) * 100
  theDiff.value = data.income + '-' + data.expenditure + '=¥' + (diff / 100).toFixed(2)
}

// 构建柱形图右边的描述信息
const buildAnalysisDesc = (data) => {
  // 用户聚合
  let userAggregations = data.userAggregations
  userAggregations.sort(function (a, b) {
    return b.amount - a.amount
  });
  // 分类聚合
  let categoryAggregation = data.categoryAggregations
  categoryAggregation.sort(function (a, b) {
    return b.amount - a.amount
  });

  let desc = '在这段时间里，'
  if (props.familyId === 0) {
    desc += '您共' + getTransactionTypeNameById(typeId.value) + data.total + '元。\n'
    desc += '其中' + categoryAggregation[0].categoryName + '占比(' + categoryAggregation[0].percent + '%)最高，共' + categoryAggregation[0].amount + '元。';

  } else {
    desc += '您的家庭共' + getTransactionTypeNameById(typeId.value) + data.total + '元。\n'
    desc += '其中' + categoryAggregation[0].categoryName + '占比(' + categoryAggregation[0].percent + '%)最高，共' + categoryAggregation[0].amount + '元。\n';
    for (let userAggregation of userAggregations) {
      desc += '\n' + getUserNameById(userAggregation.userId, users.value) + ': ' + userAggregation.amount + '元，' + '占比：' + userAggregation.percent + '%'
    }
  }
  analysisDesc.value = desc
}

// 查询具体信息
const transactionsDetail = (param) => {
  if (param.startTime === '' || param.endTime === '') {
    return
  }
  // api 调用
  queryTransaction(param).then((res) => {
    if (res.success) {
      tableDataTotal.value = res.total
      rankTableData.value = res.data
    }
  })
}

// 当数据变动时就展示数据
watch(showChartOption, () => {
      // 未初始化，且有数据时
      if (isInitChart.value === false && !noData.value) {
        // chart 初始化
        const chartDom = document.getElementById('month-chart')
        showChart.value = echarts.init(chartDom);
        isInitChart.value = true

        // window.addEventListener("resize", () => {
        //   showChart.value.resize()
        // });

        // 若dom尺寸变化，则resize
        const chartObserver = new ResizeObserver(() => {
          // setTimeOut 解决闪烁问题
          setTimeout(() => {
            // ElMessage.success("aaa")
            showChart.value.resize();
          }, 0)
        });
        chartObserver.observe(chartDom);
      }
      // 使用 nextTick 方法延迟调用 resize() 【确保在 showChart dom 渲染后执行】
      nextTick(() => {
        showChart.value.resize();
      });
      showChart.value.setOption(showChartOption)
    }
)

// 组装当前页面查询参数信息
const buildQueryParams = () => {
  let startTime
  let endTime
  if (activeTab.value === 'month') {
    if (monthSelect === '') {
      noData.value = true
      return
    }
    // 2023-10 + -01
    startTime = monthSelect.value + '-01'
    const [year, month] = startTime.split('-');
    const lastDay = new Date(year, month, 0).getDate();
    endTime = `${year}-${month}-${lastDay}`;
  }
  if (activeTab.value === 'year') {
    if (yearSelect === '') {
      noData.value = true
      return;
    }
    // 2023 + -01-01
    startTime = yearSelect.value + '-01-01'
    const lastDay = new Date(yearSelect.value, 12, 0).getDate();
    endTime = `${yearSelect.value}-12-${lastDay}`;
  }
  if (activeTab.value === 'custom') {
    if (customDateSelect.value === '') {
      noData.value = true
      return;
    }
    startTime = customDateSelect.value[0]
    endTime = customDateSelect.value[1]
  }
  let param = {}
  if (props.familyId !== '') {
    param.familyId = parseInt(props.familyId)
  }
  param.startTime = startTime
  param.endTime = endTime

  // 收入or支出
  param.type = typeId.value

  // 用户
  if (users.value.length > 0) {
    let userIds = []
    for (let user of users.value) {
      userIds.push(user.id)
    }
    param.userIds = userIds
  }

  // 排序
  param.amountDesc = true

  // 分页
  param.pageSize = pageSize.value
  param.pageIndex = currentPage.value

  return param
}

// // 图表随着容器大小而改变
window.onresize = function () {
  showChart.value.resize();
};

// 侧边栏改变
// watch(() => useSidebarStore().collapse, (newValue, oldValue) => {
//   // ElMessage.success(newValue)
//   nextTick(()=>{
//     setTimeout(() => {
//       console.log(document.getElementById('month-chart').getBoundingClientRect())
//       showChart.value.resize();
//     }, 0);
//   })
// })

</script>

<style scoped>
.time-pane {
}

.transaction-analysis {
  margin-top: 20px;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 5px;
}

.transaction-rank {
  margin-top: 20px;
  overflow: auto;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 5px;
}

.rank-text {
//font-size: 40px;
}

.rank-table {
  padding-top: 30px;
}

.month-chart {
  width: 100%;
  /*动态调整大小*/
  height: var(--height);
}

.analysis-desc {
  white-space: pre-wrap;
  font-size: 15px;
  font-family: "Weibei SC", sans-serif;
  color: #b7b7ad;
}

.the-diff {
  padding-top: 20px;
  font-family: "Weibei SC", sans-serif;
  font-size: 15px;
}
</style>