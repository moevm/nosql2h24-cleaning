<script setup lang="ts">
import { defineProps, withDefaults } from 'vue';

const props = withDefaults(defineProps<{
  id?: string,
  class?: string,
  placeholder?: string,
  type?: string,
  label: string
}>(), {
  id: undefined,
  class: undefined,
  placeholder: undefined,
  type: undefined
})
</script>

<template>
  <v-text-field
    :id="props.id"
    :class="props.class"
    :placeholder="props.placeholder"
    variant="solo"
    :type="props.type"
    :label="props.label"
  ></v-text-field>
</template>

<style scoped>
</style>
