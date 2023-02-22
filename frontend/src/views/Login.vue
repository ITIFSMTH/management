<template>
  <div class="container">
    <LoginForm/>
  </div>
</template>

<script>
  import LoginForm from '@/components/Login/LoginForm'

  export default {
    name: 'login-view',
    components: {
      LoginForm
    }
  }
</script>

<style scoped>
  .container {
    display: flex;
    justify-content: center;
  }
</style>