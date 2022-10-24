<script lang="ts">
    import { onMount } from "svelte";
    import { selectedFilePaths } from "./AppState";
    import { GetFileDetailsSingle } from "../wailsjs/go/main/App.js";

    let fileDetails;

    selectedFilePaths.subscribe((selFiles) => {
        if (!selFiles) {return}
        if (selFiles.length == 1) {
            GetFileDetailsSingle(selFiles[0]).then((details) => {
                fileDetails = details
            })
        } else if (selFiles.length > 1) {
            fileDetails = "multiple Selected";
        } else {
            fileDetails = "Nothing Selected";
        }
    });
</script>

<div class="inspector">
    {#key fileDetails}
    {#if fileDetails}
    <h3>{fileDetails.Name}</h3>
    <p>{fileDetails.MimeType}</p>
    <p>{fileDetails.CreationTime}</p>
    <p>{fileDetails.ModifiedTime}</p>
    {/if}
    {/key}
</div>

<style>
    .inspector {
        padding: 10px;
    }
</style>
