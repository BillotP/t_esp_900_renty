<template>
  <v-container>
    <v-btn
      v-if="getPrivilege == 2"
      right
      color="primary"
      v-on:click="goToCreateTicket"
    >
      Create Ticket
    </v-btn>
    <v-data-table
      :headers="headers"
      :items="anomalies"
      :items-per-page="5"
      class="elevation-1"
    >
      <template v-slot:[`item.state`]="{ item }">
        <v-avatar
          size="20"
          v-bind:color="
            item.state != states[0]
              ? item.state != states[1]
                ? 'green'
                : 'yellow'
              : 'red'
          "
        >
        </v-avatar>
      </template>
      <template v-slot:[`item.actions`]="{ item }">
        <v-icon small class="mr-2" @click="goToProfile(item)"> mdi-eye </v-icon>
        <v-icon
          v-if="getPrivilege == 1"
          small
          class="mr-2"
          @click="selectTicketAssign(item)"
        >
          mdi-account-convert
        </v-icon>
        <v-icon
          v-if="getPrivilege == 1"
          small
          class="mr-2"
          @click="selectTicketState(item)"
        >
          mdi-state-machine
        </v-icon>
      </template>
    </v-data-table>
    <v-dialog v-model="dialogSetTicketState">
      <v-card>
        <v-card-title class="headline"> Update state of ticket: </v-card-title>

        <v-card-text>
          <v-spacer></v-spacer>
          <v-select v-model="state" :items="states" label="States"></v-select>
        </v-card-text>

        <v-card-actions>
          <v-spacer></v-spacer>

          <v-btn color="error" text @click="dialogSetTicketState = false">
            Cancel
          </v-btn>

          <v-btn color="primary" text @click="setTicketState"> Update </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <v-dialog v-model="dialogAssignTicket">
      <v-card>
        <v-card-title class="headline">
          Assign ticket to estate agent:
        </v-card-title>

        <v-card-text>
          <v-spacer></v-spacer>
          <v-select
            v-model="assignedToID"
            :items="estateAgents"
            item-text="user.username"
            item-value="ID"
            label="Estate Agent"
          ></v-select>
        </v-card-text>

        <v-card-actions>
          <v-spacer></v-spacer>

          <v-btn color="error" text @click="dialogAssignTicket = false">
            Cancel
          </v-btn>

          <v-btn color="primary" text @click="setTicketAssign"> Assign </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script>
import gql from "graphql-tag";

const TICKET_QUERY = gql`
  query anomalies {
    anomalies {
      ID
      createBy {
        user {
          username
        }
      }
      assignedTo {
        user {
          username
        }
      }
      type
      description
      assignedToID
      state
    }
  }
`;

const ASSIGN_TICKET_MUTATION = gql`
  mutation updateAnomaly($id: Int!, $input: AnomalyUpdateInput) {
    updateAnomaly(id: $id, input: $input) {
      ID
      state
      assignedToID
    }
  }
`;

const ESTATE_AGENTS_QUERY = gql`
  query estateAgents {
    estateAgents {
      ID
      user {
        username
      }
      company {
        name
      }
    }
  }
`;
export default {
  apollo: {
    anomalies: {
      query: TICKET_QUERY,
      pollInterval: 4000,
    },
  },
  methods: {
    get getPrivilege() {
      return Number(localStorage.getItem("privilege")) || 0;
    },
    goToProfile(ticket) {
      this.$router.push("/ticket/" + ticket.ID);
    },

    goToCreateTicket() {
      this.$router.push("/create/ticket/");
    },

    selectTicketState(ticket) {
      this.ticket = ticket;
      this.state = ticket.state;
      this.dialogSetTicketState = true;
    },

    selectTicketAssign(ticket) {
      this.ticket = ticket;
      this.assignedToID = ticket.assignedToID;
      this.dialogAssignTicket = true;
    },

    setTicketState() {
      console.log(this.ticket);
      this.cli
        .mutate({
          mutation: ASSIGN_TICKET_MUTATION,
          variables: {
            id: this.ticket.ID,
            input: { state: this.state, assignedTo: this.ticket.assignedToID },
          },
        })
        .then((res) => {
          console.log(res);
          this.dialogSetTicketState = false;
          this.ticket = {};
          this.state = null;
          this.fetchTickets();
        })
        .catch((err) => {
          console.error(err);
        });
    },

    setTicketAssign() {
      console.log(this.ticket);
      this.cli
        .mutate({
          mutation: ASSIGN_TICKET_MUTATION,
          variables: {
            id: this.ticket.ID,
            input: { state: this.ticket.state, assignedTo: this.assignedToID },
          },
        })
        .then((res) => {
          console.log(res);
          this.dialogAssignTicket = false;
          this.ticket = {};
          this.assignedToID = -1;
          this.fetchTickets();
        })
        .catch((err) => {
          console.error(err);
        });
    },
  },
  data() {
    return {
      cli: this.$apollo.getClient(),
      state: null,
      assignedToID: -1,
      ticket: {},
      anomalies: [],
      estateAgents: [],
      dialogSetTicketState: false,
      dialogAssignTicket: false,
      headers: [
        {
          text: "ID",
          align: "start",
          sortable: false,
          value: "ID",
        },
        { text: "Type", value: "type" },
        { text: "Description", value: "description" },
        { text: "CreateBy", value: "createBy.user.username" },
        { text: "AssignedTo", value: "assignedTo.user.username" },
        { text: "State", value: "state", sortable: false },
        { text: "Actions", value: "actions", sortable: false },
      ],
      states: ["TODO", "IN_PROGRESS", "DONE"],
    };
  },
};
</script>


<style>
.v-data-table {
  margin-top: 5rem;
}
</style>
