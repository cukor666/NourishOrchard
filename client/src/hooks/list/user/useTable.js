import request from "@/axios/request.js";
import {ElMessage} from "element-plus";

export function useTable(userList, pageSize, currentPage, total) {
    const updateList = async (searchUser) => {
        try {
            let res = await request.get('/user/list', {
                params: {
                    pageSize: pageSize.value,
                    pageNum: currentPage.value,
                    ...searchUser.value
                }
            })
            if (res.code === 200) {
                total.value = res.data.total
                userList.value = res.data.users
            } else {
                console.log(res.msg)
                ElMessage({message: '参数错误', type: 'error'})
            }
        } catch (err) {
            console.error(err)
            ElMessage({message: '服务器错误', type: 'error'})
        }
    }

    const showDetail = (item, detailDialogV, user) => {
        console.log(item)
        detailDialogV.value = true
        user.value = item
    }

    const deleteUser = async (item) => {
        try {
            let res = await request.delete('/user/delete', {
                params: {
                    username: item.username
                }
            })
            if (res.code === 200) {
                ElMessage({message: '删除成功', type: 'success'})
            } else {
                ElMessage({message: '删除失败', type: 'error'})
            }
            res = await request.get('/user/list', {
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
        updateList, showDetail, deleteUser
    }
}