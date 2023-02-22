<template>
  <div class="page-block">
    <h2 class="page-block__header">Операторы</h2>

    <button class="button button-green"
      v-on:click="this.$store.commit('switchAddOperatorModal', true)">
      <div class="button-icon material-icons">add</div>
      Добавить оператора
    </button>

    <table class="table__table">
      <thead class="table__thead">
        <tr class="table__tr">
          <th class="table__th">ID</th>
          <th class="table__th">Оператор</th>
          <th class="table__th">Telegram</th>
          <th class="table__th">Активность</th>
          <th class="table__th">Действия</th>
        </tr>
      </thead>
      <tbody>
        <tr class="table__tr" 
          v-for="operator in this.$store.getters.operators(Roles.JuniorOperatorRole.ID)"
          v-bind:key="operator.id">
          <td class="table__td">{{operator.id}}</td>
          <td class="table__td">{{operator.worker.login}}</td>
          <td class="table__td">@{{operator.telegram}}</td>
          <td class="table__td">
            <OperatorStatus v-bind:onShift="operator.on_shift" v-bind:onTimeout="operator.on_timeout"/>
          </td>
          <td class="table__td">
            <OperatorActions v-bind:operator="operator"/>
          </td>
        </tr>

        <tr><td class="table__header" colspan="5">Старшие операторы</td></tr>
        <tr class="table__tr" 
          v-for="operator in this.$store.getters.operators(Roles.SeniorOperatorRole.ID)"
          v-bind:key="operator.id">
          <td class="table__td">{{operator.id}}</td>
          <td class="table__td">{{operator.worker.login}}</td>
          <td class="table__td">@{{operator.telegram}}</td>
          <td class="table__td">
            <OperatorStatus v-bind:onShift="operator.on_shift" v-bind:onTimeout="operator.on_timeout"/>
          </td>
          <td class="table__td">
            <OperatorActions v-bind:operator="operator"/>
          </td>
        </tr>
      </tbody>
    </table>

    <AddOperatorModal/>
  </div>
</template>

<script>
  import { Roles } from "@/shared"
  import OperatorStatus from "../../Operator/OperatorStatus.vue"
  import OperatorActions from "./OperatorActions.vue";
  import AddOperatorModal from "@/components/Modals/AddOperatorModal.vue";

  export default {
    name: "operators-component",
    data() {
      return {
        Roles
      }
    },
    mounted() {
        this.$store.dispatch("getOperators");
    },
    components: { OperatorStatus, OperatorActions, AddOperatorModal }
  }
</script>

<style scoped>
  .table__table {
    margin-top: 15px;
  }
</style>