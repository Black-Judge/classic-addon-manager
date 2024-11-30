// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {addon} from '../models';
import {api} from '../models';
import {github} from '../models';

export function GetAddOns():Promise<Array<addon.Addon>>;

export function GetAddonManifest(arg1:boolean):Promise<Array<addon.AddonManifest>>;

export function GetAddonManifests():Promise<Array<addon.AddonManifest>>;

export function GetConfigBool(arg1:string):Promise<boolean>;

export function GetConfigString(arg1:string):Promise<string>;

export function GetLatestAddonRelease(arg1:string):Promise<api.Release>;

export function GetLatestApplicationRelease():Promise<api.ApplicationRelease>;

export function GetReleases(arg1:string):Promise<Array<github.GithubRelease>>;

export function InstallAddon(arg1:addon.AddonManifest):Promise<boolean>;

export function IsAddonInstalled(arg1:string):Promise<boolean>;

export function SelectDirectory():Promise<string>;

export function SetConfigBool(arg1:string,arg2:boolean):Promise<void>;

export function SetConfigString(arg1:string,arg2:string):Promise<void>;

export function UninstallAddon(arg1:string):Promise<boolean>;
