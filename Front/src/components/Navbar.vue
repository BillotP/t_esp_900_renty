<template>
  <nav>
    <v-toolbar flat style="color: #42b983">
      <v-app-bar app>
        <v-toolbar-title class="text-uppercase grey--text">
          Renty
        </v-toolbar-title>

        <v-spacer></v-spacer>

        <v-tab exact to="/login/company">
          <v-btn text color="grey">
            <span>Sign out</span>
            <v-icon right>mdi-exit-to-app</v-icon>
          </v-btn>
        </v-tab>

        <template v-slot:extension>
          <v-tabs fixed-tabs horizontal align-with-title>
            <v-tab
              v-if="getPrivilege == 0 || getPrivilege == 1"
              exact
              to="/dashboard"
              >Dashboard</v-tab
            >
            <v-tab v-if="getPrivilege == 2" exact to="/EstateAgent"
              >EstateAgent</v-tab
            >
            <v-tab
              v-if="getPrivilege == 0 || getPrivilege == 2"
              exact
              to="/tenement"
              >tenement</v-tab
            >
            <v-tab
              v-if="getPrivilege == 0 || getPrivilege == 1"
              exact
              to="/tickets"
              >Tickets</v-tab
            >
          </v-tabs>
        </template>
      </v-app-bar>
    </v-toolbar>
  </nav>
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
}
</script>
