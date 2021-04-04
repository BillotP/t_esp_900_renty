<template>
  <v-main
    style="
      margin-top: 10vh;
      margin-bottom: 10vh;
      margin-right: 10vh;
      margin-left: 10vh;
    "
  >
    <v-card>
      <v-card-text>
        <v-container fluid>
          <v-row>
            <v-text-field
              v-model="modelUsername"
              label="Username"
              :error-messages="userErrorText"
            ></v-text-field>
          </v-row>
          <v-row>
            <v-text-field
              type="password"
              v-model="modelPassword"
              label="Password"
              :error-messages="passwordErrorText"
            ></v-text-field>
          </v-row>
          <v-row>
            <v-radio-group :row="$vuetify.breakpoint.mobile"  v-model="role" mandatory>
              <v-radio
                color="primary"
                label="Company"
                value="0"
                key="0"
              ></v-radio>
              <v-radio
                color="primary"
                label="Estate Agent"
                value="1"
                key="1"
              ></v-radio>
              <v-radio
                color="primary"
                label="Tenant"
                value="2"
                key="2"
              ></v-radio>
            </v-radio-group>
          </v-row>
          <v-btn
            :disabled="isLoading"
            depressed
            block
            color="primary"
            v-on:click="login()"
          >
            <span v-if="!isLoading">Login</span>
            <v-progress-circular
              indeterminate
              v-else
              color="primary"
            ></v-progress-circular>
          </v-btn>
        </v-container>
        <p>
          Dont have an account yet ?
          <router-link to="/register">register</router-link>
        </p>
      </v-card-text>
    </v-card>
    <v-snackbar v-model="snackbar" color="red" timeout="3000">
      {{ errorText }}
      <template v-slot:action="{ attrs }">
        <v-btn color="blue" text v-bind="attrs" @click="snackbar = false">
          Close
        </v-btn>
      </template>
    </v-snackbar>
  </v-main>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import gql from "graphql-tag";
import { GraphQLError } from "graphql";

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

@Component
export default class Auth extends Vue {
  public isLoading: boolean = false;
  public hasError: boolean = false;
  public snackbar: boolean = false;
  public passwordErrorText: string = "";
  public userErrorText: string = "";
  public hasOtherError: boolean = false;
  public errorText: string = "";

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
    return +this.role as number;
  }

  set modelRole(role: number) {
    this.role = +role as number;
  }
  private parseJwt(token: string): Object {
    var base64Url = token.split(".")[1];
    var base64 = base64Url.replace(/-/g, "+").replace(/_/g, "/");
    var jsonPayload = decodeURIComponent(
      atob(base64)
        .split("")
        .map(function (c) {
          return "%" + ("00" + c.charCodeAt(0).toString(16)).slice(-2);
        })
        .join("")
    );

    return JSON.parse(jsonPayload);
  }

  async login() {
    this.isLoading = true;
    this.hasError = false;
    this.passwordErrorText = "";
    this.userErrorText = "";
    this.errorText = "";
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
      if (resp.data[loginAs[this.modelRole].key].token) {
        const payload = this.parseJwt(
          resp.data[loginAs[this.modelRole].key].token
        );
        localStorage.setItem("exp", payload["exp"]);
        localStorage.setItem("username", this.modelUsername);
        localStorage.setItem("privilege", this.modelRole.toString());
        localStorage.setItem(
          "token",
          resp.data[loginAs[this.modelRole].key].token
        );
        localStorage.setItem("id", resp.data[loginAs[this.modelRole].key].ID);
        this.$router.push("/dashboard");
      }
    } catch (e) {
      this.hasError = true;
      let gqlerror = e as GraphQLError;
      let errMsq = gqlerror.message.replace("GraphQL error: ", "");
      if (errMsq.includes("record not found")) {
        this.userErrorText = "Unknow user, please register first";
      } else if (errMsq.includes("bad password")) {
        this.passwordErrorText = "Wrong password, try again";
      } else {
        this.snackbar = true;
        this.errorText = "Something wrong happen, please try again later";
      }
    }
    this.isLoading = false;
  }
}
</script>

<style>
</style>
