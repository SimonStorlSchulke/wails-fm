<script lang="ts">
    import { GetSubDirPaths } from "../wailsjs/go/main/App.js";
    import Tree from "./Tree.svelte"
	export let tree
	const {label, children} = tree

    let mounts: string[]

    GetSubDirPaths("A").then((mps) => {
        mounts = mps
    })

    let selectedMount;
</script>

<div class="sidebar">
    <fieldset class="radio">
        {#if mounts}
        {#each mounts as mount}
        <div>
            <input bind:group={selectedMount} type="radio" id="mount-{mount.substr(0, 1)}" name="a" value={mount.substr(0, 1)}>
            <label for="mount-{mount.substr(0, 1)}" >{mount.substr(0, 1)}</label>
        </div>
        {/each}
        {/if}
          
    </fieldset>

    <Tree {tree}/>
</div>

<style>


.radio {
    float: left;
    background-color: #0005;
    width: 42px;
    font-size: 18px;
    font-weight: bold;
    text-align: center;
    height: 100%;
    border: none;
    padding: 0px;
    margin: 0px 10px 0px 0px;
}

label {
    width: 30px;
    margin: 0px;
    width: 100%;
    display: block;
}

input:checked + label {
    background-color: rgba(19, 120, 214, 0.541);
    color: white;
}


input {
    display: none;
}



</style>
