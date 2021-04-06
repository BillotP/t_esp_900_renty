<template>
  <v-main>
    <v-card elevation="5" style="margin: 5vh; padding: 10px">
      <v-carousel v-if="!selection" height="auto">
        <v-carousel-item v-if="!property.photos || !property.photos.length">
          <v-img :src="mockPicture">
            <v-chip-group
              v-model="selection"
              active-class="deep-purple accent-4 white--text"
              style="padding: 5px"
              column
            >
              <v-chip>Photos</v-chip>
              <v-chip :disabled="property.model == null"> 3D </v-chip>
            </v-chip-group>
          </v-img>
        </v-carousel-item>
        <v-carousel-item v-else :key="i" v-for="(photo, i) in property.photos">
          <v-img :src="url + photo.url">
            <v-chip-group
              v-model="selection"
              active-class="deep-purple accent-4 white--text"
              style="padding: 5px"
              column
            >
              <v-chip>Photos</v-chip>
              <v-chip :disabled="!property.model"> 3D </v-chip>
            </v-chip-group>
          </v-img>
        </v-carousel-item>
      </v-carousel>
      <div v-else>
        <model-obj :src="url + property.model.url"> </model-obj>
        <v-chip-group
          v-model="selection"
          active-class="deep-purple accent-4 white--text"
          column
        >
          <v-chip>Photos</v-chip>
          <v-chip :disabled="property.model == null"> 3D </v-chip>
        </v-chip-group>
      </div>
      <v-divider style="margin: 10px" />
      <v-row style="padding: 10px">
        <v-list-item-title
          >{{ propertyTypes[property.type] }} -
          {{ property.area }} m²</v-list-item-title
        >
        <v-list-item-title>
          📍 {{ property.address }} {{ property.postalCode + " " }}
          {{ property.cityName + " " + property.country }}
        </v-list-item-title>
        <v-list-item-title>
          🔎
          <a
            target="_blank"
            :href="
              'https://www.google.com/maps/search/?api=1&query=' +
              encodeURI(
                property.address +
                  ' ' +
                  property.postalCode +
                  ' ' +
                  property.cityName
              )
            "
          >
            View on Google Maps</a
          >
        </v-list-item-title>
      </v-row>
      <v-divider style="margin: 10px" />
      <v-col>
        <v-row align="center" justify="center">
          <v-col>
            <v-list-item-title>🛋️ Furnished</v-list-item-title>
          </v-col>
          <v-col>
            <strong
            :aria-label="property.furnished ? 'true' : 'false'"
            style="margin-left: 10px">{{
              property.furnished ? "✅" : "❎"
            }}</strong>
          </v-col>
        </v-row>

        <v-row align="center" justify="center">
          <v-col>
            <v-list-item-title>🚪 Rooms</v-list-item-title>
          </v-col>
          <v-col>
            <strong
              :style="'margin-left:' + (property.rooms >= 10 ? '8px' : '10px')"
              >{{ property.rooms }}</strong
            >
          </v-col>
        </v-row>

        <v-row align="center" justify="center">
          <v-col>
            <v-list-item-title>🛌 Bedrooms</v-list-item-title>
          </v-col>
          <v-col>
            <strong style="margin-left: 10px">{{ property.bedrooms }}</strong>
          </v-col>
        </v-row>

        <v-row
          align="center"
          justify="center"
          v-if="property && property.energyRating"
        >
          <v-col>
            <v-list-item-title>🌡️ Energy Rating</v-list-item-title>
          </v-col>
          <v-col>
            <v-progress-circular
              :value="
                100 -
                ((this.property.energyRating.charCodeAt(0) - 65) / 7) * 100
              "
              :color="energyColors[property.energyRating]"
            >
              <strong>{{ property.energyRating }}</strong>
            </v-progress-circular>
          </v-col>
        </v-row>
      </v-col>
      <v-divider style="margin: 10px" />
      <v-row dense>
        <v-col dense :key="i" v-for="(badge, i) in property.badges">
          <v-card elevation="10">
            <v-col align="center" justify="center">
              <v-icon x-large color="primary">
                {{ badgeIcons[badge.type] }}
              </v-icon>
            </v-col>
            <v-col>
              <p style="text-align: center">{{ badge.type }}</p>
            </v-col>
          </v-card>
        </v-col>
      </v-row>

      <v-row>
        <p style="padding: 10px">{{ property.description }}</p>
      </v-row>
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
      mockPicture:
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
        Maison: "🏠 Independent House",
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
