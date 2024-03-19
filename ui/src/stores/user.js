import { defineStore } from 'pinia';

export const useUserStore = defineStore('user', {
    state: () => {
        return {
            account: "",
            username: "",
            desc: "",
            avatar: localStorage.getItem("avatar"),
        };
    },
    getters: {},
});