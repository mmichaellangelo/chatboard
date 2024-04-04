import { writable, type Writable } from "svelte/store";

export interface SessionData {
    username: string;
}

export const sessionStore = writable<SessionData>()


