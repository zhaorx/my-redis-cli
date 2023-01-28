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
    <div v-for="item in keys" @click="selectKeyKey(item)">
      <div v-if="item === selectKey" class="my-item">{{ item }}</div>
      <div v-else class="my-select-item">{{ item }}</div>
    </div>

  </main>
</template>

<script setup>
import {reactive, ref, watch} from 'vue'
import {KeyList} from "../../wailsjs/go/main/App.js";
import {ElNotification} from "element-plus";

let props = defineProps(['keyDB', 'keyConnIdentity'])
let emits = defineEmits(['emit-select-key'])
let keys = ref()
let selectKey = ref()

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

function selectKeyKey(key) {
  selectKey.value = key
  emits("emit-select-key",key)
}
</script>

<style scoped>

</style>
