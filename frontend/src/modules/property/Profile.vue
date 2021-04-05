<template>
  <v-main>
    <v-card elevation="3" style="margin: 5vh; padding: 10px" tile>
      <v-chip-group
        v-model="selection"
        active-class="deep-purple accent-4 white--text"
        column
      >
        <v-chip>Photos</v-chip>

        <v-tooltip top>
          <template v-slot:activator="{ on, attrs }">
            <v-chip v-bind="attrs" v-on="on" :disabled="property.model == null">
              3D
            </v-chip>
          </template>
          <span>3D visual not yet available</span>
        </v-tooltip>
      </v-chip-group>
      <v-carousel v-if="!selection" height="auto">
        <v-carousel-item :key="i" v-for="(photo, i) in property.photos">
          <v-img :src="url + photo.url"></v-img>
        </v-carousel-item>
        <v-carousel-item v-if="!property.photos || !property.photos.length">
          <v-img
            src="https://image.freepik.com/photos-gratuite/lay-plat-concept-immobilier_53876-14502.jpg"
          ></v-img>
        </v-carousel-item>
      </v-carousel>
      <model-obj v-else :src="url + property.model.url"></model-obj>
      <v-card-text>
        <v-row v-if="property && property.energyRating">
          <v-col cols="12" md="6">
            <v-list-item-title>{{
              propertyTypes[property.type]
            }}</v-list-item-title>
            <v-list-item-subtitle> {{ property.address }}</v-list-item-subtitle>
            <v-list-item-subtitle
              >{{ property.postalCode + " " }}
              <strong>{{ property.cityName }}</strong></v-list-item-subtitle
            >
            <v-list-item-subtitle
              ><u>{{ property.country }}</u></v-list-item-subtitle
            >
            <span
              ><strong>{{ property.area }} m²</strong></span
            >
          </v-col>
          <v-col align-self="center" cols="12" md="6">
            <v-row>
              <span style="margin: 5px">Energy Rating :</span>
            </v-row>
            <v-row>
              <v-progress-circular
                :value="
                  100 -
                  ((this.property.energyRating.charCodeAt(0) - 65) / 7) * 100
                "
                :color="energyColors[property.energyRating]"
                size="50"
                style="margin: 5px"
              >
                <strong>{{ property.energyRating }}</strong>
              </v-progress-circular>
            </v-row>
          </v-col>
        </v-row>
        <v-row>
          <v-col :key="i" v-for="(badge, i) in property.badges">
            <v-tooltip bottom>
              <template v-slot:activator="{ on, attrs }">
                <v-icon v-bind="attrs" v-on="on" x-large color="primary">
                  {{ badgeIcons[badge.type] }}
                </v-icon>
              </template>
              <span>{{ badge.type }}</span>
            </v-tooltip>
          </v-col>
        </v-row>
        <v-row>
          <div class="my-4 subtitle-1">
            • Rent: <strong>{{ property.rentAmount }}</strong
            >€ / Monthly<br />
            • Charges: <strong>{{ property.chargesAmount }}</strong
            >€ / Monthly
          </div>
        </v-row>
        <v-row>
          <p style="padding: 10px">{{ property.description }}</p>
        </v-row>
      </v-card-text>
    </v-card>
  </v-main>
</template>


<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import gql from "graphql-tag";

const PROPERTY_QUERY = gql`
  query property($id: Int!) {
    property(id: $id) {
      ID
      area
      country
      cityName
      address
      postalCode
      type
      photos {
        url
      }
      model {
        url
      }
      badges {
        type
      }
      description
      rooms
      bedrooms
      furnished
      constructionDate
      energyRating
      rentAmount
      chargesAmount
    }
  }
`;

@Component
export default class PropertyProfile extends Vue {
  public property: any = {};

  data() {
    return {
      banneer:
        "https://image.freepik.com/photos-gratuite/lay-plat-concept-immobilier_53876-14502.jpg",
      selection: 0,
      badgeIcons: {
        Garden: "mdi-shovel",
        Fireplace: "mdi-fireplace",
        Caretaker: "mdi-account-child",
        GreatView: "mdi-image-filter-hdr",
        Balcony: "mdi-warehouse",
        SwimmingPool: "mdi-swim",
        Lift: "mdi-elevator-passenger",
        Terrace: "mdi-account-child",
        Garage: "mdi-garage-variant",
        Orientation: "mdi-compass",
      },
      energyColors: {
        A: "rgb(000,128,000)",
        B: "rgb(000,255,000)",
        C: "rgb(128,255,000)",
        D: "rgb(255,255,000)",
        E: "rgb(255,128,000)",
        F: "rgb(255,128,064)",
        G: "rgb(255,000,000)",
      },
      url: process.env.VUE_APP_GRAPHQL_HTTP,
      propertyTypes: {
        Maison: "🏠 House",
        Appartement: "🏙️ Flat",
      },
    };
  }

  beforeMount() {
    this.$apollo
      .getClient()
      .query({
        query: PROPERTY_QUERY,
        variables: { id: this.$route.params.id },
      })
      .then((res) => {
        this.property = res.data.property;
        console.log(this.property);
      })
      .catch((err) => {
        console.error(err);
      });
  }
}
</script>

<style>
.v-card {
  margin-top: 4rem;
}
</style>
