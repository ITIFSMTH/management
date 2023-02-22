<template>
  <div class="create-poll">
    <div class="input-box">
      <label class="input-box__label" for="budget">Бюджет</label>
      <input class="input-box__input" placeholder="Бюджет" type="text" id="budget"
        v-model="budget">
    </div>

    <button class="button button-green" 
      v-on:click="this.$store.dispatch('addBudgetPoll')">
      <span class="material-icons">add</span>
      Создать
    </button>
  </div>
</template>

<script>
  export default {
    name: 'create-poll-component',
    computed: {
      budget: {
        get() {return this.$store.getters.newBudgetPoll.budget},
        set(budget) {this.$store.commit('updateBudgetPollBudget', budget)},
      }
    }
  }
</script>

<style scoped>
  
</style>