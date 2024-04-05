import request from "@/axios/request.js";
import {ElMessage} from "element-plus";
import {SupplierDelete, SupplierList} from "@/api/supplier/sup-api.js";

export function useTable(supplierList, pageSize, currentPage, total) {
    const updateList = async (searchSupplier) => {
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
            } else {
                console.log(res.msg)
                ElMessage({message: '参数错误', type: 'error'})
            }
        } catch (err) {
            console.error(err)
            ElMessage({message: '服务器错误', type: 'error'})
        }
    }

    const showDetail = (item, detailDialogV, supplier) => {
        console.log(item)
        detailDialogV.value = true
        supplier.value = item
    }

    const deleteSupplier = async (item) => {
        try {
            let res = await request.delete(SupplierDelete, {
                params: {
                    id: item.id
                }
            })
            if (res.code === 200) {
                ElMessage({message: '删除成功', type: 'success'})
            } else {
                ElMessage({message: '删除失败', type: 'error'})
            }
            res = await request.get(SupplierList, {
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
        updateList, showDetail, deleteSupplier
    }
}