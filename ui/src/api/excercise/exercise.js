import request from "@/api/request";

export function createExercise(param) {
    return request.put("/exercise", param)
}

export function listExercise(param) {
    return request.put('/exercise/list', param)
}
