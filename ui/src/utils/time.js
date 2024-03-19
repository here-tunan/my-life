export const getToday = () => {
    return new Date();
}

export const getYesterday = () => {
    const date = new Date()
    date.setTime(date.getTime() - 3600 * 1000 * 24)
    return date
}

export const getWeekAgoDay = () => {
    const date = new Date()
    date.setTime(date.getTime() - 3600 * 1000 * 24 * 7)
    return date
}

export const getShortcuts = () => {
    return [
        {
            text: 'Today',
            value: getToday(),
        },
        {
            text: 'Yesterday',
            value: getYesterday(),
        },
        {
            text: 'A week ago',
            value: getWeekAgoDay(),
        },
    ]
}

