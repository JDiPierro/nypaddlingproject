import { locationService } from '../services'

const state = {
  locations: [],
};

const getters = {
  counties: state => {
    return [...new Set(state.locations.map(item => item.county))].sort();
  }
}

const actions = {
  loadAll ({ commit, dispatch }) {
    locationService.load().then((locations) => {
      commit('load', locations)
    }).catch(() => {
      dispatch('alert/error', 'Unable to communicate with server...', {root:true})
    })
  },
  async claim({ commit, dispatch }, { location_id }) {
    await locationService.claim(location_id).then(async () => {
      dispatch('getDetails', { location_id })
    }).catch(() => {
      dispatch('alert/error', 'Unable to communicate with server...', {root:true})
    })
  },
  getDetails({commit, dispatch}, { location_id }) {
    locationService.details(location_id).then(() => {
      commit('details', location_id)
    }).catch(() => {
      dispatch('alert/error', 'Unable to communicate with server...', {root:true})
    })
  }
};

const mutations = {
  load (state, locations) {
    state.locations = locations
  },
};

export const locations = {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
};
