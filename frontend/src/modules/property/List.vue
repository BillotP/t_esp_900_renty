<template>
  <v-container>
    <v-row align="center" justify="center" style="height: 10vh">
      <v-col align="center" justify="center">
        <h3 style="text-transform: capitalize">properties</h3>
      </v-col>

      <v-row align="center" justify="center">
        <v-btn
          style="margin: 5px"
          right
          color="primary"
          v-on:click="goToCreateProperty"
        >
          +
        </v-btn>
        <v-btn
          style="margin: 5px"
          right
          color="primary"
          v-on:click="$apollo.queries['properties'].refetch()"
        >
          🔄
        </v-btn>
      </v-row>
    </v-row>

    <v-data-table
      :headers="headers"
      :items="properties"
      :items-per-page="5"
      class="elevation-1"
    >
      <template v-slot:[`item.actions`]="{ item }">
        <v-icon small class="mr-2" @click="goToProfile(item)"> mdi-eye </v-icon>
      </template>
    </v-data-table>
  </v-container>
</template>

<script>
import gql from "graphql-tag";

const PROPERTIES_QUERY = gql`
  query properties {
    properties {
      ID
      area
      address
      codeNumber
    }
  }
`;
export default {
  apollo: {
    properties: {
      query: PROPERTIES_QUERY,
      pollInterval: 3000,
    },
  },
  data() {
    return {
      headers: [
        {
          text: "ID",
          align: "start",
          sortable: false,
          value: "ID",
        },
        { text: "Postal Code", value: "codeNumber" },
        { text: "Address", value: "address" },
        { text: "Property Area (m²)", value: "area" },
        { text: "Actions", value: "actions", sortable: false },
      ],
    };
  },
  methods: {
    goToProfile(property) {
      this.$router.push("/property/" + property.ID);
    },

    goToCreateProperty() {
      this.$router.push("/create/property/");
    },
  },
};
</script>

<style>
.v-data-table {
  margin-top: 5rem;
}
</style>
