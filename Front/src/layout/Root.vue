<template>
      <router-view></router-view>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { RootState } from "@/store/types";

@Component({
})
export default class Root extends Vue {
}
</script>
