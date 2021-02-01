<template>
  <nav>
    <v-toolbar flat style="color: #42b983;">
      <v-app-bar app>

        <v-toolbar-title class="text-uppercase grey--text">
          Renty
        </v-toolbar-title>

        <v-spacer></v-spacer>

        <v-btn text color="grey">
          <span>Sign out</span>
          <v-icon right>mdi-exit-to-app</v-icon>
        </v-btn>

        <template v-slot:extension>
          <v-tabs fixed-tabs horizontal align-with-title>
            <v-tab exact to="/dashboard">Dashboard</v-tab>
            <v-tab exact to="/tenement">tenement</v-tab>
            <v-tab exact to="/tickets">Tickets</v-tab>
          </v-tabs>
        </template>
      </v-app-bar>
    </v-toolbar>
  </nav>
</template>

<script lang="ts">
import Component from "vue-class-component";
import Vue from "vue";

@Component
export default class NavBar extends Vue {
}
</script>
