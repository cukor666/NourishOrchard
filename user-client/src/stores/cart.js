import {defineStore} from 'pinia'
import {computed, ref} from "vue";

export const useCartStore = defineStore('cart', () => {
    const cartList = ref(localStorage.getItem('cart') ? JSON.parse(localStorage.getItem('cart')) : [])
    const total = computed(() => {
        let cnt = 0
        cartList.value.forEach(item => {
            cnt += item.quantity
        })
        return cnt
    })
    const addToCart = (product, quantity) => {
        let baseItem = {
            fruit: product,
            quantity: quantity,
        }
        cartList.value.push(baseItem)
        // 存入本地存储
        localStorage.setItem('cart', JSON.stringify(cartList.value))
    }

    return {cartList, total, addToCart}
})
