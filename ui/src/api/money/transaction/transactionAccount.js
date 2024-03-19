import request from '@/api/request'

export function allTransactionAccount() {
    return request.get("/money/transactionAccount")
}