<template>
  <div class="panel-layout">
    <Navbar/>
    <div class="panel-layout__content">
      <router-view/>
    </div>
  </div>
</template>


<script>
  import Navbar from '@/components/PanelLayout/Navbar.vue';

  export default {
    name: 'panel-layout',
    components: {
      Navbar
    }
  }
</script>


<style scoped>
  .panel-layout {
    display: flex;
  }

  .panel-layout__content {
    padding: 20px;
    width: 100%;
    box-sizing: border-box;
  }


  /* MEDIA */


  @media screen and (max-width: 800px) {
    .panel-layout__content {
      padding: 10px;
      margin-bottom: 74px;
    }
  }
</style>