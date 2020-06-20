import jwtDecode from 'jwt-decode';
import md5 from 'md5';

import { API_ENDPOINT, AUTH_TOKEN } from '../../constants';
import { userService } from "../services"

// Initial State
const state = (() => {
  const token = localStorage.getItem(AUTH_TOKEN);

  if (token === null || token === '') {
    return {
      email: undefined,
      id: undefined,
      name: undefined,
      token: undefined,
    };
  }

  return {
    token,
  };
})();

// Getters
const getters = {
  isAuthenticated(us) {
    return !!us.token;
  },

  getName(us) {
    return us.name ? us.name : '';
  },

  getAvatar(us) {
    const base = 'https://www.gravatar.com/avatar/';
    const query = `d=mm&r=g&s=${512}`;
    const formattedEmail = ('' + us.email).trim().toLowerCase();
    const hash = md5(formattedEmail, {encoding: 'binary'});

    return `${base}${hash}?${query}`;
  },
};

// Actions
const actions = {
  async login({ commit }, { fbAuth }) {
      //commit('SET_USER', fbAuth);
      commit('SET_TOKEN', user.id);
  },

  async get({commit}) {
    try {
      const response = await fetch(API_ENDPOINT + '/api/v1/account', {
        headers: {
          'Accept': 'application/json',
          'Authorization': 'Bearer ' + localStorage.getItem(AUTH_TOKEN),
          'Content-Type': 'application/json',
        },
        method: 'get',
      });

      const json = await response.json();

      if (response.status >= 200 && response.status < 300) {
        const payload = json;
        commit('SET_USER', payload);
      } else {
        if (json.error) {
          commit('UNSET_USER');
          throw new Error(json.message);
        }
      }
    } catch (err) {
      commit('UNSET_USER');
      throw new Error(err);
    }
  },

  logout({commit}) {
    commit('UNSET_USER');
  },
};

// Mutations
const mutations = {
  SET_USER(us, payload) {
    us.email = payload.email;
    us.id = payload.id;
    us.name = payload.name;
  },

  SET_TOKEN(us, payload) {
    localStorage.setItem(AUTH_TOKEN, payload);

    us.token = payload;
  },

  UNSET_USER(us) {
    localStorage.removeItem(AUTH_TOKEN);

    us.email = undefined;
    us.id = undefined;
    us.name = undefined;
    us.token = undefined;
  },
};

export const user = {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
