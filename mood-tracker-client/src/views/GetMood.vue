<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <GetMood :moodId="$route.params.moodId" :moodAccessCode="$route.params.moodAccessCode" />
  </div>
</template>

<script>
// @ is an alias to /src
import GetMood from '@/components/GetMood.vue'

export default {
  name: 'Mood',
  components: {
    GetMood
  }
}
</script>
