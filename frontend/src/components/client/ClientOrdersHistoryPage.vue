<script setup lang="ts">
import { ref, onMounted, computed, watch } from 'vue'
import PanelContainer from '../../ui/uikit/containers/PanelContainer.vue'
import HeaderList from '../../ui/uikit/containers/HeaderList.vue'
import OrderItem from '../../ui/uikit/items/OrderItem.vue'
import ActionButton from '../../ui/uikit/ActionButton.vue'
import MainContainer from '../../ui/uikit/containers/MainContainer.vue'
import { getAllServices, getAllOrders, getFiltredOrders } from '../../api/request'
import Order from '../../api/models/order'
import Address from '../../api/models/address'
import Service from '../../api/models/service'
import OrderFilter from '../../api/models/orderFilter'

type StatusKey = "CREATED" | "IN_PROGRESS" | "FINISHED";
const statusTranslations: Record<StatusKey, string> = {
  "CREATED": "Создан",
  "IN_PROGRESS": "В работе",
  "FINISHED": "Выполнен"
};
function translateStatus(status: StatusKey): string | null {
  return statusTranslations[status] || null;
}

const orders = ref<Order[]>([])
const allAdresses = ref<Address[]>([])
const selectedAddresses = ref(allAdresses.value.map(() => false));
watch(allAdresses, (newValue) => {
  selectedAddresses.value = newValue.map(() => false);
});

const allServices = ref<Service[]>([])
const selectedServices = ref(allServices.value.map(() => false));
watch(allServices, (newValue) => {
  selectedServices.value = newValue.map(() => false);
});

const allStatuses = ref<string[]>(["CREATED", "IN_PROGRESS", "FINISHED"])
const selectedStatuses = ref(allStatuses.value.map(() => false));

const maxPriceOrder = ref<number | null>(null);
const maxPrice = computed(() => maxPriceOrder.value ?? 0);
const orderPriceRange = ref<[number, number]>([0, maxPrice.value]);
watch(maxPrice, (newMax) => {
  orderPriceRange.value = [0, newMax];
});

const maxPollution = 10;
const orderPollutionRange = ref<[number, number]>([0, maxPollution]);

const maxRoomsOrder = ref<number | null>(null);
const maxRooms = computed(() => maxRoomsOrder.value ?? 0);
const orderRoomsRange = ref<[number, number]>([0, maxRooms.value]);
watch(maxRooms, (newMax) => {
  orderRoomsRange.value = [0, newMax];
});

const maxBathroomsOrder = ref<number | null>(null);
const maxBathrooms = computed(() => maxBathroomsOrder.value ?? 0);
const orderBathroomsRange = ref<[number, number]>([0, maxBathrooms.value]);
watch(maxBathrooms, (newMax) => {
  orderBathroomsRange.value = [0, newMax];
});

const maxAreaOrder = ref<number | null>(null);
const maxArea = computed(() => maxAreaOrder.value ?? 0);
const orderAreaRange = ref<[number, number]>([0, maxArea.value]);
watch(maxArea, (newMax) => {
  orderAreaRange.value = [0, newMax];
});

const maxWorkersOrder = ref<number | null>(null);
const maxWorkers = computed(() => maxWorkersOrder.value ?? 0);
const orderWorkersRange = ref<[number, number]>([1, maxWorkers.value]);
watch(maxWorkers, (newMax) => {
  orderWorkersRange.value = [1, newMax];
});


const initOrderFilter = ref<OrderFilter>({
  user_id: 'me',
  address_city: '',
  address_street: '',
  address_building: '',
  address_entrance: '',
  address_floor: '',
  address_door_number: '',
  price_min: 0,
  price_max: 0,
  pollution_min: 0,
  pollution_max: 0,
  required_workers_min: 0,
  required_workers_max: 0,
  area_min: 0,
  area_max: 0,
  number_of_rooms_min: 0,
  number_of_rooms_max: 0,
  number_of_bathrooms_min: 0,
  number_of_bathrooms_max: 0,

  worker_name: '',
  worker_surname: '',
  worker_patronymic: '',
  statuses: [],
  services: [],
  date_time_begin: '',
  date_time_end: ''
});


