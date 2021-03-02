<template>
  <div>
    <h1>Create Tenant User account</h1>
    <v-text-field v-model="username" label="Username"></v-text-field>
    <v-text-field type="password" v-model="password" label="Password"></v-text-field>
    <v-btn depressed color="primary" v-on:click="createUser()">
      Create
    </v-btn>
  </div>
</template>

<script>
import Vue from "vue";
import Component from 'vue-class-component';
import gql from "graphql-tag";

const CREATE_TENANT_USER_MUTATION = gql`
mutation ($input: TenantInput) {
  createTenantUser(input: $input) {
    ID
  }
}
`;

@Component
export default class CreateTenantUser extends Vue {

  data() {
    return {
      username: '',
      password: ''
    }
  }

  async createUser() {
    try {
      const resp = await this.$apollo.getClient().mutate({
        mutation: CREATE_TENANT_USER_MUTATION,
        variables: {
          input: {
            user: {
              username: this.$data.username,
              password: this.$data.password
            }
          }
        }
      });
      if (resp.data.createTenantUser.ID) {
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
