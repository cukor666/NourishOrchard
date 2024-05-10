import {defineStore} from "pinia";
import {ref} from "vue";

export const useOrderStore = defineStore('order', () => {
    const order = {
        id: '',
        name: '',
        price: 0,
        quantity: 0
    }
    const orderList = ref([{
        order
    }])
    const money = ref(0)

    return {
        orderList, money
    }
});