<script context="module">
    // retain module scoped expansion state for each tree node
    const _expansionState = {
        /* treeNodeId: expanded <boolean> */
    };
</script>

<script lang="ts">
    export let tree;
    const { label, children } = tree;

    let expanded = _expansionState[label] || false;
    const toggleExpansion = () => {
        expanded = _expansionState[label] = !expanded;
    };
    $: arrowDown = expanded;
</script>

<ul>
    <li>
        {#if children}
            <span on:click={toggleExpansion}>
                <span class="arrow" class:arrowDown>&#x25b6</span>
                {label}
            </span>
            {#if expanded}
                {#each children as child}
                    <svelte:self tree={child} />
                {/each}
            {/if}
        {:else}
            <span>
                <span class="no-arrow" />
                {label}
            </span>
        {/if}
    </li>
</ul>

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
