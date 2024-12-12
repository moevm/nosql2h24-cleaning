<script setup lang="ts">
import MainContainer from '../../ui/uikit/containers/MainContainer.vue'
import PanelContainer from '../../ui/uikit/containers/PanelContainer.vue'
import HeaderList from '../../ui/uikit/containers/HeaderList.vue'
import ActionButton from '../../ui/uikit/ActionButton.vue'
import AdminOrderItem from '../../ui/uikit/items/AdminOrderItem.vue'
import { VDateInput } from 'vuetify/labs/components'
import { computed, onMounted, ref } from 'vue'
import { getAllOrders, getAllServices } from '../../api/request'
import Order from '../../api/models/order'
import Service from '../../api/models/service'
import Dialog from '../../ui/uikit/Dialog.vue'
import InputTextField from '../../ui/uikit/inputs/InputTextField.vue'

const orders = ref<Order[]>([])
const services = ref<Service[]>([])
const isDialogVisible = ref<boolean>(false)

function openDialog(): void {
  isDialogVisible.value = true
}

function closeDialog(): void {
  isDialogVisible.value = false
}

const serviceNames = computed(() => services.value.map(service => service.name));

onMounted(() => {
  getAllOrders()
  .then((response) => {
    orders.value = response
  })

  getAllServices()
  .then((response) => {
    services.value = response
  })
})
</script>

<template>
  <MainContainer>
    <PanelContainer
      height="10%"
      width="95%"
    >
      <div class="filters-btn">
        <ActionButton
          text="Все"
          type="all"
          color="#394cc2"
        ></ActionButton>
        <ActionButton
          text="Фильтры"
          type="filters"
          color="#394cc2"
          @click="openDialog"
        ></ActionButton>
      </div>
      <v-form
        class="search-form"
        validate-on="submit lazy"
        @submit.prevent="">
        <div class="container-date">
          <v-date-input
            class="input-date"
            label="С"
            variant="solo"
            rounded="xl"
          ></v-date-input>
          <v-date-input
            class="input-date"
            label="По"
            variant="solo"
            rounded="xl"
          ></v-date-input>
        </div>
        <ActionButton
          text="Поиск"
          type="submit"
          variant="flat"
          color="#394cc2"
        ></ActionButton>
      </v-form>
    </PanelContainer>
    <HeaderList
      title="Заказы" 
      :items="orders"
      height="100%"
      width="95%"
    >
      <template #items="{ item }">
        <AdminOrderItem :order="item" />
      </template>
    </HeaderList>
    <Dialog
      title="Филтрация заказов"
      :visible="isDialogVisible"
      dialogMaxWidth="40%"
      @update:visible="closeDialog"
    >
      <template #body>
        <div class="input-data">
          <InputTextField
            placeholder="Иван"
            type="text"
            label="Имя заказчика"
          ></InputTextField>
          <InputTextField
            placeholder="Иванов"
            type="text"
            label="Фамилия заказчика"
          ></InputTextField>
          <InputTextField
            placeholder="Иванович"
            type="text"
            label="Отчество заказчика"
          ></InputTextField>
        </div>
        <div class="input-data">
          <InputTextField
            placeholder="Иван"
            type="text"
            label="Имя исполнителя"
          ></InputTextField>
          <InputTextField
            placeholder="Иванов"
            type="text"
            label="Фамилия исполнителя"
          ></InputTextField>
          <InputTextField
            placeholder="Иванович"
            type="text"
            label="Отчество исполнителя"
          ></InputTextField>
        </div>
        <div class="input-data">
          <InputTextField
            placeholder="Название города"
            type="text"
            label="Город"
          ></InputTextField>
          <InputTextField
            placeholder="Название улицы"
            type="text"
            label="Улица"
          ></InputTextField>
          <InputTextField
            placeholder="Номер дома"
            type="text"
            label="Дом"
          ></InputTextField>
        </div>
        <div class="input-data">
          <InputTextField
            placeholder="Номер квартиры"
            type="text"
            label="Квартира"
          ></InputTextField>
          <InputTextField
            placeholder="Номер подъезда"
            type="text"
            label="Подъезд"
          ></InputTextField>
          <InputTextField
            placeholder="Номер этажа"
            type="text"
            label="Этаж"
          ></InputTextField>
        </div>
        <v-select
          class="selector"
          label="Статус"
          :items="['CREATED', 'IN_PROGRESS', 'FINISHED']"
          variant="solo"
          rounded="xl"
          multiple
          persistent-hint
        ></v-select>
        <v-select
          class="selector"
          label="Услуги"
          :items="serviceNames"
          variant="solo"
          rounded="xl"
          multiple
          persistent-hint
        ></v-select>
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
          text="Поиск"
          type="search"
          variant="flat"
          color="#394cc2"
          @click="closeDialog"
        ></ActionButton>
      </template>
    </Dialog>
  </MainContainer>
</template>

<style scoped>
.input-date{
  font-weight: bold;
}
.container-date {
  width: 40%;
  display: flex;
  flex-direction: row;
  gap: 10px;
}
.filters-btn {
  width: 20%;
  display: flex;
  flex-direction: row;
  gap: 10px;
}
.search-form {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: flex-end;
  width: 80%;
  height: 100%;
  gap: 10px;
}
.input-data {
  display: flex;
  flex-direction: row;
  gap: 10px;
  width: 100%;
}
.selector {
  font-weight: bold;
}
</style>