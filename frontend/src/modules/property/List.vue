<template>
  <v-data-table
      :headers="headers"
      :items="properties"
      :items-per-page="5"
      class="elevation-1"
      @click:row="goToProfile"
  ></v-data-table>
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
    this.$apollo.getClient().query({
      query: PROPERTIES_QUERY
    }).then((res) => {
      this.properties = res.data.properties;
      console.log(res);
    }).catch((err) => {
      console.error(err);
    });
  }

  data() {
    return {
      headers: [
        {
          text: 'ID',
          align: 'start',
          sortable: false,
          value: 'ID',
        },
        {text: 'CodeNumber', value: 'codeNumber'},
        {text: 'Address', value: 'address'},
        {text: 'Area', value: 'area'},
      ]
    }
  }

  goToProfile(property: any) {
    this.$router.push("/property/" + property.ID);
  }
}
</script>

<style>
.v-data-table {
  margin-top: 5rem;
}
</style>
