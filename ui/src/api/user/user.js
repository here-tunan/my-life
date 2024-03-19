import request from "@/api/request"

export const getUserNameById = (id, users) => {
    for (let user of users) {
        if (user.id === id) {
            return user.name;
        }
    }
    return ''
}

export function getLoginUserInfo() {
    return request.get('/user/info')
}

// 更新用户信息
export function putUser(param) {
    return request.put('/user', param)
}

// 新增用户
export function postUser() {
    return request.post('/user')
}
