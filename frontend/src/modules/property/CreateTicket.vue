<template>
  <v-form v-model="valid">
    <v-container>
      <v-row>
        <v-col cols="12" md="4">
          <v-text-field v-model="type" label="Type" required></v-text-field>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" md="4">
          <v-textarea
            v-model="description"
            value="Description"
            label="Description"
            required
          ></v-textarea>
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <v-hover>
            <v-btn outlined @click="CreateTicket" :disabled="!valid"
              >Create</v-btn
            >
          </v-hover>
        </v-col>
      </v-row>
    </v-container>
  </v-form>
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
      property: this.$route.params.id,
      type: "",
    };
  }

  async CreateTicket() {
    try {
      const resp = await this.$apollo.getClient().mutate({
        mutation: CREATE_ANOMALY_MUTATION,
        variables: {
          input: {
            property: this.$data.property,
            type: this.$data.type,
            description: this.$data.description,
          },
        },
      });
      if (resp.data.createProperty) {
        console.log("coucou");
      }
    } catch (e) {
      console.error(e);
    }
  }

  // public mounted() {
  //   const resolved = this.$router.resolve({
  //     name: "SomeRouteName",
  //     params: { id: item.id },
  //   });
  // }
}
</script>

<style>
</style>