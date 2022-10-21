<script lang="ts">
    import {FormatDate, OpenWithDefaultApp } from '../wailsjs/go/main/App.js'
	import { onMount } from 'svelte';
    import { createEventDispatcher } from 'svelte';

    export let fileinfo;
    export let isDir: boolean;

    let datestring: string = ""
    onMount(async () => {
        FormatDate(fileinfo.ModDate, "2006-01-02").then(dStr => {
            datestring = dStr;
        })
	});

    const dispatch = createEventDispatcher();

    function openCard() {
        if (isDir) {
            dispatch('message', {
                foldername: fileinfo.Name
            });
        } else {
            console.log("sss")
            OpenWithDefaultApp(fileinfo.FullPath)
        }
    }
</script>

<main on:dblclick={openCard} class={isDir ? 'file-card dir' : 'file-card'}>
    <p class="filename">{fileinfo.Name}</p>
        {#if !isDir}<p class="filesize">{fileinfo.Size}</p>{/if}
    <p class="modified">{datestring}</p>
</main>

<style>
</style>