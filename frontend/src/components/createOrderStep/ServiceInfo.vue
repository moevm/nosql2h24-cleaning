<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'
import HeaderList from '../../ui/uikit/containers/HeaderList.vue'
import ServiceItem from '../../ui/uikit/items/ServiceItem.vue'
import Service from '../../api/models/service'
import { getAllServices } from '../../api/request'

const emit = defineEmits(['update-form-validity', 'update-price'])
const props = defineProps<{ currentPrice: number }>()
const services = ref<Service[]>([])
const price = ref<number>(props.currentPrice)

onMounted(() => {
  emit('update-form-validity', true)
  getAllServices()
  .then((response) => {
    services.value = response
  })
})

watch(price, (newVal) => {
  emit('update-price', newVal);
})

function addServicePrice(servicePrice: number): void {
  price.value += servicePrice;
}

function removeServicePrice(servicePrice: number): void {
  price.value -= servicePrice;
}
</script>

<template>
  <HeaderList
    title="Услуги" 
    :items="services"
    height="58%"
    width="95%"
  >
    <template #items="{ item }">
      <ServiceItem
        :isAdmin="false"
        :service="item"
        @add-service-price="addServicePrice"
        @remove-service-price="removeServicePrice"
      ></ServiceItem>
    </template>
    <template #bottom>
      <div class="service-info-price-container">
        <div class="service-info-price">
          <p>Итого: {{ price }}₽</p>
        </div>
      </div>
    </template>
  </HeaderList>
</template>

<style scoped>
.service-info-price-container {
  display: flex;
  flex-direction: row;
  width: 95%;
  height: auto;
  padding-left: 10px;
  padding-right: 10px;
  margin-top: 10px;
}
.service-info-price {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  width: 49%;
  border: 3px solid #3846c0;
  border-radius: 15px;
  padding: 10px;
  text-align: left;
}
p {
  font-size: 32px;
  font-weight: bold;
}
</style>