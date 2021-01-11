<template>
  <div>
    <h1>Test</h1>
    <p>{{ test() }}</p>
    <ticketlist-field />
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import gql from "graphql-tag";

@Component
export default class TicketList extends Vue {
  // private async loginAsCompany() {
  //   console.log("coucou");
  //   const result = await this.$apollo.mutate({
  //     // Query
  //     mutation: gql`
  //       mutation($input: String!) {
  //         addTag(input: $input) {
  //           input
  //         }
  //       }
  //     `,
  //     // Parameters
  //     variables: {
  //       input: {username : "test", password: "test"},
  //     },
  //   });
  // }

  private test() {
  return gql`query Tenants {
  tenants(where: { is_public: { _eq: false} }, order_by: { created_at: desc }) {
    ID
    createdAt
    updatedAt
    properties
    userID
    user
    estateAgentID
    EstateAgent
    documents
  }
 }`;
}

  // public mounted() {
  //   apollo: {
  //     tenants: gql`
  //       query {
  //         tenants: hello
  //       }
  //     `;
  //   }
  // }
}
</script>