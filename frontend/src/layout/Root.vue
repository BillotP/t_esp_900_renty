<template>
  <div>
    <div v-if="!getRoute()">
      <navBar-field></navBar-field>
    </div>
    <div v-else>
      <navBarUnlogged-field />
    </div>
    <div>
      <router-view></router-view>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";

@Component({})
export default class Root extends Vue {
  public getRoute() {
    return this.$route.path.toString().includes("login");
  }

  public mounted() {}
}
</script>
