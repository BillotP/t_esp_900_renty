<template>
  <v-container style="margin-top: 20vh; margin-bottom: 20vh">
    <v-card
      class="mx-auto"
      style="padding: 10px"
      max-height="40vh"
      max-width="50vw"
      outlined
    >
      <v-form v-model="valid">
        <h4>Type</h4>
        <v-row flex>
          <v-col cols="12" md="4">
            <v-select
              v-model="typeSelected"
              :items="propertyTypes"
              item-text="label"
              item-value="value"
              label="Property kind"
              hint="Select the real estate property kind (flat or house)"
              single-line
              required
            ></v-select>
          </v-col>
          <v-col cols="12" md="4">
            <v-text-field
              v-model="area"
              label="Area size (m²)"
              required
              type="number"
            ></v-text-field>
          </v-col>
        </v-row>
        <h4>Address</h4>
        <v-row>
          <v-col cols="12" md="4">
            <vue-google-autocomplete
              ref="addressSearch"
              id="map"
              style="padding: 10px"
              placeholder="Please type your address"
              v-on:placechanged="getAddressData"
              country="fr"
            >
            </vue-google-autocomplete>
          </v-col>
          <v-col cols="12" md="4">
            <v-text-field
              v-model="codeNumber"
              label="ZipCode"
              required
              disabled
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
      </v-form>
    </v-card>
  </v-container>
</template>

<script>
import VueGoogleAutocomplete from "vue-google-autocomplete";
import gql from "graphql-tag";

const CREATE_PROPERTY_MUTATION = gql`
  mutation($input: PropertyInput) {
    createProperty(input: $input) {
      ID
    }
  }
`;

export default {
  components: { VueGoogleAutocomplete },
  data() {
    return {
      area: null,
      address: null,
      addressSearch: null,
      codeNumber: 0,
      propertyTypes: [
        { label: "🏠 House", value: "Maison" },
        { label: "🏙️ Apartment", value: "Appartement" },
      ],
      typeSelected: null,
      valid: false,
      timeout: 2000,
    };
  },
  methods: {
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
          this.$router.back();
        }
      } catch (e) {
        console.error(e);
      }
    },
    getAddressData: function (addressData, _placeResultData, _id) {
      console.log(addressData);
      this.address = `${addressData.street_number} ${addressData.route}`;
      this.codeNumber = `${addressData.postal_code}`;
    },
  },
};
</script>

<style>
</style>