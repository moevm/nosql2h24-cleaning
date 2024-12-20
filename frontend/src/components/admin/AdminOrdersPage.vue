<script setup lang="ts">
import MainContainer from '../../ui/uikit/containers/MainContainer.vue'
import PanelContainer from '../../ui/uikit/containers/PanelContainer.vue'
import HeaderList from '../../ui/uikit/containers/HeaderList.vue'
import ActionButton from '../../ui/uikit/ActionButton.vue'
import AdminOrderItem from '../../ui/uikit/items/AdminOrderItem.vue'
import { VDateInput } from 'vuetify/labs/components'
import { computed, onMounted, ref } from 'vue'
import { getAllOrders, getAllServices, filterOrder } from '../../api/request'
import Order from '../../api/models/order'
import Service from '../../api/models/service'
import Dialog from '../../ui/uikit/Dialog.vue'
import InputTextField from '../../ui/uikit/inputs/InputTextField.vue'
import { FilterOrder } from '../../api/models/filterOrder'

const orders = ref<Order[]>([])
const services = ref<Service[]>([])
const isDialogVisible = ref<boolean>(false)
const statuses = ref<string[]>([])
const servicesSelect = ref<string[]>([])
const filterData = ref<FilterOrder>({
  user_name: '',
  user_surname: '',
  user_patronymic: '',
  worker_name: '',
  worker_surname: '',
  worker_patronymic: '',
  address_city: '',
  address_street: '',
  address_building: '',
  address_entrance: '',
  address_floor: '',
  address_door_number: '',
})
const startDate = ref<Date | null>(null)
const endDate = ref<Date | null>(null)

function openDialog(): void {
  isDialogVisible.value = true
}

function closeDialog(): void {
  isDialogVisible.value = false
}

const serviceNames = computed(() => services.value.map(service => service.name));
const selectedServiceIds = computed(() => {
  return services.value
    .filter(service => servicesSelect.value.includes(service.name))
    .map(service => service.id);
});

function defaultValues(): void {
  filterData.value = {
    user_name: '',
    user_surname: '',
    user_patronymic: '',
    worker_name: '',
    worker_surname: '',
    worker_patronymic: '',
    address_city: '',
    address_street: '',
    address_building: '',
    address_entrance: '',
    address_floor: '',
    address_door_number: '',
  }
  statuses.value = []
  servicesSelect.value = []
  filtrationOrder()
}

function filtrationOrder(): void {
  closeDialog()
  filterOrder({
    ...filterData.value,
    statuses: statuses.value,
    services: selectedServiceIds.value
  })
  .then((response) => {
    orders.value = response
  })
}

function submitSearch(): void {
  filterOrder({
    ...filterData.value,
    statuses: statuses.value,
    services: selectedServiceIds.value,
    date_time_begin: startDate.value?.toISOString(),
    date_time_end: endDate.value?.toISOString()
  })
  .then((response) => {
    orders.value = response
  })
}

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
          @click="defaultValues"
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
            v-model="startDate"
            class="input-date"
            label="С"
            variant="solo"
            rounded="xl"
          ></v-date-input>
          <v-date-input
            v-model="endDate"
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
          @click="submitSearch"
        ></ActionButton>
      </v-form>
    </PanelContainer>
    <HeaderList
      title="Заказы" 
      :items="orders"
      height="92%"
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
            v-model="filterData.user_name"
            placeholder="Иван"
            type="text"
            label="Имя заказчика"
          ></InputTextField>
          <InputTextField
            v-model="filterData.user_surname"
            placeholder="Иванов"
            type="text"
            label="Фамилия заказчика"
          ></InputTextField>
          <InputTextField
            v-model="filterData.user_patronymic"
            placeholder="Иванович"
            type="text"
            label="Отчество заказчика"
          ></InputTextField>
        </div>
        <div class="input-data">
          <InputTextField
            v-model="filterData.worker_name"
            placeholder="Иван"
            type="text"
            label="Имя исполнителя"
          ></InputTextField>
          <InputTextField
            v-model="filterData.worker_surname"
            placeholder="Иванов"
            type="text"
            label="Фамилия исполнителя"
          ></InputTextField>
          <InputTextField
            v-model="filterData.worker_patronymic"
            placeholder="Иванович"
            type="text"
            label="Отчество исполнителя"
          ></InputTextField>
        </div>
        <div class="input-data">
          <InputTextField
            v-model="filterData.address_city"
            placeholder="Название города"
            type="text"
            label="Город"
          ></InputTextField>
          <InputTextField
            v-model="filterData.address_street"
            placeholder="Название улицы"
            type="text"
            label="Улица"
          ></InputTextField>
          <InputTextField
            v-model="filterData.address_building"
            placeholder="Номер дома"
            type="text"
            label="Дом"
          ></InputTextField>
        </div>
        <div class="input-data">
          <InputTextField
            v-model="filterData.address_door_number"
            placeholder="Номер квартиры"
            type="text"
            label="Квартира"
          ></InputTextField>
          <InputTextField
            v-model="filterData.address_entrance"
            placeholder="Номер подъезда"
            type="text"
            label="Подъезд"
          ></InputTextField>
          <InputTextField
            v-model="filterData.address_floor"
            placeholder="Номер этажа"
            type="text"
            label="Этаж"
          ></InputTextField>
        </div>
        <v-select
          v-model="statuses"
          class="selector"
          label="Статус"
          :items="['CREATED', 'IN_PROGRESS', 'FINISHED']"
          variant="solo"
          rounded="xl"
          multiple
          persistent-hint
        ></v-select>
        <v-select
          v-model="servicesSelect"
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
          @click="filtrationOrder"
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