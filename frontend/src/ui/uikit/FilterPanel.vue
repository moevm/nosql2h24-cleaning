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
  <div class="filter-panel">
    <ActionButton
      id="list-filters-btn"
      :text="'Все'"
    ></ActionButton>
    <DateFilter @filter-dates="updateDateFilter" />
    <ActionButton
      @click="updateDateFilter"
      id="search-btn"
      :text="'Поиск'"
    ></ActionButton>
  </div>
</template>

<style scoped>
.filter-panel {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 10px;
  margin-bottom: 10px;
}
.btn {
  font-size: 16px;
  background-color: #3846c0;
  color: black;
  padding-left: 20px;
  padding-right: 20px;
  height: 40px;
}
</style>
