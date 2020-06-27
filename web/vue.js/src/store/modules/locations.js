import { locationService } from '../services'

const state = {
  locations: [],
  user_claims: []
};

const getters = {
  counties: state => {
    return [...new Set(state.locations.map(item => item.county))].sort();
  },
  initialClaims: state => {
    return state.user_claims.filter((claim) => {
      return claim['status'] === "initial"
    })
  },
  submittedClaims: state => {
    return state.user_claims.filter((claim) => {
      return claim['status'] === "submitted"
    })
  }
}

const actions = {
  async loadAll ({ commit, dispatch }) {
    await locationService.load().then((locations) => {
      commit('load', locations)
    }).catch(() => {
      dispatch('alert/error', 'Unable to communicate with server...', {root:true})
    })
  },
  async loadClaims ({ commit, dispatch }) {
    await locationService.loadClaims().then((claims) => {
      commit('loadClaims', claims)
    }).catch(() => {
      dispatch('alert/error', 'Unable to communicate with server...', {root:true})
    })
  },
  async claim({ commit, dispatch }, { location_id }) {
    await locationService.claim(location_id).then(async (new_claim) => {
      commit('claim', { new_claim })
    }).catch((e) => {
      dispatch('alert/error', 'Unable to communicate with server...', {root:true})
    })
  },
  async release({ commit, dispatch }, { location_id }) {
    await locationService.release(location_id).then(async () => {
      commit('release', { location_id })
    }).catch((e) => {
      console.log(e)
      dispatch('alert/error', 'Unable to communicate with server...', {root:true})
    })
  },
  async submit({ commit, dispatch }, { location_id, update_info }) {
    await locationService.submit(location_id, update_info).then(async (updated_claim) => {
      commit('submit', { updated_claim })
    }).catch((e) => {
      console.log(e)
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
  claim(state, { new_claim }) {
    for(const loc of state.locations) {
      if(loc["_id"] === new_claim['location_id']) {
        loc["claims"].push(new_claim)
        break
      }
    }
  },
  release(state, { location_id }) {
    for(const loc of state.locations) {
      if(loc["_id"] === location_id) {
        loc["claims"].filter(function( obj ) {
          return obj.location_id !== location_id;
        });
        break
      }
    }
    state.user_claims = state.user_claims.filter(function( obj ) {
      return obj.location_id !== location_id;
    });
  },
  submit(state, { updated_claim }) {
    // Update the claims list on the location
    const location = state.locations.find((obj) => {
      return obj['_id'] === updated_claim['location_id']
    })
    if (location) {
      const submitted_claim_idx = location["claims"].findIndex(( obj ) => {
        return obj.location_id !== updated_claim['location_id'];
      });
      location['claims'][submitted_claim_idx] = updated_claim
    }

    // Update the claim in the user's list of claims
    const user_claim_idx = state.user_claims.findIndex(( obj ) => {
      return obj['_id'] !== updated_claim['_id'];
    });
    state.user_claims[user_claim_idx] = updated_claim
  },
  loadClaims (state, claims) {
    state.user_claims = claims
  },
};

export const locations = {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
};
