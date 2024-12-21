<script setup lang="ts">
import { onMounted, computed, watch, Ref, ref } from 'vue'
import { createWorkerUser, getUsers } from '../../api/request'
import { User, UserRegisterData } from '../../api/models/user'
import HeaderList from '../../ui/uikit/containers/HeaderList.vue'
import PanelContainer from '../../ui/uikit/containers/PanelContainer.vue'
import ActionButton from '../../ui/uikit/ActionButton.vue'
import InputTextField from '../../ui/uikit/inputs/InputTextField.vue'
import MainContainer from '../../ui/uikit/containers/MainContainer.vue'
import Dialog from '../../ui/uikit/Dialog.vue'
import WorkerItem from '../../ui/uikit/items/WorkerItem.vue'
import { filterWorkers } from '../../api/request'
import { FilterWorkers } from '../../api/models/filterWorkers'


const maxOrdersWorker = ref<number | null>(null);
const maxOrders = computed(() => maxOrdersWorker.value ?? 0);
watch(maxOrders, (newMax) => {
  ordersRange.value = [0, newMax];
});

const filterWorker = ref<FilterWorkers>({
  name: '',
  surname: '',
  patronymic: '',
  email: '',
  phone_number: '',
  orders_count_min: 0,
  orders_count_max: maxOrders.value,
  created_at_begin: '',
  created_at_end: '',
});

const isOpen = ref(false);

function toggleDropdown() {
  isOpen.value = !isOpen.value;
}

const workers = ref<{
  id: string;
  email: string;
  name: string;
  surname: string;
  patronymic: string;
  phone_number: string;
  user_type: string;
  created_at: Date;
  updated_at: Date;
  orders_count: number;
}[]>([]);

const newWorker = ref<UserRegisterData>({
  name: '',
  surname: '',
  patronymic: '',
  email: '',
  phone_number: '',
  password: '',
});

const isDialogVisible: Ref<boolean> = ref(false)

function openDialog(): void {
  isDialogVisible.value = true
}

function closeDialog(): void {
  isDialogVisible.value = false
}

function formatToRFC3339(dateString: string, isEnd: boolean) {
  let date = new Date(dateString);
  if (isEnd) {
    date.setDate(date.getDate() + 1);
    date = new Date(date.getTime() - 1);
  }
  return date.toISOString(); 
}

async function fetchWorkersList() {
  try {
    const response = await getUsers('WORKER');
    workers.value = response.map(user => ({
      id: user.id,
      name: user.name,
      surname: user.surname,
      patronymic: user.patronymic,
      email: user.email,
      phone_number: user.phone_number,
      user_type: user.user_type,
      orders_count: user.orders_count,
      created_at: user.created_at,
      updated_at: user.updated_at
    }));

    if (maxOrdersWorker.value === null) {
      console.log("maxOrdersWorker null finde max orders")
      maxOrdersWorker.value = Math.max(...workers.value.map(worker => worker.orders_count));
    }
  } catch (error) {
    console.error("Failed to fetch workers list:", error);
  }
}

const ordersRange = ref<[number, number]>([0, maxOrders.value]);

async function fetchCreateUser(workerData: UserRegisterData) {
  const userPayload: Partial<User> = {
    ...workerData,
    user_type: 'WORKER',
  };
  await createWorkerUser(userPayload as User).then((response) => {
    console.log('Created User ID:', response.id);
  }).catch ((error) => {
    console.error('Error HTTP:', error);
  })
}

async function handleCreateWorkerUser() {
  closeDialog();
  try {
    await fetchCreateUser(newWorker.value);
    newWorker.value = {
      email: '',
      name: '',
      password: '',
      patronymic: '',
      phone_number: '',
      surname: ''
    };
    fetchWorkersList()
  } catch (error) {
    console.error('Error creating user:', error);
  }
}

