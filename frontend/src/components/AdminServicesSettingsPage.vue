<script setup lang="ts">
import { ref, Ref } from 'vue'
import MainContainer from '../ui/uikit/MainContainer.vue'
import PanelContainer from '../ui/uikit/PanelContainer.vue'
import HeaderList from '../ui/uikit/HeaderList.vue'
import ActionButton from '../ui/uikit/ActionButton.vue'
import InputTextField from '../ui/uikit/InputTextField.vue'
import Dialog from '../ui/uikit/Dialog.vue'
import ServiceItem from '../ui/uikit/ServiceItem.vue'

const services = ref([
  {id: 1, name: "Уборка", price: "10$"},
  {id: 2, name: "Уборка", price: "10$"}
]) // TODO DB request

const isDialogVisible: Ref<boolean> = ref(false)

function openDialog(): void {
  isDialogVisible.value = true
}

function closeDialog(): void {
  isDialogVisible.value = false
}
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
    <!-- Dialog to add worker -->
    <Dialog
      title="Добавление услуги"
      :visible="isDialogVisible"
      dialogMaxWidth="30%"
      @update:visible="closeDialog"
    >
      <template #body>
        <InputTextField
          placeholder="Введите название"
          type="text"
          label="Название услуги"
        ></InputTextField>
        <InputTextField
          placeholder="Введите стоимость"
          type="text"
          label="Стоимость услуги"
        ></InputTextField>
        <InputTextField
          placeholder="Введите количество"
          type="text"
          label="Количество исполнителей"
        ></InputTextField>
        <InputTextField
          type="text"
          label="Описание"
        ></InputTextField>
        <InputTextField
          placeholder="eancode, name, count;"
          type="text"
          label="Рассходники"
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
          @click="closeDialog"
        ></ActionButton>
      </template>
    </Dialog>
    <!-------------------------->
    <HeaderList
      title="Услуги" 
      :items="services"
      height="100%"
      width="95%"
    >
     <template #items="{ item }">
        <ServiceItem :service="item"/>
      </template>
    </HeaderList>
  </MainContainer>
</template>

<style scoped>
</style>
