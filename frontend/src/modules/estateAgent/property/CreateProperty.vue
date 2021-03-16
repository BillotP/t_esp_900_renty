<template>
  <v-form v-model="valid">
    <v-container>
      <v-row>
        <v-col cols="12" md="4">
          <v-select
            v-model="typeSelected"
            :items="type"
            label="Property"
            single-line
            required
          ></v-select>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" md="4">
          <v-text-field
            v-model="address"
            append-outer-icon="mdi-map"
            menu-props="auto"
            hide-details
            label="Address"
            single-line
            required
          ></v-text-field>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" md="4">
          <v-text-field
            v-model="area"
            label="Area size (m²)"
            required
            type="number"
          ></v-text-field>
        </v-col>
        <v-col cols="12" md="4">
          <v-text-field
            v-model="codeNumber"
            label="ZipCode"
            required
            type="number"
          ></v-text-field>
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <v-hover>
            <v-btn outlined @click="CreateEstateAgent" :disabled="!valid"
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
import gql from "graphql-tag";

const CREATE_PROPERTY_MUTATION = gql`
  mutation($input: PropertyInput) {
    createProperty(input: $input) {
      ID
    }
  }
`;

@Component
export default class CreateProperty extends Vue {
  data() {
    return {
      area: 0.0,
      address: null,
      codeNumber: 0,
      type: ["Maison", "Appartement"],
      typeSelected: null,
      valid: false,
      timeout: 2000,
    };
  }

  async CreateEstateAgent() {
    try {
      const resp = await this.$apollo.getClient().mutate({
        mutation: CREATE_PROPERTY_MUTATION,
        variables: {
          input: {
            area: this.$data.area,
            address: this.$data.address,
            codeNumber: this.$data.codeNumber,
            type: this.$data.typeSelected,
          },
        },
      });
      if (resp.data.createProperty) {
        console.log("coucou");
      }
    } catch (e) {
      console.error(e);
    }
  }
}
</script>

<style>
</style>