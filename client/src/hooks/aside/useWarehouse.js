import {useAsideStore} from "@/stores/aside.js";
import {useTagsStore} from "@/stores/tags.js";
import router from "@/router/index.js";
import {ElMessage} from "element-plus";

const asideStore = useAsideStore()
const tagsStore = useTagsStore()
const {addTag} = tagsStore
export function useWarehouse() {
    const gotoWarehouseList = () => {
        let name = 'WarehouseList'
        router.push({name}).then(r => {
            console.log('跳转到仓库列表')
            asideStore.defaultActive = name
            addTag({
                name: '仓库列表',
                routerName: name
            })
        }).catch(err => {
            console.log('跳转到仓库列表失败', err)
            ElMessage({message: '无法跳转', type: 'error'})
        })
    }
    return {
        gotoWarehouseList
    }
}