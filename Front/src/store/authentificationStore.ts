import { RootState } from '@/store/types';
import { Module, ActionTree, MutationTree, GetterTree } from "vuex";

export enum privilege{
    client,
    company,
    estateAgent
}

export interface AuthentificationState {
    connected: boolean,
    privilege: privilege
    login: string,
    token: string
}


const state: AuthentificationState = {
    connected: false,
    privilege: 0,
    login: "",
    token: ""
};

const getters: GetterTree<AuthentificationState, RootState> = {
    getConnected(state) : boolean
    {
        return state.connected;
    }
};


const actions: ActionTree<AuthentificationState, RootState> = {
    
};

const mutations: MutationTree<AuthentificationState> = {
};

export const authentificationStore: Module<AuthentificationState, RootState> = {
    namespaced: true,
    state,
    getters,
    actions,
    mutations,
};