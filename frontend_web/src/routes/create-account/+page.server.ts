import { login } from "$lib/auth/login.server";
import { fail, type Actions, redirect } from "@sveltejs/kit";
import jwt from "jsonwebtoken";

export function load({cookies}) {
    let session = cookies.get("session")
    if (session) {
        try {
            jwt.verify(session, "secret key")
            
        } catch (error: any) {
            console.log(error.message);
            cookies.delete("session", { path: '/' })
            return;
        }
        redirect(302, '/')
    }
}

export const actions: Actions = {
    default: async ({request, fetch, cookies}) => {
        console.log("SENDING POST REQUEST")
        let formData;
        try {
            formData = await request.formData();
        } catch (error: any) {
            console.log(`Error parsing form data: ${error.message}`)
        }
        if (formData) {
            if (formData.has("username") && formData.has("email") && formData.has("password")) {
                let email = formData.get("email") as string;
                let password = formData.get("password") as string;
                let username = formData.get("username") as string;
                if (username && password && email) {
                    if (username != "" && password != "" && email != "") {
                        try {
                            let response = await fetch("http://api:8080/accounts/", {
                                method: "POST",
                                mode: "cors",
                                body: formData,
                            });
                            if (!response.ok) {
                                console.log(`Failed to create account: ${response}`)
                                return fail(400);
                            }
                            const result = await login(username, password);
                            if (!result.success) {
                                return fail(400);
                            }
                            if (result.token) {
                                cookies.set("session", result.token, {
                                    path: '/',
                                    httpOnly: true,
                                })
                            }
                        } catch (error: any) {
                            console.log(`Failed to create account: ${error.message}`)
                        }
                    }
                }
            }
        }
    }
}
