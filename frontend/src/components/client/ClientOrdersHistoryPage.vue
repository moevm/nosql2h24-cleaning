<script setup lang="ts">
import { ref, onMounted } from 'vue'
import PanelContainer from '../../ui/uikit/containers/PanelContainer.vue'
import HeaderList from '../../ui/uikit/containers/HeaderList.vue'
import OrderItem from '../../ui/uikit/items/OrderItem.vue'
import ActionButton from '../../ui/uikit/ActionButton.vue'
import MainContainer from '../../ui/uikit/containers/MainContainer.vue'
import { getAllOrders } from '../../api/request'
import Order from '../../api/models/order'
import Address from '../../api/models/address'

const orders = ref<Order[]>([])

const allAdresses = ref<Address[]>([])

const filterOrder = ref<Order>({
  address: Address(),
  price: number,
  services: [],
  date_time: Date,
  number_of_rooms: number,
  number_of_bathrooms: number,
  area: number,
  pollution: number
});

function fetchAllOrders() {
  getAllOrders('me')
    .then((response) => {
      orders.value = response
      // todo add to allAdresses all uniq adresses. (uniq only by fields city, street, building, entrance, floor, door_number)
    })
    .catch((error) => {
      console.error('Fetch all orders error:', error);
    });
}

const isOpen = ref(false);

function toggleDropdown() {
  isOpen.value = !isOpen.value;
}

onMounted(() => {
  fetchAllOrders()
})

</script>

<template>
  <MainContainer>
    <PanelContainer>
      <ActionButton
        text="Все"
        color="#394cc2"
      ></ActionButton>
      <v-form
        class="search-form"
        validate-on="submit lazy"
        @submit.prevent="handleSearch">
        <div class="filter-dropdown">
          <div class="filter-header" @click="toggleDropdown">
            <h3>Фильтры</h3>
            <ActionButton
              text="Очистить"
              type="clear"
              color="#394cc2"
              @click.stop="clearFilters"
            ></ActionButton>
          </div>
          <transition name="fade">
            <div v-if="isOpen" class="filter-content">
              <InputTextField
                v-model="filterWorker.surname"
                label="Фамилия"
                type="last-name-search"
                placeholder="Введите фамилию"
              />
              <InputTextField
                v-model="filterWorker.name"
                label="Имя"
                type="first-name-search"
                placeholder="Введите имя"
              />
              <InputTextField
                v-model="filterWorker.patronymic"
                label="Отчество"
                type="patronymic-search"
                placeholder="Введите отчество"
              />
              <InputTextField
                v-model="filterWorker.email"
                label="Email"
                type="email"
                placeholder="Введите email"
              />
              <InputTextField
                v-model="filterWorker.phone_number"
                label="Номер телефона"
                type="tel"
                placeholder="Введите номер телефона"
              />
              <div class="order-filter">
                <p>Стоимость: {{ priceRange[0] }} - {{ priceRange[1] }}</p>
                <v-range-slider
                  v-model="priceRange"
                  hide-details
                  :min="0"
                  :max="maxPrice"
                  :step="1"
                  class="align-center"
                ></v-range-slider>
              </div>
              <div class="date-filter">
                <label>Зарегистрирован от: </label>
                <input type="date" v-model="startDate" placeholder="C" />
              </div>
              <div class="date-filter">
                <label>Зарегистрирован до: </label>
                <input type="date" v-model="endDate" placeholder="По" />
              </div>
            </div>
          </transition>
        </div>
        <ActionButton
          text="Поиск"
          type="submit"
          variant="flat"
          color="#394cc2"
        ></ActionButton>
      </v-form>
      <ActionButton
        text="Поиск"
        color="#394cc2"
      ></ActionButton>
    </PanelContainer>
    <HeaderList 
      title="История заказов" 
      :items="orders"
      height="92%"
      width="95%"
      >
      <template #items="{ item }">
        <OrderItem :order="item" />
      </template>
    </HeaderList>
  </MainContainer>
</template>

<style scoped>
</style>
