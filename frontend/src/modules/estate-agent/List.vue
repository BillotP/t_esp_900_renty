<template>
  <v-container>
    <v-row align="center" justify="center" style="height: 10vh">
      <v-col align="center" justify="center">
        <h3>{{ companyName }}</h3>
        <v-row align="center" justify="center"><h5>estate agents</h5></v-row>
      </v-col>

      <v-row align="center" justify="center">
        <v-btn
          style="margin: 5px"
          right
          color="primary"
          v-on:click="goToCreateEstateAgent"
        >
          +
        </v-btn>
        <v-btn
          style="margin: 5px"
          right
          color="primary"
          v-on:click="$apollo.queries['estateAgents'].refetch()"
        >
          🔄
        </v-btn>
      </v-row>
    </v-row>

    <v-data-table
      :headers="headers"
      :items="estateAgents"
      :items-per-page="5"
      class="elevation-1"
    >
      <template v-slot:[`item.createdAt`]="{ item }">
        {{ new Date(item.createdAt).toLocaleString() }}
      </template>
      <template v-slot:[`item.actions`]="{ item }">
        <v-icon small class="mr-2" @click="goToProfile(item)"> mdi-eye </v-icon>
      </template>
    </v-data-table>
    <v-snackbar
      v-model="snackbar"
      :color="hasError ? 'red' : ''"
      :timeout="timeout"
    >
      {{ text }}
      <template v-slot:action="{ attrs }">
        <v-btn color="blue" text v-bind="attrs" @click="snackbar = false">
          Close
        </v-btn>
      </template>
    </v-snackbar>
  </v-container>
</template>

<script>
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
      createdAt
      updatedAt
    }
  }
`;

export default {
  apollo: {
    estateAgents: {
      query: ESTATE_AGENT_QUERY,
      pollInterval: 3000,
    },
  },
  data() {
    return {
      snackbar: false,
      text: "",
      estateAgents: [],
      timeout: 2000,
      hasError: false,
      companyName: localStorage.getItem("username").toString().toUpperCase(),
      headers: [
        {
          text: "ID",
          align: "start",
          sortable: false,
          value: "ID",
        },
        { text: "Username", value: "user.username" },
        { text: "Registration Date", value: "createdAt" },
        { text: "Company", value: "company.name" },
        { text: "Actions", value: "actions", sortable: false },
      ],
    };
  },
  methods: {
    goToProfile(estateAgent) {
      this.$router.push("/estate-agent/" + estateAgent.ID);
    },
    goToCreateEstateAgent() {
      this.$router.push("/create/estate-agent/");
    },
  },
};

// import Vue from "vue";
// import Component from "vue-class-component";
// import gql from "graphql-tag";

// const ESTATE_AGENT_QUERY = gql`
//   query estateAgents {
//     estateAgents {
//       ID
//       user {
//         username
//       }
//       company {
//         name
//       }
//       createdAt
//       updatedAt
//     }
//   }
// `;

// @Component
// export default class EstateAgentList extends Vue {
//   async beforeMount() {
//     // this.$data.query = this.$apollo.getClient().watchQuery({
//     //   query: ESTATE_AGENT_QUERY,
//     //   pollInterval: 3000,
//     //   fetchResults: true,
//     // });

//     // this.$data.query.startPolling(3000);
//     // var res = await this.$data.query.result();
//     // if (res.errors) {
//     //   console.error(res.errors);
//     //   this.$data.text =
//     //     "⚠️ Failed to list estate agents :" + res.errors[0]["message"];
//     //   this.$data.hasError = true;
//     //   this.$data.snackbar = true;
//     // } else {
//     //   this.$data.estateAgents = res.data.estateAgents;
//     // }
//   }
//   public apollo: any = {
//     estateAgents: ESTATE_AGENT_QUERY,
//   };
//   data() {
//     return {
//       query: null,
//       snackbar: false,
//       text: "",
//       estateAgents: [],
//       timeout: 2000,
//       hasError: false,
//       companyName: "My company",
//       headers: [
//         {
//           text: "ID",
//           align: "start",
//           sortable: false,
//           value: "ID",
//         },
//         { text: "Username", value: "user.username" },
//         { text: "Registration Date", value: "createdAt" },
//         { text: "Company", value: "company.name" },
//         { text: "Actions", value: "actions", sortable: false },
//       ],
//     };
//   }

//   public goToProfile(estateAgent: any) {
//     this.$router.push("/estate-agent/" + estateAgent.ID);
//   }

//   public goToCreateEstateAgent() {
//     this.$router.push("/create/estate-agent/");
//   }
// }
</script>

<style>
.v-data-table {
  margin-top: 5rem;
}
</style>
