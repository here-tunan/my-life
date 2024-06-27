import request from "@/api/request";

export function createExercise(param) {
    return request.put("/exercise", param)
}
