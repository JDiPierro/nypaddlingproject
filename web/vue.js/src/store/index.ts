import Vue from 'vue';
import Vuex from 'vuex';

import { user } from './modules/user';
import { applications } from './modules/applications';
import { deployments } from './modules/deployments';
import { environments } from './modules/environments';

Vue.use(Vuex);

const debug = process.env.NODE_ENV !== 'production';

export default new Vuex.Store({
  modules: {
    user,
    applications: applications,
    deployments: deployments,
    environments: environments,
  },
  strict: debug,
  plugins: [],
});
