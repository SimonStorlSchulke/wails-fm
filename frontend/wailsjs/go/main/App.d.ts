// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {time} from '../models';
import {main} from '../models';

export function FormatDate(arg1:time.Time,arg2:string):Promise<string>;

export function GetFileDetailsSingle(arg1:string):Promise<main.FileDetailsSingle>;

export function GetFileMeta(arg1:string):Promise<main.FileMeta>;

export function GetFolderAPI(arg1:string):Promise<main.FolderData>;

export function GetHomeDir():Promise<string>;

export function GetListView(arg1:string):Promise<string>;

export function GetLocalFile(arg1:string):Promise<string>;

export function GetSubDirPaths(arg1:string):Promise<Array<string>>;

export function GetThumbnailAsBase64(arg1:string):Promise<string>;

export function GetTree(arg1:string):Promise<{[key: string]: main.Fs}>;

export function GetTreeHTML(arg1:string):Promise<Array<string>>;

export function OpenWithDefaultApp(arg1:string):Promise<number>;

export function RenameFile(arg1:string,arg2:string):Promise<number>;

export function SetFileMeta(arg1:string,arg2:main.FileMeta):Promise<void>;
