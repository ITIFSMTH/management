<template>
  <div>
    <ul class="operator-actions">
      <li class="operator-action operator-action__telegram material-icons"
        v-on:click="this.$store.commit('switchEditOperatorTelegramModal', {
          login: operator.worker.login,
          operatorId: operator.id,
        })">flight</li>
      <li class="operator-action operator-action__worker material-icons"
        v-on:click="this.$store.commit('switchEditWorkerModal', {
          login: operator.worker.login,
          roleId: operator.worker.worker_role.id,
          workerId: operator.worker.id,
        })">drive_file_rename_outline</li>
      <li class="operator-action operator-action__password material-icons"
        v-on:click="this.$store.commit('switchEditWorkerPasswordModal', {
          login: operator.worker.login,
          workerId: operator.worker.id,
        })">lock</li>
      <li class="operator-action operator-action__delete material-icons"
        v-on:click="this.$store.commit('switchDeleteOperatorModal', {
          login: operator.worker.login,
          workerId: operator.worker.id,
        })">delete</li>
    </ul>

    <!-- MODALS -->
    <DeleteOperatorModal/>
    <EditOperatorTelegramModal/>
    <EditWorkerModal/>
    <EditWorkerPasswordModal/>
  </div>
</template>

<script>
  import DeleteOperatorModal from "@/components/Modals/DeleteOperatorModal.vue";
  import EditOperatorTelegramModal from "@/components/Modals/EditOperatorTelegramModal.vue";
  import EditWorkerModal from "@/components/Modals/EditWorkerModal.vue";
  import EditWorkerPasswordModal from "@/components/Modals/EditWorkerPasswordModal.vue";
  
  export default {
    name: "operator-actions-component",
    props: ['operator'],
    components: { DeleteOperatorModal, EditOperatorTelegramModal, EditWorkerModal, EditWorkerPasswordModal }
}
</script>

<style scoped>
  .operator-action {
    font-size: 20px;
    padding-left: 5px;
    cursor: pointer;
    transition: .2s;
  }

  .operator-action:hover {
    opacity: .7;
  }

  .operator-action__delete {
    color: var(--lt__red);
  }

  .operator-action__worker,
  .operator-action__password,
  .operator-action__telegram {
    color: var(--lt__yellow);
  }


  /* DARK THEME */


  .dark__theme .operator-action__delete {
    color: var(--dt__red);
  }

  .dark__theme .operator-action__login,
  .dark__theme .operator-action__password,
  .dark__theme .operator-action__telegram {
    color: var(--dt__yellow);
  }
</style>