<template>
  <main>
    <el-form :inline="true" :model="form" class="demo-form-inline">
      <el-form-item>
        <el-input v-model="form.keyword" placeholder="请输入键的信息"/>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="getKeyList">查询</el-button>
      </el-form-item>
    </el-form>

    <div v-for="item in keys" class="my-item">{{ item }}</div>
  </main>
</template>

<script setup>
import {reactive, ref, watch} from 'vue'
import {KeyList} from "../../wailsjs/go/main/App.js";
import {ElNotification} from "element-plus";

let props = defineProps(['keyDB', 'keyConnIdentity'])
let keys = ref()

const form = reactive({
  keyword: '',
})

watch(props, () => {
  getKeyList()
})

function getKeyList() {
  KeyList({
    db: props.keyDB,
    conn_identity: props.keyConnIdentity,
    keyword: form.keyword,
  }).then((res) => {
    if (res.code !== 200) {
      ElNotification({
        title: res.msg,
        type: 'error'
      })

      return
    }

    keys.value = res.data
  })
}
</script>

<style scoped>

</style>
