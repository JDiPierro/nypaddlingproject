<template>
  <v-card>
    <v-card-title>
      {{claim.location.name}}
    </v-card-title>
    <v-card-subtitle>
      {{ claim.location.county }} county
    </v-card-subtitle>
    <v-card-actions>
      <v-btn color="primary" :href="claim.location.link" target="_blank">Paddling.com</v-btn>

      <v-btn color="success">Submit Change</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
  export default {
    props: ["claim"],
    name: "ClaimedLocationCard"
  }
</script>

<style scoped>

</style>
