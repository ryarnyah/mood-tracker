<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <CreateMood />
  </div>
</template>

<script>
// @ is an alias to /src
import CreateMood from '@/components/CreateMood.vue'

export default {
  name: 'Home',
  components: {
    CreateMood
  }
}
</script>
