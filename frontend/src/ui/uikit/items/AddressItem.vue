<script setup lang="ts">
import { defineProps } from 'vue'
import Address from '../../../api/models/address'

const props = defineProps<{
  address: Address;
}>()
</script>

<template>
  <div class="address-item">
    <p>
      {{ props.address.city }}, {{ props.address.street }}, 
      дом {{ props.address.building }}, подъезд {{ props.address.entrance }},
      этаж {{ props.address.floor }}, квартира {{ props.address.door_number }}
    </p>
  </div>
</template>

<style scoped>
.address-item {
  display: flex;
  flex-direction: row;
  justify-content: top;
  align-items: center;
  width: 49%;
  border: 3px solid #3846c0;
  border-radius: 15px;
  padding: 10px;
  text-align: left;
}
p {
  font-size: 28px;
  font-weight: bold;
}
</style>
