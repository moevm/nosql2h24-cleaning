<template v-if="step == 1">
  <div class="container">
    <h1 class="title">Контактная информация</h1>
    <v-form
      class="action-container-form"
      validate-on="submit lazy" 
    >

      <div class="input-row">
        <InputTextField
          class="input-field"
          placeholder="Ваше имя"
          type="text"
          label="Имя"
        />
        <InputTextField
          class="input-field"
          placeholder="Ваша фамилия"
          type="text"
          label="Фамилия"
        />
      </div>

      <div class="input-row">
        <InputTextField
          class="input-field full-width"
          placeholder="Номер телефона"
          type="text"
          label="Номер телефона"
        />
      </div>

      <div class="input-row address-row">
        <InputTextField
          class="input-field full-width"
          placeholder="Адрес"
          type="text"
          label="Адрес"
        />
        <ActionButton
          id="my-address-btn"
          text="Мои адреса"
          variant="text"
        />
      </div>

      <div class="input-row">
        <InputTextField
          class="input-field"
          placeholder="Квартира"
          type="text"
          label="Квартира"
        />
        <InputTextField
          class="input-field"
          placeholder="Подъезд"
          type="text"
          label="Подъезд"
        />
        <InputTextField
          class="input-field"
          placeholder="Этаж"
          type="text"
          label="Этаж"
        />
        <InputTextField
          class="input-field"
          placeholder="Домофон"
          type="text"
          label="Домофон"
        />
      </div>

      <div class="input-row">
        <InputTextField
          class="input-field"
          placeholder="Дата"
          type="date"
          label="Дата"
        />
        <InputTextField
          class="input-field"
          placeholder="Время"
          type="string"
          label="Время"
        />
      </div>
    </v-form>
  </div>
</template>

<script setup lang="ts">
import ActionButton from '../../ui/uikit/ActionButton.vue'
import InputTextField from '../../ui/uikit/InputTextField.vue'
</script>


<style scoped>
.container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  max-width: 100%;
  height: 450px;
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

.full-width {
  flex: 1;
}

.address-row {
  display: flex;
}

#my-address-btn {
  width: 150px;
  height: 56px;
  padding: 10px 15px;
  font-size: 14px;
  background-color: #394cc2;
  color: white;
  cursor: pointer;
}

h1 {
  font-size: 42px;
  font-weight: bold;
}
</style>