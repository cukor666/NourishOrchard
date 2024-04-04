import request from "@/axios/request.js";
import {ElMessage} from "element-plus";
import {FruitDelete, FruitDetail, FruitList} from "@/api/fruit/fruit-api.js";

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

    const deleteFruit = async (item) => {
        try {
            let res = await request.delete(FruitDelete, {
                params: {
                    id: item.id
                }
            })
            if (res.code === 200) {
                ElMessage({message: '删除成功', type: 'success'})
            } else {
                ElMessage({message: '删除失败', type: 'error'})
            }
            res = await request.get(FruitList, {
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
        updateList, showDetail, deleteFruit
    }
}