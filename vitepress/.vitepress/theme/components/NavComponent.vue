<template>

<el-container>
  <el-main>

    <el-button
      class="add-item-button"
      style="width: 100%" color="#626aef"
      @click="dialogSubmitState = true">
      Add URL
    </el-button>

    <el-table :data="linkData.data">
      <el-table-column label="URL" width="250">
        <template #default="scope">
          <el-link type="primary" :href="scope.row.link" target="_blank"> {{ scope.row.title }} </el-link>
        </template>
      </el-table-column>
      <el-table-column prop="description" label="Description"></el-table-column>
      <el-table-column fixed="right" width="130">
        <template #default="scope">
          <el-button type="primary" :icon="Edit" @click="onClickEdit(scope.row)"/>
          <el-button color="#F56C6C" :icon="Delete" @click="onClickDelete(scope.row.id)"/>
        </template>
      </el-table-column>
    </el-table>

  </el-main>
</el-container>
  
<el-dialog class="submit-dialog" v-model="dialogSubmitState" title="Add URL" width="650">
    <el-form class="dialog-form" :model="formData.data" label-width="100px">
      <el-form-item label="Title" required>
        <el-input v-model="formData.data.title" placeholder="Input title here"></el-input>
      </el-form-item>
      <el-form-item label="URL" required>
        <el-input v-model="formData.data.link" placeholder="Input url here"></el-input>
      </el-form-item>
      <el-form-item label="Description">
        <el-input v-model="formData.data.description" placeholder="Input description here"></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <span>
        <el-button @click="dialogSubmitState = false"> Cancel </el-button>
        <el-button type="primary" @click="onSubmit"> Confirm </el-button>
      </span>
    </template>
  </el-dialog>

  <el-dialog class="delete-dialog" v-model="dialogDeleteState" title="Are you sure to delete this url?" width="380">
    <template #footer>
      <span>
        <el-button @click="dialogDeleteState = false"> Cancel </el-button>
        <el-button type="primary" @click="onDelete()"> Confirm </el-button>
      </span>
    </template>
  </el-dialog>

  <el-dialog class="edit-dialog" v-model="dialogEditState" title="Edit URL" width="650">
    <el-form class="dialog-form" :model="editData.data" label-width="100px">
      <el-form-item label="Title" required>
        <el-input v-model="editData.data.title" placeholder="Input title here"></el-input>
      </el-form-item>
      <el-form-item label="URL" required>
        <el-input v-model="editData.data.link" placeholder="Input url here"></el-input>
      </el-form-item>
      <el-form-item label="Description">
        <el-input v-model="editData.data.description" placeholder="Input description here"></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <span>
        <el-button @click="dialogEditState = false"> Cancel </el-button>
        <el-button type="primary" @click="onEdit"> Confirm </el-button>
      </span>
    </template>
  </el-dialog>

</template>

<script setup>

const props = defineProps(['page'])

import { ref, reactive, onMounted} from 'vue'
import { Delete, Edit } from '@element-plus/icons-vue'
import axios from 'axios'

const dialogSubmitState = ref(false)    // 添加对话框
const dialogDeleteState = ref(false)    // 删除对话框
const dialogEditState = ref(false)      // 编辑对话框
let deleteColumnId = 0

const linkData = reactive({ data: [] })  // 页面数据
const formData = reactive({ data: {} })  // 添加暂存数据
const editData = reactive({ data: {} })  // 编辑暂存数据

function onSubmit() {
  formData.data.type = props.page
  axios.post('/link/add', formData.data)
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
  // 深拷贝，防止修改页面值
  editData.data = JSON.parse(JSON.stringify(data))
  dialogEditState.value = true
}

function onEdit() {
  axios.post('/link/update', editData.data)
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
  axios.get('/link/delete/' + deleteColumnId)
  .then(res => {
    dialogDeleteState.value = false
    getLinks(props.page)
  })
  .catch(err => {
    alert(err)
  })
}

function getLinks(page) {
  axios.get('/link/' + page)
  .then(res => {
    linkData.data = res.data
    console.log(res)
  })
  .catch(err => {
    alert(err)
  })
}

onMounted(() => { 
    getLinks(props.page)
})
</script>
  
<style>

@media screen and (width <= 500px) {
  .delete-dialog {
    min-width: 280px !important;
  }
  .submit-dialog {
    min-width: 380px !important;
  }

  .edit-dialog {
    min-width: 380px !important;
  }
}

.add-item-button {
  margin-bottom: 20px;
}

.dialog-form {
  margin-right: 20px;
}
</style>