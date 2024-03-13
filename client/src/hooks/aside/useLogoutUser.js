import router from "@/router/index.js";
import {useAsideStore} from "@/stores/aside.js";
import {useTagsStore} from "@/stores/tags.js";
import {ElMessage} from "element-plus";

const asideStore = useAsideStore()
const tagsStore = useTagsStore()
const {addTag} = tagsStore

export function useLogoutUser() {
    const gotoLogoutUserList = () => {
        let name = 'LogoutUserList'
        router.push({name}).then(r => {
            console.log('跳转到注销用户列表')
            asideStore.defaultActive = name
            addTag({
                name: '注销用户列表',
                routerName: name
            })
        }).catch(err => {
            console.log('跳转失败')
            ElMessage({
                message: '无法跳转',
                type: 'error'
            })
        })
    }

    return {
        gotoLogoutUserList
    }
}