function handleSearch(): void {
  toggleDropdown()
  const filterParams = { ...filterWorker.value };
  if (filterParams.created_at_begin) {
    filterParams.created_at_begin = formatToRFC3339(filterParams.created_at_begin, false);
  }
  if (filterParams.created_at_end) {
    filterParams.created_at_end = formatToRFC3339(filterParams.created_at_end, true);
  }
  filterParams.orders_count_min = ordersRange.value[0];
  filterParams.orders_count_max = ordersRange.value[1];
  filterWorkers(filterParams)
  .then((response) => {
    workers.value = response.map(user => ({
      id: user.id,
      name: user.name,
      surname: user.surname,
      patronymic: user.patronymic,
      email: user.email,
      phone_number: user.phone_number,
      user_type: user.user_type,
      orders_count: user.orders_count  || 0,
      created_at: user.created_at,
      updated_at: user.updated_at
    }));
    console.log('Filtered workers:', workers.value);
  }).catch ((error) => {
    console.error("Failed to fetch workers list:", error)
  })
}

function clearFilters() {
  fetchWorkersList();
  filterWorker.value = {
    name: '',
    surname: '',
    patronymic: '',
    email: '',
    phone_number: '',
    orders_count_min: 0,
    orders_count_max: maxOrders.value,
    created_at_begin: '',
    created_at_end: '',
  };
  ordersRange.value = [0, maxOrders.value];
}

onMounted(() => {
  fetchWorkersList()
})

</script>

<template>
  <MainContainer>
    <PanelContainer
      height="10%"
      width="95%"
    >
      <ActionButton
        text="Создать"
        type="create"
        color="#394cc2"
        @click="openDialog"
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
                <p>Количество выполненных заказов: {{ ordersRange[0] }} - {{ ordersRange[1] }}</p>
                <v-range-slider
                  v-model="ordersRange"
                  hide-details
                  :min="0"
                  :max="maxOrders"
                  :step="1"
                  class="align-center"
                ></v-range-slider>
              </div>
              <div class="date-filter">
                <label>Зарегистрирован от: </label>
                <input type="date" v-model="filterWorker.created_at_begin" />
              </div>
              <div class="date-filter">
                <label>Зарегистрирован до: </label>
                <input type="date" v-model="filterWorker.created_at_end" />
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
    </PanelContainer>
    <!-- Dialog to add worker -->
    <Dialog
      title="Создание исполнителя"
      :visible="isDialogVisible"
      dialogMaxWidth="30%"
      @update:visible="closeDialog"
    >
      <template #body>
        <InputTextField
          v-model="newWorker.name"
          placeholder="Введите имя"
          type="text"
          label="Имя"
        ></InputTextField>
        <InputTextField
          v-model="newWorker.surname"
          placeholder="Введите фамилию"
          type="text"
          label="Фамилия"
        ></InputTextField>
        <InputTextField
          v-model="newWorker.patronymic"
          placeholder="Введите отчество"
          type="text"
          label="Отчество"
        ></InputTextField>
        <InputTextField
          v-model="newWorker.phone_number"
          placeholder="Введите телефон"
          type="phonenumber"
          label="Номер телефона"
        ></InputTextField>
        <InputTextField
          v-model="newWorker.email"
          placeholder="Введите почту"
          type="email"
          label="Почта"
        ></InputTextField>
        <InputTextField
          v-model="newWorker.password"
          placeholder="Введите пароль"
          type="password"
          label="Пароль"
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
          @click="handleCreateWorkerUser"
        ></ActionButton>
      </template>
    </Dialog>
    <!-- -------------------- -->
    <HeaderList
      title="Исполнители" 
      :items="workers"
      height="100%"
      width="95%"
    >
      <template #items="{ item }">
        <WorkerItem :worker="item" @update-worker="fetchWorkersList" />
      </template>
    </HeaderList>
  </MainContainer>
</template>

<style scoped>
.filter-dropdown {
  position: relative;
  width: 100%;
  max-width: 400px;
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
  padding: 10px;
  border: 1px solid #ddd;
  background: #fff;
  border-radius: 30px;
  box-shadow: 0px 4px 6px rgba(0, 0, 0, 0.1);
  overflow: hidden;
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
</style>
