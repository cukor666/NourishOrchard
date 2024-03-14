import {ref} from "vue";

export function usePage() {
    const currentPage = ref(1)
    const pageSize = ref(3)
    const pageSizes = ref([3, 6, 10, 15, 20])
    const total = ref(Number(sessionStorage.getItem('nourish-user-total')) || 0)

    return {
        currentPage, pageSize, pageSizes, total
    }
}