import {useAsideStore} from "@/stores/aside.js";
import {useTagsStore} from "@/stores/tags.js";
import router from "@/router/index.js";
import {ElMessage} from "element-plus";

const asideStore = useAsideStore()
const tagsStore = useTagsStore()
const {addTag} = tagsStore

export function useOrders() {
    const gotoOrdersList = () => {
        let name = 'OrdersList'
        router.push({name}).then(r => {
            asideStore.defaultActive = name
            addTag({
                name: '订单列表',
                routerName: name
            })
        }).catch(err => {
            console.log('无法跳转到订单列表', err)
            ElMessage({message: '无法跳转', type: 'error'})
        })
    }
    return {
        gotoOrdersList
    }
}