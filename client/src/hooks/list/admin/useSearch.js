import {ref} from "vue";
import request from "@/axios/request.js";
import {ElMessage} from "element-plus";
import {UserList} from "@/api/user/user-api.js";
import {AdminList} from "@/api/admin/admin-api.js";

export function useSearch() {
    const searchDialogV = ref(false)
    const searchAdmin = ref({
        id: 0,
        username: '',
        name: '',
        email: ''
    })
    const changeSearchDialog = () => {
        searchDialogV.value = true
    }
    const closeSearchDialog = () => {
        searchDialogV.value = false
    }

    const findAdmin = async (u, pageSize, currentPage, total, adminList) => {
        searchAdmin.value = u
        console.log('searchAdmin = ', searchAdmin.value)
        searchDialogV.value = false
        // 向服务器端做查询，然后将结果放到userList中
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
        searchDialogV, searchAdmin, changeSearchDialog, closeSearchDialog, findAdmin
    }
}