import request from "@/api/request"
import {useAuthStore} from "@/stores/auth";

export function login(account, password) {
    return request.get("/user/login", {
        params: {
            account: account,
            password: password
        }
    })
}

export function validToken(token) {
    return request.get("/user/validToken", {
        params: {
            token: token
        }
    })
}

export async function validToken2(token) {
    try {
        const response = await request.get("/user/validToken", {
            params: {
                token: token
            }
        });
        return !!response.success;
    } catch (error) {
        console.error(error);
        return false;
    }
}

export function removeToken() {
    localStorage.removeItem("token")
    localStorage.removeItem("validTime")
}

export function isNeedRefresh() {
    // 有正在刷新的
    if (useAuthStore().isRefreshing) {
        return false
    }

    // 如果有token 验证token是否需要刷新
    let validTimeStr = localStorage.getItem("validTime")
    if (!validTimeStr) {
        return true
    }
    let validTime = parseInt(validTimeStr)
    let nowTime = new Date().getTime()
    // 是否低于100s
    return (validTime - nowTime) < 100 * 1000
}

export async function refreshToken() {
    useAuthStore().isRefreshing = true
    return request.get("/user/refreshToken", {
        params: {
            token: localStorage.getItem("token")
        }
    });
}