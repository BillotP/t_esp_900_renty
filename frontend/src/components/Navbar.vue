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
        <v-tab v-if="getPrivilege == 0" exact to="/estate-agents"
          >Estate Agents</v-tab
        >
        <v-tab v-if="getPrivilege == 1" exact to="/tenants">Tenants</v-tab>
        <v-tab exact to="/tickets">Tickets</v-tab>
        <v-tab v-if="getPrivilege == 1" exact to="/property">Property</v-tab>
      </v-tabs>
      <v-tab exact>
        <v-btn text v-on:click="signOut" color="white">
          <span>Sign out</span>
          <v-icon right>mdi-exit-to-app</v-icon>
        </v-btn>
      </v-tab>
    </v-toolbar>
  </div>
  <!-- <nav>
    <v-toolbar prominent style="color: #42b983">
      <v-app-bar app>
        <v-toolbar-title class="text-uppercase grey--text">
          Renty
        </v-toolbar-title>

        <v-spacer></v-spacer>

        <v-tab exact>
          <v-btn text v-on:click="signOut" color="grey">
            <span>Sign out</span>
            <v-icon right>mdi-exit-to-app</v-icon>
          </v-btn>
        </v-tab>

          <v-tabs fixed-tabs horizontal align-with-title>
            <v-tab
              exact
              to="/dashboard"
              >Dashboard</v-tab
            >
            <v-tab
              v-if="getPrivilege == 0"
              exact
              to="/estate-agents"
              >Estate Agents</v-tab
            >
            <v-tab
              v-if="getPrivilege == 1"
              exact
              to="/tenants"
              >Tenants</v-tab
            >
            <v-tab
                exact
                to="/tickets"
            >Tickets</v-tab
            >
            <v-tab
                v-if="getPrivilege == 1"
                exact
                to="/property"
            >Property</v-tab
            >
          </v-tabs>
        </template>
      </v-app-bar>
    </v-toolbar>
  </nav> -->
</template>

<script lang="ts">
import Component from "vue-class-component";
import Vue from "vue";
import { namespace } from "vuex-class";

@Component
export default class NavBar extends Vue {
  get getPrivilege() {
    return Number(localStorage.getItem("privilege")) || 0;
  }

  signOut() {
    localStorage.clear();
    this.$router.push("/login");
  }
}
</script>
