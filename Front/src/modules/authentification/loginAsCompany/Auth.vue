<template>
  <div>
    <h1>Login as company</h1>
    <v-text-field v-model="modelUsername" label="Username"></v-text-field>
    <v-text-field
      type="password"
      v-model="modelPassword"
      label="Password"
    ></v-text-field>
    <v-btn depressed color="primary" v-on:click="login()"> Login </v-btn>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import gql from "graphql-tag";
import { Login, privilege } from "@/store/authentificationStore";
import { Action, Getter, namespace } from "vuex-class";

const authModule = namespace("authentificationStore");
const LOGIN_COMPANY_MUTATION = gql`
  mutation($input: UserInput) {
    loginAsCompany(input: $input) {
      token
    }
  }
`;

@Component
export default class Auth extends Vue {
  private username: string = "";
  private password: string = "";

  get modelUsername() {
    return this.username;
  }

  set modelUsername(username: string) {
    this.username = username;
  }

  get modelPassword() {
    return this.password;
  }

  set modelPassword(password: string) {
    this.password = password;
  }

  @authModule.Action("login")
  private setInformationsLogin!: (log: Login) => void;

  async login() {
    try {
      const resp = await this.$apollo.getClient().mutate({
        mutation: LOGIN_COMPANY_MUTATION,
        variables: {
          input: {
            username: this.modelUsername,
            password: this.modelPassword,
          },
        },
      });
      if (resp.data.loginAsCompany.token) {
        //this.setInformationsLogin({ username: this.modelUsername, privilege: 0 });
        localStorage.setItem("username", this.modelUsername);
        localStorage.setItem("privilege", "0");
        localStorage.setItem("token", resp.data.loginAsCompany.token);
        this.$router.push("/");
      }
    } catch (e) {
      console.error(e);
    }
  }
}
</script>

<style>
</style>
