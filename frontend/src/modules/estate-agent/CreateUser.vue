<template>
  <v-main
      style="
        margin-top: 10vh;
        margin-bottom: 10vh;
        margin-right: 10vh;
        margin-left: 10vh;
      "
  >
    <v-card>
      <v-card-text>
        <h3>Estate Agent</h3>
        <v-container fluid>
          <v-form v-model="valid">
            <h4>My Profile</h4>
            <v-row>
              <v-text-field
                  type="phone"
                  v-model="phone"
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
                  v-model="about"
                  label="About"
              ></v-textarea>
            </v-row>
            <h4>My Account</h4>
            <v-row>
              <v-text-field
                  type="text"
                  v-model="username"
                  label="Username"
                  required
                  :rules="rules"
              ></v-text-field>
            </v-row>
            <v-row>
              <v-text-field
                  type="password"
                  v-model="password"
                  label="Password"
                  required
                  :rules="rules"
              ></v-text-field>
            </v-row>
            <v-row>
              <v-btn depressed color="primary" v-on:click="createUser()"> Create</v-btn>
              <v-snackbar
                  v-model="snackbar"
                  :color="hasError ? 'red' : ''"
                  :timeout="timeout"
              >
                {{ text }}
                <template v-slot:action="{ attrs }">
                  <v-btn color="blue" text v-bind="attrs" @click="snackbar = false">
                    Close
                  </v-btn>
                </template>
              </v-snackbar>
            </v-row>
          </v-form>
        </v-container>
      </v-card-text>
    </v-card>
  </v-main>
</template>

<script>
import Vue from "vue";
import Component from "vue-class-component";
import gql from "graphql-tag";

const CREATE_ESTATE_AGENT_USER_MUTATION = gql`
  mutation($input: EstateAgentInput) {
    createEstateAgentUser(input: $input) {
      ID
    }
  }
`;

@Component
export default class CreateEstateAgentUser extends Vue {
  data() {
    return {
      about: null,
      phone: null,
      username: "",
      password: "",
      snackbar: false,
      text: "",
      timeout: 2000,
      hasError: false,
      valid: false,
      rules: [
        v => !!v || 'field is required',
      ],
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
    };
  }

  async createUser() {
    try {
      const resp = await this.$apollo.getClient().mutate({
        mutation: CREATE_ESTATE_AGENT_USER_MUTATION,
        variables: {
          input: {
            user: {
              username: this.$data.username,
              password: this.$data.password,
            },
            about: this.$data.about,
            tel: this.$data.phone,
            skills: this.$data.skillsSelected,
            specialities: this.$data.specialitiesSelected
          },
        },
      });
      if (resp.data.createEstateAgentUser.ID) {
        this.$data.text =
            "Estate agent " + this.$data.username + " successfully created !";
        this.$data.hasError = false;
        this.$data.snackbar = true;
        // this.$router.("/estate-agents");
        this.$router.go(-1);
      }
    } catch (e) {
      this.$data.text =
          "⚠️ Failed to create estate agent :" + e["graphQLErrors"][0]["message"];
      this.$data.hasError = true;
      this.$data.snackbar = true;
    }
  }
}
</script>

<style>
</style>
