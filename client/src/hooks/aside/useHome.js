import router from "@/router/index.js";
import {useAsideStore} from "@/stores/aside.js";
import {useTagsStore} from "@/stores/tags.js";
import {ElMessage} from "element-plus";

const asideStore = useAsideStore()
const tagsStore = useTagsStore()
const {addTag} = tagsStore
export function useHome() {
    const goHome = () => {
        let name = 'Home'
        router.push({name}).then(r => {
            console.log('跳转到首页')
            asideStore.defaultActive = name
            addTag({
                name: '首页',
                routerName: name
            })
        }).catch(err => {
            console.log('路由跳转失败', err)
            ElMessage({
                message: '跳转失败',
                type: 'error'
            })
        })
    }

    return {
        goHome
    }
}