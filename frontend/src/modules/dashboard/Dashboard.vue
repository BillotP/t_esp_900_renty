<template>
  <v-container>
    <v-card>
      <h3 style="text-align: center;margin: 10px;">Welcome {{ username }}</h3>
      <p>
        Lorem ipsum dolor sit amet consectetur adipisicing elit. Magni, neque.
        Magni doloribus perspiciatis voluptatem porro minus rem explicabo modi
        quam ab iure molestiae nulla, quo sapiente repellat dicta nostrum culpa!
      </p>
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