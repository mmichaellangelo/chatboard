import type { Actions } from "@sveltejs/kit";
import { redirect } from "@sveltejs/kit";
import type { LayoutServerLoad } from "./$types";

export const actions: Actions = {
    default: async ({request}) => {
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

export const load: LayoutServerLoad = async ({cookies}) => {
    const session = cookies.get("session");

    if (!session) {
        redirect(303, "/login")
    }
} 