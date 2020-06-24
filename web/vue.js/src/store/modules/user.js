import jwtDecode from 'jwt-decode';
import md5 from 'md5';

import { API_ENDPOINT, AUTH_TOKEN } from '../../constants';
import { userService } from "../services"
import axios from "axios"
import config from 'config'

// Initial State
const state = {
  user: null,
};

// Getters
const getters = {
  isAuthenticated(us) {
    return us.user !== null
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
  async login({ commit, router }) {
    // TODO: Doesn't need to be an action.
    window.location = "/api/login"
  },

  async me({ commit, dispatch }) {
    console.log("In ME action")
    userService.loadUser().then((me) => {
      commit('SET_USER', me)
    }).catch(() => {
      dispatch('alert/error', 'Unable to communicate with server...', {root:true})
    })
  },
/*
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
  },*/

  logout({commit}) {
    commit('UNSET_USER');
  },
};

// Mutations
const mutations = {
  SET_USER(state, payload) {
    state.user = payload
  },

  UNSET_USER(state) {
    state.user = null;
  },
};

export const user = {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
