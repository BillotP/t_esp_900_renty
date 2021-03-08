<template>
  <v-card
      v-if="tenant && tenant.user && tenant.user.username"
      tile
  >
    <v-list>
      <v-list-item>
        <v-list-item-avatar>
          <v-img src="https://www.flaticon.com/premium-icon/icons/svg/1018/1018651.svg"></v-img>
        </v-list-item-avatar>
      </v-list-item>

      <v-list-item>
        <v-list-item-content>
          <v-list-item-title class="title">
            {{ tenant.user.username }}
          </v-list-item-title>
        </v-list-item-content>
      </v-list-item>
    </v-list>
    <v-divider></v-divider>
    <v-list
    >
      <v-list-group
          no-action
          sub-group
          prepend-icon="mdi-home"
      >
        <template v-slot:activator>
          <v-list-item-content>
            <v-list-item-title>Properties</v-list-item-title>
          </v-list-item-content>
        </template>

        <v-list-item
            v-for="(property, i) in tenant.properties"
            :key="i"
            link
        >
          <v-list-item-title v-text="property.address"></v-list-item-title>
          <v-list-item-subtitle v-text="property.codeNumber"></v-list-item-subtitle>

        </v-list-item>
      </v-list-group>

      <v-list-group
          no-action
          sub-group
          prepend-icon="mdi-file-document"
      >
        <template v-slot:activator>
          <v-list-item-content>
            <v-list-item-title>Documents</v-list-item-title>
          </v-list-item-content>
        </template>

        <v-list-item
            v-for="(document, i) in tenant.documents"
            v-on:click="downloadDocument(document)"
            :key="i"
            link
        >
          <v-list-item-title v-text="document.type"></v-list-item-title>
          <v-list-item-subtitle v-text="document.url"></v-list-item-subtitle>

        </v-list-item>
      </v-list-group>

    </v-list>
  </v-card>
</template>


<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import gql from "graphql-tag";

const TENANT_QUERY = gql`
  query tenant($id: Int!) {
    tenant(id: $id) {
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
        type
      }
    }
  }
`;

@Component
export default class TenantProfile extends Vue {
  public tenant: any = {};

  beforeMount() {
    this.$apollo.getClient().query({
      query: TENANT_QUERY,
      variables: {id: this.$route.params.id}
    }).then((res) => {
      this.tenant = res.data.tenant;
      console.log(res);
    }).catch((err) => {
      console.error(err);
    });
  }

  downloadDocument(myDocument: any) {
    const a = document.createElement("a");
    a.href = 'http://localhost:8080' + myDocument.url;
    a.target = "_blank";
    document.body.appendChild(a);
    a.click();
    window.URL.revokeObjectURL(myDocument.url);
    a.remove();
  }
}

</script>

<style>
.v-card {
  margin-top: 4rem;
}
</style>
