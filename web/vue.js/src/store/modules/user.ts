import jwtDecode from 'jwt-decode';
import md5 from 'md5';
import { ActionContext, ActionTree, GetterTree, Module, MutationTree } from 'vuex';

import { API_ENDPOINT, AUTH_TOKEN } from '../../constants';
import { IRootState } from './types';

// Interfaces
export interface IUserState {
  email?: string;
  id?: string;
  name?: string;
  token?: string;
}

export interface IJWTDecoded {
  exp: number;
  id: string;
  email: string;
  name: string;
}

// Initial State
const state: IUserState = (() => {
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
const getters: GetterTree<IUserState, IRootState> = {
  isAuthenticated(us: IUserState): boolean {
    return !!us.token;
  },

  getName(us: IUserState): string {
    return us.name ? us.name : '';
  },

  getAvatar(us: IUserState): string {
    const base = 'https://www.gravatar.com/avatar/';
    const query = `d=mm&r=g&s=${512}`;
    const formattedEmail = ('' + us.email).trim().toLowerCase();
    const hash = md5(formattedEmail, {encoding: 'binary'});

    return `${base}${hash}?${query}`;
  },
};

// Actions
const actions: ActionTree<IUserState, IRootState> = {
  async login({ commit }, { fbAuth }) {
      //commit('SET_USER', fbAuth);
      commit('SET_TOKEN', fbAuth.accessToken);
  },

  async get({commit}: ActionContext<IUserState, IRootState>) {
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
        const payload: IUserState = json;
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

  logout({commit}: ActionContext<IUserState, IRootState>) {
    commit('UNSET_USER');
  },
};

// Mutations
const mutations: MutationTree<IUserState> = {
  SET_USER(us: IUserState, payload: IUserState) {
    us.email = payload.email;
    us.id = payload.id;
    us.name = payload.name;
  },

  SET_TOKEN(us: IUserState, payload) {
    localStorage.setItem(AUTH_TOKEN, payload as string);

    us.token = payload as string;
  },

  UNSET_USER(us: IUserState) {
    localStorage.removeItem(AUTH_TOKEN);

    us.email = undefined;
    us.id = undefined;
    us.name = undefined;
    us.token = undefined;
  },
};

export const user: Module<IUserState, IRootState> = {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
