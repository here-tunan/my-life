import request from "@/api/request";

// 获取当前登录用户的家庭信息
export function getFamily() {
    return request.get("/family")
}

export function createFamily(param) {
    return request.post("/family", param)
}

// 增加当前用户所在家庭的家庭成员
export function addFamilyMember(userId) {
    return request.post("/family/member", {
        userId: userId
    })
}