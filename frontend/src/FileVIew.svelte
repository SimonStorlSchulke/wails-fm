<script lang="ts">
    import iconLink from './assets/images/open_folder.png'
    import { selectedFilePaths } from "./AppState"

    import { FormatDate, OpenWithDefaultApp } from "../wailsjs/go/main/App.js";
    import { onMount } from "svelte";
    import { createEventDispatcher } from "svelte";

    export let fileinfo;
    export let isDir: boolean;
    let selected: boolean = false

    let datestring: string = "";
    onMount(async () => {
        FormatDate(fileinfo.ModDate, "2006-01-02").then((dStr) => {
            datestring = dStr;
        });
        if (fileinfo.Hidden) {
            card_class += " file-hidden";
        }
        if (isDir) {
            if (fileinfo.IsLink) {
                console.log("IS DIR")
                card_class += " file-link"
            }
        }
    });

    const dispatch = createEventDispatcher();

    function openCard() {
        if (isDir) {
            dispatch("message", {
                foldername: fileinfo.FullPath,
            });
        } else {
            OpenWithDefaultApp(fileinfo.FullPath);
        }
    }

    function toggleSelectCard(e: MouseEvent) {
        selected = !selected
        if (selected) {
            (e.currentTarget as HTMLElement).classList.add("file-selected");
        } else {
            (e.currentTarget as HTMLElement).classList.remove("file-selected");
        }
        UpdateSelectedFilePaths()
    }

        // Filepath = ID of Element
    export function UpdateSelectedFilePaths() {
        let els = document.querySelectorAll(".file-selected")
        let list: Array<string> = new Array<string>(els.length)
        for (let i = 0; i < els.length; i++) {
            list[i] = els[i].id;
        }
        selectedFilePaths.set(list)
    }

    $: card_class = isDir ? "dir" : "b";
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<div id={fileinfo.FullPath} on:click={(e) => toggleSelectCard(e)} on:dblclick={openCard} class="file-card {card_class}">
    {#if fileinfo.IsLink}<img alt="icon" class="link-icon" src="{iconLink}">{/if}
    <p class="filename">{fileinfo.Name}</p>
    <i class="fa-solid fa-link"></i>
    {#if !isDir}
        <p class="filesize">{fileinfo.Size}</p>
    {/if}
    <p class="modified">{datestring}</p>

</div>

<style>
    .link-icon {
        float: right;
        width: 24px;
    }

    .file-selected {
        background-image: linear-gradient(rgba(49, 144, 240, 0.626), rgba(0, 128, 183, 0.21));
    }
</style>
