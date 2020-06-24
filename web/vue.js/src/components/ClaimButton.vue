<template>
  <v-tooltip
    top
    :disabled="claims.length === 0"
  >
    <template v-slot:activator="{ on, attrs }">
      <v-badge
        color="orange"
        :content="claims.length"
        :value="claims.length > 0"
      >
        <v-btn
          color="success"
          :disabled="claimedByActiveUser()"
          @click="claim({ location_id })"
          v-bind="attrs"
          v-on="on"
        >
          Claim
        </v-btn>
      </v-badge>
    </template>
    <span>{{claims.length}} pending claims</span>
  </v-tooltip>
</template>

<script>
  import { mapActions, mapState } from 'vuex';

  export default {
    name: "ClaimButton",
    props: ['location_id', 'claims'],
    methods: {
      ...mapActions('locations', ['claim']),
      claimedByActiveUser() {
        for(const claim of this.claims) {
          if(claim['user_id'] === this.user['_id']) {
            return true
          }
        }
        return false
      }
    },
    computed: {
      ...mapState('user', ['user'])
    }
  }
</script>

<style scoped>

</style>
