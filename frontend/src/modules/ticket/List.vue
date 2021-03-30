<template>
  <v-container>
    <v-btn right color="primary" v-on:click="goToCreateTicket">
      Create Ticket
    </v-btn>
    <v-data-table
        :headers="headers"
        :items="anomalies"
        :items-per-page="5"
        class="elevation-1"
    >
      <template v-slot:item.actions="{ item }">
        <v-icon
            small
            class="mr-2"
            @click="goToProfile(item)"
        >
          mdi-eye
        </v-icon>
      </template>
    </v-data-table>
  </v-container>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
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
    }
  }
`;

@Component
export default class TicketList extends Vue {
  public anomalies = [];

  beforeMount() {
    this.$apollo.getClient().query({
      query: TICKET_QUERY
    }).then((res) => {
      this.anomalies = res.data.anomalies;
      console.log(res);
    }).catch((err) => {
      console.error(err);
    });
  }

  data() {
    return {
      headers: [
        {
          text: 'ID',
          align: 'start',
          sortable: false,
          value: 'ID',
        },
        {text: 'Type', value: 'type'},
        {text: 'Description', value: 'description'},
        {text: 'CreateBy', value: 'createBy.user.username'},
        {text: 'AssignedTo', value: 'assignedTo.user.username'},
        {text: 'Actions', value: 'actions', sortable: false},
      ]
    }
  }

  goToProfile(ticket: any) {
    this.$router.push("/ticket/" + ticket.ID);
  }

  goToCreateTicket() {
    this.$router.push("/create/ticket/");
  }
}
</script>

<style>
.v-data-table {
  margin-top: 5rem;
}
</style>
