<template>
  <v-card
      v-if="property"
      tile
  >
    <v-list>
      <v-list-item>
        <v-list-item-avatar>
          <v-img src="https://as2.ftcdn.net/jpg/01/35/38/75/500_F_135387578_vKyGn4NM9E2ipUS9j1GRCDLs40CwRNyC.jpg"></v-img>
        </v-list-item-avatar>
      </v-list-item>

      <v-list-item>
        <v-list-item-content>
          <v-list-item-title class="title">
            {{ property.type }}
          </v-list-item-title>
        </v-list-item-content>
      </v-list-item>
    </v-list>
    <v-divider></v-divider>
    <v-list-item>
      <v-list>
        <v-list-item-title class="title">
          {{ property.address }}
        </v-list-item-title>
        <i>{{ property.codeNumber }}</i>
      </v-list>
    </v-list-item>
    <v-divider></v-divider>
    <v-alert>{{ property.area }}</v-alert>
  </v-card>
</template>


<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import gql from "graphql-tag";

const PROPERTY_QUERY = gql`
  query property($id: Int!) {
    property(id: $id) {
      ID
      area
      address
      codeNumber
      type
    }
  }
`;

@Component
export default class PropertyProfile extends Vue {
  public property: any = {};

  beforeMount() {
    this.$apollo.getClient().query({
      query: PROPERTY_QUERY,
      variables: {id: this.$route.params.id}
    }).then((res) => {
      this.property = res.data.property;
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
