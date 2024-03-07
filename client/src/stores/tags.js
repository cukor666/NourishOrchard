import {defineStore} from "pinia";
import {ref} from "vue";
import {useAsideStore} from "@/stores/aside.js";
import router from "@/router/index.js";

const asideStore = useAsideStore()

export const useTagsStore = defineStore("tagsStore", () => {
    // 路由标签
    const tags = ref([
        {
            name: '首页',
            routerName: 'Home'
        }
    ])

    const addTag = (tag) => {
        let newArr = tags.value.filter(e => {
            return e.name === tag.name && e.routerName === tag.routerName
        })
        if (newArr.length === 0) {
            tags.value.push(tag)
        }
    }

    const removeTag = (index) => {
        tags.value.splice(index, 1)
    }

    const selectTag = (tag) => {
        console.log(tag)
        asideStore.defaultActive = tag.routerName
        router.push({name: tag.routerName}).then(r => {
            console.log('路由跳转成功', r)
        }).catch(err => {
            console.log('路由跳转失败', err)
        })
    }

    return {
        tags,
        addTag,
        removeTag,
        selectTag
    }
})