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

<script lang="ts">
import { Component, Emit, Prop, Vue } from 'vue-property-decorator';
import { Action, Getter, State } from 'vuex-class';

import { IUserState } from '../store/modules/user';

@Component
export default class Sidebar extends Vue {
  @Prop(Boolean) public value!: boolean;

  @Action('logout', { namespace: 'user' }) private logout: any;
  @Getter('getName', { namespace: 'user' }) private getName!: string;
  @Getter('getAvatar', { namespace: 'user' }) private getAvatar!: string;
  @State('user') private user!: IUserState;

  private items = [
    { title: 'Home', icon: 'dashboard' },
    { title: 'About', icon: 'question_answer' },
  ];

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
