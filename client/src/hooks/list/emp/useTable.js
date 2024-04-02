import request from "@/axios/request.js";
import {ElMessage} from "element-plus";
import {EmpList} from "@/api/emp/emp-api.js";

export function useTable(empList, pageSize, currentPage, total) {
    const updateList = async (searchEmp) => {
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
            } else {
                console.log(res.msg)
                ElMessage({message: '参数错误', type: 'error'})
            }
        } catch (err) {
            console.error(err)
            ElMessage({message: '服务器错误', type: 'error'})
        }
    }

    const updateInfo = () => {
        console.log('更新员工信息')
    }

    const promotion = () => {
        console.log('晋升管理员')
    }


    return {
        updateList, updateInfo, promotion
    }
}