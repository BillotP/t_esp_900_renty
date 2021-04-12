<template>
  <v-main>
    <v-btn
        v-if="privilege == '2'"
        right
        block
        color="primary"
        v-on:click="goToCreateTicket"
    >
      Create Ticket
    </v-btn>
    <v-chip-group
        v-model="selection"
        active-class="deep-purple accent-4 white--text"
        style="padding: 5px"
        column
    >
      <v-chip>Datatable</v-chip>
      <v-chip>Board</v-chip>
    </v-chip-group>
    <v-container v-if="!selection">
      <v-data-table
          :headers="headers"
          :items="anomalies"
          :items-per-page="5"
          class="elevation-1"
      >
        <template v-slot:[`item.priority`]="{ item }">
          <span>{{ priorityIcons[item.priority] }}</span>
        </template>
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
          <v-icon small class="mr-2" @click="goToProfile(item)"> mdi-eye</v-icon>
          <v-icon
              v-if="privilege == '1'"
              small
              class="mr-2"
              @click="selectTicketAssign(item)"
          >
            mdi-account-convert
          </v-icon>
          <v-icon
              v-if="privilege == '1'"
              small
              class="mr-2"
              @click="selectTicketState(item)"
          >
            mdi-state-machine
          </v-icon>
          <v-icon
              v-if="privilege == '1'"
              small
              class="mr-2"
              @click="selectTicketPriority(item)"
          >
            mdi-alert
          </v-icon>
        </template>
      </v-data-table>
    </v-container>
    <v-container v-else>
      <v-row>
        <v-col cols="3">
          <v-card-subtitle style="text-align: center; text-transform: uppercase; color: #2c3e50">
            {{ ticketStates['TODO'] }}
          </v-card-subtitle>
          <draggable :disabled="privilege != '1'" class="list-group" :list="toDo" group="people"
                     @change="log($event, 'TODO')">
            <div
                class="list-group-item"
                v-for="(element, _) in toDo"
                :key="element.ID"
            >
              <v-alert>
                <v-icon
                    v-if="privilege == '1'"
                    small
                    class="mr-2"
                    @click="selectTicketAssign(element)"
                    right
                    style="position: absolute; right: 0;"
                >
                  mdi-account-convert
                </v-icon>
                <v-icon
                    v-if="privilege == '1'"
                    small
                    class="mr-2"
                    @click="selectTicketPriority(element)"
                    right
                    style="position: absolute; right: 20px;"
                >
                  mdi-alert
                </v-icon>
                <h4>{{ ticketTypes[element.type] }}</h4>
                <span>{{ priorityIcons[element.priority] }}</span>
                <p>{{ element.description }}</p>
                <i>👤: {{ element.createBy.user.username }}</i>
              </v-alert>
            </div>
          </draggable>
        </v-col>
        <v-col cols="3">
          <v-card-subtitle style="text-align: center; text-transform: uppercase; color: #2c3e50">
            {{ ticketStates['IN_PROGRESS'] }}
          </v-card-subtitle>
          <draggable :disabled="privilege != '1'" class="list-group" :list="inProgress" group="people"
                     @change="log($event, 'IN_PROGRESS')">
            <div
                class="list-group-item"
                v-for="(element, _) in inProgress"
                :key="element.ID"
            >
              <v-alert>
                <v-icon
                    v-if="privilege == '1'"
                    small
                    class="mr-2"
                    @click="selectTicketAssign(element)"
                    right
                    style="position: absolute; right: 0;"
                >
                  mdi-account-convert
                </v-icon>
                <v-icon
                    v-if="privilege == '1'"
                    small
                    class="mr-2"
                    @click="selectTicketPriority(element)"
                    right
                    style="position: absolute; right: 20px;"
                >
                  mdi-alert
                </v-icon>
                <h4>{{ ticketTypes[element.type] }}</h4>
                <span>{{ priorityIcons[element.priority] }}</span>
                <p>{{ element.description }}</p>
                <i>👤: {{ element.createBy.user.username }}</i>
              </v-alert>
            </div>
          </draggable>
        </v-col>
        <v-col cols="3">
          <v-card-subtitle style="text-align: center; text-transform: uppercase; color: #2c3e50">
            {{ ticketStates['DONE'] }}
          </v-card-subtitle>
          <draggable :disabled="privilege != '1'" class="list-group" :list="done" group="people"
                     @change="log($event, 'DONE')">
            <div
                class="list-group-item"
                v-for="(element, _) in done"
                :key="element.ID"
            >
              <v-alert>
                <v-icon
                    v-if="privilege == '1'"
                    small
                    class="mr-2"
                    @click="selectTicketAssign(element)"
                    right
                    style="position: absolute; right: 0;"
                >
                  mdi-account-convert
                </v-icon>
                <v-icon
                    v-if="privilege == '1'"
                    small
                    class="mr-2"
                    @click="selectTicketPriority(element)"
                    right
                    style="position: absolute; right: 20px;"
                >
                  mdi-alert
                </v-icon>
                <h4>{{ ticketTypes[element.type] }}</h4>
                <span>{{ priorityIcons[element.priority] }}</span>
                <p>{{ element.description }}</p>
                <i>👤: {{ element.createBy.user.username }}</i>
              </v-alert>
            </div>
          </draggable>
        </v-col>
      </v-row>
    </v-container>

    <v-dialog v-model="dialogSetTicketState">
      <v-card>
        <v-card-title class="headline"> Update state of ticket:</v-card-title>

        <v-card-text>
          <v-spacer></v-spacer>
          <v-select v-model="state" :items="states" label="States"></v-select>
        </v-card-text>

        <v-card-actions>
          <v-spacer></v-spacer>

          <v-btn color="error" text @click="dialogSetTicketState = false">
            Cancel
          </v-btn>

          <v-btn color="primary" text @click="setTicketState"> Update</v-btn>
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

          <v-btn color="primary" text @click="setTicketAssign"> Assign</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <v-dialog v-model="dialogTicketPriority">
      <v-card>
        <v-card-title class="headline"> Update priority of ticket:</v-card-title>

        <v-card-text>
          <v-spacer></v-spacer>
          <v-select v-model="priority" :items="priorities" label="Priorities"></v-select>
        </v-card-text>

        <v-card-actions>
          <v-spacer></v-spacer>

          <v-btn color="error" text @click="dialogSetTicketState = false">
            Cancel
          </v-btn>

          <v-btn color="primary" text @click="setTicketPriority"> Update</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-main>
