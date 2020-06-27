<template>
  <v-app>
    <Sidebar v-if="isAuthenticated === true" v-model="drawer" />
    <Header v-model="drawer" />
    <router-view></router-view>
    <v-footer
      absolute
    >
      <v-spacer></v-spacer>
      <div>favicon by <a href="https://icons8.com">icons8</a></div>
    </v-footer>
  </v-app>
</template>

<script>
import { mapGetters, mapActions } from 'vuex';

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
  methods: {
    ...mapActions('user', ['me'])
  },
  beforeMount() {
    this.me()
  },
  data: () => {
    return {
      drawer: true,
    }
  }
}
</script>
