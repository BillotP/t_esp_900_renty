<template>
  <v-main>
    <v-card v-if="this.$apollo.queries.estateAgent.loading" height="200">
      <v-layout align-center justify-center column fill-height>
        <v-flex row align-center>
          <v-progress-circular
            indeterminate
            :size="50"
            color="primary"
            class=""
          ></v-progress-circular>
        </v-flex>
      </v-layout>
    </v-card>
    <v-card v-else style="margin: 10px">
      <v-card-title class="justify-center"
        ><b style="text-transform: capitalize"
          >{{ estateAgent.user.username }} @
          {{ estateAgent.company.name }} 🖊️</b
        >
      </v-card-title>
      <v-card-text>
        <v-row>
          <v-text-field
            hint="Update your estateAgent name"
            v-model="estateAgent.user.username"
            label="Name"
          ></v-text-field>
        </v-row>
        <v-row>
          <v-text-field
            hint="Update your estateAgent password"
            v-model="userPassword"
            label="New password"
            type="password"
          ></v-text-field>
        </v-row>
        <v-btn
          depressed
          block
          color="primary"
          style="margin: 10px"
          :disabled="isLoading"
          v-on:click="update()"
        >
          {{ isLoading ? "" : "Update" }}</v-btn
        >
      </v-card-text>
    </v-card></v-main
  >
</template>

<script>
import gql from "graphql-tag";
export default {
  apollo: {
    estateAgent: {
      query: gql`
        query estateAgent($id: Int!) {
          estateAgent(id: $id) {
            ID
            createdAt
            updatedAt
            company {
              ID
              name
            }
            user {
              ID
              createdAt
              updatedAt
              username
              password
              role
            }
          }
        }
      `,
      variables() {
        return { id: this.$route.params.id };
      },
      pollInterval: 2000,
    },
  },
  methods: {
    udpate() {
      console.log("Update !");
    },
  },
  data() {
    return {
      estateAgent: {},
      userPassword: "",
      isLoading: false,
      estateAgentID: this.$route.params.id,
    };
  },
};
</script>