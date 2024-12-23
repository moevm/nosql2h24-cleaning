<script setup lang="ts">
import { ref, Ref, onMounted } from 'vue'
import MainContainer from '../../ui/uikit/containers/MainContainer.vue'
import PanelContainer from '../../ui/uikit/containers/PanelContainer.vue'
import HeaderList from '../../ui/uikit/containers/HeaderList.vue'
import ActionButton from '../../ui/uikit/ActionButton.vue'
import InputTextField from '../../ui/uikit/inputs/InputTextField.vue'
import Dialog from '../../ui/uikit/Dialog.vue'
import ServiceItem from '../../ui/uikit/items/ServiceItem.vue'
import { getAllServices, createService } from '../../api/request'
import Service from '../../api/models/service'
import InputNumber from '../../ui/uikit/inputs/InputNumber.vue'

const services = ref<Service[]>([])

let newService = ref<Service>({
  id: '',
  name: '',
  price: 0,
  workers_quantity: 1,
  description: '',
  consumables: [],
  created_at: ''
})

const isDialogVisible: Ref<boolean> = ref(false)

function openDialog(): void {
  isDialogVisible.value = true
}

function closeDialog(): void {
  isDialogVisible.value = false
  newService = ref<Service>({
    id: '',
    name: '',
    price: 0,
    workers_quantity: 1,
    description: '',
    consumables: [],
    created_at: ''
  })
}

function createNewService() {
  if (newService.value.name === '') {
    alert("Название не может быть пустым")
  } else if (newService.value.price < 1) {
    alert("Услуга не может быть бесплатной")
  } else {
    newService.value.created_at = new Date().toISOString();
    createService(newService.value).then((_) => {console.log("Service created")})
    fetchAllServices()
    closeDialog()
  }
}

function fetchAllServices() {
  getAllServices()
  .then((response) => {
    services.value = response;
  })
  .catch((error) => {
    console.log('Error load services:', error);
  });
}

onMounted(() => {
  fetchAllServices()
})
</script>

<template>
  <MainContainer>
    <PanelContainer
      height="10%"
      width="95%"
    >
      <ActionButton
        text="Добавить"
        type="add"
        color="#394cc2"
        @click="openDialog"
      ></ActionButton>
    </PanelContainer>
    <!-- Dialog to add services -->
    <Dialog
      title="Добавление услуги"
      :visible="isDialogVisible"
      dialogMaxWidth="40%"
      @update:visible="closeDialog"
    >
      <template #body>
        <InputTextField
          v-model="newService.name"
          placeholder="Введите название"
          type="text"
          label="Название услуги"
        ></InputTextField>
        <InputNumber
          v-model="newService.price"
          label="Стоимость услуги"
          :min="0"
          :max="10000"
          :step="10"
        ></InputNumber>
        <InputNumber
          v-model="newService.workers_quantity"
          label="Количество исполнителей"
          :min="1"
          :max="10"
          :step="1"
        ></InputNumber>
        <InputTextField
          v-model="newService.description"
          type="text"
          label="Описание"
        ></InputTextField>
      </template>
      <template #footer>
        <ActionButton
          text="Отмена"
          type="cancel"
          variant="flat"
          color="#f65858"
          @click="closeDialog"
        ></ActionButton>
        <ActionButton
          text="Создать"
          type="create"
          variant="flat"
          color="#394cc2"
          @click="createNewService"
        ></ActionButton>
      </template>
    </Dialog>
    <!-- -------------------- -->
    <HeaderList
      title="Услуги" 
      :items="services"
      height="92%"
      width="95%"
    >
     <template #items="{ item }">
        <ServiceItem
          :isAdmin="true"
          :service="item"
          @update-service-item="fetchAllServices">
        </ServiceItem>
      </template>
    </HeaderList>
  </MainContainer>
</template>

<style scoped>
</style>
