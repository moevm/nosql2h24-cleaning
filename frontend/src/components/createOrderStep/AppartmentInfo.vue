<script setup lang="ts">
import { computed, ref, watch } from 'vue';
import InputNumber from '../../ui/uikit/inputs/InputNumber.vue'
import InputTextArea from '../../ui/uikit/inputs/InputTextArea.vue'

const emit = defineEmits(['update-form-validity', 'update-price'])
const appartmentData = ref({
  rooms: 0,
  bathrooms: 0,
  area: 0,
  pollution: 0,
  comment: "",
})

const isFormValid = computed(() => {
  const { rooms, bathrooms, area, pollution } = appartmentData.value;
  return rooms !== 0 && bathrooms !== 0 && area !== 0 && pollution !== 0;
})

const price = computed(() => {
  const { rooms, bathrooms, area, pollution } = appartmentData.value;
  return rooms * 100 + bathrooms * 100 + area * 150 + pollution * 200;
})

watch(isFormValid, (newVal) => {
  emit('update-form-validity', newVal)
})

watch(price, (newVal) => {
  emit('update-price', newVal)
}, {deep: true})
</script>

<template>
  <div class="appartment-info-container">
    <h1>Квартира</h1>
    <v-form
      class="appartment-info-container-form"
      validate-on="submit lazy" 
    >
      <div class="input-row">
        <InputNumber
          v-model="appartmentData.rooms"
          label="Количество комнат"
          :min="1"
          :max="10"
          :step="1"
        ></InputNumber>
        <InputNumber
          v-model="appartmentData.bathrooms"
          label="Количество санузлов"
          :min="1"
          :max="10"
          :step="1"
        ></InputNumber>
      </div>
      <div class="input-row">
        <InputNumber
          v-model="appartmentData.area"
          label="Площадь квартиры в м²"
          :min="5"
          :max="120"
          :step="1"
        ></InputNumber>
        <InputNumber
          v-model="appartmentData.pollution"
          label="Загрезненность"
          :min="1"
          :max="10"
          :step="1"
        ></InputNumber>
      </div>
      <InputTextArea
        v-model="appartmentData.comment"
        label="Комментарий к заказу"
        placeholder="Введите комментарий"
        :rows="3"
        :row-height="30"
        :max-length="250"
      ></InputTextArea>
    </v-form>
  </div>
</template>

<style scoped>
.appartment-info-container {
  display: flex;
  flex-direction: column;
  text-align: center;
  height: auto;
  width: 95%;
  background-color: white;
  border: 2px solid grey;
  border-radius: 25px;
}
.appartment-info-container-form {
  max-height: 100%;
  max-width: 100%;
  display: flex;
  flex-direction: column;
  width: 100%;
  padding-left: 10px;
  padding-right: 10px;
}
.input-row {
  display: flex;
  flex-direction: row;
  align-items: flex-start;
  width: 100%;
  height: 100%;
  gap: 10px;
}
.input-field-row {
  width: 13.3%;
}
h1 {
  font-size: 42px;
  font-weight: bold;
  margin-bottom: 10px;
}
/* .container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  max-width: 100%;
  height: 270px;
  border: 2px solid grey;
  border-radius: 25px;
  padding-top: 20px;
  padding-bottom: 20px;
  padding-left: 15px;
  padding-right: 15px;
}

.input-row {
  display: flex;
  width: 1000px;
  height: 50px;
  gap: 10px;
  margin-bottom: 15px;
}

h1 {
  font-size: 42px;
  font-weight: bold;
} */
</style>