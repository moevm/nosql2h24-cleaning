<script setup lang="ts">
import { defineProps, withDefaults } from 'vue';

const props = withDefaults(defineProps<{
  id?: string,
  class?: string,
  isActive?: boolean,
  type?: string,
  text: string,
  onClick?: Function
}>(), {
  id: undefined,
  class: undefined,
  isActive: true,
  type: undefined,
  onClick: undefined
})
</script>

<template>
  <v-btn
    :id="props.id"
    :class="[props.class, { 'inactive-btn': props.isActive === false }]"
    color="white"
    rounded="xl"
    variant="outlined"
    :type="props.type"
    :text="props.text"
    @click="props.onClick"
  ></v-btn>
</template>

<style scoped>
.inactive-btn {
  opacity: 0.5;
}
</style>
