import {ref} from "vue";
import request from "@/axios/request.js";
import {ElMessage} from "element-plus";
import {SupplierList} from "@/api/supplier/sup-api.js";

export function useSearch() {
    const searchDialogV = ref(false)
    const searchSupplier = ref({
        id: 0,
        name: "",
        address: "",
        contactPerson: "",
        phone: "",
        reputation: 0
    })
    const changeSearchDialog = () => {
        searchDialogV.value = true
    }
    const closeSearchDialog = () => {
        searchDialogV.value = false
    }

    const findSupplier = async (u, pageSize, currentPage, total, supplierList) => {
        searchSupplier.value = u
        console.log('searchSupplier = ', searchSupplier.value)
        searchDialogV.value = false
        // 向服务器端做查询，然后将结果放到userList中
        try {
            let res = await request.get(SupplierList, {
                params: {
                    pageSize: pageSize.value,
                    pageNum: currentPage.value,
                    ...searchSupplier.value
                }
            })
            if (res.code === 200) {
                total.value = res.data.total
                supplierList.value = res.data.suppliers
                ElMessage({message: '查询成功', type: 'success'})
            } else {
                ElMessage({message: '查询失败', type: 'error'})
            }
        } catch (err) {
            console.error(err)
            ElMessage({message: '服务器错误', type: 'error'})
        }
    }

    return {
        searchDialogV, searchSupplier, changeSearchDialog, closeSearchDialog, findSupplier
    }
}