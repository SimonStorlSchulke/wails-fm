import { writable } from 'svelte/store';

export const selectedFilePaths = writable<Array<any>>();
export const zoomLevel = writable<number>();
export const selectedSidebarFolder = writable("");

export let selectedSidebarFolderB = "";

export function SetSidebarFolder(f) {
    selectedSidebarFolderB = f;
}