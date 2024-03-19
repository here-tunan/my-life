import { defineStore } from 'pinia';

export const useConfigStore = defineStore('config', {
    state: () => {
        return {
            mode: "dark"
        };
    },
    getters: {},
    actions: {
        toggle() {
            if (this.mode === "dark") {
                this.mode = "light";
            } else {
                this.mode = "dark"
            }
        }
    }
});