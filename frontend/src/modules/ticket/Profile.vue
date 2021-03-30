<template>
  <v-container>
    <v-card
        v-if="anomaly && anomaly.createBy && anomaly.assignedTo"
        tile
    >
      <v-list>
        <v-list-item>
          <v-list-item-avatar>
            <v-img
                src="https://as2.ftcdn.net/jpg/01/90/07/23/500_F_190072303_iN1moor8sdI21msMcD0gO6AwSeQuygMW.jpg"></v-img>
          </v-list-item-avatar>
        </v-list-item>

        <v-list-item>
          <v-list-item-content>
            <v-list-item-title class="title">
              {{ anomaly.type }}
            </v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
      <v-divider></v-divider>
      <v-list-item>
        <v-list>
          <v-list-item-title class="title">
            {{ anomaly.description }}
          </v-list-item-title>
        </v-list>
      </v-list-item>
      <v-divider></v-divider>
      <v-list-item>
        <v-list>
          <v-btn
              depressed
              color="primary"
          >
            Assigned To:
          </v-btn>
          <v-chip>
            <v-icon left>
              mdi-account
            </v-icon>
            {{ anomaly.assignedTo.user.username }}
          </v-chip>
        </v-list>
      </v-list-item>
      <v-list-item>
        <v-list>
          <v-btn
              depressed
              color="primary"
          >
            Created By:
          </v-btn>
          <v-chip>
            <v-icon left>
              mdi-account
            </v-icon>
            {{ anomaly.createBy.user.username }}
          </v-chip>
        </v-list>
      </v-list-item>
    </v-card>
  </v-container>
</template>


<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import gql from "graphql-tag";

const TICKET_QUERY = gql`
  query anomaly($id: Int!) {
    anomaly(id: $id) {
      ID
      createBy {
        user {
          username
        }
      }
      assignedTo {
        user {
          username
        }
      }
      type
      description
    }
  }
`;

@Component
export default class TicketProfile extends Vue {
  public anomaly: any = {};

  beforeMount() {
    this.$apollo.getClient().query({
      query: TICKET_QUERY,
      variables: {id: this.$route.params.id}
    }).then((res) => {
      this.anomaly = res.data.anomaly;
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
