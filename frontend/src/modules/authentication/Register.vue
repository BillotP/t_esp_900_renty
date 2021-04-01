<template>
  <v-main
    style="
      margin-top: 10vh;
      margin-bottom: 10vh;
      margin-right: 10vh;
      margin-left: 10vh;
    "
  >
    <v-card max-width="80vw" max-height="71vh">
      <v-card-text>
        <h3>My Company</h3>
        <v-container fluid>
          <v-row>
            <v-text-field
              v-model="companyName"
              label="Company Name"
            ></v-text-field>
          </v-row>
          <v-row>
            <v-text-field
              v-model="companyDescription"
              label="Company Description"
            ></v-text-field>
          </v-row>
          <v-row>
            <v-text-field
              type="phone"
              v-model="companyPhone"
              label="Company Phone Number"
            ></v-text-field>
          </v-row>
          <v-row>
            <v-img
              style="margin-right: 10px; border: 1px solid black"
              max-height="64px"
              max-width="64px"
              :src="companyLogoUrl"
              left
            ></v-img>
            <v-text-field
              type="url"
              v-model="companyLogoUrl"
              label="Company Logo URL"
            ></v-text-field>
          </v-row>
        </v-container>
        <div style="margin-top: 20px" />
        <h3>My Profile</h3>
        <v-container fluid>
          <v-row>
            <v-text-field
              v-model="companyUserName"
              label="Username"
              mandatory
            ></v-text-field>
          </v-row>
          <v-row>
            <v-text-field
              type="password"
              v-model="companyUserPwd"
              label="Password"
            ></v-text-field>
          </v-row>
          <v-row>
            <v-radio-group row v-model="role" mandatory>
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
                disabled
              ></v-radio>
              <v-radio
                color="primary"
                label="Tenant"
                value="2"
                key="2"
                disabled
              ></v-radio>
            </v-radio-group>
          </v-row>
          <v-btn
            depressed
            block
            color="primary"
            :disabled="isLoading"
            v-on:click="register()"
          >
            {{ isLoading ? "" : "Register*" }}</v-btn
          >
          <p style="margin-top: 10px">
            *By registering, you automaticaly accept our
            <router-link to="/tos">terms and conditions</router-link>
          </p>
        </v-container>
        <p>
          Already have an account ?
          <router-link to="/login">signin</router-link>
        </p></v-card-text
      >
    </v-card>
  </v-main>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import gql from "graphql-tag";

const REGISTER_AS_COMPANY = gql`
  mutation {
    signupAsCompany(
      input: {
        name: $companyName
        user: { username: $companyUserName, password: $companyUserPwd }
        logo: $companyLogoUrl
        description: $companyDescription
        tel: $companyPhone
      }
    ) {
      ID
      token
    }
  }
`;

@Component
export default class Register extends Vue {
  public companyName: string = "";
  public companyDescription: string = "";
  public companyPhone: string = "";
  public companyLogoUrl: string | null = null;
  public companyUserName: string = "";
  public companyUserPwd: string = "";
  public role: number = 0;
  public isLoading: boolean = false;
  public hasError: boolean = false;

  async register() {
    const registerAs = [
      {
        mutation: REGISTER_AS_COMPANY,
        key: "registerAsCompany",
      },
    ];
    this.hasError = false;
    this.isLoading = true;
    try {
      const resp = await this.$apollo.getClient().mutate({
        mutation: registerAs[this.role].mutation,
        variables: {
          input: {
            companyName: this.companyName,
            companyLogoUrl: this.companyLogoUrl,
            companyUserName: this.companyUserName,
            companyUserPwd: this.companyUserPwd,
          },
        },
      });
      if (resp.data[registerAs[this.role].key].token) {
        localStorage.setItem("username", this.companyUserName);
        localStorage.setItem("privilege", this.role.toString());
        localStorage.setItem(
          "token",
          resp.data[registerAs[this.role].key].token
        );
        localStorage.setItem("id", resp.data[registerAs[this.role].key].ID);
        this.$router.push("/");
      }
    } catch (e) {
      console.error(e);
      this.hasError = true;
    }
    this.isLoading = false;
  }
}
</script>

<style>
</style>
