<template>
  <v-main>
    <v-card class="mx-auto" max-width="434" tile>
      <v-img height="100%" :src="banneer"></v-img>
      <v-col>
        <v-avatar size="100" style="position: absolute; top: 130px">
          <v-img
            :src="`https://robohash.org/${estateAgent.user.username}.png?set=set5&bgset=bg1&size=100x100`"
          ></v-img>
        </v-avatar>
      </v-col>
      <v-list-item color="rgba(0, 0, 0, .4)">
        <v-list-item-content>
          <v-list-item-title style="text-transform: capitalize" class="title">{{
            estateAgent.user.username
          }}</v-list-item-title>
          <v-list-item-subtitle
            >RealEstate Agent N° {{ id }}</v-list-item-subtitle
          >
          <v-list-item-subtitle
            >@
            <strong>{{
              estateAgent.company.name
            }}</strong></v-list-item-subtitle
          >
        </v-list-item-content>
        <v-list-item-content> </v-list-item-content>
      </v-list-item>
    </v-card>
  </v-main>
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
  data() {
    return {
      id: this.$route.params.id,
      banneer:
        "https://image.freepik.com/free-photo/real-estate-agent-customer-face-mask-looking-new-project_53876-97516.jpg",
    };
  }
  beforeMount() {
    this.$apollo
      .getClient()
      .query({
        query: ESTATE_AGENT_QUERY,
        variables: { id: this.$route.params.id },
      })
      .then((res) => {
        this.estateAgent = res.data.estateAgent;
        console.log(res);
      })
      .catch((err) => {
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
