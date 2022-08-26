import { writable } from "svelte/store";
import type { Language } from "../models/Settings";
import { English } from "../constants/Language";

export const LanguageState = writable<Language>(English);
