<template>
  <v-toolbar dark prominent height="70">
    <v-col align="center" justify="center">
      <v-toolbar-title>🏡 Renty</v-toolbar-title>
    </v-col>
    <v-spacer />

    <v-tabs fixed-tabs>
      <v-tab exact to="/dashboard">Dashboard</v-tab>
      <v-tab v-if="getPrivilege == 0" exact to="/estate-agents"
        >Estate Agents</v-tab
      >
      <v-tab v-if="getPrivilege == 1" exact to="/tenants">Tenants</v-tab>
      <v-tab v-if="getPrivilege != 0" exact to="/tickets">Tickets</v-tab>
      <v-tab v-if="getPrivilege == 1" exact to="/properties">Properties</v-tab>
    </v-tabs>
    <v-tab exact style="margin-top: auto; margin-bottom: auto">
      <v-btn text v-on:click="goToEditProfile" color="white">
        <v-icon>mdi-account-circle</v-icon>
      </v-btn>
    </v-tab>
    <v-tab exact style="margin-top: auto; margin-bottom: auto">
      <v-btn text v-on:click="signOut" color="white">
        Sign out
        <v-icon right>mdi-exit-to-app</v-icon>
      </v-btn>
    </v-tab>
  </v-toolbar>
</template>

<script lang="ts">
import Component from "vue-class-component";
import Vue from "vue";
import gql from "graphql-tag";

const PROFILE_QUERY = gql`
  query profile {
    profile {
      ID
      __typename
    }
  }
`;

@Component
export default class NavBar extends Vue {
  profile: any = {};

  get getPrivilege() {
    return Number(localStorage.getItem("privilege")) || 0;
  }

  // TODO each call => token not refresh
  async beforeMount() {
    var foo = this.$apollo.getClient().watchQuery({
      query: PROFILE_QUERY,
      pollInterval: 4000,
    });
    var fon = await foo.result();
    if (fon.errors) {
      console.error(fon.errors);
      this.$router.push("/login");
    } else {
      this.profile = fon.data.profile;
      console.log(fon.data);
    }
  }

  goToEditProfile() {
    const roles = {
      Company: "companies",
      EstateAgent: "estate-agents",
      Tenant: "tenants",
    };

    this.$router.push(
      "/" + roles[this.profile.__typename] + "/" + this.profile.ID + "/edit"
    );
  }

  signOut() {
    localStorage.clear();
    this.$router.push("/login");
  }
}
</script>
