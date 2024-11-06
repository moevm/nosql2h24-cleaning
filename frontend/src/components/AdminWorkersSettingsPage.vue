<script setup lang="ts">
import { ref } from 'vue'
import HeaderList from '../ui/uikit/HeaderList.vue'
import PanelContainer from '../ui/uikit/PanelContainer.vue'
import ActionButton from '../ui/uikit/ActionButton.vue'
import InputTextField from '../ui/uikit/InputTextField.vue'

const workers = ref([]) // TODO DB request
</script>

<template>
  <div class="container">
    <PanelContainer
      width="60%"
    >
      <ActionButton
        text="Создать"
        type="create"
        color="#394cc2"
      ></ActionButton>
      <InputTextField
        class="search-field"
        label="Имя"
        type="first-name-search"
        placeholder="Введите имя"
      ></InputTextField>
      <InputTextField
        class="search-field"
        label="Фамилия"
        type="last-name-search"
        placeholder="Введите фамилию"
      ></InputTextField>
      <ActionButton
        text="Поиск"
        type="create"
        color="#394cc2"
      ></ActionButton>
    </PanelContainer>
    <HeaderList
      title="Исполнители" 
      :items="workers"
      height="100%"
      width="95%"
    >
    </HeaderList>
  </div>
</template>

<style scoped>
.container {
  max-width: 100%;
  max-height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100%;
  width: 80%;
  background-color: white;
  gap: 20px;
  padding-top: 20px;
  padding-bottom: 20px;
}
</style>
