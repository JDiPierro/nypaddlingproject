import { locationService } from '../services'

const state = {
  locations: [],
};

const actions = {
  loadAll ({ commit, dispatch }) {
    locationService.load().then((locations) => {
      commit('load', locations)
    }).catch(() => {
      dispatch('alert/error', 'Unable to communicate with server...', {root:true})
    })
  },
};

const mutations = {
  load (state, locations) {
    state.locations = locations
  },
};

export const locations = {
  namespaced: true,
  state,
  actions,
  mutations
};
