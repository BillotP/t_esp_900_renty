import { RootState } from '@/store/types';
import { Module, ActionTree, MutationTree, GetterTree } from "vuex";
import gql from "graphql-tag";

const LOGIN_COMPANY_MUTATION = gql`
mutation ($input: UserInput) {
  loginAsCompany(input: $input) {
    token
  }
}
`;

export enum privilege{
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
    privilege: privilege
    login: string,
}


const state: AuthentificationState = {
    connected: false,
    privilege: 0,
    login: "",
};

const getters: GetterTree<AuthentificationState, RootState> = {
    getConnected(state) : boolean
    {
        return state.connected;
    },
    getPrivilege(state) : privilege
    {
        return state.privilege;
    }
};


const actions: ActionTree<AuthentificationState, RootState> = {
    login({commit}, login: Login)
    {
        commit("setLogin")
    }
};

const mutations: MutationTree<AuthentificationState> = {
    setLogin(state, login: Login)
    {
        state.connected = true;
        state.login = login.username;
        state.privilege = login.privilege;
    }
};

export const authentificationStore: Module<AuthentificationState, RootState> = {
    namespaced: true,
    state,
    getters,
    actions,
    mutations,
};