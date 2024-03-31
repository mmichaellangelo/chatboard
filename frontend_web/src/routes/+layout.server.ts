import { json } from "@sveltejs/kit";
import type { LayoutServerLoad } from "./$types";
import jwt from 'jsonwebtoken';

export const load: LayoutServerLoad = async ({cookies}) => {
    const session = cookies.get("session");
    if (session) {
        try {
            const sessionData = jwt.verify(session, "secret key", {algorithms: ["HS256"]})
        } catch (error: any) {
            cookies.delete("session", {path: "/"});
        }
        
    }
}