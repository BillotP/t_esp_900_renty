<template>
  <div>
    <ticketlist-field />
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import gql from "graphql-tag";

@Component
export default class TicketList extends Vue {
}
</script>