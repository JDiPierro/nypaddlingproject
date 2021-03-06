import Vue from 'vue';
import Vuex from 'vuex';

import { user } from './modules/user';
import { locations } from './modules/locations';

Vue.use(Vuex);

const debug = process.env.NODE_ENV !== 'production';

const store = new Vuex.Store({
  modules: {
    user,
    locations,
  },
  strict: debug,
  plugins: [],
});

export default store
