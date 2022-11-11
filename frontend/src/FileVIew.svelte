<script lang="ts">
    import iconLink from "./assets/images/open_folder.png";
    import { selectedFilePaths } from "./AppState";
    import { zoomLevel } from "./AppState";
    import { FileSizeText, MenuItem } from "./FMUtils";
    import { CustomRightclickDialog } from "./FMUtils";

    import {
        FormatDate,
        OpenWithDefaultApp,
        GetThumbnailAsBase64,
        RenameFile,
        GetFileMeta
    } from "../wailsjs/go/main/App.js";
    import { onMount } from "svelte";
    import { createEventDispatcher } from "svelte";

    export let fileinfo;
    export let isDir: boolean;
    let selected: boolean = false;
    let thumbnailData: string;
    let hasThumbnail: boolean;

    let datestring: string = "";

    let FWElement: HTMLElement;

    
    onMount(async () => {

        CustomRightclickDialog(FWElement, [
            new MenuItem("Open", () => { openFile() }),
            new MenuItem("Rename", () => { RenameFile(fileinfo.FullPath, fileinfo.FullPath + ".txt") }), //TODO
        ]);

        FormatDate(fileinfo.ModDate, "2006-01-02").then((dStr) => {
            datestring = dStr;
        });
        if (fileinfo.Hidden) {
            card_class += " file-hidden";
        }
        if (isDir) {
            if (fileinfo.IsLink) {
                card_class += " file-link";
            }
        }

        let thumbnailFormats: string[] = [".jpg", ".jpeg", ".png", ".gif"];

        if (fileinfo.Extension) {
            if (
                thumbnailFormats.includes(
                    (fileinfo.Extension as string).toLowerCase()
                )
            ) {
                GetThumbnailAsBase64(fileinfo.FullPath).then((res) => {
                    thumbnailData = "data:image/png;base64," + res;
                    hasThumbnail = true;
                });
            }
        }
    });

    const dispatch = createEventDispatcher();

    function openFile() {
        if (isDir) {
            dispatch("message", {
                foldername: fileinfo.FullPath,
            });
        } else {
            OpenWithDefaultApp(fileinfo.FullPath);
        }
    }

    function toggleSelectCard(e: MouseEvent) {
        selected = !selected;
        if (selected) {
            (e.currentTarget as HTMLElement).classList.add("file-selected");
        } else {
            (e.currentTarget as HTMLElement).classList.remove("file-selected");
        }
        UpdateSelectedFilePaths();
    }

    // Filepath = ID of Element
    export function UpdateSelectedFilePaths() {
        let els = document.querySelectorAll(".file-selected");
        let list: Array<string> = new Array<string>(els.length);
        for (let i = 0; i < els.length; i++) {
            list[i] = els[i].id;
        }
        selectedFilePaths.set(list);
    }

    let card_class = isDir ? "dir" : "";

    zoomLevel.subscribe(zoom => {
        card_class = isDir ? "dir" : "";
        
		if (zoom > 18) {
            card_class += " file-card big"
        } else if (zoom > 13) {
            card_class += " file-card"
        } else if (zoom > 8) {
            card_class += " file-item big"
        } else if (zoom > 4) {
            card_class += " file-item"
        }
		else {
            card_class += " file-minimal"
        }
        if (fileinfo.Hidden) {
            card_class += " file-hidden"
        }
	});

    const handleloaded = el => {
        GetFileMeta(fileinfo.QID).then( meta => {
            if (fileinfo.IsLink || isDir) {
                (el as HTMLElement).style.backgroundImage = `linear-gradient(90deg, #0000 40%, ${meta.Color})`;
            } else {
                // el.style.color = meta.Color;
                (el as HTMLElement).style.backgroundImage = `linear-gradient(90deg, #0000 40%, ${meta.Color})`;
            }
        })
    }

</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<div
    id={fileinfo.FullPath}
    on:click={(e) => toggleSelectCard(e)}
    on:dblclick={openFile}
    bind:this={FWElement}
    use:handleloaded

    class="{card_class}"
