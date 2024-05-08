import {defineStore} from 'pinia'
import {ref} from "vue";

export const useCartStore = defineStore('cart', () => {
    const cartList = ref([])
    const addToCart = (product) => {
        cartList.value.push(product)
    }
    const removeFromCart = (product) => {
        cartList.value = cartList.value.filter(item => item.id !== product.id)
    }
    return {cartList, addToCart, removeFromCart}
})
