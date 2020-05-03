<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <GetMood />
  </div>
</template>

<script>
// @ is an alias to /src
import GetMood from '@/components/GetMood.vue'

export default {
  name: 'GetMood',
  components: {
    GetMood
  }
}
</script>
