<script lang="ts">
    import { onMount } from "svelte";
    import { selectedFilePaths } from "./AppState";
    import { GetFileDetailsSingle } from "../wailsjs/go/main/App.js";

    let fileDetails;
    let selFilesNum = 1;

    selectedFilePaths.subscribe((selFiles) => {
        if (selFiles) {
            selFilesNum = selFiles.length;
        }
        if (!selFiles) {return}
        if (selFilesNum == 1) {
            GetFileDetailsSingle(selFiles[0]).then((details) => {
                fileDetails = details
            })
        } else if (selFilesNum > 1) {
            fileDetails = "multiple Selected";
        } else {
            fileDetails = "Nothing Selected";
        }
    });
</script>

<div class="inspector">
    {#key fileDetails}
        {#if fileDetails}
            {#if fileDetails == "multiple Selected"}
                <p> {selFilesNum} Files selected</p>
            {:else if  fileDetails == "Nothing Selected"}
                <p>Nothing selected</p>
            {:else}
                <h3>{fileDetails.Name}</h3>
                <p>{fileDetails.MimeType}</p>
                <p>{fileDetails.CreationTime}</p>
                <p>Modified: {fileDetails.ModifiedTime}</p>
                <p>ID: {fileDetails.QID}</p>
            {/if}
        {/if}
    {/key}
</div>

<style>
    .inspector {
        padding: 10px;
        min-width: 300px;
        max-width: 300px;
    }

    h3 {
        word-wrap: break-word;
    }
</style>
