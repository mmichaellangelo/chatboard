import type { LayoutServerLoad } from "./$types";
import jwt from 'jsonwebtoken';
import { setContext } from "svelte";
import { sessionStore, type SessionData } from "$lib/state/userStore";


export const load: LayoutServerLoad = async ({cookies}) => {
    console.log("LAYOUT SERVER LOAD")
    const session = cookies.get("session");
    if (session) {
        try {
            const sessionData = jwt.verify(session, "secret key", {algorithms: ["HS256"]}) as SessionData
            console.log(sessionData.username)
            sessionStore.set({username: sessionData.username})
        } catch (error) {
            console.log(`LAYOUT SERVER LOAD ERROR: ${error}`)
            cookies.delete("session", {path: "/"});
            sessionStore.set({username: ""})
        }
        
    }
}