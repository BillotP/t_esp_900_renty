<template>
  <div>
    <h1>Login</h1>
    <input type="radio" id="company" value="0" v-model="role">
    <label for="company">Company</label>
    <br>
    <input type="radio" id="estate-agent" value="1" v-model="role">
    <label for="estate-agent">Estate Agent</label>
    <br>
    <input type="radio" id="tenant" value="2" v-model="role">
    <label for="tenant">Tenant</label>
    <br>
    <v-text-field v-model="modelUsername" label="Username"></v-text-field>
    <v-text-field
        type="password"
        v-model="modelPassword"
        label="Password"
    ></v-text-field>
    <v-btn depressed color="primary" v-on:click="login()"> Login</v-btn>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import gql from "graphql-tag";
import {Login, privilege} from "@/store/authentificationStore";
import {Action, Getter, namespace} from "vuex-class";

const authModule = namespace("authentificationStore");
const LOGIN_COMPANY_MUTATION = gql`
  mutation($input: UserInput) {
    loginAsCompany(input: $input) {
      token
    }
  }
`;
const LOGIN_ESTATE_AGENT_MUTATION = gql`
  mutation($input: UserInput) {
    loginAsEstateAgent(input: $input) {
      token
    }
  }
`;
const LOGIN_TENANT_MUTATION = gql`
  mutation($input: UserInput) {
    loginAsTenant(input: $input) {
      token
    }
  }
`;

@Component
export default class Auth extends Vue {
  private username: string = "";
  private password: string = "";
  private role: number = -1;

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

  get modelRole(): number {
    return (+this.role) as number;
  }

  set modelRole(role: number) {
    this.role = (+role) as number;
  }

  @authModule.Action("login")
  private setInformationsLogin!: (log: Login) => void;

  async login() {
    const loginAs = [
      {
        mutation: LOGIN_COMPANY_MUTATION,
        key: 'loginAsCompany'
      },
      {
        mutation: LOGIN_ESTATE_AGENT_MUTATION,
        key: 'loginAsEstateAgent'
      },
      {
        mutation: LOGIN_TENANT_MUTATION,
        key: 'loginAsTenant'
      }
    ];

    try {
      const resp = await this.$apollo.getClient().mutate({
        mutation: loginAs[this.modelRole].mutation,
        variables: {
          input: {
            username: this.modelUsername,
            password: this.modelPassword,
          },
        },
      });
      debugger;
      if (resp.data[loginAs[this.modelRole].key].token) {
        //this.setInformationsLogin({ username: this.modelUsername, privilege: 0 });
        localStorage.setItem("username", this.modelUsername);
        localStorage.setItem("privilege", this.modelRole.toString());
        localStorage.setItem("token", resp.data[loginAs[this.modelRole].key].token);
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
