<script setup lang="ts">
import { onMounted, Ref, ref } from 'vue'
import { createWorkerUser, getUsers, updateUser } from '../../api/request'
import { User, UserRegisterData } from '../../api/models/user'
import HeaderList from '../../ui/uikit/containers/HeaderList.vue'
import PanelContainer from '../../ui/uikit/containers/PanelContainer.vue'
import ActionButton from '../../ui/uikit/ActionButton.vue'
import InputTextField from '../../ui/uikit/inputs/InputTextField.vue'
import MainContainer from '../../ui/uikit/containers/MainContainer.vue'
import Dialog from '../../ui/uikit/Dialog.vue'
import WorkerItem from '../../ui/uikit/items/WorkerItem.vue'

const workers = ref<{ 
  id: string;
  name: string;
  surname: string;
  patronymic: string;
  email: string;
  phone_number: string;
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

async function fetchWorkersList() {
    await getUsers('WORKER').then((response) => { 
      workers.value = response.map(user => ({
        id: user.id,
        name: user.name,
        surname: user.surname,
        patronymic: user.patronymic,
        email: user.email,
        phone_number: user.phone_number
      }));
    }).catch ((error) => {
      console.error("Failed to fetch workers list:", error)
    })
  }

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
        @submit.prevent="">
        <InputTextField
          label="Имя"
          type="first-name-search"
          placeholder="Введите имя"
        ></InputTextField>
        <InputTextField
          label="Фамилия"
          type="last-name-search"
          placeholder="Введите фамилию"
        ></InputTextField>
        <ActionButton
          text="Поиск"
          type="create"
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
.search-form {
  display: flex;
  flex-direction: row;
  align-items: center;
  width: 50%;
  height: 100%;
  gap: 10px;
}
</style>
