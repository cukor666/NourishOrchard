import {ref} from "vue";
import request from "@/axios/request.js";
import {ElMessage} from "element-plus";
import {OrderList} from "@/api/order/order-api.js";

export function useSearch() {
    const searchDialogV = ref(false)
    const searchOrder = ref({
        id: 0,
        title: "",
        status: 0,
        commodityId: 0,
        quantity: 0,
        buyerId: 0,
        adminId: 0,
        warehouseId: 0,
        receiverName: "",
        receiverPhone: "",
        address: "",
        remark: "",
        createdAt: "",
        updatedAt: ""
    })
    const changeSearchDialog = () => {
        searchDialogV.value = true
    }
    const closeSearchDialog = () => {
        searchDialogV.value = false
    }

    const findOrder = async (u, pageSize, currentPage, total, orderList) => {
        searchOrder.value = u
        console.log('searchOrder = ', searchOrder.value)
        searchDialogV.value = false
        // 向服务器端做查询，然后将结果放到userList中
        try {
            let res = await request.get(OrderList, {
                params: {
                    pageSize: pageSize.value,
                    pageNum: currentPage.value,
                    ...searchOrder.value
                }
            })
            if (res.code === 200) {
                total.value = res.data.total
                orderList.value = res.data.orders
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
        searchDialogV, searchOrder, changeSearchDialog, closeSearchDialog, findOrder
    }
}