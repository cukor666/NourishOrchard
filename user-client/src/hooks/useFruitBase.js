import {ref} from "vue";

export const useFruitBase = () => {
    const fruitList = ref([{
        id: 0,
        name: '番石榴',
        price: 12.5,
        imgs: '',
        desc: '香甜可口，营养丰富，营养价值高。'
    }])

    return {
        fruitList
    }
}