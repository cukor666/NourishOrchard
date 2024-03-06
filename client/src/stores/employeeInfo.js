import {defineStore} from "pinia";
import {ref} from "vue";

export const useEmployeeInfoStore = defineStore("employeeInfoStore", () => {
    const employee = ref({
        username: "",
        name: "",
        phone: "",
        position: "",
        salary: 0
    })

    const setEmployee = (newEmp) => {
        employee.value = newEmp
    }
    return {
        employee,
        setEmployee
    }
})