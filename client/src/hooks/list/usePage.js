import {ref} from "vue";
import {useSessionKey} from "@/hooks/common/useSessionKey.js";

const { NourishUserTotal } = useSessionKey()

export function usePage() {
    const currentPage = ref(1)
    const pageSize = ref(10)
    const pageSizes = ref([10, 15, 20])
    const total = ref(Number(sessionStorage.getItem(NourishUserTotal)) || 0)

    return {
        currentPage, pageSize, pageSizes, total
    }
}