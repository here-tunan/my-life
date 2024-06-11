import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import Dashboard from "@/views/Dashboard.vue";
import TransactionRecord from "../views/money/TransactionRecord.vue";
import TransactionAccount from "@/views/money/TransactionAccount.vue";
import TransactionCategory from "@/views/money/TransactionCategory.vue";
import Weight from "@/views/health/Weight.vue";
import Login from "@/views/login/Login.vue"
import MyBill from "@/views/money/MyBill.vue";
import Family from "@/views/family/Family.vue";
import FamilyBill from "@/views/family/FamilyBill.vue";
import User from "@/views/User.vue";
import BalanceDiet from "@/components/health/BalanceDiet.vue";
import Sport from '../views/health/Sport.vue';
import LegTrain from '../views/train/LegTrain.vue';
import ChestTrain from '../views/train/ChestTrain.vue';

const router = createRouter({
    history: createWebHashHistory(),
    routes: [
        {
            path: '/',
            redirect: '/dashboard',
        },
        {
            path: '/login',
            name: 'login',
            component: Login,
        },
        {
            path: '/',
            name: 'home',
            component: HomeView,
            children: [
                {
                    path: '/dashboard',
                    name: 'dashboard',
                    meta: {
                        title: '系统首页',
                    },
                    component: Dashboard,
                },
                {
                    path: '/my-bill',
                    name: 'my-bill',
                    component: MyBill,
                    meta: {
                        title: "账单统计"
                    }
                },
                {
                    path: '/transaction-record',
                    name: 'transaction-record',
                    component: TransactionRecord,
                    meta: {
                        title: "记账"
                    }
                },
                {
                    path: '/transaction-account',
                    name: 'transaction-account',
                    component: TransactionAccount,
                    meta: {
                        title: "账户管理"
                    }
                },
                {
                    path: '/transaction-category',
                    name: 'transaction-category',
                    component: TransactionCategory,
                    meta: {
                        title: "账目分类"
                    }
                },
                {
                    path: '/weight',
                    name: '/weight',
                    component: Weight,
                    meta: {
                        title: "体重管理"
                    }
                },
                {
                    path: '/family',
                    name: 'family',
                    component: Family,
                    meta: {
                        title: "温馨小家"
                    }
                },
                {
                    path: '/family-bill',
                    name: 'familyBill',
                    component: FamilyBill,
                    meta: {
                        title: "家庭账单"
                    }
                },
                {
                    path: '/user',
                    name: 'user',
                    component: User,
                    meta: {
                        title: "用户中心"
                    }
                },
                {
                    path: '/balance-diet',
                    name: 'balance-diet',
                    component: BalanceDiet,
                    meta: {
                        title: '均衡膳食',
                    }
                },
                {
                    path: '/sport',
                    name: 'sport',
                    component: Sport,
                    meta: {
                        title: '运动健康',
                    }
                },
                {
                    path: '/leg-train',
                    name: 'leg-train',
                    component: LegTrain,
                    meta: {
                        title: '腿部训练',
                    }
                },
                {
                    path: '/chest-train',
                    name: 'chest-train',
                    component: ChestTrain,
                    meta: {
                        title: '胸部训练',
                    }
                },
            ]
        }
    ]
})

export default router
