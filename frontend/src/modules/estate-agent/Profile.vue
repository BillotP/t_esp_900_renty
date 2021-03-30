<template>
  <v-container>
    <v-card
        v-if="estateAgent"
        tile
    >
      <v-list>
        <v-list-item>
          <v-list-item-avatar>
            <v-img
                src="https://www.flaticon.com/svg/vstatic/svg/1029/1029022.svg?token=exp=1617065368~hmac=568d1fdbc9fbf86cf32adff388f85872"></v-img>
          </v-list-item-avatar>
        </v-list-item>

        <v-list-item>
          <v-list-item-content>
            <v-list-item-title class="title">
              {{ estateAgent.user.username }}
            </v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
      <v-divider></v-divider>
      <v-list-item>
        <v-list>
          <v-list-item-title class="title">
            {{ estateAgent.company.name }}
          </v-list-item-title>
        </v-list>
      </v-list-item>
    </v-card>
  </v-container>
</template>


<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import gql from "graphql-tag";

const ESTATE_AGENT_QUERY = gql`
  query estateAgent($id: Int!) {
    estateAgent(id: $id) {
      ID
      user {
        username
      }
      company {
        name
      }
    }
  }
`;

@Component
export default class EstateAgentProfile extends Vue {
  public estateAgent: any = {};

  beforeMount() {
    this.$apollo.getClient().query({
      query: ESTATE_AGENT_QUERY,
      variables: {id: this.$route.params.id}
    }).then((res) => {
      this.estateAgent = res.data.estateAgent;
      console.log(res);
    }).catch((err) => {
      console.error(err);
    });
  }
}

</script>

<style>
.v-card {
  margin-top: 4rem;
}
</style>
