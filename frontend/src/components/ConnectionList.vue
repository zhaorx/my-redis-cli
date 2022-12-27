<template>
  <main>
    <div class="demo-collapse">
      <el-collapse accordion>
        <el-collapse-item :name="item.name" v-for="item in list">
          <template #title>

            <div class="item">
              <span>{{ item.addr }}</span>
              <span><ConnectionManage :data="item" title="编辑" btn-type="text" @click.stop/></span>
            </div>

          </template>
        </el-collapse-item>
      </el-collapse>
    </div>

  </main>
</template>

<script setup>
import {ref, watch} from "vue";
import {ConnectionList} from "../../wailsjs/go/main/App.js";
import {ElNotification} from 'element-plus'
import ConnectionManage from "./ConnectionManage.vue";

let flush = defineProps(['flush'])
watch(flush, (val) => {
  connectionList()
})

let list = ref()

function connectionList() {
  ConnectionList().then(result => {
    if (result.code !== 200) {
      ElNotification({
        title: result.msg,
        type: 'error'
      })
    }
    list.value = result.data
  })
}

connectionList()


</script>

<style scoped>
.item {
  display: flex;
  width: 100%;
  justify-content: space-between;
}
</style>
