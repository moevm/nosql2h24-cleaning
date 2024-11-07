<script setup lang="ts">
import { Ref, ref, defineProps } from 'vue'
import ActionButton from './ActionButton.vue'
import InputTextField from './InputTextField.vue'
import Dialog from './Dialog.vue'

const props = defineProps<{
  worker: {
    id: string;
    firstName: string;
    lastName: string;
    email: string;
  }
}>()

const isDialogVisible: Ref<boolean> = ref(false)

function openDialog(): void {
  isDialogVisible.value = true
}

function closeDialog(): void {
  isDialogVisible.value = false
}
</script>

<template>
  <div class="worker-item">
    <p>{{ props.worker.firstName }} {{ props.worker.lastName }}</p>
    <p>{{ props.worker.email }}</p>
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
          placeholder="Введите имя"
          type="text"
          label="Имя"
        ></InputTextField>
        <InputTextField
          placeholder="Введите фамилию"
          type="text"
          label="Фамилия"
        ></InputTextField>
        <InputTextField
          placeholder="Введите отчество"
          type="text"
          label="Отчество"
        ></InputTextField>
        <InputTextField
          placeholder="Введите телефон"
          type="phonenumber"
          label="Номер телефона"
        ></InputTextField>
        <InputTextField
          placeholder="Введите почту"
          type="email"
          label="Почта"
        ></InputTextField>
        <InputTextField
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
          text="Удалить"
          type="delete"
          variant="flat"
          color="#f65858"
          @click="closeDialog"
        ></ActionButton>
        <ActionButton
          text="Подтвердить"
          type="submit"
          variant="flat"
          color="#394cc2"
          @click="closeDialog"
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