const orderFilter = ref<OrderFilter>(initOrderFilter.value)

function fetchAllOrders() {
  getAllOrders('me')
    .then((response) => {
      orders.value = response

      if (maxPriceOrder.value === null) {
        maxPriceOrder.value = Math.max(...orders.value.map(order => order.price));
      }
      if (maxWorkersOrder.value === null) {
        maxWorkersOrder.value = Math.max(...orders.value.map(order => order.required_workers));
      }
      if (maxRoomsOrder.value === null) {
        maxRoomsOrder.value = Math.max(...orders.value.map(order => order.number_of_rooms));
      }
      if (maxAreaOrder.value === null) {
        maxAreaOrder.value = Math.max(...orders.value.map(order => order.area));
      }
      if (maxBathroomsOrder.value === null) {
        maxBathroomsOrder.value = Math.max(...orders.value.map(order => order.number_of_bathrooms));
      }

      const addresses = orders.value.map(order => order.address);
      const uniqueAddresses = Array.from(new Set(addresses.map(address => JSON.stringify(address))))
        .map(addressStr => JSON.parse(addressStr));
      allAdresses.value = uniqueAddresses;

      getAllServices()
        .then(allServicesResponse => {
          allServices.value = allServicesResponse;
          selectedServices.value = allServicesResponse.map(() => false);
        })
        .catch(error => {
          console.error('Ошибка при загрузке сервисов:', error);
        });

    })
    .catch((error) => {
      console.error('Fetch all orders error:', error);
    });
}

function handleSearch() {
  orderFilter.value.price_min = orderPriceRange.value[0];
  orderFilter.value.price_max = orderPriceRange.value[1];
  orderFilter.value.pollution_min = orderPollutionRange.value[0];
  orderFilter.value.pollution_max = orderPollutionRange.value[1];
  orderFilter.value.required_workers_min = orderWorkersRange.value[0];
  orderFilter.value.required_workers_max = orderWorkersRange.value[1];
  orderFilter.value.area_min = orderAreaRange.value[0];
  orderFilter.value.area_max = orderAreaRange.value[1];
  orderFilter.value.number_of_rooms_min = orderRoomsRange.value[0];
  orderFilter.value.number_of_rooms_max = orderRoomsRange.value[1];
  orderFilter.value.number_of_bathrooms_min = orderBathroomsRange.value[0];
  orderFilter.value.number_of_bathrooms_max = orderBathroomsRange.value[1];


  const updatedOrderFilter = {
    ...orderFilter.value
  };

  updatedOrderFilter.statuses = allStatuses.value.filter((_, index) => selectedStatuses.value[index]);

  if (orderFilter.value.date_time_begin !== '') {
    updatedOrderFilter.date_time_begin = formatToRFC3339(orderFilter.value.date_time_begin, false)
  }
  if (orderFilter.value.date_time_end !== '') {
    updatedOrderFilter.date_time_end = formatToRFC3339(orderFilter.value.date_time_end, true)
  }

  updatedOrderFilter.services = allServices.value
    .filter((_, index) => selectedServices.value[index])
    .map(service => service.id);
    
  const selectedIndex = selectedAddresses.value.indexOf(true);
  if (selectedIndex !== -1) {
    const selectedAddress = allAdresses.value[selectedIndex];
    updatedOrderFilter.address_city = selectedAddress.city;
    updatedOrderFilter.address_street = selectedAddress.street;
    updatedOrderFilter.address_building = selectedAddress.building;
    updatedOrderFilter.address_entrance = selectedAddress.entrance;
    updatedOrderFilter.address_floor = selectedAddress.floor;
    updatedOrderFilter.address_door_number = selectedAddress.door_number;
  }

  getFiltredOrders(updatedOrderFilter)
  .then((response) => {
    orders.value = response
  })
  .catch((error) => {
    console.log("Error filter orders:", error);
  })
}

