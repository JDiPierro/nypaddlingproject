<template id="app">
  <v-main>
    <v-container>
      <h1>Claimed Locations</h1>
      <ClaimList
        :claims="initialClaims"
      >
        <h3>
          Visit the <v-btn small color="primary" to="locations"><v-icon left>mdi-map-marker</v-icon> Locations</v-btn> page to claim a location!
        </h3>
      </ClaimList>
      <div v-if="initialClaims.length > 0">
        <v-divider class="my-10"></v-divider>
        <h1>Submitted Claims</h1>
        <ClaimList
          :claims="submittedClaims"
        >
          Scout a location, submit an update to Paddling.com, and then tell us about your awesome work!
        </ClaimList>
      </div>
    </v-container>
  </v-main>
</template>

<script>
  import { default as LocationTable } from '../components/LocationTable.vue'
  import { mapGetters, mapActions } from 'vuex'
  import ClaimedLocationCard from "./ClaimedLocationCard"
  import ClaimList from './ClaimList'

  export default{
    components: {
      ClaimedLocationCard,
      LocationTable,
      ClaimList
    },
    computed: {
      ...mapGetters('locations', ['initialClaims', 'submittedClaims'])
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
