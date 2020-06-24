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
  import { mapActions } from 'vuex';

  export default {
    name: "ClaimButton",
    props: ['location_id', 'claims'],
    methods: {
      ...mapActions('locations',['claim']),
    }
  }
</script>

<style scoped>

</style>
