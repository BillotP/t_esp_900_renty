<template>
  <v-data-table
      :headers="headers"
      :items="tenants"
      :items-per-page="5"
      class="elevation-1"
      @click:row="goToProfile"
  ></v-data-table>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import gql from "graphql-tag";

const TENANT_QUERY = gql`
  query tenants {
    tenants {
      ID
      user {
        username
      }
      properties {
        area
        address
        codeNumber
      }
      documents {
        url
      }
    }
  }
`;

@Component
export default class TenantList extends Vue {
  public tenants = [];

  beforeMount() {
    this.$apollo.getClient().query({
      query: TENANT_QUERY
    }).then((res) => {
      this.tenants = res.data.tenants;
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
        {text: 'Username', value: 'user.username'},
        {text: 'CodeNumber', value: 'properties[0].codeNumber'},
        {text: 'Address', value: 'properties[0].address'},
      ]
    }
  }

  goToProfile(tenant: any) {
    this.$router.push("/tenants/" + tenant.ID);
  }
}
</script>

<style>
.v-data-table {
  margin-top: 5rem;
}
</style>
