<script setup lang="ts">
import { ref } from 'vue';
import { defineEmits } from 'vue';

const emit = defineEmits<{
  (event: 'filter-dates', payload: { startDate: string | null; endDate: string | null }): void;
}>();

const startDate = ref<string | null>(null);
const endDate = ref<string | null>(null);

function updateDateFilter() {
  emit("filter-dates", { startDate: startDate.value, endDate: endDate.value });
}
</script>

<template>
  <div class="date-filter">
    <input type="date" v-model="startDate" placeholder="C" @change="updateDateFilter" />
    <input type="date" v-model="endDate" placeholder="По" @change="updateDateFilter" />
  </div>
</template>

<style scoped>
.date-filter {
  border: 2px solid grey;
  border-radius: 25px;
  padding: 5px;
  padding-left: 10px;
}
</style>
