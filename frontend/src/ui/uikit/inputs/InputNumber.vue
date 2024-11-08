<script setup lang="ts">
import { defineProps, withDefaults } from 'vue'
import NumberField from '../../model/numberfield'
import { VNumberInput } from 'vuetify/labs/components'

const props = withDefaults(defineProps<NumberField>(), {
    id: undefined,
    class: undefined,
    type: undefined
})
</script>

<template>
  <v-number-input
    :id="props.id"
    :class="[props.class, 'input-field']"
    :type="props.type"
    :label="props.label"
    :min="props.min"
    :max="props.max"
    :step="props.step"
    variant="outlined"
    control-variant="split"
  ></v-number-input>
</template>

<style scoped>
.input-field {
    font-weight: bold;
}
</style>