>
    {#if fileinfo.IsLink}<img
            alt="icon"
            class="link-icon"
            src={iconLink}
        />{/if}
    {#if hasThumbnail}
        <img alt="a" class="thumbnail" src={thumbnailData} />
    {/if}
    <div class="file-info">
    <p class="filename">{fileinfo.Name}</p>
    {#if !isDir}
        <p class="filesize">{FileSizeText(fileinfo.Size)}</p>
    {/if}
    <p class="modified">{datestring}</p>
</div>
</div>

<style>

    .file-minimal {
        display: inline-block;
        position: relative;
        margin: 0px;
        background-color: rgba(0, 0, 0, 0.521);
        width: 110px;
        height: 16px;
        max-height: 16px;
        border-width: 0px 0px 0px 5px;
        border-style: none;
        padding: 2px 10px;
        white-space: nowrap;
        z-index: 0;
        outline: 1px solid rgba(0, 0, 0, 0.333);
        font-size: 12px;
    }

    .file-minimal .filename {
        overflow: hidden;
        text-overflow: ellipsis;
    }

    .file-minimal .filesize, .modified {
        display: none;
    }
    
    .file-item {
        font-size: 19px;
        display: block;
        position: relative;
        background-color: rgba(0, 0, 0, 0.521);
        width: calc(100% - 20px);
        height: 14px;
        max-height: 14px;
        overflow: hidden;
        border-width: 0px 0px 0px 5px;
        border-style: none;
        padding: 2px 10px;
        white-space: nowrap;
        z-index: 0;
        outline: 1px solid rgba(0, 0, 0, 0.598);
    }

    .file-item.big {
        height: 35px;
        max-height: 35px;
        padding: 10px;
    }

    .file-item .thumbnail {
        float: right;
        width: 40px;
        height: 100%;
        top: 0px;
        z-index: -1;
        margin-left: -10px;
        object-fit: cover;
    }

    .file-minimal .thumbnail {
        float: right;
        width: 30px;
        height: 100%;
        top: 0px;
        z-index: -1;
        margin-left: -10px;
        object-fit: cover;
    }

    .file-item.big .thumbnail {
        width: 60px;
    }

    .file-item .file-info {
        position: absolute;
        left: 50px;
        bottom: 0px;
        width: 100%;
        font-size: 12px;
    }

    .file-item.big .file-info {
        position: absolute;
        left: 20px;
        bottom: 0px;
        width: 100%;
        padding: 14px 6px;
        font-size: 14px;
    }
    
    .file-item p {
        min-width: 30%;
        display: inline-block;
    }
    
    .file-item-big p {
        min-width: 30%;
        display: inline-block;
    }

    
    .file-minimal .link-icon, .file-item .link-icon {
        float: right;
        width: 17px;
        margin-top: -2px;
    }
    
    .file-item.big .link-icon {
        float: right;
        width: 30px;
        margin-top: 4px;
    }

    .link-icon {
        float: right;
        width: 24px;
        opacity: 0.7;
    }

    .file-card .thumbnail {
        position: absolute;
        float: left;
        top: 0px;
        z-index: -1;
        width: 100%;
        height: 100%;
        margin-left: -10px;
        object-fit: cover;
    }

    .file-card.file-selected {
        outline: 2px solid rgb(82, 144, 214);
    }

    .file-card {
        font-size: 15px;
        display: inline-block;
        position: relative;
        margin: 5px;
        background-color: black;
        width: 145px;
        height: 120px;
        overflow: hidden;
        border-width: 0px 0px 0px 5px;
        border-style: none;
        padding: 10px;
        white-space: nowrap;
        box-shadow: 0px 0px 12px rgba(0, 0, 0, 0.402);
        z-index: 0;
    }

    .file-card.big {
        width: 280px;
        height: 170px;
    }

    .file-card .file-info {
        position: absolute;
        bottom: -12px;
        margin: -10px;
        width: 100%;
        height: 70px;
        padding-top: 20%;
        background-color: linear-gradient(#0000, rgba(0, 0, 0, 0.651) 30px);
        padding: 14px 6px;
    }

    .file-card.dir .file-info {
        height: 40px;
    }

    .file-selected .file-info {
        background-image: linear-gradient(#0000, rgba(24, 110, 215, 0.228) 50px);
    }

    .file-selected.file-minimal, .file-selected.file-item {
        background-image: linear-gradient(rgb(49, 78, 119), rgb(45 90 140)) !important;
    }

    .file-card .filename {
        max-width: calc(100% - 20px);
        overflow: hidden;
        text-overflow: ellipsis;
        font-weight: bold;
        color: white;
    }

    .file-hidden {
        opacity: 0.7;
    }

    .dir {
        background-color: rgba(255, 255, 255, 0.205);
        height: 40px !important;
        border-color: transparent;
    }


    .dir:after {
        content: '';
        width: 0;
        height: 0;
        border-style: solid;
        border-width: 7px 7px 0px 0px;
        border-color: rgb(24, 38, 59) transparent transparent transparent;
        top: 0;
        left: 0px;
        position: absolute;
    }

    .file-item.big:after {
        border-width: 15px 15px 0px 0px; 
    }

    .file-card.dir.big {
        width: 125px !important;
    }
</style>
