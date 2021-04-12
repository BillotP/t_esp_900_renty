<template>
  <v-main>
    <v-card
        v-if="tenant && tenant.user && tenant.user.username"
        class="mx-auto" max-width="434" tile
    >
      <v-img height="100%" :src="banner"></v-img>
      <v-card-text>
        <v-list>
          <v-list-item>
            <v-list-item-avatar>
              <v-img :src="'https://avatars.dicebear.com/api/' + gender + '/' + tenant.user.username + '.svg'"></v-img>
            </v-list-item-avatar>
            <v-list-item-title class="title">
              {{ tenant.user.username }}
            </v-list-item-title>
          </v-list-item>

          <v-list-item>
            <v-list-item-content>
              <v-list-item-subtitle>
                <v-icon style="margin-right: 10px">mdi-phone</v-icon>
                <span>{{ tenant.tel }}</span>
              </v-list-item-subtitle>
              <v-list-item-subtitle>
                <v-icon style="margin-right: 10px">mdi-cake</v-icon>
                <span>{{ birthday }}</span>
              </v-list-item-subtitle>
              <v-list-item-subtitle>
                <v-icon style="margin-right: 10px">mdi-barcode</v-icon>
                <strong>{{ tenant.customerRef }}</strong>
              </v-list-item-subtitle>
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
                v-on:click="goToProperty(property)"
                :key="i"
                link
            >
              <v-list-item-title v-text="property.address"></v-list-item-title>
              <v-list-item-subtitle v-text="property.postalCode"></v-list-item-subtitle>

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
      </v-card-text>
    </v-card>
  </v-main>
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
        ID
        area
        address
        postalCode
      }
      documents {
        url
        type
      }
      customerRef
      birthday
      tel
      gender
    }
  }
`;

@Component
export default class TenantProfile extends Vue {
  public tenant: any = {};
  public gender: string = '';
  public birthday: string = '';

  beforeMount() {
    this.$apollo.getClient().query({
      query: TENANT_QUERY,
      variables: {id: this.$route.params.id}
    }).then((res) => {
      const genders = {
        'MAN': 'male',
        'WOMAN': 'female',
        'OTHER': 'bottts',
      };
      this.tenant = res.data.tenant;
      this.gender = genders[this.tenant.gender];
      this.birthday = this.tenant.birthday.split('T')[0];
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

  goToProperty(property: any) {
    this.$router.push('/property/' + property.ID);
  }

  data() {
    return {
      banner: 'https://image.freepik.com/photos-gratuite/agent-immobilier-presentant-consultant-client-pour-prise-decision-signe-contrat-formulaire-assurance_1150-15023.jpg'
    }
  }
}

</script>

<style>
.v-card {
  margin-top: 4rem;
}
</style>
