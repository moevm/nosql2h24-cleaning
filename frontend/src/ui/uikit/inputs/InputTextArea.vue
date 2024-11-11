<script setup lang="ts">
import { defineProps, withDefaults } from 'vue'
import type TextArea from '../../model/textarea'

const props = withDefaults(defineProps<TextArea>(), {
  id: undefined,
  class: undefined,
  placeholder: undefined,
  type: undefined
})
</script>

<template>
  <v-textarea
    :id="props.id"
    :class="[props.class, 'input-field']"
    :placeholder="props.placeholder"
    variant="solo"
    rounded="xl"
    :type="props.type"
    :label="props.label"
    :rows="props.rows"
    :row-height="props.rowHeight"
    :maxlength="props.maxLength"
    clearable
    no-resize
    counter
  ></v-textarea>
</template>

<style scoped>
.input-field {
  font-weight: bold;
}
</style>
