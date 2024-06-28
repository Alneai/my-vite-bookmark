<template>
  <el-table :data="tableData.data" style="width: 100%" :show-header="false">
    <el-table-column prop="data" label="Date"/>
  </el-table>
</template>

<script setup>

const props = defineProps(['page'])

import { reactive, onMounted } from 'vue'
import axios from 'axios'

const tableData = reactive({ data: [] })

function getLinks(page) {
  axios.get('/data/' + page)
  .then(res => {
    tableData.data = res.data
  }).catch(err => {
    alert(err)
  })
}

onMounted(() => { 
    getLinks(props.page)
})
</script>

<style> 
.el-table__inner-wrapper::before {
  height: 0 !important;
}
</style>