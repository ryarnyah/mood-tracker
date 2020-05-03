<template>
  <div class="home">
      <img alt="Vue logo" src="../assets/logo.png">
      <AddEntry :moodId="$route.params.moodId" :entryAccessCode="$route.params.entryAccessCode" />
  </div>
</template>

<script>
// @ is an alias to /src
import AddEntry from '@/components/AddEntry.vue'

export default {
  name: 'Entry',
  components: {
    AddEntry
  }
}
</script>
