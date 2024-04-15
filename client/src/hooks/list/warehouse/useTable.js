import request from "@/axios/request.js";
import {ElMessage} from "element-plus";
import {UserDelete, UserList} from "@/api/user/user-api.js";
import {WarehouseDelete, WarehouseList} from "@/api/warehouse/warehouse-api.js";

export function useTable(warehouseList, pageSize, currentPage, total) {
    const updateList = async (searchWarehouse) => {
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
            } else {
                console.log(res.msg)
                ElMessage({message: '参数错误', type: 'error'})
            }
        } catch (err) {
            console.error(err)
            ElMessage({message: '服务器错误', type: 'error'})
        }
    }

    const showDetail = (item, detailDialogV, warehouse) => {
        console.log(item)
        detailDialogV.value = true
        warehouse.value = item
    }

    const deleteWarehouse = async (item) => {
        try {
            let res = await request.delete(WarehouseDelete, {params: {id: item.id}})
            if (res.code === 200) {
                ElMessage({message: '删除成功', type: 'success'})
            } else {
                ElMessage({message: '删除失败', type: 'error'})
            }
            res = await request.get(WarehouseList, {
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
        updateList, showDetail, deleteWarehouse
    }
}