<template>
  <v-form v-model="valid">
    <v-container>
      <v-row>
        <h1>Create Property</h1>
      </v-row>
      <v-row>
        <v-col cols="12" md="4">
          <v-text-field
            v-model="name"
            :rules="nameRules"
            label="Name"
            required
          ></v-text-field>
        </v-col>
        <v-col cols="12" md="4">
          <v-text-field
            v-model="city"
            append-outer-icon="mdi-map"
            menu-props="auto"
            hide-details
            label="City"
            single-line
          ></v-text-field>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" md="4">
          <v-text-field
            v-model="email"
            :rules="emailRules"
            label="E-mail"
            type="email"
            required
          ></v-text-field>
        </v-col>
        <v-col cols="12" md="4">
          <v-text-field
            v-model="repeatEmail"
            :rules="emailRules"
            label="Repeat E-mail"
            type="email"
            required
          ></v-text-field>
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <v-hover>
            <v-btn
              outlined
              @click="CreateEstateAgent"
              :disabled="!(email === repeatEmail && email.length > 0)"
              >Create</v-btn
            >
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
  name = "";
  repeatEmail = "";
  city = "";
  valid = false;
  nameRules = [(v) => !!v || "Name is required"];
  email = "";
  emailRules = [
    (v) => !!v || "E-mail is required",
    (v) => /.+@.+/.test(v) || "E-mail must be valid",
  ];

  CreateEstateAgent() {
    console.debug(this.firstname, "  " + this.password);
  }
}
</script>

<style>
</style>