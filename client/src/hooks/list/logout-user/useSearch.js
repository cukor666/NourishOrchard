import {ref} from "vue";
import request from "@/axios/request.js";
import {ElMessage} from "element-plus";
import {LogoutUserList} from "@/api/logout-user/logout-user-api.js";

export function useSearch() {
    const searchDialogV = ref(false)
    const searchUser = ref({
        id: 0,
        username: '',
        name: '',
        gender: '',
        phone: '',
        address: '',
        birthday: ''
    })
    const changeSearchDialog = () => {
        searchDialogV.value = true
    }
    const closeSearchDialog = () => {
        searchDialogV.value = false
    }

    const findUser = async (u, pageSize, currentPage, total, userList) => {
        searchUser.value = u
        console.log('searchUser = ', searchUser.value)
        searchDialogV.value = false
        // 向服务器端做查询，然后将结果放到userList中
        try {
            let res = await request.get(LogoutUserList, {
                params: {
                    pageSize: pageSize.value,
                    pageNum: currentPage.value,
                    ...searchUser.value
                }
            })
            if (res.code === 200) {
                total.value = res.data.total
                userList.value = res.data.users
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
        searchDialogV, searchUser, changeSearchDialog, closeSearchDialog, findUser
    }
}