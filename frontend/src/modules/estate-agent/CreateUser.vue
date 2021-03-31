<template>
  <v-container>
    <h1>Create Estate Agent User account</h1>
    <v-text-field v-model="username" label="Username"></v-text-field>
    <v-text-field
      type="password"
      v-model="password"
      label="Password"
    ></v-text-field>
    <v-btn depressed color="primary" v-on:click="createUser()"> Create </v-btn>
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
import Vue from "vue";
import Component from "vue-class-component";
import gql from "graphql-tag";

const CREATE_ESTATE_AGENT_USER_MUTATION = gql`
  mutation($input: EstateAgentInput) {
    createEstateAgentUser(input: $input) {
      ID
    }
  }
`;

@Component
export default class CreateEstateAgentUser extends Vue {
  data() {
    return {
      username: "",
      password: "",
      snackbar: false,
      text: "",
      timeout: 2000,
      hasError: false,
    };
  }

  async createUser() {
    try {
      const resp = await this.$apollo.getClient().mutate({
        mutation: CREATE_ESTATE_AGENT_USER_MUTATION,
        variables: {
          input: {
            user: {
              username: this.$data.username,
              password: this.$data.password,
            },
          },
        },
      });
      if (resp.data.createEstateAgentUser.ID) {
        this.$data.text =
          "Estate agent " + this.$data.username + " successfully created !";
        this.$data.hasError = false;
        this.$data.snackbar = true;
        // this.$router.("/estate-agents");
        this.$router.go(-1);
      }
    } catch (e) {
      this.$data.text =
        "⚠️ Failed to create estate agent :" + e["graphQLErrors"][0]["message"];
      this.$data.hasError = true;
      this.$data.snackbar = true;
    }
  }
}
</script>

<style>
</style>
