<template>
  <v-card>
    <v-card-title>
      <v-text-field
        v-model="search"
        append-icon="mdi-magnify"
        label="Rechercher..."
        single-line
        hide-details
      ></v-text-field>
    </v-card-title>
    <v-data-table
      :search="search"
      :headers="headers"
      :items="ticketReceived.items"
      :options.sync="options"
      :server-items-length="ticketReceived.items.length"
      :loading="loading"
      loading-text="Loading... Please wait"
      group-by="priority"
      sort-by="priority"
    >
    </v-data-table>
  </v-card>
</template>

<script lang="ts">
import Vue from "vue"
import Component from "vue-class-component"
import { namespace } from "vuex-class";
import { DataTicketReceived, Ticket } from './type'
import { ticketStore } from './store'

const ticketModule = namespace("ticketStore");
@Component
export default class TicketList extends Vue {
  private search:String = '';
  private totalTickets:Number = 0;
  private tickets = [];
  private options = {};
  // private { sortBy, sortDesc, page, itemsPerPage } = options;
  private headers = [
    {
      text: 'Client',
      align: 'start',
      sortable: true,
      value: 'name',
    },
    { text: 'Besoin', value: 'needs' },
    { text: 'Priorité', value: 'priority' },
    { text: 'Assigné à', value: 'responsible' },
  ];
  @ticketModule.Getter("getTickets")
  private ticketReceived!: DataTicketReceived;

  @ticketModule.Getter("getLoading")
  private loading!: boolean;

  // private getDataFromApi() {
  //   this.loading = true
  //   this.fakeApiCall().then(data!: DataTicketReceived => {
  //     this.tickets = data.items
  //     this.totalTickets = data.total
  //     this.loading = false
  //   })
  // }
  // private fakeApiCall() {
  //   return new Promise((resolve, reject) => {
  //     const { sortBy, sortDesc, page, itemsPerPage } = this.options
  //     let items:Ticket[] = this.getTickets()
  //     const total = items.length
  //     if (sortBy.length === 1 && sortDesc.length === 1) {
  //       items = items.sort((a, b) => {
  //         const sortA = a[sortBy[0]]
  //         const sortB = b[sortBy[0]]
  //         if (sortDesc[0]) {
  //           if (sortA < sortB) return 1
  //           if (sortA > sortB) return -1
  //           return 0
  //         } else {
  //           if (sortA < sortB) return -1
  //           if (sortA > sortB) return 1
  //           return 0
  //         }
  //       })
  //     }
  //     if (itemsPerPage > 0) {
  //       items = items.slice((page - 1) * itemsPerPage, page * itemsPerPage)
  //     }
  //     setTimeout(() => {
  //       resolve({
  //         items,
  //         total,
  //       })
  //     }, 1000)
  //   })
  // }
  public beforeCreate() {
    if (this.$store.state.ticketStore === undefined) {
      this.$store.registerModule("ticketStore", ticketStore);
    }
  }
}
</script>