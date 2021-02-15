<template>
  <div>
    <div v-if="!getRoute()">
      <navBar-field></navBar-field>
    </div>
    <div>
      <router-view></router-view>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { RootState } from "@/store/types";

@Component({})
export default class Root extends Vue {

  private getRoute() {
    return this.$route.path.toString().includes("login");
  }

  public mounted()
  {
    
  }
}
</script>
