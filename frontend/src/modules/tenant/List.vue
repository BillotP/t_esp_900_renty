<template>
  <v-container>
    <v-row align="center" justify="center" style="height: 10vh">
      <v-col align="center" justify="center">
        <h3 style="text-transform: capitalize">tenants</h3>
      </v-col>

      <v-row align="center" justify="center">
        <v-btn
          style="margin: 5px"
          right
          color="primary"
          v-on:click="goToCreateTenant"
        >
          +
        </v-btn>
        <v-btn
          style="margin: 5px"
          right
          color="primary"
          v-on:click="$apollo.queries['tenants'].refetch()"
        >
          🔄
        </v-btn>
      </v-row>
    </v-row>

    <v-data-table
      :headers="headers"
      :items="tenants"
      :items-per-page="5"
      class="elevation-1"
    >
      <template v-slot:[`item.actions`]="{ item }">
        <v-icon small class="mr-2" @click="goToProfile(item)"> mdi-eye </v-icon>
        <v-icon small @click.stop="selectTenant(item)"> mdi-home </v-icon>
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

          <v-btn color="error" text @click="dialog = false"> Cancel </v-btn>

          <v-btn color="primary" text @click="assignProperty"> Assign </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
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
      postalCode
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
        postalCode
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
export default {
  apollo: {
    tenants: {
      query: TENANT_QUERY,
      pollInterval: 3000,
    },
    properties: {
      query: PROPERTIES_QUERY,
    },
  },
  data() {
    return {
      tenants: [],
      properties: [],
      headers: [
        {
          text: "ID",
          align: "start",
          sortable: false,
          value: "ID",
        },
        { text: "Username", value: "user.username" },
        { text: "Postal Code", value: "properties[0].postalCode" },
        { text: "Address", value: "properties[0].address" },
        { text: "Actions", value: "actions", sortable: false },
      ],
      dialog: false,
      property: {},
    };
  },
  methods: {
    goToProfile(tenant) {
      this.$router.push("/tenants/" + tenant.ID);
    },

    goToCreateTenant() {
      this.$router.push("/create/tenant/");
    },

    selectTenant(tenant) {
      this.dialog = true;
      this.tenant = tenant;
      console.log(this.tenant);
    },

    assignProperty() {
      console.log(this.property);
      this.$apollo
        .getClient()
        .mutate({
          mutation: ASSIGN_PROPERTY_MUTATION,
          variables: { tenantId: this.tenant.ID, propertyId: this.property.ID },
        })
        .then((res) => {
          this.dialog = false;
          console.log(res);
        })
        .catch((err) => {
          console.error(err);
        });
    },
  },
};
</script>
<style>
.v-data-table {
  margin-top: 5rem;
}
</style>
