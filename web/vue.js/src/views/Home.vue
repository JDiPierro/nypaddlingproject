<template>
  <div>
    <LandingPage v-if="isAuthenticated === false" />
    <Dashboard v-if="isAuthenticated === true" />
  </div>
</template>

<script lang="ts">
import { mapGetters } from 'vuex';

import Dashboard from '../components/Dashboard.vue';
import LandingPage from '../components/LandingPage.vue';

export default {
  components: {
    Dashboard,
    LandingPage,
  },
  computed: {
    ...mapGetters('user', ['isAuthenticated'])
  }
}
</script>
