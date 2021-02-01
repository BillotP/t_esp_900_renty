import { RootState } from '@/store/types';
import { Module, ActionTree, MutationTree, GetterTree } from "vuex";

// tslint:disable-next-line: no-empty-interface
export interface BoardState {
  board: number[][][];
  boardChoose: number[][];
  tour: number;
  score: number;
}

const state: BoardState = {
  board: [],
  boardChoose: [],
  tour: 1,
  score: 4
};

const getters: GetterTree<BoardState, RootState> = {
  getScore(state) {
    return state.score;
  },
  getTour(state) {
    return state.tour;
  },
  getSave(state) {
    return state.board.length
  },
  getBoard(state, save: number) {
    if (state.boardChoose.length === 0) {
      state.boardChoose =
        [[0, 0, 0, 0],
        [0, 0, 0, 0],
        [0, 0, 0, 0],
        [0, 0, 0, 0]];
      state.boardChoose[Math.floor(Math.random() * Math.floor(4))].splice(
        Math.floor(Math.random() * Math.floor(4)),
        1,
        2
      );
      state.boardChoose[Math.floor(Math.random() * Math.floor(4))].splice(
        Math.floor(Math.random() * Math.floor(4)),
        1,
        2
      );
      return state.boardChoose;
    }
    return state.boardChoose;
  }
}

const actions: ActionTree<BoardState, RootState> = {
  save({ commit }, board: number[][]) {
    commit("addSave", board)
  },
  changeSave({ commit }, changeSave: number) {
    commit("changeSave", changeSave)
  },
  addScore({ commit }, score: number) {
    commit("scoreAdded", score);
  },
  addTour({ commit }, tour: number) {
    commit("tourAdded", tour);
  }
};

const mutations: MutationTree<BoardState> = {
  addSave(state, board: number[][]) {
    state.board.push([...board]);
  },
  changeSave(state, changeSave: number) {
    state.boardChoose = [...state.board[changeSave]];
  },
  tourAdded(state, tour: number) {
    state.tour += tour;
  },
  scoreAdded(state, score: number) {
    state.score += score;
  }
}
export const board: Module<BoardState, RootState> = {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
};
