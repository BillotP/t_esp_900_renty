<template>
  <v-main
      style="
        margin-top: 10vh;
        margin-bottom: 10vh;
        margin-right: 10vh;
        margin-left: 10vh;
      "
  >
    <v-card>
      <v-card-text>
        <h3>Ticket</h3>
        <v-container fluid>
          <v-form v-model="valid">
            <h4>Informations</h4>
            <v-row>
              <v-col cols="12">
                <v-select
                    v-model="type"
                    :items="ticketTypes"
                    item-text="label"
                    item-value="value"
                    label="Ticket kind"
                    hint="Select the appropriate issue type"
                    required
                ></v-select>
              </v-col>
            </v-row>
            <v-row>
              <v-col cols="12">
                <v-textarea
                    v-model="description"
                    value="Description"
                    label="Description"
                    hint="Describe your issue"
                    required
                ></v-textarea>
              </v-col>
            </v-row>
            <v-row>
              <v-col>
                <v-hover>
                  <v-btn outlined @click="CreateTicket" :disabled="!valid"
                  >Create
                  </v-btn
                  >
                </v-hover>
              </v-col>
            </v-row>
          </v-form>
        </v-container>
      </v-card-text>
    </v-card>
  </v-main>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import gql from "graphql-tag";

const CREATE_ANOMALY_MUTATION = gql`
  mutation($input: AnomalyInput) {
    createAnomaly(input: $input) {
      ID
    }
  }
`;

@Component
export default class CreateTicket extends Vue {
  data() {
    return {
      description: "",
      valid: false,
      timeout: 2000,
      type: "",
      ticketTypes: [
        {label: "🔨 Maintenance", value: "MAINTENANCE"},
        {label: "💳 Rent payment method", value: "PAYMENT"},
        {label: "⚠️Rent issue", value: "RENT"},
        {label: "📄 Document request", value: "DOCUMENTS"},
        {label: "⚠️Accommodation issue", value: "ACCOMMODATION"},
        {label: "🥜 Other", value: "OTHER"},
      ],
    };
  }

  async CreateTicket() {
    try {
      const resp = await this.$apollo.getClient().mutate({
        mutation: CREATE_ANOMALY_MUTATION,
        variables: {
          input: {
            type: this.$data.type,
            description: this.$data.description,
          },
        },
      });
      if (resp.data.createProperty) {
        this.$router.back();
      }
    } catch (e) {
      console.error(e);
    }
  }
}
</script>

<style>
</style>
