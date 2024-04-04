import type { Actions } from "@sveltejs/kit";
import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

export const actions: Actions = {
    default: async ({request, fetch}) => {
        let formData = await request.formData()

        try {
            const response = await fetch("http://api:8080/messages/", {
                method: "POST",
                mode: "cors",
                body: formData,
            })
        } catch (error) {
            console.log(error)
        }
    }
}

export const load: PageServerLoad = async ({cookies}) => {
    console.log("ROOT SERVER LOAD")
    const session = cookies.get("session");

    if (!session) {
        redirect(303, "/login")
    }
} 