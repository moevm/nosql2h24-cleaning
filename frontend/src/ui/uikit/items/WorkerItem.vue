<script setup lang="ts">
import { Ref, ref, defineProps } from 'vue'
import { User, UserRegisterData } from '../../../api/models/user'
import { updateUser, deleteUser } from '../../../api/request'
import ActionButton from '../ActionButton.vue'
import InputTextField from '../inputs//InputTextField.vue'
import Dialog from '../Dialog.vue'

const emit = defineEmits(['update-worker']);

const isCorrectUpdateForm = ref<boolean>(false);
const rules = {
  required: (value: any) => {
    if (
      editWorker.value.name &&
      editWorker.value.surname &&
      editWorker.value.phone_number
    ) {
      isCorrectUpdateForm.value = true;
    } else {
      isCorrectUpdateForm.value = false;
    }
    return !!value || 'Поле является обязательным'
  },
  email: (value: any) => {
    const emailRegex = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,6}$/;
    if (!emailRegex.test(value)) {
      isCorrectUpdateForm.value = false;
    }
    return emailRegex.test(value) || 'Введите корректный email'
  },
}

const props = defineProps<{
  worker: {
    id: string;
    name: string;
    surname: string;
    patronymic: string;
    email: string;
    phone_number: string;
    orders_count: number;
    created_at: Date;
  }
}>()

const editWorker = ref<UserRegisterData>({
  name: props.worker.name,
  surname: props.worker.surname,
  patronymic: props.worker.patronymic,
  email: props.worker.email,
  phone_number: props.worker.phone_number,
  password: '',
});

async function handleUpdateWorker() {
  const workerData: Partial<User> = {
    ...editWorker.value,
    id: props.worker.id,
    user_type: 'WORKER',
    orders_count: props.worker.orders_count,
    created_at: props.worker.created_at,
  };
    
  await updateUser(props.worker.id, workerData as User).then( _ => {
    console.log('User updated successfully');
    closeDialog();
    emit('update-worker')
  }).catch((error) => {
    console.error('Error updating user:', error);
  })
}

async function handleDeleteWorker() {
  await deleteUser(props.worker.id).then((_) => {
    console.log('User deleted successfully');
    closeDialog();
    emit('update-worker');
  }).catch((error) => {
    console.error('Error deleting user:', error);
  })
}

function resetEditableWorker() {
  editWorker.value = {
    name: props.worker.name,
    surname: props.worker.surname,
    patronymic: props.worker.patronymic,
    email: props.worker.email,
    phone_number: props.worker.phone_number,
    password: ''
  };
}

const isDialogVisible: Ref<boolean> = ref(false)

function openDialog(): void {
  resetEditableWorker()
  isDialogVisible.value = true
}

function closeDialog(): void {
  isDialogVisible.value = false
  resetEditableWorker();
}

function formatDate(date: Date) {
  const months = [
    'января', 'февраля', 'марта', 'апреля', 'мая', 'июня',
    'июля', 'августа', 'сентября', 'октября', 'ноября', 'декабря'
  ];

  const day = date.getDate();
  const monthIndex = date.getMonth();
  const year = date.getFullYear();

  return `${day} ${months[monthIndex]} ${year}`;
}
</script>

<template>
  <div class="worker-item">
    <p>{{ props.worker.surname }} {{ props.worker.name }} {{ props.worker.patronymic }}</p>
    <p>E-mail: {{ props.worker.email }}</p>
    <p>Номер телефона: {{ props.worker.phone_number }}</p>
    <p>Выполнено заказов: {{ isNaN(props.worker.orders_count) ? 0 : props.worker.orders_count }}</p>
    <p>На сервисе с {{ formatDate(new Date(props.worker.created_at)) }} года</p>
    <div class="worker-edit">
      <ActionButton
        text="Редактировать"
        type="edit"
        color="#394cc2"
        variant="flat"
        @click="openDialog"
      ></ActionButton>
    </div>
    <!-- Dialog to edit worker -->
    <Dialog
      title="Редактирование исполнителя"
      :visible="isDialogVisible"
      dialogMaxWidth="30%"
      @update:visible="closeDialog"
    >
      <template #body>
        <InputTextField
          v-model="editWorker.name"
          placeholder="Введите имя"
          type="text"
          label="Имя"
          :rules="[rules.required]"
        ></InputTextField>
        <InputTextField
          v-model="editWorker.surname"
          placeholder="Введите фамилию"
          type="text"
          label="Фамилия"
          :rules="[rules.required]"
        ></InputTextField>
        <InputTextField
          v-model="editWorker.patronymic"
          placeholder="Введите отчество"
          type="text"
          label="Отчество"
        ></InputTextField>
        <InputTextField
          v-model="editWorker.phone_number"
          placeholder="Введите телефон"
          type="phonenumber"
          label="Номер телефона"
          :rules="[rules.required]"
        ></InputTextField>
        <InputTextField
          v-model="editWorker.email"
          placeholder="Введите почту"
          type="email"
          label="Почта"
          :rules="[rules.required, rules.email]"
        ></InputTextField>
        <InputTextField
          v-model="editWorker.password"
          placeholder="Новый пароль"
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
          text="Удалить"
          type="delete"
          variant="flat"
          color="#f65858"
          @click="handleDeleteWorker"
        ></ActionButton>
        <ActionButton
          text="Подтвердить"
          type="submit"
          variant="flat"
          color="#394cc2"
          @click="handleUpdateWorker"
          :disabled="!isCorrectUpdateForm"
        ></ActionButton>
      </template>
    </Dialog>
    <!--------------------------->
  </div>
</template>

<style scoped>
.worker-item {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  width: 49%;
  border: 3px solid #3846c0;
  border-radius: 15px;
  padding: 10px;
  text-align: left;
}
.worker-edit {
  display: flex;
  flex-direction: row;
  justify-content: flex-end;
  width: 100%;
}
p {
  font-size: 32px;
  font-weight: bold;
}
</style>
