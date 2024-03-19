// 该文件必须挂载到main.js中才能生效

import router from './router'
import {isNeedRefresh, refreshToken, removeToken, validToken, validToken2} from './api/auth/auth'


// import {useAuthStore} from './stores/auth';


import {ElMessage} from 'element-plus'

router.beforeEach((to, from, next) => {

    // 取到用户的token
    let token = localStorage.getItem("token")

    if (token == null || token === '') {
        // 没有token
        if (to.path !== '/login') {
            ElMessage.warning('当前用户未登录，跳转到登录页面')
            next('/login')
        } else {
            next()
        }
        return
    }

    // 有token，判断token有效性
    validToken2(token).then(isValid => {
        if (!isValid) {
            // token 无效
            if (to.path === '/login') {
                next()
                return;
            }
            ElMessage.warning('当前用户登录状态已失效，跳转到登录页面')
            removeToken()
            next('/login')
            return;
        }

        // token 有效
        if (isNeedRefresh()) {
            // 需要刷新token
            refreshToken().then(res => {
                if (res.success) {
                    // 刷新成功 保存新的token
                    localStorage.setItem("token", res.data.token)
                    let expiredSeconds = res.data.expiredTime
                    let now = new Date()
                    now.setTime(now.getTime() + expiredSeconds * 1000)
                    localStorage.setItem("validTime", now.getTime())
                    if (to.path === '/login') {
                        ElMessage.success('当前已登录，跳转到首页')
                        next("/")
                    } else {
                        next()
                    }
                } else {
                    // 刷新失败，重新登录
                    removeToken()
                    if (to.path === '/login') {
                        next()
                    } else {
                        ElMessage.warning('当前用户登录状态已失效，跳转到登录页面')
                        next('/login')
                    }
                }
            })
        } else {
            // 不需要刷新token
            if (to.path === '/login') {
                ElMessage.success('当前已登录，跳转到首页')
                next("/")
            } else {
                next()
            }
        }
    });
})