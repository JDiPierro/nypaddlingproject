<template>
  <v-app-bar app color="primary" :clipped-left="$vuetify.breakpoint.lgAndUp" dark>
    <v-toolbar-title class="headline header">
      <v-app-bar-nav-icon v-if="isAuthenticated === true" @click="toggleDrawer()"></v-app-bar-nav-icon>
      <span @click="goTo('/')" v-bind:class="{ 'hidden-xs-only': isAuthenticated === true }">Upstate NY Paddling Project</span>
    </v-toolbar-title>
    <v-spacer></v-spacer>
  </v-app-bar>
</template>

<script>
import { mapActions, mapGetters } from 'vuex';

export default {
  props: ["value"],
  methods: {
    ...mapActions('user', ['logout']),
    toggleDrawer: function () {
      let newValue = !this.value
      this.$emit('input', newValue);
    },
    doLogout: function() {
      this.$emit('input', false);
      this.logout();
      this.$router.push({ path: '/' });
    },
    goTo: function(path) {
      this.$router.push({ path });
    },
  },
  computed: {
    ...mapGetters('user', ['isAuthenticated'])
  },
  data: () => {
    return {
      drawer: false
    }
  }
}
</script>

<style scoped>
.header {
  cursor: pointer;
}
</style>
