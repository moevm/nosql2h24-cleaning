<script setup lang="ts">
import ActionButton from '../../ui/uikit/ActionButton.vue'
import InputTextField from '../../ui/uikit/inputs/InputTextField.vue'
</script>

<template>
  <div class="contact-info-container">
    <h1>Контактная информация</h1>
    <v-form
      class="contact-info-container-form"
      validate-on="submit lazy" 
    >
      <div class="input-row">
        <InputTextField
          placeholder="Введите имя"
          type="text"
          label="Имя"
        ></InputTextField>
        <InputTextField
          placeholder="Введите фамилию"
          type="text"
          label="Фамилия"
        ></InputTextField>
      </div>
        <InputTextField
          placeholder="Введите номер телефона"
          type="text"
          label="Номер телефона"
        ></InputTextField>
      <div class="input-row">
        <InputTextField
          placeholder="Введите адрес"
          type="text"
          label="Адрес"
        ></InputTextField>
        <ActionButton
          id="my-address-btn"
          text="Мои адреса"
          type="my-address"
          variant="flat"
          color="#394cc2"
        ></ActionButton>
      </div>
      <div class="input-row">
        <InputTextField
          placeholder="Введите квартиру"
          type="text"
          label="Квартира"
        ></InputTextField>
        <InputTextField
          placeholder="Введите подъезд"
          type="text"
          label="Подъезд"
        ></InputTextField>
        <InputTextField
          placeholder="Введите этаж"
          type="text"
          label="Этаж"
        ></InputTextField>
        <InputTextField
          placeholder="Введите домофон"
          type="text"
          label="Домофон"
        ></InputTextField>
      </div>
      <div class="input-row">
        <InputTextField
          placeholder="Введите дату"
          type="date"
          label="Дата"
        ></InputTextField>
        <InputTextField
          placeholder="Введите время"
          type="string"
          label="Время"
        ></InputTextField>
      </div>
    </v-form>
  </div>
</template>

<style scoped>
.contact-info-container {
  display: flex;
  flex-direction: column;
  text-align: center;
  height: auto;
  width: 95%;
  background-color: white;
  border: 2px solid grey;
  border-radius: 25px;
}
.contact-info-container-form {
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
  gap: 10px;
}
#my-address-btn {
  height: 72%;
}
h1 {
  font-size: 42px;
  font-weight: bold;
  margin-bottom: 10px;
}
</style>