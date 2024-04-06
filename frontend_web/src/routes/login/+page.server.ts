import type { Actions } from "@sveltejs/kit";
import { fail, redirect } from "@sveltejs/kit"
import pkg from 'jsonwebtoken'
import type { PageServerLoad } from "../$types";
import { login } from "$lib/auth/login.server";
const {verify} = pkg
export const actions: Actions = {
    default: async ({request, cookies}) => {
        console.log("getting form data")
        const formData = await request.formData();
        if (formData.has("username") && formData.has("password")) {
            let username = formData.get("username") as string;
            let password = formData.get("password") as string;
            try {
                const result = await login(username, password)
                if (!result.success) {
                    return fail(400);
                }
                if (result.token) {
                    cookies.set("session", result.token, {
                        path: '/',
                        httpOnly: true
                    })
                }
                
            } catch (error: any) {
    
            }
        }
        
    }
}



export const load: PageServerLoad = async ({cookies}) => {
    console.log("LOGIN SERVER LOAD")
    const session = cookies.get("session");
    if (session) {
        console.log("Found session cookie!")
        redirect(302, "/")
    }
}