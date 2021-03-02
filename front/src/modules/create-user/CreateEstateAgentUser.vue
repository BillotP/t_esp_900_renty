<template>
  <div>
    <h1>Create Estate Agent User account</h1>
    <v-text-field v-model="username" label="Username"></v-text-field>
    <v-text-field type="password" v-model="password" label="Password"></v-text-field>
    <v-btn depressed color="primary" v-on:click="createUser()">
      Create
    </v-btn>
    <v-snackbar v-model="snackbar" :timeout="timeout">
      {{ text }}
      <template v-slot:action="{ attrs }">
        <v-btn color="blue" text v-bind="attrs" @click="snackbar = false">
          Close
        </v-btn>
      </template>
    </v-snackbar>
  </div>
</template>

<script>
import Vue from "vue";
import Component from 'vue-class-component';
import gql from "graphql-tag";

const CREATE_ESTATE_AGENT_USER_MUTATION = gql`
mutation ($input: EstateAgentInput) {
  createEstateAgentUser(input: $input) {
    ID
  }
}
`;

@Component
export default class CreateEstateAgentUser extends Vue {

  data() {
    return {
      username: '',
      password: '',
      snackbar: false,
      text: '',
      timeout: 2000,
    }
  }

  async createUser() {
    try {
      const resp = await this.$apollo.getClient().mutate({
        mutation: CREATE_ESTATE_AGENT_USER_MUTATION,
        variables: {
          input: {
            user: {
              username: this.$data.username,
              password: this.$data.password
            }
          }
        }
      });
      if (resp.data.createEstateAgentUser.ID) {
        this.$data.text = 'User ' + this.$data.username + ' create succesfully !';
        this.$data.snackbar = true;
      }
    } catch (e) {
      console.error(e);
    }
  }
}
</script>

<style>

</style>
