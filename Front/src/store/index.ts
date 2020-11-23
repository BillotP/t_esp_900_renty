import Vue from "vue";
import Vuex, { StoreOptions, ActionTree, MutationTree, GetterTree } from "vuex";
import { board } from "@/modules/board/store";
import { RootState } from './types';

Vue.use(Vuex);

const state: RootState = {
};

const getters: GetterTree<RootState, any> = {

};


const actions: ActionTree<RootState, any> = {

};

const mutations: MutationTree<RootState> = {
};

const store: StoreOptions<any> = {
  state,
  actions,
  mutations,
  getters,
  modules: {
    board
  },
};

export default new Vuex.Store<any>(store);
