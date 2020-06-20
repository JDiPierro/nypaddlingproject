<template>
  <v-app-bar app color="primary" :clipped-left="$vuetify.breakpoint.lgAndUp" dark>
    <v-toolbar-title class="headline header" @click="goTo('/')">
      <v-app-bar-nav-icon v-if="isAuthenticated === true" @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
      <span v-bind:class="{ 'hidden-xs-only': isAuthenticated === true }">Upstate NY Paddling Project</span>
    </v-toolbar-title>
    <v-spacer></v-spacer>
  </v-app-bar>
</template>

<script>
import { mapActions, mapGetters } from 'vuex';

export default {
  props: ["value"],
  methonds: {
    ...mapActions('user', ['logout']),
  },
  calculated: {
    ...mapGetters('user', ['isAuthenticated']),
    goTo: function(path) {
      this.$router.push({ path });
    },
    doLogout: function() {
      this.drawer = false;
      this.logout();
      this.$router.push({ path: '/' });
    }
  }
}
</script>

<style scoped>
.header {
  cursor: pointer;
}
</style>
