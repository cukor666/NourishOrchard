import router from "@/router/index.js";
import {ElMessage} from "element-plus";
import {useAsideStore} from "@/stores/aside.js";
import {useTagsStore} from "@/stores/tags.js";

const asideStore = useAsideStore()
const tagsStore = useTagsStore()
const {addTag} = tagsStore

export function useFruit() {
    const gotoFruitList = () => {
        let name = 'FruitList'
        router.push({name}).then(r => {
            console.log('跳转到水果列表')
            asideStore.defaultActive = name
            addTag({
                name: '水果列表',
                routerName: name
            })
        }).catch(err => {
            console.log('调整到水果列表失败', err)
            ElMessage({
                message: '无法跳转',
                type: 'error'
            })
        })
    }

    return {
        gotoFruitList
    }
}