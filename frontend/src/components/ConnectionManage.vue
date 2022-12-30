<template>
  <main>
    <el-button :type="btnType" @click="dialogVisible = true">{{ title }}</el-button>

    <el-dialog v-model="dialogVisible" :title="title" width="60%">
      <el-form :model="form" label-width="120px">
        <el-form-item label="连接名称">
          <el-input v-model="form.name" placeholder="请输入连接名称"/>
        </el-form-item>
        <el-form-item label="连接地址">
          <el-input v-model="form.addr" placeholder="请输入连接地址"/>
        </el-form-item>
        <el-form-item label="端口">
          <el-input v-model="form.port" placeholder="请输入端口号"/>
        </el-form-item>
        <el-form-item label="用户名">
          <el-input v-model="form.username" placeholder="请输入用户名"/>
        </el-form-item>
        <el-form-item label="密码">
          <el-input v-model="form.password" placeholder="请输入密码"/>
        </el-form-item>

        <el-form-item>
          <el-button @click="editConnection" v-if="data">编辑</el-button>
          <el-button @click="createConnection" v-else>创建</el-button>
          <el-button @click="dialogVisible = false">取消</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
  </main>
</template>

<script setup>
import {reactive, ref} from "vue";
import {ConnectionCreate, ConnectionEdit} from "../../wailsjs/go/main/App.js";
import {ElNotification} from "element-plus";

let props = defineProps(['title', 'btnType', 'data'])
let emits = defineEmits(['emit-connection-list'])
let dialogVisible = ref(false)
let form = reactive({
  name: '',
  addr: '',
  port: '',
  username: '',
  password: ''
})

if (props.data) {
  form = props.data
}

// 创建连接
function createConnection() {
  ConnectionCreate(form).then(result => {
    if (result.code !== 200) {
      ElNotification({
        title: result.msg,
        type: 'error'
      })

      return
    }

    ElNotification({
      title: result.msg,
      type: 'success'
    })

    // 获取新的连接地址
    emits('emit-connection-list')
    dialogVisible.value = false

  })
}

// 编辑连接
function editConnection() {
  ConnectionEdit(form).then(result => {
    if (result.code !== 200) {
      ElNotification({
        title: result.msg,
        type: 'error'
      })

      return
    }

    ElNotification({
      title: result.msg,
      type: 'success'
    })

    // 获取新的连接地址
    emits('emit-connection-list')
    dialogVisible.value = false

  })
}


</script>

<style scoped>

</style>
