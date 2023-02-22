import axios from 'axios'
import Vuex from 'vuex'
import VueCookies from 'vue-cookies'
import jwtDecode from "jwt-decode";
import { useToast } from "vue-toastification"
import router from '@/router/index'

let API_URL = process.env.VUE_APP_API_URL
if (process.env.NODE_ENV == "production") {
  API_URL = "http://" + location.host + "/api"
}

const toast = useToast()

export default new Vuex.Store({
  state: {
    /* SITE SETTINGS STATE */
    siteSettings: {
      isNavbarOpen: false,
    },
    /* USER STATE */
    user: {
      jwt: "",
      decodedJwt: {
        login: "",
        role_id: -1,
      },
      theme: {
        id: 1,
        theme: "Светлая",
      },
    },
    /* LOGIN PAGE STATE */
    loginForm: {
      login: '',
      password: ''
    },
    /* ADMIN PAGE STATE */
    budgetPoll: {
      id: 0,
    },
    newBudgetPoll: {
      budget: "",
    },
    operators: [],
    operatorsStatistics: [],
    modals: {
      addOperator: {
        isActive: false,
        login: "",
        password: "",
        telegram: "",
        roleId: 0,
      },
      deleteOperator: {
        isActive: false,
        login: "",
        workerId: 0,
      },
      editOperatorTelegram: {
        isActive: false,
        login: "",
        newTelegram: "",
        operatorId: 0,
      },
      editWorker: {
        isActive: false,
        login: "",
        newLogin: "",
        newRoleId: 0,
        workerId: 0,
      },
      editWorkerPassword: {
        isActive: false,
        login: "",
        newPassword: "",
        workerId: 0,
      }
    },
    /* OPERATOR PAGE STATE */
    myStatistic: [],
  },
  getters: {
    /* SITE SETTINGS GETTERS */
    isNavbarOpen: (state) => state.siteSettings.isNavbarOpen,
    isModalActive: (state) => {
      return Object.entries(state.modals).find(mod => mod[1].isActive) ? true : false
    },
    isThemeDark: (state) => {
      return state.user.theme.id === 2
    },
    /* USER GETTERS */
    jwt: (state) => state.user.jwt,
    decodedJwt: (state) => state.user.decodedJwt,
    isUserAdmin: (state) => state.user.decodedJwt.role_id == 1,
    isUserJuniorOperator: (state) => state.user.decodedJwt.role_id == 2,
    isUserSeniorOperator: (state) => state.user.decodedJwt.role_id == 3,
    isUserOperator: ((state) => state.user.decodedJwt.role_id == 2 || state.user.decodedJwt.role_id == 3),
    isUserAuthenticated: (state) => Date.now() < state.user.decodedJwt.exp * 1000,
    /* LOGIN PAGE GETTERS */
    loginFormLogin: (state) => state.loginForm.login,
    loginFormPassword: (state) => state.loginForm.password,
    /* ADMIN PAGE GETTERS */
    budgetPoll: (state) => state.budgetPoll,
    newBudgetPoll: (state) => state.newBudgetPoll,
    operators: (state) => (roleId = 0) => {
      if (roleId === 0) return state.operators
      return state.operators.filter((operator) => operator.worker.worker_role.id === roleId)
    },
    operatorsStatistics: (state) => {
      return state.operatorsStatistics
    },
    addOperatorModal: (state) => state.modals.addOperator,
    deleteOperatorModal: (state) => state.modals.deleteOperator,
    editOperatorTelegramModal: (state) => state.modals.editOperatorTelegram,
    editWorkerModal: (state) => state.modals.editWorker,
    editWorkerPasswordModal: (state) => state.modals.editWorkerPassword,
    /* OPERATOR PAGE GETTERS */
    myStatistic: (state) => state.myStatistic,
  },
  mutations: {
    /* SITE SETTINGS MUTATIONS */
    switchNavbar: (state) => state.siteSettings.isNavbarOpen = !state.siteSettings.isNavbarOpen,
    /* LOGIN PAGE MUTATIONS */
    updateLoginFormLogin: (state, login) => state.loginForm.login = login,
    updateLoginFormPassword: (state, password) => state.loginForm.password = password,
    updateJWT(state, jwt) {
      const d = new Date()
      d.setTime(d.getTime() + 3 * 24 * 60 * 60 * 1000)
      VueCookies.set("jwt", jwt, d.toUTCString())

      state.user.decodedJwt = jwtDecode(jwt)

      state.loginForm.login = ""
      state.loginForm.password = ""
    },
    updateUser(state) {
      const jwtCookie = VueCookies.get('jwt')
      state.user.jwt = jwtCookie
      state.user.decodedJwt = jwtCookie ? jwtDecode(jwtCookie) : {}
    },
    /* USER MUTATIONS */
    setTheme(state, theme) {
      state.user.theme = theme
      if (theme.id === 2) document.body.classList.add('dark__theme')
      else document.body.classList.remove('dark__theme')
    },
    /* ADMIN PAGE MUTATIONS */
    setBudgetPoll: (state, poll) => state.budgetPoll = poll, 
    setOperators: (state, operators) => state.operators = operators,
    setOperatorsStatistics: (state, statistics) => state.operatorsStatistics = statistics,
    updateBudgetPollBudget: (state, budget) => state.newBudgetPoll.budget = budget,
    switchAddOperatorModal(state, modalData) {
      if (modalData) {
        state.modals.addOperator.isActive = true
        return
      }

      state.modals.addOperator = {
        isActive: false,
        login: "",
        password: "",
        telegram: "",
        roleId: 0,
      }
    },
    switchDeleteOperatorModal(state, modalData) {
      if (modalData) {
        state.modals.deleteOperator = {
          isActive: true,
          login: modalData.login,
          workerId: modalData.workerId,
        }
        return
      }

      state.modals.deleteOperator = {
        isActive: false,
        login: "",
        workerId: 0,
      }
    },
    switchEditOperatorTelegramModal(state, modalData) {
      if (modalData) {
        state.modals.editOperatorTelegram = {
          isActive: true,
          login: modalData.login,
          newTelegram: "",
          operatorId: modalData.operatorId,
        }
        return
      }

      state.modals.editOperatorTelegram = {
        isActive: false,
        login: "",
        newTelegram: "",
        workerId: 0,
      }
    },
    switchEditWorkerModal(state, modalData) {
      if (modalData) {
        state.modals.editWorker = {
          isActive: true,
          login: modalData.login,
          newLogin: "",
          newRoleId: modalData.roleId,
          workerId: modalData.workerId,
        }
        return
      }

      state.modals.editWorker = {
        isActive: false,
        login: "",
        newLogin: "",
        newRoleId: 0,
        workerId: 0,
      }
    },
    switchEditWorkerPasswordModal(state, modalData) {
      if (modalData) {
        state.modals.editWorkerPassword = {
          isActive: true,
          login: modalData.login,
          newPassword: "",
          workerId: modalData.workerId,
        }
        return
      }

      state.modals.editWorkerPassword = {
        isActive: false,
        login: "",
        newPassword: "",
        workerId: 0,
      }
    },
    updateAddOperatorModalLogin: (state, login) => state.modals.addOperator.login = login,
    updateAddOperatorModalPassword: (state, password) => state.modals.addOperator.password = password,
    updateAddOperatorModalTelegram: (state, telegram) => state.modals.addOperator.telegram = telegram,
    updateAddOperatorModalRole: (state, roleId) => state.modals.addOperator.roleId = roleId,
    updateEditOperatorTelegramModalTelegram: (state, telegram) => state.modals.editOperatorTelegram.newTelegram = telegram,
    updateEditWorkerModalLogin: (state, login) => state.modals.editWorker.newLogin = login,
    updateEditWorkerModalRole: (state, roleId) => state.modals.editWorker.newRoleId = roleId,
    updateEditWorkerPasswordModalPassword: (state, password) => state.modals.editWorkerPassword.newPassword = password,
    addOperator(state, operator) {
      state.operators.push(operator)
    },
    deleteOperator(state, workerId) {
      state.operators = state.operators.filter((operator) => operator.worker.id != workerId)
    },
    updateOperatorTelegram(state, data) {
      state.operators = state.operators.map((operator) => {
        if (operator.id === data.operatorId) operator.telegram = data.newTelegram 
        return operator
      })
    },
    updateWorker(state, data) {
      state.operators = state.operators.map((operator) => {
        if (operator.worker.id === data.workerId) {
          if (data.newLogin != "") operator.worker.login = data.newLogin
          operator.worker.worker_role.id = data.newRoleId
        }
        return operator
      })
    },
    /* OPERATOR PAGE MUTATIONS */
    setMyStatistic: (state, statistic) => state.myStatistic = statistic
  },
  actions: {
    /* LOGIN ACTIONS */
    async login(ctx) {
      const { data } = await axios.post(API_URL + '/auth/login', {
        login: ctx.state.loginForm.login,
        password: ctx.state.loginForm.password
      }, { validateStatus: () => true })

      if (data.error) {
        toast.error("Неверные данные")
        return false;
      }
      toast.success("Успешно")

      ctx.commit('updateJWT', data.data.token)

      router.push({name: 'admin-route'})
      return true;
    },
    /* USER ACTIONS */
    async getTheme(ctx) {
      const { data } = await axios.get(API_URL + '/user/theme', {
        headers: {
          'Authorization': `Bearer ${VueCookies.get('jwt')}`
        }
      }, { validateStatus: () => true })

      if (data.error) {
        toast.error("Ошибка при получении темы")
        return false
      }

      ctx.commit('setTheme', data.data.theme)
      return true
    },
    async editTheme(ctx) {
      const { data } = await axios.patch(API_URL + '/user/theme', {
        theme_id: ctx.state.user.theme.id === 1 ? 2 : 1
      }, {
        headers: {
          'Authorization': `Bearer ${VueCookies.get('jwt')}`
        }, 
        validateStatus: () => true 
      })

      if (data.error) {
        toast.error("Ошибка при изменении темы")
        return false
      }

      ctx.commit('setTheme', data.data.theme)
      return true
    },
    /* ADMIN PAGE ACTIONS */
    async getBudgetPoll(ctx) {
      const { data } = await axios.get(API_URL + '/admin/poll/budget', {
        headers: {
          'Authorization': `Bearer ${VueCookies.get('jwt')}`
        },
        validateStatus: false
      })

      if (data.error === "not_exists") {
        return false
      }

      if (data.error) {
        toast.error("Ошибка при получении голосования")
        return false
      }

      ctx.commit('setBudgetPoll', data.data.poll)
      return true
    },
    async getOperators(ctx) {
      const { data } = await axios.get(API_URL + '/admin/operators', {
        headers: {
          'Authorization': `Bearer ${VueCookies.get('jwt')}`
        },
        validateStatus: () => true
      })

      if (data.error) {
        toast.error("Ошибка при получении операторов")
        return false
      }

      ctx.commit('setOperators', data.data.operators)
      return true
    },
    async getOperatorsStatistics(ctx) {
      const { data } = await axios.get(API_URL + '/admin/operators/statistic', {
        headers: {
          'Authorization': `Bearer ${VueCookies.get('jwt')}`
        },
        validateStatus: () => true
      })

      if (data.error) {
        toast.error("Ошибка при получении статистики операторов")
        return false
      }

      ctx.commit('setOperatorsStatistics', data.data.statistic)
      return true
    },
    async addBudgetPoll(ctx) {
      const { data } = await axios.post(API_URL + '/admin/poll/budget', {
        "budget": Number(ctx.state.newBudgetPoll.budget),
      }, {
        headers: {
          'Authorization': `Bearer ${VueCookies.get('jwt')}`
        },
        validateStatus: () => true
      })

      if (data.error) {
        toast.error("Ошибка при создании голосвания (Проверьте поля ввода)")
        return false
      }

      ctx.commit('setBudgetPoll', data.data.poll)
      return true
    },
    async addOperator(ctx) {
      const { data } = await axios.post(API_URL + '/admin/operator', {
        "login": ctx.state.modals.addOperator.login,
        "password": ctx.state.modals.addOperator.password,
        "telegram": ctx.state.modals.addOperator.telegram,
        "role_id": ctx.state.modals.addOperator.roleId,
      }, {
        headers: {
          'Authorization': `Bearer ${VueCookies.get('jwt')}`
        },
        validateStatus: () => true
      })

      if (data.error) {
        toast.error("Ошибка при создании оператора (Проверьте поля ввода)")
        return false
      }
      toast.success("Оператор успешно добавлен")

      ctx.commit('addOperator', data.data.operator)
      return true
    },
    async deleteOperator(ctx) {
      const { data } = await axios.delete(API_URL + '/admin/operator', {
        data: {
          "worker_id": ctx.state.modals.deleteOperator.workerId,
        },
        headers: {
          'Authorization': `Bearer ${VueCookies.get('jwt')}`
        },
        validateStatus: () => true
      })

      if (data.error) {
        toast.error("Ошибка при удалении оператора")
        return false
      }
      toast.success("Оператор успешно удалён")

      ctx.commit('deleteOperator', ctx.state.modals.deleteOperator.workerId)
      return true
    },
    async editOperatorTelegram(ctx) {
      const { data } = await axios.patch(API_URL + '/admin/operator', {
        "operator_id": ctx.state.modals.editOperatorTelegram.operatorId,
        "telegram": ctx.state.modals.editOperatorTelegram.newTelegram
      }, {
        headers: {
          'Authorization': `Bearer ${VueCookies.get('jwt')}`
        },
        validateStatus: () => true
      })

      if (data.error) {
        toast.error("Ошибка при изменении Telegram оператора (Проверьте поля ввода)")
        return false
      }
      toast.success("Telegram оператора успешно изменён")

      ctx.commit('updateOperatorTelegram', {
        operatorId: ctx.state.modals.editOperatorTelegram.operatorId,
        newTelegram: ctx.state.modals.editOperatorTelegram.newTelegram,
      })
      return true
    },
    async editWorker(ctx) {
      const { data } = await axios.patch(API_URL + '/admin/worker', {
        "worker_id": ctx.state.modals.editWorker.workerId,
        "role_id": ctx.state.modals.editWorker.newRoleId, 
        "login": ctx.state.modals.editWorker.newLogin,
      }, {
        headers: {
          'Authorization': `Bearer ${VueCookies.get('jwt')}`
        },
        validateStatus: () => true
      })

      if (data.error) {
        toast.error("Ошибка при изменении оператора (Проверьте поля ввода)")
        return false
      }
      toast.success("Данные оператора успешно изменены")

      ctx.commit('updateWorker', {
        workerId: ctx.state.modals.editWorker.workerId,
        newLogin: ctx.state.modals.editWorker.newLogin,
        newRoleId: ctx.state.modals.editWorker.newRoleId,
      })
      return true
    },
    async editWorkerPassword(ctx) {
      const { data } = await axios.patch(API_URL + '/admin/worker/password', {
        "worker_id": ctx.state.modals.editWorkerPassword.workerId,
        "password": ctx.state.modals.editWorkerPassword.newPassword
      }, {
        headers: {
          'Authorization': `Bearer ${VueCookies.get('jwt')}`
        },
        validateStatus: () => true
      })

      if (data.error) {
        toast.error("Ошибка при изменении пароля оператора (Проверьте поля ввода)")
        return false
      }
      toast.success("Пароль оператора успешно изменён")
      
      return true
    },
    /* OPERATOR PAGE ACTIONS */
    async getMyStatistic(ctx) {
      const { data } = await axios.get(API_URL + '/operator/statistic', {
        headers: {
          'Authorization': `Bearer ${VueCookies.get('jwt')}`
        },
        validateStatus: () => true
      })

      if (data.error) {
        toast.error("Ошибка при получении статистики")
        return false
      }

      ctx.commit('setMyStatistic', data.data.statistic)
      return true
    }
  }
})