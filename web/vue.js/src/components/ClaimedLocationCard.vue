<template>
  <v-card class="fill-height d-flex flex-column pa-2">
    <v-btn
      v-if="canRelease"
      dark
      fab
      absolute
      top
      right
      x-small
      class="mt-7"
      color="red"
      @click="release({ location_id: claim.location_id })"
    >
      <v-icon>mdi-close</v-icon>
    </v-btn>
    <v-card-title class="mr-6">
      {{claim.location.name}}
    </v-card-title>
    <v-card-subtitle>
      {{ claim.location.county }} county
    </v-card-subtitle>
    <v-spacer></v-spacer>
    <v-card-actions>
      <v-btn color="primary" :href="claim.location.link" target="_blank">Paddling.com</v-btn>
      <SubmitClaimButton :claim="claim"></SubmitClaimButton>
    </v-card-actions>
  </v-card>
</template>

<script>
  import { mapActions } from 'vuex';
  import { default as SubmitClaimButton } from './SubmitClaimButton'

  export default {
    props: ["claim"],
    name: "ClaimedLocationCard",
    components: {
      SubmitClaimButton
    },
    methods: {
      ...mapActions('locations', ['release']),
    },
    computed: {
      canRelease() {
        return this.claim.status === "initial"
      }
    }
  }
</script>

<style scoped>
  .v-card__text, .v-card__title {
    word-break: normal;
  }
</style>
