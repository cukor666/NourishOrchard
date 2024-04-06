import {ref} from "vue";

export function useEmployee() {
    const employee = ref({
        id: 0,
        username: "",
        name: "",
        phone: "",
        position: "",
        salary: 0,
        salaryA: 0,
        salaryB: 0
    })

    return {
        employee
    }
}