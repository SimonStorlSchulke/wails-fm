<script lang="ts">
    import iconLink from "./assets/images/open_folder.png";
    import { selectedFilePaths } from "./AppState";
    import { zoomLevel } from "./AppState";


    import {
        FormatDate,
        OpenWithDefaultApp,
        GetThumbnailAsBase64,
    } from "../wailsjs/go/main/App.js";
    import { onMount } from "svelte";
    import { createEventDispatcher } from "svelte";

    export let fileinfo;
    export let isDir: boolean;
    let selected: boolean = false;
    let thumbnailData: string;
    let hasThumbnail: boolean;

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
        
		if (zoom > 15) {
            card_class += " file-card big"
        } else if (zoom > 10) {
            card_class += " file-card"
        } else if (zoom > 5) {
            card_class += " file-item big"
        }
		else {
            card_class += " file-item"
        }
	});

</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<div
    id={fileinfo.FullPath}
    on:click={(e) => toggleSelectCard(e)}
    on:dblclick={openCard}
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
        <p class="filesize">{fileinfo.QID}</p>
    {/if}
    <p class="modified">{datestring}</p>
</div>
</div>

<style>
    
    .file-item {
        display: block;
        position: relative;
        margin: 5px;
        background: rgba(0, 0, 0, 0.521);
        width: calc(100% - 50px);
        height: 14px;
        max-height: 14px;
        overflow: hidden;
        border-width: 0px 0px 0px 5px;
        border-style: none;
        padding: 2px 10px;
        white-space: nowrap;
        box-shadow: 0px 0px 12px rgba(0, 0, 0, 0.402);
        z-index: 0;
    }

    .file-item.big {
        height: 35px;
        max-height: 35px;
        padding: 10px;
    }

    .file-item .thumbnail {
        position: absolute;
        float: left;
        width: 40px;
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
        font-size: 13px;
    }

    .file-item.big .file-info {
        position: absolute;
        left: 80px;
        bottom: 0px;
        width: 100%;
        padding: 14px 6px;
    }
    
    .file-item p {
        min-width: 30%;
        display: inline-block;
    }
    
    .file-item-big p {
        min-width: 30%;
        display: inline-block;
    }

    .link-icon {
        float: right;
        width: 24px;
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

    .file-selected {
        outline: 2px solid rgb(82, 144, 214);
    }

    .file-card {
        display: inline-block;
        position: relative;
        margin: 5px;
        background: black;
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
        background-image: linear-gradient(#0000, rgba(0, 0, 0, 0.651) 30px);
        padding: 14px 6px;
    }

    .file-selected .file-info {
        background-image: linear-gradient(#0000, rgba(24, 110, 215, 0.228) 50px);
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
        background-color: rgb(103, 86, 35);
        height: 40px !important;
        border-color: transparent;
        background-image: none !important;
    }

    .dir .file-info {
        background-image: none !important;
    }

    .file-card.dir.big {
        width: 125px !important;
    }
</style>
