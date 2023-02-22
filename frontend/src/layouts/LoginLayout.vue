<template>
  <div class="login-layout">
    <router-view/>
  </div>
</template>


<script>
  export default {
    name: 'login-layout'
  }
</script>


<style scoped>
  .login-layout {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 100%;
    height: 100vh;
  }
</style>