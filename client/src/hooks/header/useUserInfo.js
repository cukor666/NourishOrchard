import request from "@/axios/request";
import { ElMessage } from "element-plus";

export function useUserInfo(user, selfInfoDialogVisible) {
  const updateUserInfo = () => {
    request
      .put("/account/update", {
        ...user.value,
      })
      .then((res) => {
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
    updateUserInfo,
  };
}
