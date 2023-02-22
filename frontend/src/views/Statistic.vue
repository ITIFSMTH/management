<template>
  <div class="page statistic-page">
   <Statistic v-for="monthOperatorsStatistic in this.$store.getters.operatorsStatistics"
    v-bind:key="monthOperatorsStatistic.date"
    v-bind:statistic="monthOperatorsStatistic"/>
  </div>
</template>

<script>
  import Statistic from "@/components/Admin/Statistic/Statistic.vue";
  export default {
    name: "statistic-view",
    components: { Statistic },
    mounted() {
      this.$store.dispatch("getOperatorsStatistics")
    },
}
</script>

<style scoped>
</style>