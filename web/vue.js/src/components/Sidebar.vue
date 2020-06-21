<template>
  <v-navigation-drawer :clipped="$vuetify.breakpoint.lgAndUp" v-model="drawer" fixed app disable-resize-watcher>

    <v-list dense>
      <v-list-item @click="goTo('/')">
        <v-list-item-action>
          <v-icon>home</v-icon>
        </v-list-item-action>
        <v-list-item-content>
          <v-list-item-title>Home</v-list-item-title>
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
      drawer: false,
      items: [
        { title: 'Home', icon: 'dashboard' },
        { title: 'About', icon: 'question_answer' },
      ]
    }
  }
}
</script>
