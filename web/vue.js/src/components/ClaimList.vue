<template>
  <v-row v-if="claims.length > 0">
    <v-col
      xs="12"
      sm="6"
      md="4"
      lg="4"
      xl="3"
      v-for="claim in claims"
      :key="claim._id"
    >
      <ClaimedLocationCard
        :claim="claim"
        :key="claim._id"
      ></ClaimedLocationCard>
    </v-col>
  </v-row>
  <v-sheet v-else-if="claims.length == 0" class="mt-4">
    <slot></slot>
  </v-sheet>
</template>

<script>
  import { default as ClaimedLocationCard } from './ClaimedLocationCard'

  export default {
    name: "ClaimList",
    components: {
      ClaimedLocationCard
    },
    props: ['claims']
  }
</script>

<style scoped>

</style>
