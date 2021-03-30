<template>
  <v-container>
    <v-btn right color="primary" v-on:click="goToCreateEstateAgent">
      Create EstateAgent
    </v-btn>
    <v-data-table
        :headers="headers"
        :items="estateAgents"
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

const ESTATE_AGENT_QUERY = gql`
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

@Component
export default class EstateAgentList extends Vue {
  public estateAgents: any[] = [];

  beforeMount() {
    this.$apollo.getClient().query({
      query: ESTATE_AGENT_QUERY
    }).then((res) => {
      this.estateAgents = res.data.estateAgents;
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
        {text: 'Username', value: 'user.username'},
        {text: 'Company', value: 'company.name'},
        {text: 'Actions', value: 'actions', sortable: false},
      ]
    }
  }

  goToProfile(estateAgent: any) {
    this.$router.push("/estate-agent/" + estateAgent.ID);
  }

  goToCreateEstateAgent() {
    this.$router.push("/create/estate-agent/");
  }
}
</script>

<style>
.v-data-table {
  margin-top: 5rem;
}
</style>
