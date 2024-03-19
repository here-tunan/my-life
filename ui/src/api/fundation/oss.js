import request from "@/api/request";
import axios from "axios";

const ossService = axios.create({
    timeout: 50000, // request timeout
    headers: {
        'Content-Type': 'multipart/form-data', // header config
    },
})

export function getOssInfo() {
    return request.get("/oss")
}

export function upload(url, file) {
    return ossService.post(url, file)
}