import request from "@/axios/request.js";
import {ElMessage} from "element-plus";
import {UserDelete, UserList} from "@/api/user/user-api.js";
import {OrderDelete, OrderList} from "@/api/order/order-api.js";

export function useTable(orderList, pageSize, currentPage, total) {
    const updateList = async (s) => {
        try {
            let res = await request.get(OrderList, {
                params: {
                    pageSize: pageSize.value,
                    pageNum: currentPage.value,
                    ...s.value
                }
            })
            if (res.code === 200) {
                total.value = res.data.total
                orderList.value = res.data.orders
            } else {
                console.log(res.msg)
                ElMessage({message: '参数错误', type: 'error'})
            }
        } catch (err) {
            console.error(err)
            ElMessage({message: '服务器错误', type: 'error'})
        }
    }

    const showDetail = (item, detailDialogV, order) => {
        console.log(item)
        detailDialogV.value = true
        order.value = item
    }

    const deleteOrder = async (item) => {
        try {
            let res = await request.delete(OrderDelete, {
                params: {
                    id: item.id
                }
            })
            if (res.code === 200) {
                ElMessage({message: '删除成功', type: 'success'})
            } else {
                ElMessage({message: '删除失败', type: 'error'})
            }
            res = await request.get(OrderList, {
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
        updateList, showDetail, deleteOrder
    }
}