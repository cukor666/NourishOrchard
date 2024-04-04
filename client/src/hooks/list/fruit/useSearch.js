import {ref} from "vue";
import request from "@/axios/request.js";
import {ElMessage} from "element-plus";
import {FruitList} from "@/api/fruit/fruit-api.js";

export function useSearch() {
    const searchDialogV = ref(false)
    const searchFruit = ref({
        id: 0,
        name: '',
        water: 0.0,
        sugar: 0.0,
        shelfLife: 0,
        origin: '',
        supplierId: 0
    })
    const changeSearchDialog = () => {
        searchDialogV.value = true
    }
    const closeSearchDialog = () => {
        searchDialogV.value = false
    }

    const findFruit = async (u, pageSize, currentPage, total, fruitList) => {
        searchFruit.value = u
        console.log('searchFruit = ', searchFruit.value)
        searchDialogV.value = false
        // 向服务器端做查询，然后将结果放到userList中
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
        searchDialogV, searchFruit, changeSearchDialog, closeSearchDialog, findFruit
    }
}