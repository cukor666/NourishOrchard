import {ref} from "vue";
import request from "@/axios/request.js";
import {ElMessage} from "element-plus";
import {WarehouseList} from "@/api/warehouse/warehouse-api.js";

export function useSearch() {
    const searchDialogV = ref(false)
    const searchWarehouse = ref({
        id: 0,
        address: '',
        capacity: 0,
        status: 0
    })
    const changeSearchDialog = () => {
        searchDialogV.value = true
    }
    const closeSearchDialog = () => {
        searchDialogV.value = false
    }

    const findWarehouse = async (u, pageSize, currentPage, total, warehouseList) => {
        searchWarehouse.value = u
        console.log('searchWarehouse = ', searchWarehouse.value)
        searchDialogV.value = false
        // 向服务器端做查询，然后将结果放到userList中
        try {
            let res = await request.get(WarehouseList, {
                params: {
                    pageSize: pageSize.value,
                    pageNum: currentPage.value,
                    ...searchWarehouse.value
                }
            })
            if (res.code === 200) {
                total.value = res.data.total
                warehouseList.value = res.data.warehouses
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
        searchDialogV, searchWarehouse, changeSearchDialog, closeSearchDialog, findWarehouse
    }
}