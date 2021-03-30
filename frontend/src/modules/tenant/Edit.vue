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
            v-on:click="goToProperty(property)"
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
            <v-list-item-title>
              Documents
            </v-list-item-title>
            <v-dialog v-model="dialog">
              <v-card>
                <v-card-title class="headline">
                  Add document to your profile:
                </v-card-title>

                <v-card-text>
                  <v-spacer></v-spacer>
                  <v-text-field prepend-icon="mdi-file-question-outline" label="Main input"
                                hide-details="auto" v-model="docNameToUpload"></v-text-field>
                  <v-file-input prepend-icon="mdi-paperclip" label="File input"
                                show-size v-model="docToUpload"></v-file-input>
                </v-card-text>

                <v-card-actions>
                  <v-spacer></v-spacer>

                  <v-btn color="error" text @click="dialog = false">
                    Cancel
                  </v-btn>

                  <v-btn color="primary" text @click="uploadDocument">
                    Add
                  </v-btn>
                </v-card-actions>
              </v-card>
            </v-dialog>
          </v-list-item-content>
          <v-icon id="uploadDoc"
                  right
                  color="grey"
                  @click.stop="dialog = true"
          >
            mdi-plus-circle-outline
          </v-icon>

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
        ID
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

const UPLOAD_DOCUMENT = gql`
  mutation uploadDocument($file: Upload!, $title: String!) {
    uploadDocument(file: $file, title: $title)
  }
`;

@Component
export default class TenantEdit extends Vue {
  public tenant: any = {};
  public docToUpload: any = {};
  public docNameToUpload: string = '';
  public dialog = false;

  beforeMount() {
    this.fetchTenant();
  }

  fetchTenant() {
    this.$apollo.getClient().query({
      query: TENANT_QUERY,
      variables: {id: this.$route.params.id},
      fetchPolicy: "network-only"
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

  uploadDocument() {
    this.dialog = false;
    this.$apollo.getClient().mutate({
      mutation: UPLOAD_DOCUMENT,
      variables: {file: this.docToUpload, title: this.docNameToUpload}
    }).then((res) => {
      console.log(res);
      this.docToUpload = {};
      this.docNameToUpload = '';
      this.fetchTenant();
    }).catch((err) => {
      console.error(err);
    });
  }

  goToProperty(property: any) {
    this.$router.push('/property/' + property.ID);
  }
}

</script>

<style>
template > .v-card {
  margin-top: 4rem;
}

#uploadDoc:hover {
  opacity: 0.6;
  zoom: 1.5;
}
</style>
