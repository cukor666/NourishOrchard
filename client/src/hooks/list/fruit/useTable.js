import request from "@/axios/request.js";
import {ElMessage, ElMessageBox} from "element-plus";
import {FruitDelete, FruitDetail, FruitList} from "@/api/fruit/fruit-api.js";
import searchFruit from "@/components/fruit/SearchFruit.vue";

export function useTable(fruitList, pageSize, currentPage, total) {
    const updateList = async (searchFruit) => {
        try {
            let res = await request.get(FruitList, {
                params: {
                    pageSize: pageSize.value,
                    pageNum: currentPage.value,
                    ...searchFruit.value
                }
            })
            if (res.code === 200) {
                total.value = res.data.total
                fruitList.value = res.data.fruits
            } else {
                console.log(res.msg)
                ElMessage({message: '参数错误', type: 'error'})
            }
        } catch (err) {
            console.error(err)
            ElMessage({message: '服务器错误', type: 'error'})
        }
    }

    const showDetail = async (item, detailDialogV, fruit) => {
        console.log(item)
        detailDialogV.value = true

        try {
            let res = await request.get(FruitDetail, {
                params: {
                    id: item.id
                }
            })
            if (res.code === 200) {
                fruit.value = res.data
                ElMessage({message: '查询水果详情成功', type: 'success'})
            } else {
                console.log('查询失败')
                ElMessage({message: '查询水果详情失败', type: 'error'})
            }
        } catch (e) {
            console.log(e)
            ElMessage({message: '获取水果详细信息失败', type: 'error'})
        }

    }

    const delConfirm = async (item) => {
        try {
            let res = await request.delete(FruitDelete, {
                params: {
                    id: item.id
                }
            })
            if (res.code === 200) {
                ElMessage({message: '删除成功', type: 'success'})
                // 更新列表
                updateList(searchFruit)
            } else {
                ElMessage({message: '删除失败', type: 'error'})
            }
        } catch {
            ElMessage({message: '系统错误', type: 'error'})
        }
    }

    const deleteFruit = async (item) => {
        // 删之前再次询问一下
        ElMessageBox.confirm(
            '您确定要删除该水果信息吗？',
            'Warning',
            {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }
        ).then(() => {
            delConfirm(item)
        }).catch(() => {
            ElMessage({message: '取消删除', type: 'info'})
        })
    }


    return {
        updateList, showDetail, deleteFruit
    }
}