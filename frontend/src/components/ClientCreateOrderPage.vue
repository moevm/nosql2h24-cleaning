<template>
  <div class="container">
    <h1 class="title">Контактная информация</h1>
    <v-form
      class="action-container-form"
      validate-on="submit lazy"
      @submit.prevent="submitForm"
    >
      <!-- Имя и Фамилия -->
      <div class="input-row">
        <InputTextField
          class="input-field"
          placeholder="Ваше имя"
          type="text"
        />
        <InputTextField
          class="input-field"
          placeholder="Ваша фамилия"
          type="text"
        />
      </div>

      <!-- Номер телефона -->
      <div class="input-row">
        <InputTextField
          class="input-field full-width"
          placeholder="Номер телефона"
          type="tel"
        />
      </div>

      <!-- Адрес и кнопка "Мои адреса" -->
      <div class="input-row address-row">
        <InputTextField
          class="input-field full-width"
          placeholder="Адрес"
          type="text"
        />
        <ActionButton
          id="my-address-btn"
          color="primary"
          variant="outlined"
          text="Мои адреса"
        />
      </div>

      <!-- Квартира, Подъезд, Этаж, Домофон -->
      <div class="input-row">
        <InputTextField
          class="input-field"
          placeholder="Квартира"
          type="text"
        />
        <InputTextField
          class="input-field"
          placeholder="Подъезд"
          type="text"
        />
        <InputTextField
          class="input-field"
          placeholder="Этаж"
          type="text"
        />
        <InputTextField
          class="input-field"
          placeholder="Домофон"
          type="text"
        />
      </div>

      <!-- Дата и Время -->
      <div class="input-row">
        <InputTextField
          class="input-field"
          placeholder="Дата"
          type="date"
        />
        <InputTextField
          class="input-field"
          placeholder="Время"
          type="time"
        />
      </div>
    </v-form>
    <ActionButton
      id="next-btn"
      color="primary"
      variant="contained"
      text="Далее"
      @click="submitForm"
    />
  </div>
</template>

<script setup lang="ts">
function submitForm(): void {
  console.log("Форма отправлена");
}
</script>

<style scoped>
.container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100%;
  width: 80%;
  padding-top: 20px;
  padding-bottom: 20px;
}

.title {
  text-align: center;
  font-size: 20px;
  font-weight: bold;
  margin-bottom: 20px;
}

.input-row {
  display: flex;
  width: 1000px;
  height: 50px;
  gap: 10px;
  margin-bottom: 15px;
}

.input-field {
  flex: 1;
  padding: 10px;
  border: 1px solid grey;
  border-radius: 8px;
  font-size: 14px;
}

.full-width {
  flex: 1;
}

.address-row {
  display: flex;
}

#my-address-btn {
  width: 150px;
  padding: 10px 15px;
  font-size: 14px;
  border-radius: 8px;
  background-color: #394cc2;
  cursor: pointer;
}

.button-row {
  text-align: center;
}

#next-btn {
  width: 1000px;
  padding: 12px;
  font-size: 16px;
  background-color: #394cc2;
  color: white;
  border-radius: 8px;
  cursor: pointer;
  border: none;
}

#next-btn:hover {
  background-color: #394cc2;
}
</style>
