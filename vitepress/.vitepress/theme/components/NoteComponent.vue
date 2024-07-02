<template>
    <el-row>
    <el-col :span="6" :offset="20">
      <div class="mb-4">
        <el-button @click="dialogSubmitState = true">Add</el-button>
      </div>
    </el-col>
  </el-row>
  

  <el-table :data="tableData.data" style="width: 100%" :show-header="false">
    <el-table-column prop="data" label="Date"/>
    <el-table-column fixed="right" width="50">
        <template #default="scope">
          <el-button color="#F56C6C" :icon="Delete" @click="onClickDelete(scope.row.id)"/>
        </template>
      </el-table-column>
  </el-table>

  <el-dialog class="submit-dialog" v-model="dialogSubmitState" title="Add Note" width="650">
    <el-form class="dialog-form" :model="formData.data">
      <el-form-item label="Data">
        <el-input v-model="formData.data.data" placeholder="Input data here"></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <span>
        <el-button @click="dialogSubmitState = false"> Cancel </el-button>
        <el-button type="primary" @click="onSubmit"> Confirm </el-button>
      </span>
    </template>
  </el-dialog>

  <el-dialog class="delete-dialog" v-model="dialogDeleteState" title="Are you sure to delete this data?" width="380">
    <template #footer>
      <span>
        <el-button @click="dialogDeleteState = false"> Cancel </el-button>
        <el-button type="primary" @click="onDelete()"> Confirm </el-button>
      </span>
    </template>
  </el-dialog>


</template>

<script setup>

const props = defineProps(['page'])

import { ref, reactive, onMounted } from 'vue'
import { Delete, Edit } from '@element-plus/icons-vue'
import axios from 'axios'

const dialogSubmitState = ref(false)
const dialogDeleteState = ref(false)
const dialogEditState = ref(false)
let deleteColumnId = 0

const tableData = reactive({ data: [] })
const formData = reactive({ data: {} })
const editData = reactive({ data: {} })

function onSubmit() {
  formData.data.type = props.page
  axios.post('/data/' + props.page + '/add/', formData.data)
  .then(res => {
    formData.data = {}
    dialogSubmitState.value = false
    getLinks(props.page)
  })
  .catch(err => {
    alert(err)
  })
}

function onClickEdit(data) {
  editData.data = JSON.parse(JSON.stringify(data))
  dialogEditState.value = true
}

function onEdit() {
  axios.post('/data/' + props.page + '/update/', editData.data)
  .then(res => {
    dialogEditState.value = false
    getLinks(props.page)
  })
  .catch(err => {
    alert(err)
  })
}

function onClickDelete(id) {
  deleteColumnId = id
  dialogDeleteState.value = true
}

function onDelete() {
  axios.get('/data/' + props.page + '/delete/' + deleteColumnId)
  .then(res => {
    dialogDeleteState.value = false
    getLinks(props.page)
  })
  .catch(err => {
    alert(err)
  })
}

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

.mb-4 {
  margin-top: 5px;
}
</style>