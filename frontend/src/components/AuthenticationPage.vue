<script setup lang="ts">
import ActionButton from '../ui/uikit/ActionButton.vue'
import InputTextField from '../ui/uikit/InputTextField.vue'
import { Ref, ref, computed } from 'vue'

const isRegistrationAuthorization: Ref<boolean, boolean> = ref(true)

function handleAuthorizationClick(): void {
  isRegistrationAuthorization.value = true
}

function handleRegistrationClick(): void {
  isRegistrationAuthorization.value = false
}

const actionContainerStyle: any = computed(() => {
  return {
    height: isRegistrationAuthorization.value ? '350px' : '700px'
  }
})
</script>

<template>
  <div class="container">
    <div
      class="action-container"
      :style="actionContainerStyle"
      >
      <div class="action-container-btns">
        <ActionButton
          id="authorization-btn"
          color="white"
          variant="outlined"
          :isActive="isRegistrationAuthorization"
          type="authorization"
          text="Авторизация"
          :onClick="handleAuthorizationClick"
        ></ActionButton>
        <ActionButton
          id="registration-btn"
          color="white"
          variant="outlined"
          :isActive="!isRegistrationAuthorization"
          type="registration"
          text="Регистрация"
          :onClick="handleRegistrationClick"
        ></ActionButton>
      </div>
      <v-form
        class="action-container-form"
        validate-on="submit lazy"
        @submit.prevent=""
      >
        <template v-if="!isRegistrationAuthorization">
          <InputTextField
            class="input-field"
            placeholder="Введите имя"
            type="text"
            label="Имя"
          ></InputTextField>
          <InputTextField
            class="input-field"
            placeholder="Введите фамилию"
            type="text"
            label="Фамилия"
          ></InputTextField>
          <InputTextField
            class="input-field"
            placeholder="Введите отчество"
            type="text"
            label="Отчество"
          ></InputTextField>
          <InputTextField
            class="input-field"
            placeholder="Введите телефон"
            type="phonenumber"
            label="Номер телефона"
          ></InputTextField>
        </template>
        <InputTextField
          class="input-field"
          placeholder="Введите почту"
          type="email"
          label="Почта"
        ></InputTextField>
        <InputTextField
          class="input-field"
          placeholder="Введите пароль"
          type="password"
          label="Пароль"
        ></InputTextField>
        <ActionButton
          id="submit-btn"
          color="white"
          variant="outlined"
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
  transition: all 0.3s ease;
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
.action-container-form {
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
}
.input-field {
  width: 95%;
}
#submit-btn {
  max-width: 100%;
  width: 95%;
}
</style>
