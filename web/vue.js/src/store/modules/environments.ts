import { ActionContext, ActionTree, GetterTree, Module, MutationTree } from 'vuex';

import { API_ENDPOINT, AUTH_TOKEN } from '../../constants';
import { IRootState } from './types';

// Interfaces
export interface IEnvironment {
  id: string;
  name?: string;
}

// Initial State
export class EnvironmentStoreState {
  all: Array<IEnvironment> = [];
}

// Getters
const getters: GetterTree<EnvironmentStoreState, IRootState> = {
  all(state: EnvironmentStoreState) {
    return state.all;
  },
};

// Actions
const actions: ActionTree<EnvironmentStoreState, IRootState> = {
  async getAll({commit}: ActionContext<EnvironmentStoreState, IRootState>) {
    try {
      const response = await fetch(API_ENDPOINT + '/api/v1/environments', {
        headers: {
          'Accept': 'application/json',
          'Authorization': 'Bearer ' + localStorage.getItem(AUTH_TOKEN),
          'Content-Type': 'application/json',
        },
        method: 'get',
      });

      const json = await response.json();

      if (response.status >= 200 && response.status < 300) {
        const payload: Array<IEnvironment> = json;
        commit('SET_ALL', payload);
      } else {
        if (json.error) {
          throw new Error(json.message);
        }
      }
    } catch (err) {
      throw new Error(err);
    }
  },
};

// Mutations
const mutations: MutationTree<EnvironmentStoreState> = {
  SET_ALL(state: EnvironmentStoreState, payload: Array<IEnvironment>) {
    // Attempt to enforce order:
    const expected_envs = ['live', 'test', 'int', 'qa', 'dev'];
    let envs = new Set<IEnvironment>();
    // Find the expected environments
    for (const env_name of expected_envs) {
      for (const env of payload) {
        if (env.name === env_name) {
          envs.add(env);
        }
      }
    }
    // Try to add all envs to the set. Existing ones will fail, extras will show up at the end.
    for (const env of payload) {
      envs.add(env);
    }


    state.all = Array.from(envs);
  },
};

export const environments: Module<EnvironmentStoreState, IRootState> = {
  namespaced: true,
  state: new EnvironmentStoreState(),
  getters,
  actions,
  mutations,
};
