<template>
  <div class="poll-status">
    <p class="poll-status__text">Начало: 
      <span>{{formatDate(this.$store.getters.budgetPoll.start_date)}}</span>
    </p>
    <p class="poll-status__text">Конец: 
      <span>{{formatDate(new Date(this.$store.getters.budgetPoll.start_date).addDays(1))}}</span>
    </p>
    <p class="poll-status__text">Бюджет для распределения: 
      <span>{{this.$store.getters.budgetPoll.budget}}₽</span>
    </p>
  </div>
</template>

<script>
  import moment from "moment"

  Date.prototype.addDays = function(days) {
    var date = new Date(this.valueOf());
    date.setDate(date.getDate() + days);
    return date;
  }

  export default {
    name: 'poll-status-component',
    methods: {
      formatDate(date) {
        return moment(date).format("DD.MM.YYYY HH:mm")
      }
    }
  }
</script>

<style scoped>
  .poll-status__text {
    font-size: 18px;
    font-weight: 500;
  }

  .poll-status__text > span {
    font-weight: 600;
    color: var(--lt__main);
  }


  /* DARK THEME */


  .dark__theme .poll-status__text {
    color: var(--dt__header);
  }

  .dark__theme .poll-status__text > span {
    color: var(--dt__text-lighter);
  }
</style>