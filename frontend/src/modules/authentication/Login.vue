<template>
  <v-card flat>
    <v-card-title>Login</v-card-title>
    <v-card-text>
      <v-container fluid>
        <v-row>
          <v-radio-group v-model="role" mandatory>
            <v-radio color="primary" label="Company" value="0" key="0"></v-radio>
            <v-radio color="primary" label="Estate Agent" value="1" key="1"></v-radio>
            <v-radio color="primary" label="Tenant" value="2" key="2"></v-radio>
          </v-radio-group>
        </v-row>
        <v-divider></v-divider>
        <v-row>
          <v-text-field v-model="modelUsername" label="Username"></v-text-field>
        </v-row>
        <v-row>
          <v-text-field
              type="password"
              v-model="modelPassword"
              label="Password"
          ></v-text-field>
        </v-row>
        <v-row>
          <v-btn depressed color="primary" v-on:click="login()"> Login</v-btn>
        </v-row>
      </v-container>
    </v-card-text>
  </v-card>
  <!-- <v-form v-model="valid">
    <v-container>
      <h1>Login</h1>
      <v-row>
        <v-col cols="12" md="4">
          <v-row>
            <v-col cols="12" md="4">
              <input
                type="radio"
                id="company"
                value="0"
                v-model="roleSelected"
              />
              <label for="company">Company</label>
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12" md="4">
              <input
                type="radio"
                id="estate-agent"
                value="1"
                v-model="roleSelected"
              />
              <label for="estate-agent">Estate Agent</label>
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12" md="4">
              <input
                type="radio"
                id="tenant"
                value="2"
                v-model="roleSelected"
              />
              <label for="tenant">Tenant</label>
            </v-col>
          </v-row>
        </v-col>
      </v-row>
      <v-text-field v-model="modelUsername" label="Username"></v-text-field>
      <v-text-field
        type="password"
        v-model="modelPassword"
        label="Password"
      ></v-text-field>
      <v-btn depressed color="primary" v-on:click="login()"> Login</v-btn>
    </v-container>
  </v-form> -->
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

  data() {
    return {
      type: ["Company", "Estate-agent", "Tenant"],
      roleSelected: 0,
      valid: false,
    };
  }

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
    const loginAs = [
      {
        mutation: LOGIN_COMPANY_MUTATION,
        key: "loginAsCompany",
      },
      {
        mutation: LOGIN_ESTATE_AGENT_MUTATION,
        key: "loginAsEstateAgent",
      },
      {
        mutation: LOGIN_TENANT_MUTATION,
        key: "loginAsTenant",
      },
    ];

    try {
      const resp = await this.$apollo.getClient().mutate({
        mutation: loginAs[this.$data.roleSelected].mutation,
        variables: {
          input: {
            username: this.modelUsername,
            password: this.modelPassword,
          },
        },
      });
<<<<<<< HEAD
      if (resp.data[loginAs[this.modelRole].key].token) {
        //this.setInformationsLogin({ username: this.modelUsername, privilege: 0 });
        localStorage.setItem("username", this.modelUsername);
        localStorage.setItem("privilege", this.modelRole.toString());
        localStorage.setItem("token", resp.data[loginAs[this.modelRole].key].token);
        localStorage.setItem("id", resp.data[loginAs[this.modelRole].key].ID);
=======
      if (resp.data[loginAs[this.$data.roleSelected].key].token) {
        //this.setInformationsLogin({ username: this.modelUsername, privilege: 0 });
        localStorage.setItem("username", this.modelUsername);
        localStorage.setItem("privilege", this.$data.roleSelected.toString());
        localStorage.setItem(
          "token",
          resp.data[loginAs[this.$data.roleSelected].key].token
        );
>>>>>>> 2c9126d (create ticket + responsive #WIP)
        this.$router.push("/");
      }
    } catch (e) {
      console.error(e);
    }
  }

  public mounted()
  {
    if (localStorage.getItem('token')) {
        this.$router.push("/");
    }
  }

}
</script>

<style>
</style>
