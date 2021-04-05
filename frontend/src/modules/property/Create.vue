<template>
  <v-form v-model="valid">
    <v-container style="margin-top: 5vh; margin-bottom: 5vh;">
      <v-card
          class="mx-auto"
          style="padding: 10px"
          max-width="50vw"
          outlined
      >
        <v-form v-model="valid">
          <h4>Type</h4>
          <v-row flex>
            <v-col cols="12" md="6">
              <v-select
                  v-model="typeSelected"
                  :items="propertyTypes"
                  item-text="label"
                  item-value="value"
                  label="Property kind"
                  hint="Select the real estate property kind (flat or house)"
                  single-line
                  required
                  :rules="rules"
              ></v-select>
            </v-col>
            <v-col cols="12" md="6">
              <v-switch
                  v-model="furnished"
                  label="Furnished"
              ></v-switch>
            </v-col>
          </v-row>
          <h4>Infos</h4>
          <v-row>
            <v-col cols="12" md="6">
              <v-select
                  v-model="badgesSelected"
                  :items="propertyBadges"
                  item-text="label"
                  item-value="value"
                  label="Property badges"
                  multiple
              ></v-select>
            </v-col>
            <v-col cols="12" md="6">
              <v-text-field
                  v-model="area"
                  label="Area size (m²)"
                  required
                  :rules="rules"
                  type="number"
              ></v-text-field>
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12" md="6">
              <v-text-field
                  v-model="rooms"
                  label="Number of rooms"
                  required
                  :rules="rules"
                  type="number"
              ></v-text-field>
            </v-col>
            <v-col cols="12" md="6">
              <v-text-field
                  v-model="bedrooms"
                  label="Number of bedrooms"
                  required
                  :rules="rules"
                  type="number"
              ></v-text-field>
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12" md="6">
              <v-text-field
                  v-model="rent"
                  label="Rent (€)"
                  required
                  :rules="rules"
                  type="number"
              ></v-text-field>
            </v-col>
            <v-col cols="12" md="6">
              <v-text-field
                  v-model="charges"
                  label="Charges (€)"
                  required
                  :rules="rules"
                  type="number"
              ></v-text-field>
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12" md="6">
              <v-select
                  v-model="energy"
                  :items="energyTypes"
                  item-text="label"
                  item-value="value"
                  label="Energy types"
                  required
                  :rules="rules"
                  single-line
              ></v-select>
            </v-col>
            <v-col cols="12" md="6">
              <v-text-field
                  v-model="construction"
                  label="Construction date"
                  type="date"
              ></v-text-field>
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12" md="6">
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
            <v-col cols="12" md="6">
              <v-text-field
                  v-model="codeNumber"
                  label="ZipCode"
                  required
                  :rules="rules"
                  disabled
              ></v-text-field>
            </v-col>
          </v-row>
          <h4>View</h4>
          <v-row>
            <v-col cols="12" md="6">
              <v-file-input
                  v-model="photos"
                  accept="image/*"
                  label="Photos"
                  multiple
              ></v-file-input>
            </v-col>
            <v-col cols="12" md="6">
              <v-file-input
                  v-model="model"
                  accept="application/object"
                  label="Model"
              ></v-file-input>
            </v-col>
          </v-row>
          <v-row>
            <v-col
                cols="12"
                md="12"
            >
              <v-textarea
                  v-model="description"
                  label="Description"
              ></v-textarea>
            </v-col>
          </v-row>
          <v-row>
            <v-col>
              <v-hover>
                <v-btn outlined @click="CreateEstateAgent" :disabled="!valid"
                >Create
                </v-btn
                >
              </v-hover>
            </v-col>
          </v-row>
        </v-form>
      </v-card>
    </v-container>
  </v-form>
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
  components: {VueGoogleAutocomplete},
  data() {
    return {
      area: null,
      address: null,
      addressSearch: null,
      country: null,
      cityName: null,
      codeNumber: null,
      description: null,
      propertyTypes: [
        {label: "🏠 House", value: "Maison"},
        {label: "🏙️ Apartment", value: "Appartement"},
      ],
      propertyBadges: [
        {label: "🌿 Garden", value: "Garden"},
        {label: "🔥 Fireplace", value: "Fireplace"},
        {label: "💂 Caretaker", value: "Caretaker"},
        {label: "🌆️ Great view", value: "GreatView"},
        {label: "🏙️ Balcony", value: "Balcony"},
        {label: "🏊 Swimming pool", value: "SwimmingPool"},
        {label: "🚠 Lift", value: "Lift"},
        {label: "🏟️ Terrace", value: "Terrace"},
        {label: "🚘️ Garage", value: "Garage"},
        {label: "🧭 Orientation", value: "Orientation"},
      ],
      energy: null,
      energyTypes: [
        {label: "🟢 A", value: "A"},
        {label: "🟡 B", value: "B"},
        {label: "🟠 C", value: "C"},
        {label: "🔴 D", value: "D"},
        {label: "🟣 E", value: "E"},
        {label: "🟤 F", value: "F"},
        {label: "⚫ G", value: "G"},
      ],
      photos: [],
      model: null,
      rooms: 0,
      bedrooms: 0,
      furnished: false,
      construction: null,
      rent: 0,
      charges: 0,
      badgesSelected: [],
      typeSelected: null,
      valid: false,
      rules: [
          v => !!v || 'field is required',
      ],
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
              country: this.$data.country,
              cityName: this.$data.cityName,
              address: this.$data.address,
              postalCode: this.$data.codeNumber,
              type: this.$data.typeSelected,
              photos: this.$data.photos,
              model: this.$data.model,
              badges: this.$data.badgesSelected,
              description: this.$data.description,
              rooms: this.$data.rooms,
              bedrooms: this.$data.bedrooms,
              furnished: this.$data.furnished,
              constructionDate: new Date(this.$data.construction).toISOString(),
              energyRating: this.$data.energy,
              rentAmount: this.$data.rent,
              chargesAmount: this.$data.charges
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
      this.cityName = `${addressData.locality}`
      this.country = `${addressData.country}`
      this.address = `${addressData.street_number} ${addressData.route}`;
      this.codeNumber = `${addressData.postal_code}`;
    },
  },
};
</script>

<style>
</style>
