<script setup lang="ts">
import { Ref, ref, defineProps } from 'vue'
import ActionButton from '../ActionButton.vue'
import InputTextField from '../inputs/InputTextField.vue'
import InputNumber from '../inputs/InputNumber.vue'
import Dialog from '../Dialog.vue'
import Service from '../../../api/models/service'
import { deleteService, updateService } from '../../../api/request'

const emit = defineEmits(['update-service-item']);

const props = defineProps<{
  isAdmin: boolean;
  service: {
    id: string;
    name: string;
    price: number;
    workers_quantity: number;
    description: string;
    consumables: [];
    created_at: ''
  }
}>()

let editService = ref<Service>({
  id: props.service.id,
  name: props.service.name,
  price: props.service.price,
  workers_quantity: props.service.workers_quantity,
  description: props.service.description,
  consumables: props.service.consumables,
  created_at: props.service.created_at
})

const isDialogVisible: Ref<boolean> = ref(false)
const isServiceAdded: Ref<boolean> = ref(false)

function openDialog(): void {
  isDialogVisible.value = true
}

function closeDialog(): void {
  isDialogVisible.value = false
  editService = ref<Service>({
    id: props.service.id,
    name: props.service.name,
    price: props.service.price,
    workers_quantity: props.service.workers_quantity,
    description: props.service.description,
    consumables: props.service.consumables,
    created_at: props.service.created_at
  })
}

function toggleServicePrice(): void {
  isServiceAdded.value = !isServiceAdded.value;
}

async function fetchUpdateService() {
  if (editService.value.name === '') {
    alert("Название не может быть пустым");
  } else if (editService.value.price < 1) {
    alert("Услуга не может быть бесплатной");
  } else {
    try {
      await updateService(editService.value);
      isDialogVisible.value = false
      emit('update-service-item');
    } catch (error) {
      console.error('Ошибка при обновлении услуги:', error);
    }
  }
}

function fetchDeleteService() {
  deleteService(editService.value.id)
  .then((_) => {
    closeDialog()
    emit('update-service-item')
  })
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
          v-model="editService.name"
          placeholder="Введите название"
          type="text"
          label="Название услуги"
        ></InputTextField>
        <InputNumber
          v-model="editService.price"
          label="Стоимость услуги"
          :min="0"
          :max="10000"
          :step="10"
        ></InputNumber>
        <InputNumber
          v-model="editService.workers_quantity"
          label="Количество исполнителей"
          :min="1"
          :max="10"
          :step="1"
        ></InputNumber>
        <InputTextField
          v-model="editService.description"
          type="text"
          label="Описание"
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
          @click="fetchDeleteService"
        ></ActionButton>
        <ActionButton
          text="Подтвердить"
          type="submit"
          variant="flat"
          color="#394cc2"
          @click="fetchUpdateService"
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
  overflow: hidden;
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
  width: 100%;
  box-sizing: border-box;
  font-size: 32px;
  font-weight: bold;
}
</style>
