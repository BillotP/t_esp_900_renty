<template>
  <div>
    <navBar-field></navBar-field>
    <router-view></router-view>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { RootState } from "@/store/types";

@Component({})
export default class Root extends Vue {}
</script>
