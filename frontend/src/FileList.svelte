<script lang="ts">
    let folderData;
    export let showhiddenFiles: boolean = false;
    import FileVIew from "./FileVIew.svelte";
    import { selectedSidebarFolder } from "./AppState";

    import { GetFolderAPI } from "../wailsjs/go/main/App.js";

    export let fileViewAddress: string;

    function sanitizePath(path: string): string {
        path = path.replace(/([^:]\/)\/+/g, "$1");
        path = path.replaceAll("\\", "/");
        if (!path.endsWith("/")) {
            path += "/";
        }
        return path;
    }

    selectedSidebarFolder.subscribe((path) => {
            fileViewAddress = path;
            submitPath(false);
    });


    export function submitPath(moveUp: boolean): void {
        fileViewAddress = sanitizePath(fileViewAddress);
        if (moveUp) {
            let els: string[] = fileViewAddress.split("/");
            let num: number = els[els.length - 2].length + 1;
            fileViewAddress = fileViewAddress.slice(
                0,
                fileViewAddress.length - num
            );
        }

        GetFolderAPI(fileViewAddress).then((fi) => {
            if (fi.Files) {
                fi.Files.sort((a, b) => a.Name.localeCompare(b.Name));
            }
            if (fi.Folders) {
                fi.Folders.sort((a, b) => a.Name.localeCompare(b.Name));
            }
            folderData = fi;
        });
    }

    export function Home() {}

    function folderDoubleclicked(foldername) {
        fileViewAddress = foldername.detail.foldername;
        submitPath(false);
    }
    window.onload = function () {
        submitPath(false);
    };
</script>

<div class="file-area">
    {#key folderData}
        <section class="folders">
            {#if folderData}
                {#if folderData.Folders}
                    {#each folderData.Folders as folderInfo}
                        {#if !folderInfo.Hidden || showhiddenFiles}
                            <FileVIew
                                on:message={folderDoubleclicked}
                                fileinfo={folderInfo}
                                isDir={true}
                            />
                        {/if}
                    {/each}
                {/if}
            {/if}
        </section>

        <section class="files">
            {#if folderData}
                {#if folderData.Files}
                    {#each folderData.Files as fileInfo}
                        {#if !fileInfo.Hidden || showhiddenFiles}
                            <FileVIew fileinfo={fileInfo} isDir={false}/>
                        {/if}
                    {/each}
                {/if}
            {/if}
        </section>
    {/key}
</div>

<style>
    
.file-area {
    display: block;
    flex-grow: 1;
    flex-shrink: 1;
    flex-basis: auto;
    align-self: auto;
    order: 0;
    text-align: left;
    max-height: 100%;
    overflow-y: scroll;
    scrollbar-width: thin;
    scrollbar-color: black rgb(203, 27, 27) !important;
    scrollbar-arrow-color: black;
    font-size:0px;
}

.file-area .files {
    margin-top: 10px;
}
</style>
