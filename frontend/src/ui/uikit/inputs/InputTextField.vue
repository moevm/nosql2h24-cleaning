<script setup lang="ts">
import { defineProps, withDefaults } from 'vue'
import type TextField from '../../model/textfield'

const props = withDefaults(defineProps<TextField>(), {
  id: undefined,
  class: undefined,
  placeholder: undefined,
  type: undefined,
  rules: undefined,
})
</script>

<template>
  <v-text-field
    :id="props.id"
    :class="[props.class, 'input-field']"
    :placeholder="props.placeholder"
    variant="solo"
    rounded="xl"
    :type="props.type"
    :label="props.label"
    :rules="props.rules"
    clearable
  >
    <slot
      name="body"
    ></slot>
  </v-text-field>
</template>

<style scoped>
.input-field {
  font-weight: bold;
}
</style>
