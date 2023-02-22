<template>
  <div class="page admin-page">
    <!-- <Poll/> -->
    <Operators/>
  </div>
</template>

<script>
import Operators from "@/components/Admin/Operators/Operators.vue";
// import Poll from "@/components/Admin/Poll/Poll.vue";
  export default {
    name: "admin-view",
    components: { Operators }
}
</script>

<style scoped>
</style>