function clearFilters() {
  fetchAllOrders();
  orderFilter.value = initOrderFilter.value
}

function formatToRFC3339(dateString: string, isEnd: boolean) {
  let date = new Date(dateString);
  if (isEnd) {
    date.setDate(date.getDate() + 1);
    date = new Date(date.getTime() - 1);
  }
  console.log(date.toISOString());
  return date.toISOString(); 
}

const isOpen = ref(false);
function toggleDropdown() {
  isOpen.value = !isOpen.value;
}

function toggleAddressesSelection(index: number) {
  if (selectedAddresses.value[index]) {
    selectedAddresses.value = selectedAddresses.value.map(() => false);
  } else {
    selectedAddresses.value = selectedAddresses.value.map((_, i) => i === index);
  }
}

function toggleServiceSelection(index: number) {
  selectedServices.value[index] = !selectedServices.value[index];
}

function toggleStatusSelection(index: number) {
  selectedStatuses.value[index] = !selectedStatuses.value[index];
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
        validate-on="submit lazy">
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
              <div class="scroll-container">
                <div class="selection-container">
                  <p class="selection-label">Адреса:</p>
                  <div class="selection-list">
                    <div
                      class="selection-item"
                      v-for="(address, index) in allAdresses"
                      :key="address.id"
                      @click="toggleAddressesSelection(index)"
                      :class="{ selected: selectedAddresses[index] }"
                    >
                      <p><strong>City:</strong> {{ address.city }}</p>
                      <p><strong>Street:</strong> {{ address.street }}</p>
                      <p><strong>Building:</strong> {{ address.building }}</p>
                      <p><strong>Entrance:</strong> {{ address.entrance }}</p>
                      <p><strong>Floor:</strong> {{ address.floor }}</p>
                      <p><strong>Door Number:</strong> {{ address.door_number }}</p>
                    </div>
                  </div>
                </div>
                <div class="selection-container">
                  <p class="selection-label">Услуги:</p>
                  <div class="selection-list">
                    <div
                      class="selection-item"
                      v-for="(service, index) in allServices"
                      :key="service.id"
                      @click="toggleServiceSelection(index)"
                      :class="{ selected: selectedServices[index] }"
                    >
                      <p><strong>Услуга:</strong> {{ service.name }}</p>
                      <p><strong>Стоимость:</strong> {{ service.price }}</p>
                    </div>
                  </div>
                </div>
                <div class="selection-container">
                  <p class="selection-label">Статус:</p>
                  <div class="selection-list">
                    <div
                      class="selection-item"
                      v-for="(status, index) in allStatuses"
                      :key="status"
                      @click="toggleStatusSelection(index)"
                      :class="{ selected: selectedStatuses[index] }"
                    >
                      <p>{{ translateStatus(status as StatusKey) }}</p>
                    </div>
                  </div>
                </div>
                <div class="order-filter" v-if="0 !== maxPrice">
                  <p>Стоимость: от {{ orderPriceRange[0] }} - до {{ orderPriceRange[1] }}</p>
                  <v-range-slider
                    v-model="orderPriceRange"
                    hide-details
                    :min="0"
                    :max="maxPrice"
                    :step="1"
                    class="align-center"
                  ></v-range-slider>
                </div>
                <div class="order-filter">
                  <p>Загрязненность: от {{ orderPollutionRange[0] }} - до {{ orderPollutionRange[1] }}</p>
                  <v-range-slider
                    v-model="orderPollutionRange"
                    hide-details
                    :min="0"
                    :max="10"
                    :step="1"
                    class="align-center"
                  ></v-range-slider>
                </div>
                <div class="order-filter" v-if="1 !== maxWorkers">
                  <p>Количество исплнителей: от {{ orderWorkersRange[0] }} - до {{ orderWorkersRange[1] }}</p>
                  <v-range-slider
                    v-model="orderWorkersRange"
                    hide-details
                    :min="1"
                    :max="maxWorkers"
                    :step="1"
                    class="align-center"
                  ></v-range-slider>
                </div>
                <div class="order-filter" v-if="0 !== maxRooms">
                  <p>Количество комнат: от {{ orderRoomsRange[0] }} - до {{ orderRoomsRange[1] }}</p>
                  <v-range-slider
                    v-model="orderRoomsRange"
                    hide-details
                    :min="0"
                    :max="maxRooms"
                    :step="1"
                    class="align-center"
                  ></v-range-slider>
                </div>
                <div class="order-filter" v-if="0 !== maxBathrooms">
                  <p>Количество санузлов: от {{ orderBathroomsRange[0] }} - до {{ orderBathroomsRange[1] }}</p>
                  <v-range-slider
                    v-model="orderBathroomsRange"
                    hide-details
                    :min="0"
                    :max="maxBathrooms"
                    :step="1"
                    class="align-center"
                  ></v-range-slider>
                </div>
                <div class="order-filter" v-if="0 !== maxArea">
                  <p>Площадь: от {{ orderAreaRange[0] }} - до {{ orderAreaRange[1] }}</p>
                  <v-range-slider
                    v-model="orderAreaRange"
                    hide-details
                    :min="0"
                    :max="maxArea"
                    :step="1"
                    class="align-center"
                  ></v-range-slider>
                </div>
                <div class="date-filter">
                  <label>Создан от: </label>
                  <input type="date" v-model="orderFilter.date_time_begin" placeholder="C" />
                </div>
                <div class="date-filter">
                  <label>Создан до: </label>
                  <input type="date" v-model="orderFilter.date_time_end" placeholder="По" />
                </div>
              </div>
            </div>
          </transition>
        </div>
      </v-form>
      <ActionButton
        text="Поиск"
        color="#394cc2"
        @click="handleSearch"
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
.filter-dropdown {
  position: relative;
  width: 100%;
  max-width: 500px;
}

.filter-header {
  background: #f7f7f7;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 25px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.filter-content {
  position: absolute;
  top: 120%;
  left: 0;
  right: 0;
  z-index: 2;
  border: 1px solid #ddd;
  background: #fff;
  border-radius: 30px;
  box-shadow: 0px 4px 6px rgba(0, 0, 0, 0.1);
  overflow: hidden;
}

.scroll-container {
  max-height: 720px;
  overflow-y: auto;
  padding: 10px;
  -ms-overflow-style: none;
  scrollbar-width: none;
}

.scroll-container::-webkit-scrollbar {
  display: none;
  padding: 10px;
}

input[type="range"] {
  flex-grow: 1;
}

.fade-enter-active, .fade-leave-active {
  transition: opacity 0.5s;
}

.fade-enter, .fade-leave-to {
  opacity: 0;
}

.search-form {
  display: flex;
  flex-direction: row;
  align-items: center;
  width: 50%;
  height: 100%;
  gap: 10px;
}

.filter-header {
  background: #f7f7f7;
  padding: 10px;
  border-bottom: 1px solid #ddd;
}

.order-filter, .date-filter {
  margin-top: 10px;
}

.selection-container {
  display: flex;
  overflow-x: auto;
  white-space: nowrap;
  flex-direction: column;
  align-items: flex-start;
  margin-bottom: 8px;
}

.selection-label {
  font-weight: bold;
  margin-bottom: 4px;
  margin-top: 8px;
}

.selection-list {
  display: flex;
  gap: 20px;
}

.selection-item {
  display: inline-block;
  min-width: 200px;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 20px;
  background-color: #f9f9f9;
  cursor: pointer;
}

.selection-item.selected {
  background-color: #5269ff;
}
</style>
