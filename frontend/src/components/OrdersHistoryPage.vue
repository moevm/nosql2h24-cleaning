<script setup lang="ts">
import { Ref, ref, computed } from 'vue'
import FilterPanel from "../ui/uikit/FilterPanel.vue"
import HeaderList from "../ui/uikit/HeaderList.vue"
const startDate: Ref<string | null> = ref(null)
const endDate: Ref<string | null> = ref(null) 

const orders = ref([
        { id: 1, date: "28.02.2024", time: "10:00", address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4", price: 37 },
        { id: 2, date: "01.01.2024", time: "17:00", address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4", price: 102 },
        { id: 4, date: "01.01.2024", time: "17:00", address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4", price: 102 },
        { id: 5, date: "01.01.2024", time: "17:00", address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4", price: 102 },
        { id: 6, date: "01.01.2024", time: "17:00", address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4", price: 102 },
        { id: 7, date: "01.01.2024", time: "17:00", address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4", price: 102 },
        { id: 8, date: "01.01.2024", time: "17:00", address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4", price: 102 },
        { id: 9, date: "01.01.2024", time: "17:00", address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4", price: 102 },
        { id: 10, date: "01.01.2024", time: "17:00", address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4", price: 102 },
        { id: 11, date: "01.01.2024", time: "17:00", address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4", price: 102 },
        { id: 12, date: "15.12.2023", time: "17:00", address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4", price: 42 }
      ], // TODO DB request
  );  

const filteredOrders = computed(() => {
  return orders.value.filter(order => {
    const orderDate = new Date(order.date);
    const start = startDate.value ? new Date(startDate.value) : null;
    const end = endDate.value ? new Date(endDate.value) : null;

    return (
      (!start || orderDate >= start) &&
      (!end || orderDate <= end)
    );
  });
});

const emit = defineEmits<{
  (event: 'filter-dates', payload: { startDate: string | null; endDate: string | null }): void;
  (event: 'reset-filters'): void;
}>();

function resetFilters() {
  startDate.value = null;
  endDate.value = null;
  emit('reset-filters');
}

</script>

<template>
  <div class="orders-history">
    <FilterPanel @filter-dates="filteredOrders" @reset-filters="resetFilters" />
    <HeaderList title="История заказов" :orders="filteredOrders" />
  </div>
</template>

<style scoped>
.orders-history {
  display: flex;
  flex-direction: column;
  height: 100%;
  width: 100vw;
  padding: 20px;
}
</style>
