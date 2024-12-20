<script setup lang="ts">
import { Ref, ref, computed } from 'vue'
import MainContainer from '../../ui/uikit/containers/MainContainer.vue'
import ActionButton from '../../ui/uikit/ActionButton.vue'
import Step1 from '../createOrderStep/ContactInfo.vue'
import Step2 from '../createOrderStep/AppartmentInfo.vue'
import Step3 from '../createOrderStep/ServiceInfo.vue'
import { createOrder, getUserInfo } from '../../api/request'
import Order from '../../api/models/order'
import Service from '../../api/models/service'
import { useRouter } from 'vue-router'

const router = useRouter();

const currentStep: Ref<number> = ref(1)
const steps: any = {
  1: Step1,
  2: Step2,
  3: Step3,
}
const currentStepComponent = computed(() => steps[currentStep.value])
const currentPrice = ref<number>(0)
const isFormValid = ref<boolean>(false)
const userId = ref<string>('')
const appartmentData = ref({
  rooms: 0,
  bathrooms: 0,
  area: 0,
  pollution: 0,
  comment: "",
})
const addressData = ref({
  city: "",
  street: "",
  building: "",
  entrance: "",
  floor: "",
  doorNumber: "",
  date: new Date(""),
  time: ""
})
const selectedServices = ref<Service[]>([])

function changeStep(): void {
  if (currentStep.value >= 3) {
    getUserInfo('me').then((user) => {
      userId.value = user.id
      const [hours, minutes] = addressData.value.time.split(':').map(Number)
      const dateWithTime = new Date(addressData.value.date)
      dateWithTime.setHours(hours)
      dateWithTime.setMinutes(minutes)
      const requiredWorkers = Math.max(
        ...selectedServices.value.map((service) => service.workers_quantity)
      )
      const newOrder: Order = {
        user_id: userId.value,
        address: {
          city: addressData.value.city,
          street: addressData.value.street,
          building: addressData.value.building,
          entrance: addressData.value.entrance,
          floor: addressData.value.floor,
          door_number: addressData.value.doorNumber,
        },
        date_time: dateWithTime,
        area: appartmentData.value.area,
        number_of_rooms: appartmentData.value.rooms,
        number_of_bathrooms: appartmentData.value.bathrooms,
        pollution: appartmentData.value.pollution,
        comment: appartmentData.value.comment,
        price: currentPrice.value,
        required_workers: requiredWorkers,
        services: selectedServices.value.map((service) => service.id),
      }
      createOrder(newOrder).then(() => {
        router.push({ name: 'client-orders-history' })
      })
    })
  } else {
    currentStep.value += 1;
    isFormValid.value = false
  }
}

function updateFormValidity(isValid: boolean): void {
  isFormValid.value = isValid
}

function updateAppartmentData(data: any): void {
  appartmentData.value = data;

}
function updateAddressData(data: any): void {
  addressData.value = data;
}

function updatePrice(num: number): void {
  currentPrice.value = num
}

function updateSelectedServices(services: Service[]): void {
  selectedServices.value = services;
}
</script>

<template>
  <MainContainer
    justifyContent="flex-start"
  >
    <component
      :is="currentStepComponent"
      @update-form-validity="updateFormValidity"
      @update-appartment-data="updateAppartmentData"
      @update-address-data="updateAddressData"
      @update-price="updatePrice"
      @update-selected-services="updateSelectedServices"
      :currentPrice="currentPrice"
    ></component>
    <ActionButton
      id="next-btn"
      text="Далее"
      type="next"
      variant="flat"
      color="#394cc2"
      :onClick="changeStep"
      :disabled="!isFormValid"
    ></ActionButton>
  </MainContainer>
</template>

<style scoped>
#next-btn {
  width: 95%;
}
</style>
