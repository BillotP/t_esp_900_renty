<template>
  <v-main>
    <v-card
        v-if="anomaly && anomaly.createBy && anomaly.createBy.user.username"
        class="mx-auto" max-width="434" tile
    >
      <v-chip color="secondary" style="text-transform: uppercase;">
        <i style="color: white">{{ ticketStates[anomaly.state] }}</i>
      </v-chip>
      <v-img height="100%" :src="banner"></v-img>
      <v-card-text>
        <v-list>
          <v-list-item>
            <v-list-item-content>
              <v-list-item-title class="title">
                {{ ticketTypes[anomaly.type] }}
              </v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list>
        <v-divider></v-divider>
        <v-list-item>
          <v-list>
            <v-list-item-title class="title">
              <v-tooltip left>
                <template v-slot:activator="{ on, attrs }">
                  <v-btn
                      depressed
                      v-if="anomaly.priority"
                      v-bind="attrs"
                      v-on="on"
                  >
                    {{ priorityIcons[anomaly.priority] }}
                  </v-btn>
                </template>
                <span>{{ anomaly.priority }}</span>
              </v-tooltip>
              <p>{{ anomaly.description }}</p>
            </v-list-item-title>
          </v-list>
        </v-list-item>
        <v-divider></v-divider>
        <v-list-item>
          <v-list>
            <v-btn x-small depressed color="primary"> Assigned To:</v-btn>
            <v-chip style="margin-left: 10px">
              <v-icon left> mdi-account</v-icon>
              {{ anomaly.assignedTo.user.username }}
            </v-chip>
          </v-list>
        </v-list-item>
        <v-list-item>
          <v-list>
            <v-btn x-small depressed color="primary"> Created By:</v-btn>
            <v-chip style="margin-left: 10px">
              <v-icon left> mdi-account</v-icon>
              {{ anomaly.createBy.user.username }}
            </v-chip>
          </v-list>
        </v-list-item>
      </v-card-text>
    </v-card>
  </v-main>
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
      priority
      state
    }
  }
`;

@Component
export default class TicketProfile extends Vue {
  public anomaly: any = {};

  beforeMount() {
    this.$apollo
        .getClient()
        .query({
          query: TICKET_QUERY,
          variables: {id: this.$route.params.id},
        })
        .then((res) => {
          this.anomaly = res.data.anomaly;
          console.log(res);
        })
        .catch((err) => {
          console.error(err);
        });
  }

  data() {
    return {
      banner: 'https://image.freepik.com/vecteurs-libre/service-clientele-illustration-icones_53876-66281.jpg',
      ticketTypes: {
        MAINTENANCE: "🔨 Maintenance",
        PAYMENT: "💳 Rent payment method",
        RENT: "⚠️Rent issue",
        DOCUMENTS: "📄 Document request",
        ACCOMMODATION: "⚠️Accommodation issue",
        OTHER: "🥜 Other",
      },
      ticketStates: {
        TODO: "📋 To do",
        IN_PROGRESS: "🔄 In progress",
        DONE: "✅ Done",
      },
      priorityIcons: {
        MAJOR: "⬆️⬆️",
        HIGHEST: "⬆️",
        HIGH: "↕️",
        MEDIUM: "⬇️",
        LOW: "⬇️⬇️",
      },
    }
  }
}
</script>

<style>
.v-card {
  margin-top: 4rem;
}
</style>
