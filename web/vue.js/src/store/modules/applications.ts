import { IDeployment } from '@/store/modules/deployments';
import { ActionContext, ActionTree, GetterTree, Module, MutationTree } from 'vuex';

import { API_ENDPOINT, AUTH_TOKEN } from '../../constants';
import { IRootState } from './types';

// Interfaces
export interface IApplication {
  id?: string;
  name?: string;

  // Not returned by the API
  latest_deploy?: IDeployment;
}

// Initial State
export class ApplicationStoreState {
  all: Array<IApplication> = [];
}

// Getters
const getters: GetterTree<ApplicationStoreState, IRootState> = {
  all(state: ApplicationStoreState) {
    return state.all;
  },
};

// Actions
const actions: ActionTree<ApplicationStoreState, IRootState> = {
  async getAll({commit}: ActionContext<ApplicationStoreState, IRootState>) {
    try {
      const response = await fetch(API_ENDPOINT + '/api/v1/applications', {
        headers: {
          'Accept': 'application/json',
          'Authorization': 'Bearer ' + localStorage.getItem(AUTH_TOKEN),
          'Content-Type': 'application/json',
        },
        method: 'get',
      });

      const json = await response.json();
      if(json == null) {
        throw new Error("no response")
      }

      if (response.status >= 200 && response.status < 300) {
        const payload: Array<IApplication> = json;
        commit('SET_ALL', payload);

        for (const app of payload) {
          await this.dispatch("deployments/latestDeployForApp", {app_id: app.id})
        }
      } else {
        if (json.error) {
          throw new Error(json.message);
        }
      }
    } catch (err) {
      throw new Error(err);
    }
  },
  set_all({commit}: ActionContext<ApplicationStoreState, IRootState>, {applications}) {
    commit('SET_ALL', applications)
  }
};

// Mutations
const mutations: MutationTree<ApplicationStoreState> = {
  SET_ALL(state: ApplicationStoreState, payload: Array<IApplication>) {
    state.all = payload
  },
};

export const applications: Module<ApplicationStoreState, IRootState> = {
  namespaced: true,
  state: new ApplicationStoreState(),
  getters,
  actions,
  mutations,
};
