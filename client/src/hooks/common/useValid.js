import { usePromise } from "@/hooks/login/usePromise";

const { promiseList } = usePromise();

export function useValid(errWord) {
  const validAccount = (account) => {
    if (account == "") {
      errWord.account = "账号不能为空";
      return
    } else if (account.length < 3) {
      errWord.account = "账号字符长度不能小于3";
    } else {
      errWord.account = "";
    }
  };

  const validPassword = (password) => {
    let minLen = 3;
    if (password == "") {
      errWord.password = "密码不能为空";
      return
    } else if (password.length < minLen || password.length > 100) {
      errWord.password = "密码长度至少" + minLen + "位";
    } else {
      errWord.password = "";
    }
  };

  const validPassword2 = (password, password2) => {
    let str = password !== password2 ? "两个密码不一致" : "";
    errWord.password = str;
    errWord.password2 = str;
  };

  const validPromise = (promise) => {
    let exists = false;
    for (let i = 0; i < promiseList.value.length; i++) {
      if (promise === promiseList.value[i].label) {
        exists = true;
        break;
      }
    }
    if (!exists) {
      errWord.promise = "未知权限";
    } else {
      errWord.promise = "";
    }
  };

  const validName = (name) => {
    let maxLen = 50;
    if (name == "") {
      errWord.name = "姓名不能为空";
    } else if (name.length > maxLen) {
      errWord.name = "名字不能超过" + maxLen + "个字";
    } else {
      errWord.name = "";
    }
  };

  const validGender = (gender) => {
    if (gender === "男" || gender === "女") {
      errWord.gender = "";
    } else {
      errWord.gender = "性别错误";
    }
  };

  const validPhone = (phone) => {
    let regs =
      /^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$/;
    if (!regs.test(phone)) {
      errWord.phone = "请输入正确的电话号码";
    } else {
      errWord.phone = "";
    }
  };

  const validAddress = (address) => {
    let minLen = 3;
    if (address === "") {
      errWord.address = "地址不能为空";
      return
    } else if (address.length < minLen) {
      errWord.address = "地址长度不能小于" + minLen;
    } else {
      errWord.address = "";
    }
  };

  const isAtleast18YearsOld = (dataString) => {
    const birthdayDate = new Date(dataString);
    const todayDate = new Date();
    let age = todayDate.getFullYear() - birthdayDate.getFullYear();
    const m = todayDate.getMonth() - birthdayDate.getMonth();
    if (m < 0 || (m === 0 && todayDate.getDate() < birthdayDate.getDate())) {
      age--;
    }
    return age >= 18;
  };

  const validBirthday = (birthday) => {
    if (birthday === "") {
      errWord.birthday = "出生日期不能为空";
      return
    } else if (!isAtleast18YearsOld(birthday)) {
      errWord.birthday = "您未满18岁";
    } else {
      errWord.birthday = "";
    }
  };

  return {
    validAccount,
    validPassword,
    validPassword2,
    validPromise,
    validName,
    validGender,
    validPhone,
    validAddress,
    validBirthday,
  };
}
