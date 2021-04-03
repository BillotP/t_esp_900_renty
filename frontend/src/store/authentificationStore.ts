import { RootState } from '@/store/types';
import { Module, ActionTree, MutationTree, GetterTree } from "vuex";


export enum privilege {
    client,
    company,
    estateAgent
}

export interface Login {
    username: string,
    privilege: privilege
}


export interface AuthentificationState {
    connected: boolean,
}


const state: AuthentificationState = {
    connected: false,
};

const getters: GetterTree<AuthentificationState, RootState> = {
    getConnected(state): boolean {
        return state.connected;
    },
};


const actions: ActionTree<AuthentificationState, RootState> = {
    login({ commit }, login: Login) {
        commit("setLogin")
    }
};

const mutations: MutationTree<AuthentificationState> = {
    setLogin(state, _login: Login) {
        state.connected = true;
    }
};

export const authentificationStore: Module<AuthentificationState, RootState> = {
    namespaced: true,
    state,
    getters,
    actions,
    mutations,
};
