import {ref} from "vue";
import {useSessionKey} from "@/hooks/common/useSessionKey.js";

const { NourishLogoutUserTotal } = useSessionKey()

export function usePage() {
    const currentPage = ref(1)
    const pageSize = ref(3)
    const pageSizes = ref([3, 6, 10, 15, 20])
    const total = ref(Number(sessionStorage.getItem(NourishLogoutUserTotal)) || 0)

    return {
        currentPage, pageSize, pageSizes, total
    }
}