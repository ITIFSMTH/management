<template>
  <div class="page-block">
    <h2 class="page-block__header">
      <template v-if="this.$store.getters.budgetPoll.id === 0">Создать голосвание за бюджет</template>
      <template v-else>Голосование идёт</template>
    </h2>

    <CreatePoll v-if="this.$store.getters.budgetPoll.id === 0"/>
    <PollStatus v-else/>
  </div>
</template>

<script>
  import CreatePoll from "./CreatePoll.vue"
  import PollStatus from "./PollStatus.vue"

  export default {
    name: "poll-component",
    mounted() {
        this.$store.dispatch("getBudgetPoll");
    },
    components: { CreatePoll, PollStatus }
}
</script>

<style>

</style>