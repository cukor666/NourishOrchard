import { ref, watch } from "vue";
import request from "@/axios/request";
import { ElMessage } from "element-plus";

export function useUserInfo(user) {
  const selfInfoDialogVisiable = ref(false);
  const updateUserInfo = () => {
    console.log(user);

    request
      .put("/account/update", {
        ...user,
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

    selfInfoDialogVisiable.value = false;
  };

  const gender = ref(true);
  watch(gender, (isMan) => {
    user.gender = isMan ? "男" : "女";
  });
  return {
    selfInfoDialogVisiable,
    gender,
    updateUserInfo,
  };
}
