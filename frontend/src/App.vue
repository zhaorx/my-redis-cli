<script setup>
import ConnectionList from './components/ConnectionList.vue'
import ConnectionManage from "./components/ConnectionManage.vue";
import {ref} from "vue";
import {DbList} from "../wailsjs/go/main/App.js";

let flushFlag = ref(true)
let dbList = ref([])

function flushConnectionList() {
  flushFlag.value = !flushFlag.value
}

DbList("87e1b075-8787-11ed-a178-00ffaabbccdd").then((res) => {
  dbList.value = res.data
})
</script>


<template>
  <el-row>
    <el-col :span="6" style="height: 100vh;margin: 12px">
      <div style="margin-bottom: 12px">
        <ConnectionManage title="新建连接" btn-type="primary" @emit-connection-list="flushConnectionList"/>
      </div>
      <ConnectionList :flush="flushFlag"></ConnectionList>
    </el-col>
    ###########3 {{ dbList.length }}
  </el-row>
</template>

<style>
</style>
