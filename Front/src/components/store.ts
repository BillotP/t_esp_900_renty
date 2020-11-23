import { DataTicketReceived } from './type';
import { RootState } from '@/store/types';
import { Module, ActionTree, MutationTree, GetterTree } from "vuex";

// tslint:disable-next-line: no-empty-interface
export interface TicketState {
    data: DataTicketReceived | null;
    loading: boolean;
}

const state: TicketState = {
    data: null,
    loading: false
};

const getters: GetterTree<TicketState, RootState> = {
getTickets() {
    state.loading = true;
    state.data = {
        items: [
          {
            name: 'Jean Mescouilles',
            needs: 'Achat',
            priority: 3,
            responsible: 'Mireille Laboutteille',
          },
          {
            name: 'Jédi Wolah',
            needs: 'Location',
            priority: 1,
            responsible: 'Jean-Luc Lassodeaux',
          },
          {
            name: 'Foo Bar',
            needs: 'Location',
            priority: 2,
            responsible: 'Jean-Luc Lassodeaux',
          },
          {
            name: 'Jessica Parmoi',
            needs: 'Achat',
            priority: 1,
            responsible: 'Jean-Luc Lassodeaux',
          },
          {
            name: 'Buzz L\'Eclair',
            needs: 'Colocation',
            priority: 2,
            responsible: 'Pierre Cailloux',
          },
        ]
    }
    state.loading = false;
    return state.data;
  },
getLoading() {
    return state.loading;
}
}

const actions: ActionTree<TicketState, RootState> = {
 
};

const mutations: MutationTree<TicketState> = {
 
}
export const ticketStore: Module<TicketState, RootState> = {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
};