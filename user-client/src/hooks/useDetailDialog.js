import {ref} from "vue";

import {useCartStore} from "@/stores/cart.js"

const {addToCart} = useCartStore()

export const useDetailDialog = () => {
    const dialogVisible = ref(false)
    const baseObj = {
        id: 0,
        name: '番石榴',
        price: 12.5,
        imgs: ['https://cukor-resource-1318313222.cos.ap-guangzhou.myqcloud.com/img/番石榴.jpg'],
        desc: '香甜可口，营养丰富，营养价值高。'
    }
    const detailObj = ref({
        baseObj
    })
    const currentImgIndex = ref(0)
    const quantity = ref(0.5)

    const showDetail = (fruit) => {
        console.log(fruit)
        dialogVisible.value = true
        detailObj.value = fruit
    }

    const handleClose = () => {
        dialogVisible.value = false
        detailObj.value = baseObj
    }

    const leftImg = () => {
        currentImgIndex.value = (currentImgIndex.value - 1 + detailObj.value.imgs.length) % detailObj.value.imgs.length
    }
    const rightImg = () => {
        currentImgIndex.value = (currentImgIndex.value + 1) % detailObj.value.imgs.length
    }

    const addCart = (item, quantity) => {
        console.log('add to cart, item:', item)
        console.log('add to cart, quantity:', quantity)
        addToCart(item, quantity)
    }


    return {
        dialogVisible,
        detailObj,
        currentImgIndex,
        quantity,
        showDetail,
        handleClose,
        leftImg,
        rightImg,
        addCart
    }
}