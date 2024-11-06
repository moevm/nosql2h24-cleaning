<script setup lang="ts">
import ActionButton from "./ActionButton.vue";
import DateFilter from "./DateFilter.vue";
import { defineEmits } from 'vue';

const emit = defineEmits<{
  (event: 'filter-dates', payload: { startDate: string | null; endDate: string | null }): void;
}>();

function updateDateFilter(dates: { startDate: string | null; endDate: string | null }) {
  emit("filter-dates", dates);
}
</script>

<template>
  <div class="filter-panel-container">
    <ActionButton
      id="list-filters-btn"
      color="#394cc2"
      :text="'Все'"
    ></ActionButton>
    <DateFilter @filter-dates="updateDateFilter" />
    <ActionButton
      @click="updateDateFilter"
      id="search-btn"
      color="#394cc2"
      :text="'Поиск'"
    ></ActionButton>
  </div>
</template>

<style scoped>
.filter-panel-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 8%;
  gap: 10px;
}
</style>
