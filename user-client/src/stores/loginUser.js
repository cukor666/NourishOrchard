import {ref} from 'vue'
import {defineStore} from 'pinia'

export const useLoginUserStore = defineStore('loginUser', () => {
    const username = ref('')
    const promise = ref('user')
    const setUsername = (value) => {
        username.value = value
    }

    const setPromise = (value) => {
        promise.value = value
    }

    return {username, promise, setUsername, setPromise}
})
