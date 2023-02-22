<template>
  <div class="modal editWorkerPasswordModal"
    :class="{'modal-open': this.$store.getters.editWorkerPasswordModal.isActive}">
    <h2 class="modal__header">Изменение пароля оператора {{this.$store.getters.editWorkerPasswordModal.login}}</h2>

    <div class="input-box">
      <label class="input-box__label" for="password">Пароль</label>
      <input class="input-box__input" placeholder="Пароль" type="password" id="password"
        v-model="password">
    </div>

    <div class="modal-footer">
      <button class="modal__button modal__button-cancel"
        v-on:click="closeEditWorkerPassword()">Отмена</button>
      <button class="modal__button modal__button-agree"
        v-on:click="confirmEditWorkerPassword()">Подтвердить</button>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'edit-worker-password-modal',
    computed: {
      password: {
        get() {return this.$store.getters.editWorkerPasswordModal.newPassword},
        set(password) {this.$store.commit('updateEditWorkerPasswordModalPassword', password)},
      },
    },
    methods: {
      closeEditWorkerPassword() {
        this.$store.commit('switchEditWorkerPasswordModal', false)
      },
      async confirmEditWorkerPassword() {
        await this.$store.dispatch('editWorkerPassword')
        this.$store.commit('switchEditWorkerPasswordModal', false)
      }
    }
  }
</script>