import request from "@/axios/request";
import {ElMessage} from "element-plus";

export function useEmployeeInfo(employee, selfInfoDialogVisible) {
    const updateEmployeeInfo = () => {
        request
            .put("/account/update", {
                ...employee.value,
            })
            .then(res => {
                console.log(res)
                if (res.code === 200) {
                    ElMessage({
                        message: "更新成功",
                        type: "success",
                    });
                } else {
                    ElMessage({
                        message: "更新失败",
                        type: "error",
                    });
                }
            })
            .catch((err) => {
                ElMessage({
                    message: "系统错误",
                    type: "error",
                });
                console.log(err);
            });

        selfInfoDialogVisible.value = false;
    };

    return {
        updateEmployeeInfo
    }
}