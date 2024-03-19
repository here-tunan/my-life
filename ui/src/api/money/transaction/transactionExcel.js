import request from '@/api/request'

export function postTransactionExcel(file) {
    return request.post("/money/transaction/excel", file, {
        headers: {"Content-Type": "multipart/form-data"}
    })
}