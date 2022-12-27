<script setup>
import ConnectionList from './components/ConnectionList.vue'
import ConnectionManage from "./components/ConnectionManage.vue";
import {ref} from "vue";

let flushFlag = ref(true)

function flushConnectionList() {
  flushFlag.value = !flushFlag.value
}
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
