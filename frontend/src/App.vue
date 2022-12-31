<script setup>
import ConnectionList from './components/ConnectionList.vue'
import ConnectionManage from "./components/ConnectionManage.vue";
import {ref} from "vue";
import {KeyList} from "../wailsjs/go/main/App.js";

let flushFlag = ref(true)

function flushConnectionList() {
  flushFlag.value = !flushFlag.value
}

// 获取基本信息
function test() {
  // 获取数据库列表
  KeyList({
    conn_identity: "87e1b075-8787-11ed-a178-00ffaabbccdd",
    db: 0
  }).then((result) => {
    console.log('result', result)
  })
}

test()

</script>


<template>
  <el-row>
    <el-col :span="6" style="height: 100vh;margin: 12px">
      <div style="margin-bottom: 12px">
        <ConnectionManage title="新建连接" btn-type="primary" @emit-connection-list="flushConnectionList"/>
      </div>
      <ConnectionList :flush="flushFlag"></ConnectionList>
    </el-col>
  </el-row>
</template>

<style>
</style>
