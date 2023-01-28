<template>
  <el-row>
    <el-col :span="5" style="height: 100vh;margin: 12px">
      <div style="margin-bottom: 12px">
        <ConnectionManage title="新建连接" btn-type="primary" @emit-connection-list="flushConnectionList"/>
      </div>
      <ConnectionList :flush="flushFlag" @emit-select-db="selectDB"></ConnectionList>
    </el-col>
    <el-col :span="7" style="padding:12px">
      <Keys :keyDB="keyDB" :keyConnIdentity="keyConnIdentity" @emit-select-key="selectKey"/>
    </el-col>
    <el-col :span="12" style="padding:12px">
      <KeyValue :keyDB="keyDB" :keyConnIdentity="keyConnIdentity" :keyKey="keyKey" />
    </el-col>
  </el-row>
</template>

<script setup>
import ConnectionList from './components/ConnectionList.vue'
import ConnectionManage from "./components/ConnectionManage.vue";
import {ref} from "vue";
import {KeyList} from "../wailsjs/go/main/App.js";
import Keys from "./components/Keys.vue";
import KeyValue from "./components/KeyValue.vue";

let flushFlag = ref(true)
let keyDB = ref()
let keyKey = ref()
let keyConnIdentity = ref()

function flushConnectionList() {
  flushFlag.value = !flushFlag.value
}

// 选中数据库
function selectDB(db, connIdentity) {
  keyDB.value = db
  keyConnIdentity.value = connIdentity
}
// 选中key
function selectKey(key) {
  keyKey.value = key
}

function GetKeyValue(){

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

<style>
</style>
