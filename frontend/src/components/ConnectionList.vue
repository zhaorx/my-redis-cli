<template>
  <main>
    <div class="demo-collapse">
      <el-collapse accordion>
        <el-collapse-item :name="item.name" v-for="item in list">
          <template #title>
            {{ item.addr }}
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

</style>
