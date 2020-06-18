import { IApplication } from '@/store/modules/applications';
import { EnvironmentStoreState } from '@/store/modules/environments';
import { ActionContext, ActionTree, GetterTree, Module, MutationTree } from 'vuex';

import { API_ENDPOINT, AUTH_TOKEN } from '../../constants';
import { IRootState } from './types';

// Interfaces
export interface IDeployment {
  id: string;
  app_id: string;
  env_id: string;
  version?: string;
  user?: string;
  branch?: string;
  message?: string;
  timestamp?: string;
}

export interface IEnvironmentOverview {
  applications?: Array<IApplication>;
  deployments?: Array<IDeployment>;
}

// Initial State
export class DeploymentStoreState {
  all: Array<IDeployment> = [];
  latest_by_app: Map<string, IDeployment> = new Map<string, IDeployment>();
}

// Getters
const getters: GetterTree<DeploymentStoreState, IRootState> = {
  all(state: DeploymentStoreState) {
    return state.all;
  },
  latest_app_deploy(state: DeploymentStoreState) {
    return (app_id: any) => state.latest_by_app.get(app_id)
  }
};

// Actions
const actions: ActionTree<DeploymentStoreState, IRootState> = {
  async getAll({commit}: ActionContext<DeploymentStoreState, IRootState>) {
    try {
      const response = await fetch(API_ENDPOINT + '/api/v1/deployments', {
        headers: {
          'Accept': 'application/json',
          'Authorization': 'Bearer ' + localStorage.getItem(AUTH_TOKEN),
          'Content-Type': 'application/json',
        },
        method: 'get',
      });

      const json = await response.json();

      if (response.status >= 200 && response.status < 300) {
        const payload: any = json;
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
  async latestDeployForApp({commit}: ActionContext<DeploymentStoreState, IRootState>, {app_id}) {
    try {
      const response = await fetch(API_ENDPOINT + '/api/v1/deployment/' + app_id + '/latest', {
        headers: {
          'Accept': 'application/json',
          'Authorization': 'Bearer ' + localStorage.getItem(AUTH_TOKEN),
          'Content-Type': 'application/json',
        },
        method: 'get',
      });

      const json = await response.json();

      if (response.status >= 200 && response.status < 300) {
        const payload: IDeployment = json;
        commit('SET_LATEST_FOR_APP', payload);
      } else {
        if (json.error) {
          throw new Error(json.message);
        }
      }
    } catch (err) {
      throw new Error(err);
    }
  },
  async environmentOverview({commit, dispatch}: ActionContext<EnvironmentStoreState, IRootState>, {env_id}) {
    try {
      console.log("About to load overview of " + env_id);
      const response = await fetch(API_ENDPOINT + '/api/v1/environment/' + env_id + '/overview', {
        headers: {
          'Accept': 'application/json',
          'Authorization': 'Bearer ' + localStorage.getItem(AUTH_TOKEN),
          'Content-Type': 'application/json',
        },
        method: 'get',
      });

      const json = await response.json();

      if (response.status >= 200 && response.status < 300) {
        const payload: IEnvironmentOverview = json;
        commit('SET_ALL', payload.deployments);
        dispatch('applications/set_all', {applications: payload.applications}, {root: true});
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
const mutations: MutationTree<DeploymentStoreState> = {
  SET_ALL(state: DeploymentStoreState, payload: Array<IDeployment>) {
    state.all = payload;
    for (const deploy of payload) {
      state.latest_by_app.set(deploy.app_id, deploy)
    }
  },
  SET_LATEST_FOR_APP(state: DeploymentStoreState, payload: IDeployment) {
    state.latest_by_app.set(payload.app_id, payload)
  }
};

export const deployments: Module<DeploymentStoreState, IRootState> = {
  namespaced: true,
  state: new DeploymentStoreState(),
  getters,
  actions,
  mutations,
};
