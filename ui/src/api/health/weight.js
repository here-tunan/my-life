import request from "@/api/request";

// 添加体重
export function putWeight(weight) {
    return request.put("/health/weight", weight)
}

export function getWeightList(param) {
    return request.post("/health/weight/query", param)
}