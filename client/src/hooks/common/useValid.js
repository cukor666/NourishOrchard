import {usePromise} from "@/hooks/login/usePromise";

const {promiseList} = usePromise();

export function useValid(errWord) {
    // 校验账号
    const validAccount = (account) => {
        console.log('生效')
        if (account === "") {
            errWord.account = "账号不能为空";
        } else if (account.length < 3) {
            errWord.account = "账号字符长度不能小于3";
        } else {
            errWord.account = "";
        }
    };

    // 校验密码
    const validPassword = (password) => {
        let minLen = 3;
        if (password === "") {
            errWord.password = "密码不能为空";
        } else if (password.length < minLen || password.length > 100) {
            errWord.password = "密码长度至少" + minLen + "位";
        } else {
            errWord.password = "";
        }
    };

    // 校验再次确认密码
    const validPassword2 = (password, password2) => {
        let str = password !== password2 ? "两个密码不一致" : "";
        errWord.password = str;
        errWord.password2 = str;
    };

    // 校验权限
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

    // 校验姓名
    const validName = (name) => {
        console.log(typeof name)
        let maxLen = 50;
        if (name === "") {
            errWord.name = "姓名不能为空";
            return false
        } else if (name.length > maxLen) {
            errWord.name = "名字不能超过" + maxLen + "个字";
            return false
        } else {
            errWord.name = "";
            return true
        }
    };

    // 校验性别
    const validGender = (gender) => {
        if (gender === "男" || gender === "女") {
            errWord.gender = "";
        } else {
            errWord.gender = "性别错误";
        }
    };

    // 校验电话号码
    const validPhone = (phone) => {
        let regs =
            /^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$/;
        if (!regs.test(phone)) {
            errWord.phone = "请输入正确的电话号码";
        } else {
            errWord.phone = "";
        }
    };

    // 校验家庭地址
    const validAddress = (address) => {
        let minLen = 3;
        if (address === "") {
            errWord.address = "地址不能为空";
        } else if (address.length < minLen) {
            errWord.address = "地址长度不能小于" + minLen;
        } else {
            errWord.address = "";
        }
    };

    // 判断用户是否满18岁
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

    // 校验用户出生日期
    const validBirthday = (birthday) => {
        if (birthday === "") {
            errWord.birthday = "出生日期不能为空";
        } else if (!isAtleast18YearsOld(birthday)) {
            errWord.birthday = "您未满18岁";
        } else {
            errWord.birthday = "";
        }
    };

    // 校验管理员邮箱
    const validEmail = (email) => {
        let regStr = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
        if (!regStr.test(String(email).toLowerCase())) {
            errWord.email = "请输入正确的邮箱";
        } else {
            errWord.email = "";
        }
    };

    // 校验员工职位
    const validPosition = (position) => {
        if (position.trim() === '采购员' || position.trim() === '仓库管理员') {
            errWord.position = ''
        } else {
            errWord.position = '职位不正确'
        }
    }

    // 校验员工薪资
    const validSalary = (salary) => {
        if (salary < 3000 || salary > 7900) {
            errWord.salary = '工资范围不正确'
        } else {
            errWord.salary = ''
        }
    }

    const validId = (id) => {
        let typeStr = typeof (id)
        if (typeStr !== 'number') {
            errWord.id = 'id的类型不是number类型'
        } else if (id < 0) {
            errWord.id = 'id的值不能为负数'
        } else {
            errWord.id = ''
        }
    }

    return {
        validAccount, validPassword, validPassword2, validPromise, validName, validGender,
        validPhone, validAddress, validBirthday, validEmail, validPosition, validSalary, validId
    };
}
