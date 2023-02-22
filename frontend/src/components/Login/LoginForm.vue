<template>
  <form class="login-form"
    @submit.prevent="onLogin">
    <h1 class="login-form__header">Авторизация</h1>
    
    <div class="input-box">
      <label class="input-box__label" for="login">Логин</label>
      <input class="input-box__input" placeholder="Логин" type="text" id="login"
        v-model="login">
    </div>

    <div class="input-box">
      <label class="input-box__label" for="password">Пароль</label>
      <input class="input-box__input" placeholder="Пароль" type="password" id="password"
        v-model="password">
    </div>
    
    <button class="button" type="submit">Войти</button>
  </form>
</template>

<script>
  export default {
    name: 'login-form-component',
    methods: {
      async onLogin() {
        if (this.$store.getters.loginFormLogin === "" || this.$store.getters.loginFormPassword === "") return
        await this.$store.dispatch("login")
      }
    },
    computed: {
      login: {
        get() { return this.$store.getters.loginFormLogin },
        set(login) { return this.$store.commit("updateLoginFormLogin", login) } 
      },
      password: {
        get() { return this.$store.getters.loginFormPassword },
        set(password) { this.$store.commit("updateLoginFormPassword", password) }
      }
    }
  }
</script>

<style scoped>
  .login-form {
    width: 400px;
    padding: 20px;
    border-radius: 5px;
    background-color: var(--white);
  }

  .login-form__header {
    padding-bottom: 35px;
    text-align: center;
    color: var(--lt__main);
  }

  .login-form .button {
    margin-top: 20px;
  }
</style>