</template>

<script>
import gql from "graphql-tag";
import draggable from 'vuedraggable'

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
      priority
    }
  }
`;

const UPDATE_TICKET_MUTATION = gql`
  mutation updateAnomaly($id: Int!, $input: AnomalyUpdateInput) {
    updateAnomaly(id: $id, input: $input) {
      ID
      state
      assignedToID
      priority
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
  components: {
    draggable
  },
  apollo: {
    anomalies: {
      query: TICKET_QUERY,
      pollInterval: 4000,
    },
    estateAgents: {
      query: ESTATE_AGENTS_QUERY,
      pollInterval: 4000,
    }
  },
  methods: {
    log(evt, state) {
      if (evt.added) {
        this.ticket = evt.added.element;
        this.state = state;
        this.setTicketState();
      }
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

    selectTicketPriority(ticket) {
      this.ticket = ticket;
      this.priority = ticket.priority;
      this.dialogTicketPriority = true;
    },

    setTicketState() {
      console.log(this.ticket);
      this.cli
          .mutate({
            mutation: UPDATE_TICKET_MUTATION,
            variables: {
              id: this.ticket.ID,
              input: {state: this.state},
            },
          })
          .then((res) => {
            console.log(res);
            this.dialogSetTicketState = false;
            this.ticket = {};
            this.state = null;
          })
          .catch((err) => {
            console.error(err);
            this.dialogSetTicketState = false;
            this.ticket = {};
            this.state = null;
          });
    },

    setTicketAssign() {
      console.log(this.ticket);
      this.cli
          .mutate({
            mutation: UPDATE_TICKET_MUTATION,
            variables: {
              id: this.ticket.ID,
              input: {assignedTo: this.assignedToID},
            },
          })
          .then((res) => {
            console.log(res);
            this.dialogAssignTicket = false;
            this.ticket = {};
            this.assignedToID = -1;
          })
          .catch((err) => {
            console.error(err);
            this.dialogAssignTicket = false;
            this.ticket = {};
            this.assignedToID = -1;
          });
    },
    setTicketPriority() {
      console.log(this.ticket);
      this.cli
          .mutate({
            mutation: UPDATE_TICKET_MUTATION,
            variables: {
              id: this.ticket.ID,
              input: {priority: this.priority},
            },
          })
          .then((res) => {
            console.log(res);
            this.dialogTicketPriority = false;
            this.ticket = {};
            this.priority = null;
          })
          .catch((err) => {
            console.error(err);
            this.dialogTicketPriority = false;
            this.ticket = {};
            this.priority = null;
          });
    },
  },
  mounted() {
    this.$apollo.queries.anomalies.observer.subscribe(() => {
      this.$data.toDo = this.anomalies.filter((_anomaly) => _anomaly.state === 'TODO');
      this.$data.inProgress = this.anomalies.filter((_anomaly) => _anomaly.state === 'IN_PROGRESS');
      this.$data.done = this.anomalies.filter((_anomaly) => _anomaly.state === 'DONE');
      console.log(this.$data);
    });
  },
  data() {
    return {
      cli: this.$apollo.getClient(),
      state: null,
      priority: null,
      assignedToID: -1,
      ticket: {},
      anomalies: [],
      estateAgents: [],
      dialogSetTicketState: false,
      dialogAssignTicket: false,
      dialogTicketPriority: false,
      headers: [
        {
          text: "ID",
          align: "start",
          sortable: false,
          value: "ID",
        },
        {text: "Type", value: "type"},
        {text: "Description", value: "description"},
        {text: "CreateBy", value: "createBy.user.username"},
        {text: "AssignedTo", value: "assignedTo.user.username"},
        {text: "State", value: "state", sortable: false},
        {text: "Priority", value: "priority", sortable: false},
        {text: "Actions", value: "actions", sortable: false},
      ],
      states: ["TODO", "IN_PROGRESS", "DONE"],
      priorities: ["MAJOR", "HIGHEST", "HIGH", "MEDIUM", "LOW"],
      privilege: localStorage.getItem("privilege"),
      priorityIcons: {
        MAJOR: "⬆️⬆️",
        HIGHEST: "⬆️",
        HIGH: "↕️",
        MEDIUM: "⬇️",
        LOW: "⬇️⬇️",
      },
      done: [],
      inProgress: [],
      toDo: [],
      selection: 0,
      ticketTypes: {
        MAINTENANCE: "🔨 Maintenance",
        PAYMENT: "💳 Rent payment method",
        RENT: "⚠️Rent issue",
        DOCUMENTS: "📄 Document request",
        ACCOMMODATION: "⚠️Accommodation issue",
        OTHER: "🥜 Other",
      },
      ticketStates: {
        TODO: "📋 To do",
        IN_PROGRESS: "🔄 In progress",
        DONE: "✅ Done",
      },
    };
  },
};
</script>


<style>
.v-main {
  padding: 10px !important;
}

.container {
  height: 75vh;
}

.container .row,
.container .row .col,
.container .row .col .list-group {
  height: 100%;
}

.container .row .col {
  border: solid 1px #DDDDDD;
  border-radius: 16px;
  background-color: #DDDDDD;
  margin: 0 3vw;
}
</style>
