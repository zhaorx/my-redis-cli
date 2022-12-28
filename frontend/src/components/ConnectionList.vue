<template>
  <main>
    <div class="demo-collapse">
      <el-collapse accordion>
        <el-collapse-item :name="item.name" v-for="item in list">
          <template #title>

            <div class="item">
              <div>{{ item.addr }}</div>
              <div style="display: flex">
                <ConnectionManage :data="item" title="编辑" btn-type="text" @click.stop/>
                <el-popconfirm title="确认删除？" @confirm="connectionDelete(item.identity)" confirm-button-text="是" cancel-button-text="否">
                  <template #reference>
                    <el-button link type=danger @click.stop>删除</el-button>
                  </template>
                </el-popconfirm>
              </div>
            </div>

          </template>
        </el-collapse-item>
      </el-collapse>
    </div>

  </main>
</template>

<script setup>
import {ref, watch} from "vue";
import {ConnectionDelete, ConnectionList} from "../../wailsjs/go/main/App.js";
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

function connectionDelete(identity) {
  ConnectionDelete(identity).then(result => {
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

    connectionList()
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
