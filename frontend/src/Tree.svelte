<script lang="ts">
    import TreeView from 'svelte-tree-view'
    import { GetTree } from "../wailsjs/go/main/App.js";


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
    
</script>

<div style="float: left;">
<TreeView data={data}/>
</div>

<style>
ul {
    list-style-type: none !important;
    display: block;
    padding: 0px 20px;
}

li {
    display: block;
}

span {
    min-width: 100%;
    padding: 0px 0px;
    margin-top: 20px;
}

span:hover {
    color: white
}
</style>
