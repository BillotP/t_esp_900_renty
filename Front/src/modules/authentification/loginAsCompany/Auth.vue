<template>
  <div>

    <h1>Login as company</h1>
    <v-text-field v-model="username" label="Username"></v-text-field>
    <v-text-field type="password" v-model="password" label="Password"></v-text-field>
    <v-btn depressed color="primary" v-on:click="login()">
      Login
    </v-btn>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from 'vue-class-component';
import gql from "graphql-tag";

const LOGIN_COMPANY_MUTATION = gql`
mutation ($input: UserInput) {
  loginAsCompany(input: $input) {
    token
  }
}
`;

@Component
export default class Auth extends Vue {

  data() {
    return {
      username: '',
      password: ''
    }
  }

  async login() {
    try {
      const resp = await this.$apollo.getClient().mutate({
        mutation: LOGIN_COMPANY_MUTATION,
        variables: {
          input: {
            username: this.$data.username,
            password: this.$data.password
          }
        }
      });
      if (resp.data.loginAsCompany.token) {
        localStorage.setItem('token', resp.data.loginAsCompany.token);
        this.$router.push('/');
      }
    } catch (e) {
      console.error(e);
    }
  }
}
</script>

<style>

</style>
