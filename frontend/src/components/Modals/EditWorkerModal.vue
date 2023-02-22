<template>
  <div class="modal editWorkerModal"
    :class="{'modal-open': this.$store.getters.editWorkerModal.isActive}">
    <h2 class="modal__header">Изменение оператора {{this.$store.getters.editWorkerModal.login}}</h2>

    <div class="inputs__wrapper">
      <div class="input-box">
        <label class="input-box__label" for="login">Логин</label>
        <input class="input-box__input" placeholder="Логин" type="text" id="login"
          v-model="login">
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
        v-on:click="closeEditWorker()">Отмена</button>
      <button class="modal__button modal__button-agree"
        v-on:click="confirmEditWorker()">Подтвердить</button>
    </div>
  </div>
</template>

<script>
  import { Roles } from "@/shared"

  export default {
    name: 'edit-worker-modal',
    data() {
      return {
        Roles
      }
    },
    computed: {
      login: {
        get() {return this.$store.getters.editWorkerModal.newLogin},
        set(login) {this.$store.commit('updateEditWorkerModalLogin', login)},
      },
      role: {
        get() {return this.$store.getters.editWorkerModal.newRoleId},
        set(roleId) {this.$store.commit('updateEditWorkerModalRole', roleId)}
      }
    },
    methods: {
      closeEditWorker() {
        this.$store.commit('switchEditWorkerModal', false)
      },
      async confirmEditWorker() {
        await this.$store.dispatch('editWorker')
        this.$store.commit('switchEditWorkerModal', false)
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