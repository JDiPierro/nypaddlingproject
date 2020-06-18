<template id="app">
  <v-content>
  </v-content>
</template>

<script lang="ts">
  import { DeploymentStoreState, IDeployment } from "@/store/modules/deployments"
  import { EnvironmentStoreState, IEnvironment } from "@/store/modules/environments"
  import { Component, Vue } from 'vue-property-decorator';
  import { Action, Getter, State } from 'vuex-class';
  import { ApplicationStoreState, IApplication } from '../store/modules/applications';

  @Component
  export default class Dashboard extends Vue {
  }
</script>

<style scoped lang="scss">
</style>
