import {ref} from "vue";

export function useFruit() {
    const fruit = ref({
        id: 0,
        name: '',
        water: 0.0,
        sugar: 0.0,
        shelfLife: 0,
        origin: '',
        supplierId: 0
    })

    return {
        fruit
    }
}