<template>
  <v-form v-model="valid">
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
            :error-messages="
              this.password === this.repeatPassword ? '' : 'Password must match'
            "
            :rules="passwordRules"
            label="password"
            required
          ></v-text-field>
        </v-col>
        <v-col cols="12" md="4">
          <v-text-field
            v-model="repeatPassword"
            label="Repeat password"
            :rules="passwordRules"
            required
          ></v-text-field>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" md="4">
          <v-text-field
            v-model="city"
            append-outer-icon="mdi-map"
            menu-props="auto"
            hide-details
            label="Estate City"
            single-line
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
        <v-col cols="12" md="4">
          <v-select
            v-model="value"
            :items="items"
            chips
            label="Chips"
            multiple
            outlined
            solo
            dense
          ></v-select>
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <v-hover>
            <v-btn outlined @click="CreateEstateAgent" :disabled="!(password === repeatPassword && password.length > 0)">Create</v-btn>
          </v-hover>
        </v-col>
      </v-row>
    </v-container>
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
  city = "";
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
  items = ["foo", "bar", "fizz", "buzz"];
  value = [];
  passwordRules = [(v) => !!v || "Password is required"];

  CreateEstateAgent()
  {
    console.debug(this.firstname, "  " + this.password);
  }

}

</script>

<style>
</style>