<template>
  <v-main>
    <v-card v-if="this.$apollo.queries.company.loading" height="200">
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
        ><b style="text-transform: capitalize">{{ company.name }} 🖊️</b>
      </v-card-title>
      <v-card-text>
        <v-row>
          <v-text-field
            hint="Update your company name"
            v-model="company.name"
            label="Name"
          ></v-text-field>
        </v-row>
        <v-row>
          <v-text-field
            hint="Update your company description"
            v-model="company.description"
            label="Description"
          ></v-text-field>
        </v-row>
        <v-row>
          <v-img
            v-if="company.logo"
            style="margin-right: 10px; border: 1px solid black"
            max-height="64px"
            max-width="64px"
            :src="company.logo.url"
            left
          ></v-img>
          <v-text-field
            hint="Update your company logo URL"
            v-model="company.logo.url"
            label="Logo URL"
          ></v-text-field>
        </v-row>
        <v-row>
          <v-text-field
            hint="Update your company phone number"
            v-model="company.tel"
            label="Phone number"
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
          {{ isLoading ? "" : "Update" }}</v-btn
        >
      </v-card-text>
    </v-card></v-main
  >
</template>

<script>
import gql from "graphql-tag";
export default {
  apollo: {
    company: {
      query: gql`
        query company($id: Int!) {
          company(id: $id) {
            ID
            logo {
              url
            }
            logoID
            name
            createdAt
            updatedAt
            description
            tel
          }
        }
      `,
      variables() {
        return { id: this.$route.params.id };
      },
      pollInterval: 2000,
    },
  },
  methods: {
    udpate() {
      console.log("Update !");
    },
  },
  data() {
    return {
      company: {},
      companyDescription: "",
      companyLogoUrl: "",
      companyPhoneNumber: "",
      companyName: localStorage.getItem("username"),
      isLoading: false,
      companyID: this.$route.params.id,
    };
  },
};
</script>