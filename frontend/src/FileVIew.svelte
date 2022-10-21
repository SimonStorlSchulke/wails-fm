<script lang="ts">
    import {FormatDate} from '../wailsjs/go/main/App.js'
	import { onMount } from 'svelte';
    import { createEventDispatcher } from 'svelte';

    export let fileinfo;
    export let isDir: boolean;

    let datestring: string = ""
    onMount(async () => {
        console.log(fileinfo)
        FormatDate(fileinfo.ModDate, "2006-01-02").then(dStr => {
            datestring = dStr;
        })
	});

    const dispatch = createEventDispatcher();

    function enterFolder() {
        dispatch('message', {
            foldername: fileinfo.Name
        });
    }

</script>

<main on:dblclick={enterFolder} class={isDir ? 'file-card dir' : 'file-card'}>
    <p class="filename">{fileinfo.Name}</p>
        {#if !isDir}<p class="filesize">{fileinfo.Size}</p>{/if}
    <p class="modified">{datestring}</p>
</main>

<style>
</style>