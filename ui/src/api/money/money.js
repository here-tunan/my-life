
export const getCategoryNameById = (id, categories) => {
    for (let category of categories) {
        if (category.id === id) {
            return category.name;
        }
    }
    return ''
}

export const getAccountNameById = (id, accounts) => {
    for (let account of accounts) {
        if (account.id === id) {
            return account.name;
        }
    }
    return ''
}