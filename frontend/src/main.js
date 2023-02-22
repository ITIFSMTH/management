import { createApp } from 'vue'
import Store from '@/store/index'
import Router from '@/router/index'
import Toast from "vue-toastification"
import App from './App.vue'
import "vue-toastification/dist/index.css"
import moment from 'moment'
import momentTz from 'moment-timezone'
moment.locale("RU")
momentTz.tz.setDefault('Europe/Moscow')

createApp(App)
.use(Store)
.use(Router)
.use(Toast, {
    timeout: 2 * 1000
})
.mount('#app')