import axios from "axios";
import {ElMessage} from "element-plus";
import router from "@/router";
import {isNeedRefresh, refreshToken, removeToken} from "@/api/auth/auth";
import {useAuthStore} from "@/stores/auth";

console.log(import.meta.env)

// create an axios instance
const service = axios.create({
    baseURL: import.meta.env.VITE_APP_URL, // base url, see it in .env.[mode] file
    timeout: 50000, // request timeout
    headers: {
        'Content-Type': 'application/json', // header config
    },
})

// request interceptor
service.interceptors.request.use(
    config => {
        // do something before request is sent
        const token = localStorage.getItem('token');
        if (token) {
            config.headers['Authorization'] = token;

            if (!config.url.endsWith('/user/refreshToken') && isNeedRefresh()) {
                refreshToken().then(res => {
                    if (res.success) {
                        // 刷新成功，保存新的 token 和有效时间
                        localStorage.setItem('token', res.data.token);
                        const expiredSeconds = res.data.expiredTime;
                        const now = new Date();
                        now.setTime(now.getTime() + expiredSeconds * 1000);
                        localStorage.setItem('validTime', now.getTime());

                        // 更新请求头中的 token
                        config.headers['Authorization'] = `${res.data.token}`;
                        useAuthStore().isRefreshing = false
                        return config;
                    } else {
                        // 刷新失败，清除 token 并重定向到登录页面
                        removeToken()
                        router.push('/login');
                        useAuthStore().isRefreshing = false
                        return Promise.reject(new Error('Token refresh failed'));
                    }
                })
            }
        }
        return config
    },
    error => {
        // do something with request error
        console.log("i am an axios instance request interceptors, error: ", error)
        return Promise.reject(error)
    }
)


// response interceptor
service.interceptors.response.use(
    response => {
        let url = response.config.url

        // console.log("axios instance response interceptor, response: ", response)
        // http code check
        if (response.status !== 200) {

        }

        // 后端接口返回数据
        // response.data中 统一格式定义为{"code": xxx, "success":true, "data": xxx, "error": xxx}
        const data = response.data

        // 处理一些统一的错误code
        if (data.code !== '' && data.code !== 200) {
            if (data.code === 401) {
                // validToken 接口不参与这项处理，validToken 只在页面跳转时触发）
                if (!url.endsWith("/validToken")) {
                    // 为了避免同一时间发生重复消息提醒问题
                    if (!useAuthStore().isAxios401Failing) {
                        useAuthStore().isAxios401Failing = true
                        // 登录问题, 直接重定向到登录页面
                        ElMessage.warning('当前登录已过期，请重新登录')
                        router.push('/login')

                        // 延时更改
                        setTimeout(() => {
                            useAuthStore().isAxios401Failing = false
                        }, 300)
                    }
                }
            }
        }

        // 如果success 为false 交由业务代码去处理吧
        return data
    },

    error => {
        console.log('axios instance response interceptors, error', error) // for debug
        ElMessage({
            message: 'Ohhhhh, 出错了呢～',
            type: 'error',
        })
        return Promise.reject(error)
    }
)

export default service