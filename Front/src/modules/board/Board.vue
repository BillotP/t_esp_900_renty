<template>
  <div>
    <hello-world></hello-world>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { board } from "./store";
import { Action, Getter, namespace } from "vuex-class";
import HelloWorldVue from "@/components/HelloWorld.vue";

const boardModule = namespace("board");

@Component({
  components: {
    HelloWorld: HelloWorldVue
  }
})

export default class Board extends Vue {
  private nbrCaseByLign = 5;

  @boardModule.Action("save")
  private saveBoard!: (board: number[][]) => void;

  @boardModule.Action("changeSave")
  private changeSave!: (changeSave: number) => void;

  @boardModule.Getter("getBoard")
  private board!: number[][];

  @boardModule.Getter("getSave")
  private nbrSave!: number;

  @boardModule.Getter("getScore")
  private score!: number;

  @boardModule.Getter("getTour")
  private tour!: number;

  @boardModule.Action("addScore")
  private addScore!: (score: number) => void;

  @boardModule.Action("addTour")
  private addTour!: (tour: number) => void;

  private save() {
    this.saveBoard(this.board);
  }

  private addJeton(ligne: number[]) {
    for (let i = 0; i < ligne.length; i++) {
      if (ligne[i] === 0) {
        const jeton = Math.pow(2, Math.floor(Math.random() * 10));
        ligne.splice(i, 1, jeton);
        this.addScore(jeton);
        this.addTour(1);
        return;
      }
    }
  }
}
</script>

<style>
</style>