<template>
  <div class="container">
    <HeaderList 
      title="Опции" 
      :items="services"
      height="82%"
      width="95%"
      >
      <template #items="{ item }">
        <ServiceItem :service="item" />
      </template>
    </HeaderList>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import HeaderList from "../../ui/uikit/HeaderList.vue"
import ServiceItem from "../../ui/uikit/ServiceItem.vue"

const services = ref([
  { name: "Уборка балкона", price: 10 },
  { name: "Мойка окон", price: 10 },
  { name: "Чистка холодильника", price: 10 },
  { name: "Уборка лотка", price: 10 },
  { name: "Глажка", price: 10 },
  { name: "Мойка унитаза", price: 10 },
]) // TODO DB request
</script>

<style scoped>
.header-container {
  padding-bottom: 15px;
  max-width: 1500px;
}
</style>