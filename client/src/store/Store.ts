import type { File } from "../models/Store";
import { writable } from "svelte/store";

export const userFiles = writable<File[]>([]);
export const permittedFiles = writable<File[]>([]);
