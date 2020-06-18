<template>
  <v-app-bar app color="primary" :clipped-left="$vuetify.breakpoint.lgAndUp" dark>
    <v-toolbar-title class="headline header" @click="goTo('/')">
      <v-app-bar-nav-icon v-if="isAuthenticated === true" @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
      <span v-bind:class="{ 'hidden-xs-only': isAuthenticated }">Upstate NY Paddling Project</span>
    </v-toolbar-title>
    <v-spacer></v-spacer>
  </v-app-bar>
</template>

<script lang="ts">
import { Component, Emit, Prop, Vue } from 'vue-property-decorator';
import { Action, Getter } from 'vuex-class';

@Component
export default class Header extends Vue {
  @Prop(Boolean) public value!: boolean;

  @Action('logout', { namespace: 'user' }) private logout: any;
  @Getter('isAuthenticated', { namespace: 'user' }) private isAuthenticated!: boolean;

  private get drawer(): boolean {
    return this.value;
  }

  private set drawer(newVal: boolean) {
    this.changeValue(newVal);
  }

  @Emit('input')
  public changeValue(val: boolean) { /* */ }

  private goTo(path: string) {
    this.$router.push({ path });
  }

  private doLogout() {
    this.drawer = false;
    this.logout();
    this.$router.push({ path: '/' });
  }
}
</script>

<style scoped>
.header {
  cursor: pointer;
}
</style>
