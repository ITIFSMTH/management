<template>
  <div class="modal editOperatorTelegramModal"
    :class="{'modal-open': this.$store.getters.editOperatorTelegramModal.isActive}">
    <h2 class="modal__header">Изменение Telegram оператора {{this.$store.getters.editOperatorTelegramModal.login}}</h2>

    <div class="input-box">
      <label class="input-box__label" for="telegram">Telegram</label>
      <input class="input-box__input" placeholder="Telegram" type="text" id="telegram"
        v-model="telegram">
    </div>

    <div class="modal-footer">
      <button class="modal__button modal__button-cancel"
        v-on:click="closeEditOperatorTelegram()">Отмена</button>
      <button class="modal__button modal__button-agree"
        v-on:click="confirmEditOperatorTelegram()">Подтвердить</button>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'edit-operator-telegram-modal',
    computed: {
      telegram: {
        get() {return this.$store.getters.editOperatorTelegramModal.newTelegram},
        set(telegram) {this.$store.commit('updateEditOperatorTelegramModalTelegram', telegram)},
      },
    },
    methods: {
      closeEditOperatorTelegram() {
        this.$store.commit('switchEditOperatorTelegramModal', false)
      },
      async confirmEditOperatorTelegram() {
        await this.$store.dispatch('editOperatorTelegram')
        this.$store.commit('switchEditOperatorTelegramModal', false)
      }
    }
  }
</script>