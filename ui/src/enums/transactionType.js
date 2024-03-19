export const TRANSACTION_TYPE_ID = {
    "INCOME": 1,
    "EXPEND": 2
}

export const TRANSACTION_TYPE_NAME = {
    [TRANSACTION_TYPE_ID.INCOME]: '收入',
    [TRANSACTION_TYPE_ID.EXPEND]: '支出'
}

export const TRANSACTION_TYPES = [
    {
        id: TRANSACTION_TYPE_ID.INCOME,
        name: TRANSACTION_TYPE_NAME[TRANSACTION_TYPE_ID.INCOME]
    },
    {
        id: TRANSACTION_TYPE_ID.EXPEND,
        name: TRANSACTION_TYPE_NAME[TRANSACTION_TYPE_ID.EXPEND]
    }
]

export const getTransactionTypeNameById = (id) => {
    if (id === TRANSACTION_TYPE_ID.INCOME) {
        return TRANSACTION_TYPE_NAME[TRANSACTION_TYPE_ID.INCOME]
    }

    if (id === TRANSACTION_TYPE_ID.EXPEND) {
        return TRANSACTION_TYPE_NAME[TRANSACTION_TYPE_ID.EXPEND]
    }

    return ''
}