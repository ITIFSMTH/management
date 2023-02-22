<template>
  <div class="modal addOperatorModal"
    :class="{'modal-open': this.$store.getters.addOperatorModal.isActive}">
    <h2 class="modal__header">Добавление оператора</h2>

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

    <div class="inputs__wrapper">
      <div class="input-box">
        <label class="input-box__label" for="telegram">Telegram</label>
        <input class="input-box__input" placeholder="Telegram" type="text" id="telegram"
          v-model="telegram">
      </div>

      <div class="input-box role-box">
        <label class="input-box__label" for="role">Роль</label>
        <select class="input-box__input" id="role"
          v-model="role">
          <option value="0" disabled>Не выбрано</option>
          <option 
            v-for="role in [Roles.JuniorOperatorRole, Roles.SeniorOperatorRole]"
            v-bind:key="role.ID" 
            :value="role.ID">{{role.Role}}</option>
        </select>
      </div>
      
    </div>

    <div class="modal-footer">
      <button class="modal__button modal__button-cancel"
        v-on:click="closeAddOperator()">Отмена</button>
      <button class="modal__button modal__button-agree"
        v-on:click="confirmAddOperator()">Подтвердить</button>
    </div>
  </div>
</template>

<script>
  import { Roles } from "@/shared"
  
  export default {
    name: 'add-operator-modal',
    data() {
      return {
        Roles
      }
    },
    computed: {
      login: {
        get() {return this.$store.getters.addOperatorModal.login},
        set(login) {this.$store.commit('updateAddOperatorModalLogin', login)},
      },
      password: {
        get() {return this.$store.getters.addOperatorModal.password},
        set(password) {this.$store.commit('updateAddOperatorModalPassword', password)},
      }, 
      telegram: {
        get() {return this.$store.getters.addOperatorModal.telegram},
        set(telegram) {this.$store.commit('updateAddOperatorModalTelegram', telegram)},
      }, 
      role: {
        get() {return this.$store.getters.addOperatorModal.roleId},
        set(roleId) {this.$store.commit('updateAddOperatorModalRole', roleId)}
      }
    },
    methods: {
      closeAddOperator() {
        this.$store.commit('switchAddOperatorModal', false)
      },
      async confirmAddOperator() {
        await this.$store.dispatch('addOperator')
        this.$store.commit('switchAddOperatorModal', false)
      }
    }
  }
</script>

<style scoped>
  .inputs__wrapper {
    display: flex;
  }

  .role-box {
    margin-left: 10px;
  }
</style>