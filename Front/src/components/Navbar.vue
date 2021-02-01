<template>
  <nav>
    <v-toolbar flat style="color: #42b983">
      <v-app-bar app>
        <v-toolbar-title class="text-uppercase grey--text">
          Renty
        </v-toolbar-title>

        <v-spacer></v-spacer>

        <v-btn text color="grey">
          <v-tab exact to="/login/company">
            <span>Sign out</span>
          </v-tab>
          <v-icon right>mdi-exit-to-app</v-icon>
        </v-btn>

        <template v-slot:extension>
          <v-tabs fixed-tabs horizontal align-with-title>
            <v-tab v-if="status == 0 || status == 1" exact to="/dashboard">Dashboard</v-tab>
            <v-tab v-if="status == 2" exact to="/EstateAgent">EstateAgent</v-tab>
            <v-tab v-if="status == 0 || status == 2" exact to="/tenement">tenement</v-tab>
            <v-tab v-if="status == 0 || status == 1" exact to="/tickets">Tickets</v-tab>
          </v-tabs>
        </template>
      </v-app-bar>
    </v-toolbar>
  </nav>
</template>

<script lang="ts">
import Component from "vue-class-component";
import Vue from "vue";
import { Action, Getter, namespace } from "vuex-class";
import { privilege } from "@/store/authentificationStore";

const authModule = namespace("authentificationStore");


@Component
export default class NavBar extends Vue {

  @authModule.Getter("getPrivilege")
  private status!: privilege;
}
</script>
