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
        RenameFile
    } from "../wailsjs/go/main/App.js";
    import { onMount } from "svelte";
    import { createEventDispatcher } from "svelte";


    var temp_dropPath = "";//open window, C# will change 「temp_dropPath」 

    /**
     * Get the dragged file path 
     * @returns 
     */
    async function getDropPath() {

        let _dropPath = "";
        for (let i = 0; i < 100; i++) {

            if (temp_dropPath !== "") {
                _dropPath = temp_dropPath;
                _dropPath = decodeURIComponent(temp_dropPath)
                if (_dropPath.indexOf("file:///") === 0) {//  file:///C:/b.jpg -> C:/b.jpg
                    _dropPath = _dropPath.substring(8);
                } else if (_dropPath.indexOf("file://") === 0) {// file://Desktop-abc/AA/b.jpg -> //Desktop-abc/AA/b.jpg
                    _dropPath = _dropPath.substring(5);
                }
                break;
            }
            await Sleep(10);
        }
        temp_dropPath = "";
        _dropPath = _dropPath.replace(/[/]/g, "\\");
        return _dropPath;
    }

    window.onload = function() {
        addDropEvent(document.getElementById("box1"))
    }

    function addDropEvent(dom) {
        console.log("Adding Event")
        dom.addEventListener("dragenter", dragenter, false);
        dom.addEventListener("dragover", dragover, false);
        dom.addEventListener("drop", drop, false);

        function dragenter(e) {
            console.log("dragenter");
            e.stopPropagation();
            e.preventDefault();
        }
        function dragover(e) {
            console.log("dragover");
            e.stopPropagation();
            e.preventDefault();
        }
        async function drop(e) {
            e.stopPropagation();
            e.preventDefault();
            console.log("drop");
            if (e.dataTransfer === null) { return; }

            let files: File[] = e.dataTransfer.files;
            console.log("File: ",files);

            let _dropPath = await getDropPath();
            console.log("droppath: ", _dropPath);
            //if (_dropPath === "") { return; }

            if (files.length > 1) {
                console.log("It");
                let dirPath = GetDirectoryName(_dropPath);
                let html = "";
                for (let i = 0; i < files.length; i++) {
                    const item: File = files[i];
                    console.log("File: ", Combine([dirPath, item.name]));
                }
                dom.innerHTML = html;
            } else {
                dom.innerHTML = _dropPath;
            }

        }

    }

        async function Sleep(ms) {
            await new Promise((resolve, reject) => {
                setTimeout(function () {
                    resolve(0);
                }, ms);
            })
        }


        const GetDirectoryName = (path) => {
            path = path.replace(/[/]/ig, "\\");

            let count = path.split("\\").length - 1;// count 「\」

            if (count === 0) { return null; }
            if (path.length <= 2) { return null; }

            if (path.substring(0, 2) === "\\\\") {//ex \\Desktop-abc\AA
                if (count <= 3) {//At least 2 layers 
                    return null;
                }
                if (count === 4) {
                    if (path.lastIndexOf("\\") === path.length - 1) {//ex \\Desktop-abc\AA\
                        return null;
                    }
                }
            } else {
                if (count === 1) {
                    if (path.lastIndexOf("\\") === path.length - 1) {//ex D:/
                        return null;
                    }
                }
            }

            let name = path.substring(0, path.lastIndexOf("\\"));//get Dir path

            //avoid  D:\ to D:
            count = name.split("\\").length - 1;
            if (count === 0) {
                name = name + "\\";
            }

            return name;
        }


        const Combine = (arPath) => {
            if (arPath.length === 0) { return "" }
            if (arPath.length === 1) { return arPath[0] }

            let sum = arPath[0];
            sum = sum.replace(/[\\]+$/, '');
            sum += "\\"

            for (let i = 1; i < arPath.length; i++) {
                let item = arPath[i];
                item = item.replace(/^([\\])+/, '');
                item = item.replace(/[\\]+$/, '');
                sum += item
                if (i != arPath.length - 1) { sum += "\\"; }
            }
            return sum
        }
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<div>
    <div class="box" id="box1"> box1</div>
</div>

<style>
.box {
    width: 200px;
    height: 200px;
}
</style>
