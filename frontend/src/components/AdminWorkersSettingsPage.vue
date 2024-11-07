<script setup lang="ts">
import { Ref, ref } from 'vue'
import HeaderList from '../ui/uikit/HeaderList.vue'
import PanelContainer from '../ui/uikit/PanelContainer.vue'
import ActionButton from '../ui/uikit/ActionButton.vue'
import InputTextField from '../ui/uikit/InputTextField.vue'
import MainContainer from '../ui/uikit/MainContainer.vue'
import Dialog from '../ui/uikit/Dialog.vue'
import WorkerItem from '../ui/uikit/WorkerItem.vue'

const workers = ref([
  {id: 1, firstName: "Иван", lastName: "Иванов", email: "example@mail.com"},
  {id: 2, firstName: "Иван", lastName: "Иванов", email: "example@mail.com"}
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
      title="Исполнители" 
      :items="workers"
      height="100%"
      width="95%"
    >
      <template #items="{ item }">
        <WorkerItem :worker="item" />
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
