<template v-if="step == 1">
  <div class="container">
    <h1 class="title">Квартира</h1>
    <v-form
      class="action-container-form"
      validate-on="submit lazy" 
    >

      <div class="input-row">
        <InputTextField
          class="input-field"
          placeholder="Количество комнат"
          type="text"
          label="Количество комнат"
        />
        <InputTextField
          class="input-field"
          placeholder="Количество санузлов"
          type="text"
          label="Количество санузлов"
        />
      </div>

      <div class="input-row">
        <InputTextField
          class="input-field"
          placeholder="Площадь квартиры"
          type="text"
          label="Площадь квартиры"
        />
        <InputTextField
          class="input-field"
          placeholder="Загрязненность"
          type="text"
          label="Загрязненность"
        />
      </div>

      <div class="input-row">
        <InputTextField
          class="input-field"
          placeholder="Комментарий к заказу"
          type="string"
          label="Комментарий"
        />
      </div>
    </v-form>
  </div>
</template>

<script setup lang="ts">
import InputTextField from '../../ui/uikit/InputTextField.vue'
</script>


<style scoped>
.container {
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
}
</style>