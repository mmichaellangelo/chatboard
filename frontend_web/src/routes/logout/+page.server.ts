import { sessionStore } from "$lib/state/userStore.js";
import { redirect } from "@sveltejs/kit";

export function load({cookies}) {
    try {
        cookies.delete("session", { path: '/' });
        sessionStore.set({username: ""});
    } catch (error: any) {
        console.log(error)
    }
    redirect(302, "/login")
}