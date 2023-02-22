<template>
  <div class="modal deleteOperatorModal"
    :class="{'modal-open': this.$store.getters.deleteOperatorModal.isActive}">
    <h2 class="modal__header">Удаление оператора {{this.$store.getters.deleteOperatorModal.login}}</h2>

    <p class="modal__text">Вы уверены? Это действие нельзя отменить!</p>

    <div class="modal-footer">
      <button class="modal__button modal__button-agree"
        v-on:click="closeDeleteOperator()">Отмена</button>
      <button class="modal__button modal__button-cancel"
        v-on:click="confirmDeleteOperator()">Подтвердить</button>
    </div>
  </div>
</template>

<script>  
  export default {
    name: 'delete-operator-modal',
    methods: {
      closeDeleteOperator() {
        this.$store.commit('switchDeleteOperatorModal', false)
      },
      async confirmDeleteOperator() {
        await this.$store.dispatch('deleteOperator')
        this.$store.commit('switchDeleteOperatorModal', false)
      }
    }
  }
</script>