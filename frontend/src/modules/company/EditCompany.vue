<template>
  <v-main>
    <v-card style="margin: 10px">
      <v-card-title class="justify-center"
        ><b style="text-transform: capitalize">{{ companyName }} 🖊️</b>
      </v-card-title>
      <v-card-text>
        <v-row>
          <v-text-field
            hint="Update your company name"
            v-model="companyName"
            label="Name"
          ></v-text-field>
        </v-row>
        <v-row>
          <v-text-field
            hint="Update your company description"
            v-model="companyDescription"
            label="Description"
          ></v-text-field>
        </v-row>
        <v-row>
          <v-img
            v-if="companyLogoUrl"
            style="margin-right: 10px; border: 1px solid black"
            max-height="64px"
            max-width="64px"
            :src="companyLogoUrl"
            left
          ></v-img>
          <v-text-field
            hint="Update your company logo URL"
            v-model="companyLogoUrl"
            label="Logo URL"
          ></v-text-field>
        </v-row>
        <v-row>
          <v-text-field
            hint="Update your company phone number"
            v-model="companyPhoneNumber"
            label="Phone number"
          ></v-text-field>
        </v-row>
        <v-btn
          depressed
          block
          color="primary"
          style="margin: 10px;"
          :disabled="isLoading"
          v-on:click="update()"
        >
          {{ isLoading ? "" : "Update" }}</v-btn
        >
      </v-card-text>
    </v-card></v-main
  >
</template>

<script>
export default {
  methods: {
    udpate() {
      console.log("Update !");
    },
  },
  data() {
    return {
      companyDescription: "",
      companyLogoUrl: "",
      companyPhoneNumber: "",
      companyName: localStorage.getItem("username"),
      isLoading: false,
    };
  },
};
</script>