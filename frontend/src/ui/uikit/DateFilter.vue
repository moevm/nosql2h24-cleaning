<script setup lang="ts">
import { ref } from 'vue'

const startDate = ref<string | null>(null);
const endDate = ref<string | null>(null);

</script>

<template>
  <div class="date-filter">
    <input type="date" v-model="startDate" placeholder="C" />
    <input type="date" v-model="endDate" placeholder="По" />
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
