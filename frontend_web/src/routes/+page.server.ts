import type { Actions } from "@sveltejs/kit";

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