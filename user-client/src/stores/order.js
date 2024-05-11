import {defineStore} from "pinia";
import {ref} from "vue";

export const useOrderStore = defineStore('order', () => {
    const orderList = ref([])
    const money = ref(0)
    const selectedIndexes = ref([])

    return {
        orderList, money, selectedIndexes
    }
});