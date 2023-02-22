<template>
  <div class="operator-status" 
    :class="{'ontimeout': onTimeout, 'onshift': onShift && !onTimeout, 'offline': !onShift && !onTimeout}">
    <p class="operator-status__ontimeout" v-if="onTimeout">На перерыве</p>
    <p class="operator-status__onshift" v-else-if="onShift">На смене</p>
    <p class="operator-status__offline" v-else>Не в сети</p>
  </div>
</template>

<script>
  export default {
    name: 'operator-status-component',
    props: {
      onShift: Boolean,
      onTimeout: Boolean,
    },
  }
</script>

<style scoped>
  .operator-status {
    display: inline-block;
    position: relative;
    margin-left: 14px;
  }

  .operator-status::before {
    content: '';
    display: block;
    position: absolute;
    width: 8px;
    height: 8px;
    left: -15px;
    top: 50%;
    transform: translateY(-50%);
    border-radius: 50%;
  }

  .operator-status.offline {
    color: var(--lt__red);
  }

  .operator-status.offline::before {
    background-color: var(--lt__red);
  }

  .operator-status.onshift {
    color: var(--lt__green);
  }

  .operator-status.onshift::before {
    background-color: var(--lt__green);
  }

  .operator-status.ontimeout {
    color: var(--lt__yellow);
  }

  .operator-status.ontimeout::before {
    background-color: var(--lt__yellow);
  }


  /* DARK THEME */


  .dark__theme .operator-status.offline {
    color: var(--dt__red);
  }

  .dark__theme .operator-status.offline::before {
    background-color: var(--dt__red);
  }

  .dark__theme .operator-status.onshift {
    color: var(--dt__green);
  }

  .dark__theme .operator-status.onshift::before {
    background-color: var(--dt__green);
  }

  .dark__theme .operator-status.ontimeout {
    color: var(--dt__yellow);
  }

  .dark__theme .operator-status.ontimeout::before {
    background-color: var(--dt__yellow);
  }
</style>