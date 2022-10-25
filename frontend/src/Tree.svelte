<script lang="ts">
  import { GetTree, GetTreeHTML } from "../wailsjs/go/main/App.js";
  import App from "./App.svelte";

  let treeStruct: Array<any> = [];

  function WalkTree() {}


  GetTree("C:/").then((mp) => {
  });
</script>

<div class="wrapper" style="float: left;" />

<style>
  .wrapper {
    --tree-view-base00: #363755;
    --tree-view-base01: #604d49;
    --tree-view-base02: #6d5a55;
    --tree-view-base03: #d1929b;
    --tree-view-base04: #b79f8d;
    --tree-view-base05: #f9f8f2;
    --tree-view-base06: #f7f4f1;
    --tree-view-base07: #faf8f5;
    --tree-view-base08: #fa3e7e;
    --tree-view-base09: #fd993c;
    --tree-view-base0A: #f6bf81;
    --tree-view-base0B: #f60000;
    --tree-view-base0C: #b4efe4;
    --tree-view-base0D: #85d9ef;
    --tree-view-base0E: #be87ff;
    --tree-view-base0F: #d6724c;
  }
</style>
