import request from "@/axios/request.js";
import {ElMessage} from "element-plus";
import {useSessionKey} from "@/hooks/common/useSessionKey.js";

const {NourishLogoutUserTotal, NourishLogoutUserList, NourishLogoutUserInfoUpdated} = useSessionKey()

export function useTable(userList, pageSize, currentPage, total) {
    // 更新用户列表
    const updateList = async (searchUser) => {
        try {
            let res = await request.get('/user/logout-list', {
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

    // 恢复用户
    const recoverUser = async (item) => {
        try {
            let res = await request.post('/user/recover', {
                username: item.username
            })
            if (res.code === 200) {
                console.log(res.data)
                ElMessage({message: '恢复用户成功', type: 'success'})
            } else {
                ElMessage({message: '恢复用户失败', type: 'error'})
            }
            res = await request.get('/user/logout-list', {
                params: {
                    pageSize: pageSize.value,
                    pageNum: currentPage.value
                }
            })
            if (res.code === 200) {
                console.log('更新注销用户列表成功')
                let v = res.data
                total.value = v.total
                userList.value = v.users
                sessionStorage.setItem(NourishLogoutUserList, JSON.stringify(userList.value))
                sessionStorage.setItem(NourishLogoutUserTotal, total.value.toString())
                sessionStorage.removeItem(NourishLogoutUserInfoUpdated)
            } else {
                ElMessage({message: '参数错误', type: 'error'})
            }
        } catch (err) {
            console.error(err)
            ElMessage({message: '系统错误', type: 'error'})
        }
    }

    // 彻底删除用户
    const deleteUser = async (item) => {
        console.log(item)
        try {
            let res = await request.delete('/user/remove', {
                params: {
                    id: item.id,
                    username: item.username
                }
            })
            if (res.code === 200) {
                ElMessage({message: '删除成功', type: 'success'})
            } else {
                console.log(res.msg)
                ElMessage({message: '参数错误，删除失败', type: 'error'})
            }
        } catch (e) {
            console.error(e)
            ElMessage({message: '服务器错误', type: 'error'})
        }
    }

    return {
        updateList, recoverUser, deleteUser
    }
}