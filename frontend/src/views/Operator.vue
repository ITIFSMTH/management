<template>
  <div class="page operator-page">
   <div class="page-block">
    <h2 class="page-block__header welcome__header">Добро пожаловать, {{this.$store.getters.decodedJwt.login}}!</h2>
   </div>
   <Statistic v-for="monthOperatorsStatistic in this.$store.getters.myStatistic"
      v-bind:key="monthOperatorsStatistic.date"
      v-bind:statistic="monthOperatorsStatistic"/>
  </div>
</template>

<script>
  import Statistic from "@/components/Admin/Statistic/Statistic.vue"

  export default {
    name: 'operator-view',
    components: { Statistic },
    mounted() {
      this.$store.dispatch('getMyStatistic') 
    }
  }
</script>

<style scoped>
 .welcome__header {
  padding-bottom: 0;
 }
</style>