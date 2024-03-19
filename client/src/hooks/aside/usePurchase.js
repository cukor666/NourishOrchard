import {useAsideStore} from "@/stores/aside.js";
import {useTagsStore} from "@/stores/tags.js";
import router from "@/router/index.js";
import {ElMessage} from "element-plus";

const asideStore = useAsideStore()
const tagsStore = useTagsStore()
const {addTag} = tagsStore
export function usePurchase() {
    const gotoPurchaseList = () => {
        let name = 'PurchaseList'
        router.push({name}).then(r => {
            asideStore.defaultActive = name
            addTag({
                name: '采购列表',
                routerName: name
            })
        }).catch(err => {
            console.log('跳转到采购列表失败')
            ElMessage({message: '无法跳转', type: 'error'})
        })
    }
    return {
        gotoPurchaseList
    }
}