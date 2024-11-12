import { defineStore } from 'pinia'
import { ref, Ref } from 'vue'
import { User } from '../api/models/user'

export const useUserStore = defineStore('userStore', () => {
  const user: Ref<User | undefined> = ref(undefined)
  
  function setUser(newUser: User) {
    console.log(newUser)
    user.value = newUser
  }

  function getUser() {
    return user.value
  }

  return {
    user,
    setUser,
    getUser
  }
})