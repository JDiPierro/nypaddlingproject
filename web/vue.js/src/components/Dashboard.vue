<template id="app">
  <v-main>
    <v-container>
      <v-alert type="warning" prominent>
        <h2>Under Construction!</h2>
        There's not much here right now but check back later to find out how you can help!
      </v-alert>
      <h1>Claimed Locations</h1>
      <v-row v-if="user_claims.length > 0">
        <v-col
          xs="12"
          sm="6"
          md="4"
          lg="4"
          xl="3"
          v-for="claim in user_claims">
            <ClaimedLocationCard
              :claim="claim"
            ></ClaimedLocationCard>
        </v-col>
      </v-row>
      <v-sheet v-if="user_claims.length == 0">
        <h3>
          Visit the <v-btn small color="primary" to="locations"><v-icon left>mdi-map-marker</v-icon> Locations</v-btn> page to claim a location!
        </h3>

      </v-sheet>
    </v-container>
  </v-main>
</template>

<script>
  import { default as LocationTable } from '../components/LocationTable.vue'
  import { mapState, mapActions } from 'vuex'
  import ClaimedLocationCard from "./ClaimedLocationCard"

  export default{
    components: {
      ClaimedLocationCard,
      LocationTable,
    },
    computed: {
      ...mapState('locations', ['user_claims'])
    },
    methods: {
      ...mapActions('locations', ['loadClaims']),
    },
    beforeMount () {
      this.loadClaims()
    }
  }
</script>

<style scoped lang="scss">
</style>
