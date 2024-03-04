import {ref} from 'vue'

export function usePromise() {
    const promiseList = ref([
        {
            label: 'user',
            content: '用户'
        },
        {
            label: 'admin',
            content: '管理员'
        },
        {
            label: 'employee',
            content: '员工'
        }
    ])

    return {
        promiseList
    }
}