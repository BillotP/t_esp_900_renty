<template>
  <v-main>
    <v-card v-if="this.$apollo.queries.estateAgent.loading" height="200">
      <v-layout align-center justify-center column fill-height>
        <v-flex row align-center>
          <v-progress-circular
              indeterminate
              :size="50"
              color="primary"
              class=""
          ></v-progress-circular>
        </v-flex>
      </v-layout>
    </v-card>
    <v-card v-else style="margin: 10px">
      <v-card-title class="justify-center"
      ><b style="text-transform: capitalize"
      >{{ estateAgent.user.username }} @
        {{ estateAgent.company.name }} 🖊️</b
      >
      </v-card-title>
      <v-card-text>
        <v-row>
          <v-text-field
              type="phone"
              v-model="estateAgent.tel"
              label="Phone Number"
              required
              :rules="rules"
          ></v-text-field>
        </v-row>
        <v-row>
          <v-col cols="12" md="6">
            <v-select
                v-model="skillsSelected"
                :items="skillTypes"
                item-text="label"
                item-value="value"
                label="Skills"
                multiple
            ></v-select>
          </v-col>
          <v-col cols="12" md="6">
            <v-select
                v-model="specialitiesSelected"
                :items="specialityTypes"
                item-text="label"
                item-value="value"
                label="Specialities"
                multiple
            ></v-select>
          </v-col>
        </v-row>
        <v-row>
          <v-textarea
              hint="Let me know you"
              v-model="estateAgent.about"
              label="About"
          ></v-textarea>
        </v-row>
        <v-row>
          <v-text-field
              hint="Update your estateAgent name"
              v-model="estateAgent.user.username"
              label="Name"
          ></v-text-field>
        </v-row>
        <v-row>
          <v-text-field
              hint="Update your estateAgent password"
              v-model="userPassword"
              label="New password"
              type="password"
          ></v-text-field>
        </v-row>
        <v-btn
            depressed
            block
            color="primary"
            style="margin: 10px"
            :disabled="isLoading"
            v-on:click="update()"
        >
          {{ isLoading ? "" : "Update" }}
        </v-btn
        >
      </v-card-text>
    </v-card>
  </v-main
  >
</template>

<script>
import gql from "graphql-tag";

export default {
  apollo: {
    estateAgent: {
      query: gql`
        query estateAgent($id: Int!) {
          estateAgent(id: $id) {
            ID
            company {
              name
            }
            tel
            about
            skills {
              type
            }
            specialities {
              type
            }
            user {
              ID
              createdAt
              updatedAt
              username
              password
              role
            }
          }
        }
      `,
      variables() {
        return {id: this.$route.params.id};
      },
      pollInterval: 2000,
    },
  },
  mounted() {
    this.$apollo.queries.estateAgent.observer.subscribe(() => {
      for (const skill of this.estateAgent.skills) {
        this.$data.skillsSelected.push(skill.type);
      }
      for (const speciality of this.estateAgent.specialities) {
        this.$data.specialitiesSelected.push(speciality.type);
      }
    });
  },
  methods: {
    udpate() {
      console.log("Update !");
    },
  },
  data() {
    return {
      estateAgent: {},
      userPassword: "",
      isLoading: false,
      estateAgentID: this.$route.params.id,
      skillsSelected: [],
      skillTypes: [
        {label: "🏴󠁧󠁢󠁥󠁮󠁧󠁿 English", value: "ENGLISH"},
        {label: "🇪🇸 Spanish", value: "SPANISH"},
        {label: "🇩🇪 German", value: "GERMAN"},
        {label: "🇫🇷 French", value: "FRENCH"},
        {label: "⚙ Software", value: "SOFTWARE"},
        {label: "👔 Hard working", value: "HARD_WORKING"},
        {label: "🧑‍💻 Remote working", value: "REMOTE_WORKING"},
        {label: "😔 Pensive", value: "PENSIVE"},
        {label: "👂 Listening", value: "LISTENING"},
        {label: "💬 Communicating", value: "COMMUNICATING"},
        {label: "🗂️ Organizing", value: "ORGANIZING"},
        {label: "💱 Negociation", value: "NEGOCIATION"},
        {label: "🔥 Responsiveness", value: "RESPONSIVENESS"},
      ],
      specialitiesSelected: [],
      specialityTypes: [
        {label: "🏘️ Residential", value: "RESIDENTIAL"},
        {label: "🏬 Commercial", value: "COMMERCIAL"},
        {label: "🏚️ Property management", value: "PROPERTY_MANAGEMENT"},
        {label: "🏗️ New construction", value: "NEW_CONSTRUCTION"},
        {label: "💍️ Luxury", value: "LUXURY"},
        {label: "🚜 Farms", value: "FARMS"},
      ],
      valid: false,
      rules: [
        v => !!v || 'field is required',
      ],
    };
  },
};
</script>
