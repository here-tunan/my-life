import { defineStore } from 'pinia';

export const useFamilyStore = defineStore('family', {
    state: () => {
        return {
            id: 0,
            name: "暂未组建/加入家庭",
            desc: "诶，我一人吃饱，全家不饿～🎉🎉🎉",
            avatar: "https://yaodao-images.oss-cn-hangzhou.aliyuncs.com/my-vue-lif/1700552942808.jpg",
        };
    },
    getters: {},
});