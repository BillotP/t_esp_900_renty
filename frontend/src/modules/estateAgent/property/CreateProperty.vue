<template>
  <v-form v-model="valid">
    <h1>{{this.password}} = {{ this.repeatPassword}}</h1>
    <h1>{{passwordRules}}</h1>
    <v-container>
      <v-row>
        <h1>Create Property</h1>
      </v-row>
      <v-row>
        <v-col cols="12" md="4">
          <v-text-field
            v-model="firstname"
            :rules="nameRules"
            :counter="10"
            label="Estate name"
            required
          ></v-text-field>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" md="4">
          <v-text-field
            v-model="password"
            label="password"
            required
          ></v-text-field>
        </v-col>
        <v-col cols="12" md="4">
          <v-text-field
            v-model="repeatPassword"
            :rules="passwordRules"
            label="Repeat password"
            required
          ></v-text-field>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" md="4">
          <v-text-field
            v-model="email"
            :rules="emailRules"
            label="Estate E-mail"
            required
          ></v-text-field>
        </v-col>
      </v-row>

      <v-row>
        <v-select
          v-model="value"
          :items="items"
          chips
          label="Chips"
          multiple
          outlined
        ></v-select>
      </v-row>
    </v-container>
    <v-btn>qsdsqd</v-btn>
  </v-form>
</template>

<script>
import Vue from "vue";
import Component from "vue-class-component";

@Component
export default class CreateProperty extends Vue {
  firstname = "";
  password = "";
  repeatPassword = "";
  valid = false;
  nameRules = [
    (v) => !!v || "Name is required",
    (v) => v.length <= 10 || "Name must be less than 10 characters",
  ];

  email = "";
  emailRules = [
    (v) => !!v || "E-mail is required",
    (v) => /.+@.+/.test(v) || "E-mail must be valid",
  ];
  passwordRules = [
    (v) => !!v || "Password is required",
    this.password === this.repeatPassword || "Password must match",
  ];
  items = ["foo", "bar", "fizz", "buzz"];
  value = [];
}
</script>

<style>
</style>