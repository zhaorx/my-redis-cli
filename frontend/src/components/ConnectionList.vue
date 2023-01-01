<template>
  <main>
    <div class="demo-collapse">
      <el-collapse accordion>
        <el-collapse-item :name="item.name" v-for="item in list" @click="getInfo(item.identity)">
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
          <div v-for="db in infoDbList">
            <div :class="{'my-select-item':selectDbKey === db.key,'my-item':selectDbKey !== db.key}" @click="selectDb(db,item.identity)">
              {{ db.key }} ( {{ db.number }} )
            </div>
          </div>
        </el-collapse-item>
      </el-collapse>
    </div>

  </main>
</template>

<script setup>
import {ref, watch} from "vue";
import {ConnectionDelete, ConnectionList, DbList} from "../../wailsjs/go/main/App.js";
import {ElNotification} from 'element-plus'
import ConnectionManage from "./ConnectionManage.vue";

let flush = defineProps(['flush'])
watch(flush, () => {
  connectionList()
})

let list = ref()
let infoDbList = ref()
let selectDbKey = ref()
let emits = defineEmits(['emit-select-db'])


function selectDb(db, connIdentity) {
  selectDbKey = db.key

  emits('emit-select-db', Number(db.key.substring(2)), connIdentity)

}

// 连接列表
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

// 删除连接
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

// 获取基本信息
function getInfo(identity) {
  // 获取数据库列表
  DbList(identity).then((result) => {
    if (result.code !== 200) {
      ElNotification({
        title: result.msg,
        type: 'error'
      })

      return
    }

    infoDbList.value = result.data
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
