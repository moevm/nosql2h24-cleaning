<script setup lang="ts">
import { computed, ref, watch, onMounted } from 'vue';
import Dialog from '../../ui/uikit/Dialog.vue'
import Address from '../../api/models/address'
import { getClientAddresses, getUserInfo } from '../../api/request';
import ActionButton from '../../ui/uikit/ActionButton.vue'
import InputTextField from '../../ui/uikit/inputs/InputTextField.vue'
import { VTimePicker} from 'vuetify/labs/VTimePicker'
import { VDateInput } from 'vuetify/labs/components'

const emit = defineEmits(['update-form-validity', 'update-address-data'])
const clientAddresses = ref<Address[]>([])
const isDialogOpen = ref(false);

function openSelectAddressDialog() {
  isDialogOpen.value = !isDialogOpen.value;
}

const menu = ref<boolean>(false)
const contactData = ref({
  city: "",
  street: "",
  building: "",
  entrance: "",
  floor: "",
  doorNumber: "",
  date: new Date(""),
  time: null
})
const minDate = computed(() => {
  const today = new Date();
  const tomorrow = new Date(today);
  tomorrow.setDate(today.getDate() + 1);
  return tomorrow.toISOString().substr(0, 10);
});
const isFormValid = computed(() => {
  return Object.values(contactData.value).every(value => value !== "" && value !== null);
})

watch(isFormValid, (newVal) => {
  emit('update-form-validity', newVal);
})

watch(contactData, (newVal) => {
  emit('update-address-data', newVal);
}, { deep: true })

function onAddressSelect(address: Address) {
  contactData.value.city = address.city
  contactData.value.street = address.street
  contactData.value.building = address.building
  contactData.value.entrance = address.entrance
  contactData.value.floor = address.floor
  contactData.value.doorNumber = address.door_number
  openSelectAddressDialog()
}

onMounted(() => {
  getUserInfo('me')
  .then((response) => {
    getClientAddresses(response.id)
    .then((response) => clientAddresses.value = response)
  })
  .catch((error) => {
    console.error("Error get worker info:", error)
  })
})
</script>

<template>
  <div class="contact-info-container">
    <h1>Контактная информация</h1>
    <v-form
      class="contact-info-container-form"
      validate-on="submit lazy" 
    >
      <div class="input-row">
        <InputTextField
          v-model="contactData.city"
          placeholder="Введите город"
          type="text"
          label="Город"
        ></InputTextField>
        <InputTextField
          v-model="contactData.street"
          placeholder="Введите улицу"
          type="text"
          label="Улица"
        ></InputTextField>
        <InputTextField
          v-model="contactData.building"
          placeholder="Введите номер дома"
          type="text"
          label="Дом"
          hint="Пример: 4к1 или 8AE"
        ></InputTextField>
        <ActionButton
          id="my-address-btn"
          text="Мои адреса"
          type="button"
          variant="flat"
          color="#394cc2"
          @click="openSelectAddressDialog"
        ></ActionButton>
        <Dialog
          title="Выбор адресса"
          :visible="isDialogOpen"
          dialogMaxWidth="40%"
          @update:visible="openSelectAddressDialog">
          <template #body>
            <ul>
              <li 
                v-for="address in clientAddresses"
                :key="address.id"
                class="address-item"
                @click="onAddressSelect(address)">
                г.{{ address.city }}, ул. {{ address.street }}, д.{{ address.building }}, кв.{{ address.door_number }}
              </li>
            </ul>
          </template>
          <template #footer>
            <ActionButton
              text="Закрыть"
              type="cancel"
              variant="flat"
              color="#394cc2"
              @click="openSelectAddressDialog"
            ></ActionButton>
          </template>
        </Dialog>
      </div>
      <div class="input-row">
        <InputTextField
          v-model="contactData.entrance"
          placeholder="Введите квартиру"
          type="text"
          label="Квартира"
        ></InputTextField>
        <InputTextField
          v-model="contactData.floor"
          placeholder="Введите подъезд"
          type="text"
          label="Подъезд"
        ></InputTextField>
        <InputTextField
          v-model="contactData.doorNumber"
          placeholder="Введите этаж"
          type="text"
          label="Этаж"
        ></InputTextField>
      </div>
      <div class="input-row">
        <v-date-input
          v-model="contactData.date"
          class="input-date"
          label="Дата"
          variant="solo"
          rounded="xl"
          prepend-icon=""
          :min="minDate"
        ></v-date-input>
        <InputTextField
          v-model="contactData.time"
          placeholder="Введите время"
          type="string"
          label="Время"
          :active="menu"
          :focus="menu"
          readonly
        >
          <template #body>
            <v-menu
              v-model="menu"
              :close-on-content-click="false"
              activator="parent"
              transition="scale-transition"
            >
              <v-time-picker
                v-if="menu"
                v-model="contactData.time"
                width="auto"
                format="24hr"
                max="18:00"
                min="6:00"
                scrollable
              ></v-time-picker>
            </v-menu>
          </template>
        </InputTextField>
      </div>
    </v-form>
  </div>
</template>

<style scoped>
.contact-info-container {
  display: flex;
  flex-direction: column;
  text-align: center;
  height: auto;
  width: 95%;
  background-color: white;
  border: 2px solid grey;
  border-radius: 25px;
}
.contact-info-container-form {
  max-height: 100%;
  max-width: 100%;
  display: flex;
  flex-direction: column;
  width: 100%;
  padding-left: 10px;
  padding-right: 10px;
}
.input-row {
  display: flex;
  flex-direction: row;
  align-items: flex-start;
  gap: 10px;
}
.input-date{
  font-weight: bold;
}
#my-address-btn {
  height: 72%;
}
h1 {
  font-size: 42px;
  font-weight: bold;
  margin-bottom: 10px;
}

.address-item {
  background-color: #959FDE;
  border-radius: 20px;
  padding: 10px;
  margin-bottom: 8px;
  cursor: pointer;
  list-style-type: none;
}

.address-item:hover {
  background-color: #8a95d4;
}
</style>