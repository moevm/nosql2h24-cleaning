<script setup lang="ts">
import { RouterView } from 'vue-router'
import { Ref, provide, ref } from 'vue'
import SideBar from '../ui/uikit/SideBar.vue'
import type Button from '../ui/model/button'

const sideBarButtons: Ref<Button[]> = ref([])
const userCard: Ref<{
  name?: string;
  email?: string;
}> = ref({})

function addSideBarButtons(button: Button): void {
  sideBarButtons.value.push(button)
}

function setUserCard(name: string, email: string): void {
  userCard.value = { name, email }
}

// buttons for the sidebar
provide('addSideBarButtons', addSideBarButtons)
provide('sideBarButtons', sideBarButtons)

// user card
provide('setUserCard', setUserCard)
provide('userCard', userCard)
</script>

<template>
  <div class="container">
    <SideBar />
    <RouterView />
  </div>
</template>

<style scoped>
.container {
  display: flex;
  height: 100vh;
}
</style>
