<template>
  <v-navigation-drawer :clipped="$vuetify.breakpoint.lgAndUp" v-model="value" fixed app disable-resize-watcher>

    <v-list dense>
      <v-list-item to="/">
        <v-list-item-action>
          <v-icon>mdi-clipboard-list-outline</v-icon>
        </v-list-item-action>
        <v-list-item-content>
          <v-list-item-title>My Claimed Locations</v-list-item-title>
        </v-list-item-content>
      </v-list-item>
      <v-list-item to="locations">
        <v-list-item-action>
          <v-icon>mdi-map-marker</v-icon>
        </v-list-item-action>
        <v-list-item-content>
          <v-list-item-title>Locations</v-list-item-title>
        </v-list-item-content>
      </v-list-item>
    </v-list>

    <v-divider></v-divider>

    <v-list-item @click="doLogout">
      <v-list-item-action>
        <v-icon>power_settings_new</v-icon>
      </v-list-item-action>
      <v-list-item-content>
        <v-list-item-title>Logout</v-list-item-title>
      </v-list-item-content>
    </v-list-item>
  </v-navigation-drawer>
</template>

<script>
import { mapActions, mapState, mapGetters } from 'vuex';

export default {
  props: ["value"],
  methods: {
    ...mapActions('user', ['logout']),
    ...mapGetters('user', ['getName', 'getAvatar']),
    goTo(path) {
      this.$router.push({ path });
    },
    doLogout() {
      this.$emit("input", false)
      this.logout();
      this.$router.push({ path: '/' });
    }
  },
  computed: {
    ...mapState('user', ['user']),
  },
  data: () => {
    return {
      items: [
        { title: 'Home', icon: 'dashboard' },
        { title: 'About', icon: 'question_answer' },
      ]
    }
  }
}
</script>
