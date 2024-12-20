<script setup lang="ts">
import { Ref, ref, defineProps } from 'vue'
import ActionButton from '../ActionButton.vue'
import InputTextField from '../inputs/InputTextField.vue'
import Consumable from '../../../api/models/consumable'
import Dialog from '../Dialog.vue'

const props = defineProps<{
  isAdmin: boolean;
  service: {
    id: string;
    name: string;
    price: number;
    workers_quantity: number;
    description: string;
    consumables: Consumable[];
  }
}>()

const emit = defineEmits(['add-service-price', 'remove-service-price'])

const isDialogVisible: Ref<boolean> = ref(false)
const isServiceAdded: Ref<boolean> = ref(false)

function openDialog(): void {
  isDialogVisible.value = true
}

function closeDialog(): void {
  isDialogVisible.value = false
}

function toggleServicePrice(): void {
  if (isServiceAdded.value) {
    emit('remove-service-price', props.service);
  } else {
    emit('add-service-price', props.service);
  }
  isServiceAdded.value = !isServiceAdded.value;
}
</script>

<template>
  <div class="service-item">
    <p>{{ props.service.name }}</p>
    <p>{{ props.service.price }}₽</p>
    <div class="service-edit">
      <ActionButton
        v-if="props.isAdmin"
        text="Редактировать"
        type="edit"
        color="#394cc2"
        variant="flat"
        @click="openDialog"
      ></ActionButton>
      <ActionButton
        v-else
        :text="isServiceAdded ? 'Убрать' : 'Добавить'"
        type="add-remove"
        :color="isServiceAdded ? '#f65858' : '#394cc2'"
        variant="flat"
        @click="toggleServicePrice"
      ></ActionButton>
    </div>
    <!-- Dialog to edit service -->
    <Dialog
      v-if="props.isAdmin"
      title="Редактирование услуги"
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
.service-item {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  width: 49%;
  border: 3px solid #3846c0;
  border-radius: 15px;
  padding: 10px;
  text-align: left;
}
.service-edit {
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
