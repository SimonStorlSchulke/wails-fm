<script lang="ts">
    import {TreeView} from 'svelte-tree-view'
    import { GetTree } from "../wailsjs/go/main/App.js";

    function OnClickItem(e) {
        let el = (e.target as HTMLElement)
        console.log(el.innerText)
    }

    let data = {
		a: [1, 2, 3],
		b: new Map([
		['c', { d: null }],
			['e', { f: [9, 8, 7] }]
		])
	}

    GetTree("C:/").then((mp) => {
        data = mp
    })

    function UrlFromIDPath(tv: TreeView, n: any) {
        console.log(tv)
    }
    
</script>

<div class="wrapper" style="float: left;">
<TreeView data={data}
recursionOpts={{
    shouldExpandNode: (n) => {UrlFromIDPath(this, n.path); return false}
  }}/>
</div>

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
