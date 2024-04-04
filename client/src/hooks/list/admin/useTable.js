import request from "@/axios/request.js";
import {ElMessage} from "element-plus";
import {AdminList} from "@/api/admin/admin-api.js";

export function useTable(adminList, pageSize, currentPage, total) {
    const updateList = async (searchAdmin) => {
        try {
            let res = await request.get(AdminList, {
                params: {
                    pageSize: pageSize.value,
                    pageNum: currentPage.value,
                    ...searchAdmin.value
                }
            })
            if (res.code === 200) {
                total.value = res.data.total
                adminList.value = res.data.admins
            } else {
                console.log(res.msg)
                ElMessage({message: '参数错误', type: 'error'})
            }
        } catch (err) {
            console.error(err)
            ElMessage({message: '服务器错误', type: 'error'})
        }
    }




    return {
        updateList
    }
}