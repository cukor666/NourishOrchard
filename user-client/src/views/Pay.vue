<template>
  <div class="container">
    <h1>支付页面</h1>
    <el-card>
      <div class="qrcode"></div>
      <el-button type="primary" :disabled="disabled" @click="pay">模拟支付</el-button>
    </el-card>
  </div>

  <el-dialog title="接收人信息" v-model="showDialog">
    <el-form :model="receiver" label-width="80px" ref="receiverRef" :rules="rules">
      <el-form-item label="姓名" prop="name">
        <el-input v-model="receiver.name"></el-input>
      </el-form-item>
      <el-form-item label="手机号" prop="phone">
        <el-input v-model="receiver.phone"></el-input>
      </el-form-item>
      <el-form-item label="地址" prop="address">
        <el-input v-model="receiver.address"></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button type="primary" @click="confirmReceiver">确定</el-button>
    </template>
  </el-dialog>
</template>


<script setup>

import {ElMessage} from "element-plus";
import {reactive, ref} from "vue";
import request from "@/axios/request.js";
import {AddOrder} from "@/api/order-api.js";
import router from "@/router/index.js";

const disabled = ref(false)

const receiver = reactive({
  name: '张三',
  phone: '18577659826',
  address: '北京市海淀区西二旗北路10号'
})

const receiverRef = ref(null)

const validName = (rule, value, callback) => {
  if (value === '') {
    callback(new Error('姓名不能为空'))
  } else {
    if (value.length < 2 || value.length > 10) {
      callback(new Error('姓名长度必须在2-10个字符之间'))
    } else {
      callback()
    }
  }
}

const validPhone = (rule, value, callback) => {
  if (value === '') {
    callback(new Error('手机号不能为空'))
  } else {
    if (!/^1[34578]\d{9}$/.test(value)) {
      callback(new Error('手机号格式不正确'))
    } else {
      callback()
    }
  }
}

const rules = ref({
  name: [{ required: true, message: '请输入姓名', trigger: 'blur' },
    {validator: validName, trigger: 'blur'}],
  phone: [{ required: true, message: '请输入手机号', trigger: 'blur' },
    {validator: validPhone, trigger: 'blur'}],
  address: [{ required: true, message: '请输入地址', trigger: 'blur' }]
})

const showDialog = ref(false)
const pay = () => {
  disabled.value = true
  showDialog.value = true
  setTimeout(() =>  {
    disabled.value = false
  }, 2000)
}

const confirmReceiver = () => {
  receiverRef.value.validate((valid) => {
    if (valid) {
      // 模拟支付成功
      // 向后端发送添加订单的请求
      request.post(AddOrder, {
        title: '支付测试订单标题',
        commodityId: 6,
        quantity: 1,
        receiverName: receiver.name,
        receiverPhone: receiver.phone,
        address: receiver.address
      }).then(res => {
        console.log(res)
        if (res.code === 200) {
          ElMessage.success('支付成功')
          showDialog.value = false
          router.push('/order')
        } else  {
          throw new Error('支付失败')
        }
      }).catch(err => {
        console.log(err)
        ElMessage.error('支付失败')
      })
    } else {
      ElMessage.error('请检查输入信息')
    }
  })
}

</script>



<style scoped lang="scss">
.container {
  width: 100%;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
}
.el-card {
  border-radius: 15px;
  margin-top: 50px;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  .el-button {
    transform: translate(154px, 10px);
  }
}
.qrcode {
  width: 400px;
  height: 300px;
  background-color: #f2f2f2;
  background-image: url("@/assets/收付款.png");
  background-size: 400px;
  background-position: 0px -85px;
  background-repeat: no-repeat;
}
</style>