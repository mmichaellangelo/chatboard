import type { Actions } from "@sveltejs/kit";
import { fail, redirect } from "@sveltejs/kit"
import pkg from 'jsonwebtoken'
import type { PageServerLoad } from "../$types";
const {verify} = pkg
export const actions: Actions = {
    default: async ({request, cookies}) => {
        console.log("getting form data")
        const formData = await request.formData();
        try {
            let response = await fetch("http://api:8080/login/", {
                method: "POST",
                mode: "cors",
                body: formData,
            });

            console.log(`Response: ${response}`)

            if (!response.ok) {
                console.error(`API call failed with status: ${response.status}`)
                return fail(400)
            }
            
            try {
                const data = await response.json();
                console.log(data)

                const token = data.session;

                const ver = verify(token, "secret key", {algorithms: ['HS256']})

                console.log(ver)

                cookies.set("session", token, {
                    path: '/',
                    httpOnly: true
                })
            } catch (error) {
                console.log(error)
            }
            
            
        } catch (error) {
            console.log(error)
            return fail(400)
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