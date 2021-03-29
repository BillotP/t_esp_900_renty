<template>
  <v-container>
    <v-data-table
        :headers="headers"
        :items="tenants"
        :items-per-page="5"
        class="elevation-1"
    >
      <template v-slot:item.actions="{ item }">
        <v-icon
            small
            class="mr-2"
            @click="goToProfile(item)"
        >
          mdi-eye
        </v-icon>
        <v-icon
            small
            @click.stop="selectTenant(item)"
        >
          mdi-home
        </v-icon>
      </template>
    </v-data-table>

    <v-dialog v-model="dialog">
      <v-card>
        <v-card-title class="headline">
          Assign property to tenant:
        </v-card-title>

        <v-card-text>
          <v-spacer></v-spacer>
          <v-select
              v-model="property"
              :items="properties"
              item-text="address"
              item-value="ID"
              label="Properties"
              return-object
          ></v-select>
        </v-card-text>

        <v-card-actions>
          <v-spacer></v-spacer>

          <v-btn color="error" text @click="dialog = false">
            Cancel
          </v-btn>

          <v-btn color="primary" text @click="assignProperty">
            Assign
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
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

const ASSIGN_PROPERTY_MUTATION = gql`
  mutation assignProperty($tenantId: Int!, $propertyId: Int!) {
    assignProperty(tenantId: $tenantId, propertyId: $propertyId) {
      ID
    }
  }
`;

@Component
export default class TenantList extends Vue {
  public property: any = {};
  public properties: any[] = [];
  public tenants: any[] = [];
  public tenant: any = {};
  public dialog = false;

  beforeMount() {
    this.$apollo.getClient().query({
      query: TENANT_QUERY
    }).then((res) => {
      this.tenants = res.data.tenants;
      console.log(res);
    }).catch((err) => {
      console.error(err);
    });
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
        {text: 'Username', value: 'user.username'},
        {text: 'CodeNumber', value: 'properties[0].codeNumber'},
        {text: 'Address', value: 'properties[0].address'},
        {text: 'Actions', value: 'actions', sortable: false},
      ]
    }
  }

  goToProfile(tenant: any) {
    this.$router.push("/tenants/" + tenant.ID);
  }

  selectTenant(tenant: any) {
    this.dialog = true;
    this.tenant = tenant;
    console.log(this.tenant);
  }

  assignProperty() {
    console.log(this.property);
    this.$apollo.getClient().mutate({
      mutation: ASSIGN_PROPERTY_MUTATION,
      variables: {tenantId: this.tenant.ID, propertyId: this.property.ID}
    }).then((res) => {
      console.log(res);
    }).catch((err) => {
      console.error(err);
    });
  }
}
</script>

<style>
.v-data-table {
  margin-top: 5rem;
}
</style>
