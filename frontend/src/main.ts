// vuetify
import 'vuetify/lib/styles/main.css'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'

// pinia
import { createPinia } from 'pinia'

// app
import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './routing/router';

const vuetify = createVuetify({
    components,
    directives,
})
const pinia = createPinia()
const app = createApp(App)

app.use(vuetify)
app.use(pinia)
app.use(router)

app.mount('#app')
