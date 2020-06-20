<template>
  <v-app>
    <Sidebar v-if="isAuthenticated === true" v-model="drawer" />
    <Header v-model="drawer" />
      <v-container fluid>
        <router-view></router-view>
      </v-container>
  </v-app>
</template>

<script>
import { mapGetters } from 'vuex';

import Header from './components/Header.vue';
import Sidebar from './components/Sidebar.vue';

export default {
  components: {
    Header,
    Sidebar
  },
  computed: {
    ...mapGetters('user', ['isAuthenticated'])
  },
  data: () => {
    return {
      drawer: false,
    }
  }
}
</script>
