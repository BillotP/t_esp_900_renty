<template>
  <div>
    <h1>Company</h1>
    <h2>{{connected}}</h2>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Action, Getter, namespace } from "vuex-class";

const authModule = namespace("authentificationStore");

@Component
export default class Auth extends Vue {
  @authModule.Getter("getConnected")
  private connected!: boolean;
}
</script>

<style>
</style>