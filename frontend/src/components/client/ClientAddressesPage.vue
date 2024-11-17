<script setup lang="ts">
import { Ref, ref, onMounted } from 'vue'
import HeaderList from '../../ui/uikit/containers/HeaderList.vue'
import AddressItem from '../../ui/uikit/items/AddressItem.vue'
import MainContainer from '../../ui/uikit/containers/MainContainer.vue'
import Dialog from '../../ui/uikit/Dialog.vue'
import ActionButton from '../../ui/uikit/ActionButton.vue'
import InputTextField from '../../ui/uikit/inputs/InputTextField.vue'
import { getClientAddresses, createNewAddress } from '../../api/request'
import Address from '../../api/models/address'

onMounted(() => {
  fetchClientAddresses("10")
})

const addresses = ref<Address[]>([])
const isDialogVisible: Ref<boolean> = ref(false)
const newAddress = ref({
  city: '',
  street: '',
  building: '',
  entrance: '',
  floor: '',
  door_number: ''
})


function openDialog(): void {
  isDialogVisible.value = true
}

function closeDialog(): void {
  isDialogVisible.value = false
}

function formatAddress(item: any): string {
  return `${item.city}, ${item.street}, ${item.building}, Подъезд ${item.entrance}, Этаж ${item.floor}, Квартира ${item.door_number}`;
}

async function fetchClientAddresses(clientId: string) {
  getClientAddresses(clientId)
  .then((response) => {
    addresses.value = response
  })
  .catch((error) => {
    console.error("Failed to fetch client addresses:", error)
  })
}

async function addNewAddress(clientId: string) {
  closeDialog()
  const addressData: Address = {
    id: clientId.toString(),
    city: newAddress.value.city,
    street: newAddress.value.street,
    building: newAddress.value.building,
    entrance: newAddress.value.entrance,
    floor: newAddress.value.floor,
    door_number: newAddress.value.door_number,
    created_at: new Date(),
    updated_at: new Date()
  };
  addresses.value.push(addressData);

  createNewAddress(clientId, addressData)
  .then((newAddressId) => {
    /* раскомментировать после добавления авторизации
    addresses.value.push({
      ...addressData, 
      id: newAddressId
    });
    */
  })
  .catch((error) => {
    console.error("Failed to createNewAddress:", error)
    // addresses.value.pop() раскомментировать после добавления авторизации
  })
}
</script>

<template>
  <MainContainer>
    <HeaderList
      title="Мои адреса" 
      :items="addresses"
      height="100%"
      width="95%"
    >
      <template #items="{ item }">
        <AddressItem :address="formatAddress(item)" />
      </template>
    </HeaderList>
  </MainContainer>
  <ActionButton
    text="Добавить"
    type="add"
    color="#394cc2"
    @click="openDialog" 
    class="action-button"
  />
  <Dialog
    title="Добавление адреса"
    :visible="isDialogVisible"
    dialogMaxWidth="30%"
    @update:visible="closeDialog">
    <template #body>
      <InputTextField
        v-model="newAddress.city"
        placeholder="Город"
        type="text"
        label="Город"
      ></InputTextField>
      <InputTextField
        v-model="newAddress.street"
        placeholder="Улица"
        type="text"
        label="Улица"
      ></InputTextField>
      <InputTextField
        v-model="newAddress.building"
        placeholder="Дом"
        type="number"
        label="Дом"
      ></InputTextField>
      <InputTextField
        v-model="newAddress.entrance"
        placeholder="Подъезд"
        type="text"
        label="Подъезд"
      ></InputTextField>
      <InputTextField
        v-model="newAddress.floor"
        placeholder="Этаж"
        type="text"
        label="Этаж"
      ></InputTextField>
      <InputTextField
        v-model="newAddress.door_number"
        placeholder="Квартира"
        type="text"
        label="Квартира"
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
        text="Добавить"
        type="create"
        variant="flat"
        color="#394cc2"
        @click="addNewAddress"
      ></ActionButton>
    </template>
  </Dialog>
</template>

<style scoped>
.action-button {
  position: fixed;
  bottom: 30px;
  right: 45px;
  z-index: 1000;
}
</style>
