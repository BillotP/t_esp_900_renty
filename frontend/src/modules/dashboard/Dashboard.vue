<template>
  <v-container>
    <v-card>
      <h3 style="text-align: center; margin: 10px">Welcome {{ username }}</h3>
      <p style="text-align: center; margin-bottom: 10px">Use the above menu to access app functionality 🔝</p>
      <v-img position="sticky" src="/img/pres.png"> </v-img>
    </v-card>
  </v-container>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
@Component
export default class Dashboard extends Vue {
  public username = localStorage.getItem("username");
  public userID = localStorage.getItem("id");
}
</script>