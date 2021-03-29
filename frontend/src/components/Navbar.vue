<template>
  <div>
    <v-toolbar
        dark
        prominent
        src="https://cdn.vuetifyjs.com/images/backgrounds/vbanner.jpg"
        height="70"
    >
      <v-row>
        <v-col>
          <v-toolbar-title fixed-tabs>Renty</v-toolbar-title>
        </v-col>
      </v-row>
      <v-spacer />

      <v-tabs fixed-tabs height="70">
        <v-tab exact to="/dashboard">Dashboard</v-tab>
        <v-tab v-if="getPrivilege == 0" exact to="/estate-agents">Estate Agents</v-tab>
        <v-tab v-if="getPrivilege == 1" exact to="/tenants">Tenants</v-tab>
        <v-tab exact to="/tickets">Tickets</v-tab>
        <v-tab v-if="getPrivilege == 1" exact to="/property">Property</v-tab>
      </v-tabs>
      <v-tab exact>
        <v-btn text v-on:click="goToEditProfile" color="white">
          <v-icon>mdi-account-circle</v-icon>
        </v-btn>
      </v-tab>
      <v-tab exact>
        <v-btn text v-on:click="signOut" color="white">
          <span>Sign out</span>
          <v-icon right>mdi-exit-to-app</v-icon>
        </v-btn>
      </v-tab>
    </v-toolbar>
  </div>
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
}`;

@Component
export default class NavBar extends Vue {
  profile: any = {};

  get getPrivilege() {
    return Number(localStorage.getItem("privilege")) || 0;
  }

  // TODO each call => token not refresh
  beforeMount() {
    this.$apollo.getClient().query({
      query: PROFILE_QUERY,
    }).then((res) => {
      this.profile = res.data.profile;
      console.log(res);
    }).catch((err) => {
      console.error(err);
    });
  }

  goToEditProfile() {
    const roles = {
      Company: 'companies',
      EstateAgent: 'estate-agents',
      Tenant: 'tenants'
    };

    this.$router.push('/' + roles[this.profile.__typename] + '/' + this.profile.ID + '/edit');
  }

  signOut() {
    localStorage.clear();
    this.$router.push("/login");
  }
}
</script>
