<template>
  <v-container>
    <v-btn right color="primary" v-on:click="goToCreateProperty">
      Create Property
    </v-btn>
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

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
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

@Component
export default class PropertyList extends Vue {
  public properties = [];

  beforeMount() {
    this.$apollo
      .getClient()
      .query({
        query: PROPERTIES_QUERY,
      })
      .then((res) => {
        this.properties = res.data.properties;
        console.log(res);
      })
      .catch((err) => {
        console.error(err);
      });
  }

  data() {
    return {
      headers: [
        {
          text: "ID",
          align: "start",
          sortable: false,
          value: "ID",
        },
        { text: "CodeNumber", value: "codeNumber" },
        { text: "Address", value: "address" },
        { text: "Area", value: "area" },
        { text: "Actions", value: "actions", sortable: false },
      ],
    };
  }

  goToProfile(property: any) {
    this.$router.push("/property/" + property.ID);
  }

  goToCreateProperty() {
    this.$router.push("/create/property/");
  }
}
</script>

<style>
.v-data-table {
  margin-top: 5rem;
}
</style>
