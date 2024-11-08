<script setup lang="ts">
import { ref } from 'vue'
import HeaderList from '../../ui/uikit/containers/HeaderList.vue'
import ServiceItem from '../../ui/uikit/items/ServiceItem.vue'

const services = ref([
  {id: 1, name: "Уборка", price: "10$"},
  {id: 2, name: "Уборка", price: "10$"},
  {id: 3, name: "Уборка", price: "10$"}
]) // TODO DB request
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
        :service="item">
      </ServiceItem>
    </template>
    <template #bottom>
      <div class="service-info-price-container">
        <div class="service-info-price">
          <p>Итого: 0$</p>
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