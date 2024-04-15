import request from "@/axios/request.js";
import {ElMessage} from "element-plus";
import {UserDelete, UserList} from "@/api/user/user-api.js";
import {WarehouseDelete, WarehouseList} from "@/api/warehouse/warehouse-api.js";
import {InventoryDelete, InventoryList} from "@/api/inventory/inventory-api.js";

export function useTable(inventoryList, pageSize, currentPage, total) {
    const updateList = async (searchObj) => {
        try {
            let res = await request.get(InventoryList, {
                params: {
                    pageSize: pageSize.value,
                    pageNum: currentPage.value,
                    ...searchObj.value
                }
            })
            if (res.code === 200) {
                total.value = res.data.total
                inventoryList.value = res.data.inventors
            } else {
                console.log(res.msg)
                ElMessage({message: '参数错误', type: 'error'})
            }
        } catch (err) {
            console.error(err)
            ElMessage({message: '服务器错误', type: 'error'})
        }
    }

    const showDetail = (item, detailDialogV, inventory) => {
        console.log(item)
        detailDialogV.value = true
        inventory.value = item
    }

    const deleteInventory = async (item) => {
        try {
            let res = await request.delete(InventoryDelete, {params: {id: item.id}})
            if (res.code === 200) {
                ElMessage({message: '删除成功', type: 'success'})
            } else {
                ElMessage({message: '删除失败', type: 'error'})
            }
            res = await request.get(InventoryList, {
                params: {
                    pageSize: pageSize.value,
                    pageNum: currentPage.value
                }
            })
            if (res.code !== 200) {
                ElMessage({message: '未能及时更新列表，请刷新', type: 'warning'})
            }
        } catch {
            ElMessage({message: '系统错误', type: 'error'})
        }
    }


    return {
        updateList, showDetail, deleteInventory
    }
}