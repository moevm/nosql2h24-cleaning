<script setup lang="ts">
import { defineProps, withDefaults, defineEmits } from 'vue'

const props = withDefaults(defineProps<{
  title: string;
  visible: boolean;
  dialogMaxWidth?: string;
  dialogMaxHeight?: string;
}>(), {
  dialogMaxWidth: 'auto',
  dialogMaxHeight: 'auto',
})

const emit = defineEmits<{
  (event: 'update:visible'): void
}>()

function closeDialog(): void {
  emit('update:visible')
}
</script>

<template>
  <v-dialog
    v-model="props.visible"
    :max-width="props.dialogMaxWidth"
    :max-height="props.dialogMaxHeight"
    @click:outside="closeDialog"
  >
    <v-card
      class="card"
      rounded="xl"
      height="auto"
      width="auto"
    >
      <v-card-title
        class="card-title"
      >
        <h3>{{ props.title }}</h3>
      </v-card-title>
      <v-card-text>
        <slot name="body"></slot>
      </v-card-text>
      <v-card-actions>
        <slot name="footer"></slot>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<style scoped>
.card-title {
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
