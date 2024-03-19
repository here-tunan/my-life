import { defineStore } from 'pinia';

export const useAuthStore = defineStore('auth', {
    state: () => {
        return {
            token: "",
            validTime: null,
            isRefreshing: false,
            isAxios401Failing: false,
        };
    },
    getters: {},
});