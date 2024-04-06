import {useAsideStore} from "@/stores/aside.js";
import {useTagsStore} from "@/stores/tags.js";
import router from "@/router/index.js";
import {ElMessage} from "element-plus";

const asideStore = useAsideStore()
const tagsStore = useTagsStore()
const {addTag} = tagsStore
export function useInventory() {
    const gotoInventoryList = () => {
        let name = 'InventoryList'
        router.push({name}).then(r => {
            asideStore.defaultActive = name
            addTag({
                name: '库存列表',
                routerName: name
            })
        }).catch(err => {
            console.log('跳转到库存列表失败', err)
            ElMessage({message: '无法跳转', type: 'error'})
        })
    }
    return {
        gotoInventoryList
    }
}