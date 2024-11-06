<script setup lang="ts">
import ActionButton from './ActionButton.vue'
import type Button from '../model/button'
import { inject } from 'vue'

const sideBarButtons: Array<Button> | undefined = inject('sideBarButtons')
const userCard: {
  name?: string;
  email?: string;
} | undefined = inject('userCard')
</script>

<template>
  <div class="side-bar-container">
    <div class="side-bar-upper-container">
      <ActionButton
        v-for="button in sideBarButtons"
        id="new-btns"
        color="white"
        variant="outlined"
        :type="button.type"
        :text="button.text"
        :to="button.to"
      ></ActionButton>
    </div>
    <div class="side-bar-lower-container">
      <div class="side-bar-user-card">
        <p>{{ userCard?.name }}</p>
        <p>{{ userCard?.email }}</p>
      </div>
      <ActionButton
        id="exit-btn"
        color="white"
        variant="outlined"
        type="exit"
        text="Выход"
        to="/login"
      ></ActionButton>
    </div>
  </div>
</template>

<style scoped>
.side-bar-container {
  max-width: 100%;
  max-height: 100%;
  display: flex;
  flex-direction: column;
  width: 20%;
  height: 100%;
  background-color: #3846c0;
}
.side-bar-upper-container {
  display: flex;
  flex-direction: column;
  justify-content: start;
  align-items: center;
  width: 100%;
  height: 70%;
  margin-top: 20px;
  gap: 20px;
}
.side-bar-lower-container {
  display: flex;
  flex-direction: column;
  justify-content: end;
  align-items: center;
  width: 100%;
  height: 30%;
  margin-bottom: 20px;
  gap: 20px;
}
.side-bar-user-card {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  width: 95%;
  height: 42%;
  background-color: white;
  border-radius: 20px;
  font-size: 28px;
}
#exit-btn, #new-btns {
  max-width: 100%;
  width: 95%;
}
</style>
