import {ref} from "vue";
import request from "@/axios/request.js";
import {ElMessage} from "element-plus";
import {InventoryList} from "@/api/inventory/inventory-api.js";

export function useSearch() {
    const searchDialogV = ref(false)
    const searchInventory = ref({
        id: 0,
        commodityId: 0,
        quantity: 0,
        employeeId: 0,
        warehouseId: 0,
        createdAt: '',
        updatedAt: ''
    })
    const changeSearchDialog = () => {
        searchDialogV.value = true
    }
    const closeSearchDialog = () => {
        searchDialogV.value = false
    }

    const findInventory = async (u, pageSize, currentPage, total, inventoryList) => {
        searchInventory.value = u
        console.log('searchInventory = ', searchInventory.value)
        searchDialogV.value = false
        // 向服务器端做查询，然后将结果放到userList中
        try {
            let res = await request.get(InventoryList, {
                params: {
                    pageSize: pageSize.value,
                    pageNum: currentPage.value,
                    ...searchInventory.value
                }
            })
            if (res.code === 200) {
                total.value = res.data.total
                inventoryList.value = res.data.inventors
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
        searchDialogV, searchInventory, changeSearchDialog, closeSearchDialog, findInventory
    }
}