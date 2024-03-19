import request from '@/api/request'

export function allTransactionCategory() {
    return request.get("/money/transactionCategory")
}