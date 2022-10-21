<script lang="ts">
    import { onMount } from "svelte";
    import { selectedFilePaths } from "./AppState";
    import { GetFileDetailsSingle } from "../wailsjs/go/main/App.js";

    let fileDetails: string = ""


    selectedFilePaths.subscribe((selFiles) => {
        if (!selFiles) {return}
        if (selFiles.length == 1) {
            GetFileDetailsSingle(selFiles[0]).then((details) => {
                fileDetails = details.MimeType
            })
        } else if (selFiles.length > 1) {
            fileDetails = "multiple Selected";
        } else {
            fileDetails = "Nothing Selected";
        }
    });
</script>

<div class="inspector">
    {fileDetails}
</div>

<style>
</style>
