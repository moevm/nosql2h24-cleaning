<script setup lang="ts">
import { defineProps, withDefaults } from 'vue'
import type Button from '../model/button'

const props = withDefaults(defineProps<Button>(), {
  id: undefined,
  class: undefined,
  isActive: true,
  type: undefined,
  to: undefined,
  onClick: undefined
})
</script>

<template>
  <v-btn
    :id="props.id"
    :class="[props.class, { 'inactive-btn': props.isActive === false }, 'btn']"
    color="white"
    rounded="xl"
    variant="outlined"
    :type="props.type"
    :text="props.text"
    :to="props.to"
    @click="props.onClick"
  ></v-btn>
</template>

<style scoped>
.inactive-btn {
  opacity: 0.5;
}
.btn {
  font-size: 16px;
  font-weight: bold;
}
</style>
