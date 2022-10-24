import { writable } from 'svelte/store';

export const selectedFilePaths = writable<Array<any>>();
export const zoomLevel = writable<number>();