<script setup lang="ts">
import ActionButton from '../ui/uikit/ActionButton.vue'
import InputTextField from '../ui/uikit/InputTextField.vue'
import { ref } from 'vue'

const isRegistrationAuthorization = ref(true);

const handleAuthorizationClick = () => {
  isRegistrationAuthorization.value = true;
};

const handleRegistrationClick = () => {
  isRegistrationAuthorization.value = false;
};
</script>

<template>
  <div class="container">
    <div class="action-container">
      <div class="action-container-btns">
        <ActionButton
          id="authorization-btn"
          :isActive="isRegistrationAuthorization"
          type="authorization"
          text="Авторизация"
          :onClick="handleAuthorizationClick"
        ></ActionButton>
        <ActionButton
          id="registration-btn"
          :isActive="!isRegistrationAuthorization"
          type="registration"
          text="Регистрация"
          :onClick="handleRegistrationClick"
        ></ActionButton>
      </div>
      <v-form
        class="action-form"
        validate-on="submit lazy"
        @submit.prevent=""
      >
        <InputTextField
          class="input-field"
          placeholder="example@mail.com"
          variant="solo"
          type="email"
          label="Почта"
        ></InputTextField>
        <InputTextField
          class="input-field"
          placeholder="********"
          variant="solo"
          type="password"
          label="Пароль"
        ></InputTextField>
        <ActionButton
          id="submit-btn"
          type="submit"
          :text="isRegistrationAuthorization ? 'Войти' : 'Зарегистрироваться'"
        ></ActionButton>
      </v-form>
    </div>
  </div>
</template>

<style scoped>
.container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}
.action-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 350px;
  width: 600px;
  background-color: #394cc2;
  border-radius: 50px;
}
.action-container-btns {
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  width: 100%;
  gap: 10px;
  margin-bottom: 30px;
}
.action-form {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: 10px;
  width: 100%;
}
#authorization-btn, #registration-btn {
  max-width: 100%;
  width: 47%;
  font-size: 16px;
  font-weight: bold;
}
.input-field {
  width: 95%;
  border-radius: 4px;
  font-weight: bold;
}
#submit-btn {
  max-width: 100%;
  width: 95%;
  font-size: 16px;
  font-weight: bold;
}
</style>
