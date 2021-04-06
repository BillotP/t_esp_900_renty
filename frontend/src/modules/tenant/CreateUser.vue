<template>
  <v-main
      style="
        margin-top: 10vh;
        margin-bottom: 10vh;
        margin-right: 10vh;
        margin-left: 10vh;
      "
  >
    <v-card>
      <v-card-text>
        <h3>Tenant</h3>
        <v-container fluid>
          <v-form v-model="valid">
            <h4>My Profile</h4>
            <v-row>
              <v-col cols="6">
                <v-text-field
                    prepend-icon="mdi-barcode"
                    type="text"
                    v-model="customerRef"
                    label="Customer Ref°"
                    required
                    :rules="rules"
                ></v-text-field>
              </v-col>
              <v-col cols="2">
                <v-radio-group v-model="gender">
                  <v-radio
                      label="👦"
                      value="0"
                  ></v-radio>
                  <v-radio
                      label="👩"
                      value="1"
                  ></v-radio>
                  <v-radio
                      label="😈"
                      value="2"
                  ></v-radio>
                </v-radio-group>
              </v-col>
            </v-row>
            <v-row>
              <v-col cols="6">
                <v-text-field
                    prepend-icon="mdi-phone"
                    type="phone"
                    v-model="phone"
                    label="Phone"
                    required
                    :rules="rules"
                ></v-text-field>
              </v-col>
              <v-col cols="6">
                <v-text-field
                    prepend-icon="mdi-cake"
                    v-model="birthday"
                    label="Birthday"
                    type="date"
                    required
                    :rules="rules"
                ></v-text-field>
              </v-col>
            </v-row>
            <h4>My Account</h4>
            <v-row>
              <v-text-field
                  prepend-icon="mdi-account"
                  v-model="username"
                  label="Username"
                  required
                  :rules="rules"
              ></v-text-field>
            </v-row>
            <v-row>
              <v-text-field
                  prepend-icon="mdi-form-textbox-password"
                  type="password"
                  v-model="password"
                  label="Password"
                  required
                  :rules="rules"
              ></v-text-field>
            </v-row>
            <v-row>
              <v-btn depressed color="primary" :disabled="!valid" v-on:click="createUser()"> Create</v-btn>
            </v-row>
          </v-form>
        </v-container>
      </v-card-text>
    </v-card>
  </v-main>
</template>

<script>
import Vue from "vue";
import Component from "vue-class-component";
import gql from "graphql-tag";

const CREATE_TENANT_USER_MUTATION = gql`
  mutation($input: TenantInput) {
    createTenantUser(input: $input) {
      ID
    }
  }
`;

@Component
export default class CreateTenantUser extends Vue {
  data() {
    return {
      phone: null,
      birthday: null,
      customerRef: null,
      gender: '0',
      username: null,
      password: null,
      valid: false,
      rules: [
        v => !!v || 'field is required',
      ],
    };
  }

  async createUser() {
    try {
      const genders = [
        'MAN',
        'WOMAN',
        'OTHER'
      ];
      const resp = await this.$apollo.getClient().mutate({
        mutation: CREATE_TENANT_USER_MUTATION,
        variables: {
          input: {
            user: {
              username: this.$data.username,
              password: this.$data.password,
            },
            tel: this.$data.phone,
            birthday: new Date(this.$data.birthday).toISOString(),
            gender: genders[+this.$data.gender],
            customerRef: this.$data.customerRef,
          },
        },
      });
      if (resp.data.createTenantUser.ID) {
        this.$data.text =
            "User " + this.$data.username + " create successfully !";
        this.$data.snackbar = true;
        this.$router.back();
      }
    } catch (e) {
      console.error(e);
    }
  }
}
</script>

<style>
.v-radio > label {
  font-size: 30px;
  color: rgba(0, 0, 0, 1) !important;
}
</style>
