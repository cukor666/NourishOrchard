
import {ref} from "vue";
import request from "@/axios/request.js";
import {ElMessage} from "element-plus";
import {EmpList} from "@/api/emp/emp-api.js";

export function useSearch() {
    const searchDialogV = ref(false)
    const searchEmp = ref({
        id: 0,
        username: '',
        name: '',
        phone: '',
        position: '',
        salary: 0
    })

    const changeSearchDialog = () => {
        searchDialogV.value = true
    }
    const closeSearchDialog = () => {
        searchDialogV.value = false
    }

    const findEmp = async (u, pageSize, currentPage, total, empList) => {
        searchEmp.value = u
        console.log('searchEmp = ', searchEmp.value)
        searchDialogV.value = false
        // 向服务器端做查询，然后将结果放到userList中
        try {
            let res = await request.get(EmpList, {
                params: {
                    pageSize: pageSize.value,
                    pageNum: currentPage.value,
                    ...searchEmp.value
                }
            })
            if (res.code === 200) {
                total.value = res.data.total
                empList.value = res.data.employees
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
        searchDialogV, searchEmp, changeSearchDialog, closeSearchDialog, findEmp
    }
}