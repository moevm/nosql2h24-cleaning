<script setup lang="ts">
import ActionButton from '../ui/uikit/ActionButton.vue'
import InputTextField from '../ui/uikit/inputs/InputTextField.vue'
import { postRegister, postLogin } from '../api/request'
import { useRouter } from 'vue-router'
import { Ref, ref, computed } from 'vue'

const router = useRouter()
const isRegistrationAuthorization: Ref<boolean> = ref(true)
const loading: Ref<boolean> = ref(false)

// user data
const email: Ref<string> = ref('')
const name: Ref<string> = ref('')
const password: Ref<string> = ref('')
const patronymic: Ref<string> = ref('')
const phoneNumber: Ref<string> = ref('')
const surname: Ref<string> = ref('')

function defaultUserData(): void {
  email.value = ''
  name.value = ''
  password.value = ''
  patronymic.value = ''
  phoneNumber.value = ''
  surname.value = ''
}

function handleAuthorizationClick(): void {
  defaultUserData()
  isRegistrationAuthorization.value = true
}

function handleRegistrationClick(): void {
  defaultUserData()
  isRegistrationAuthorization.value = false
}

const actionContainerStyle: any = computed(() => {
  return {
    height: isRegistrationAuthorization.value ? '350px' : '700px'
  }
})

function handleSubmit(): void {
  loading.value = true
  if (isRegistrationAuthorization.value) {
    postLogin({
      email: email.value,
      password: password.value
    }).then((success) => {
      defaultUserData()
      loading.value = false
      router.push('/cleaning/')
    }).catch((error) => {
      loading.value = false
    })
  } else {
    postRegister({
      email: email.value,
      name: name.value,
      password: password.value,
      patronymic: patronymic.value,
      phone_number: phoneNumber.value,
      surname: surname.value
    }).then((success) => {
      defaultUserData()
      isRegistrationAuthorization.value = true
      loading.value = false
    }).catch((error) => {
      loading.value = false
    })
  }
}
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
        @submit.prevent="handleSubmit"
      >
        <template v-if="!isRegistrationAuthorization">
          <InputTextField
            v-model="name"
            class="input-field"
            placeholder="Введите имя"
            type="text"
            label="Имя"
          ></InputTextField>
          <InputTextField
            v-model="surname"
            class="input-field"
            placeholder="Введите фамилию"
            type="text"
            label="Фамилия"
          ></InputTextField>
          <InputTextField
            v-model="patronymic"
            class="input-field"
            placeholder="Введите отчество"
            type="text"
            label="Отчество"
          ></InputTextField>
          <InputTextField
            v-model="phoneNumber"
            class="input-field"
            placeholder="Введите телефон"
            type="phonenumber"
            label="Номер телефона"
          ></InputTextField>
        </template>
        <InputTextField
          v-model="email"
          class="input-field"
          placeholder="Введите почту"
          type="email"
          label="Почта"
        ></InputTextField>
        <InputTextField
          v-model="password"
          class="input-field"
          placeholder="Введите пароль"
          type="password"
          label="Пароль"
        ></InputTextField>
        <ActionButton
          id="submit-btn"
          :loading="loading"
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
