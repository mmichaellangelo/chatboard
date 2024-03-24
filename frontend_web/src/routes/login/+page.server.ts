import type { Actions } from "@sveltejs/kit";

export const actions: Actions = {
    default: async ({request}) => {
        const formData = await request.formData();
        console.log(formData)
        try {
            const response = await fetch("http://api:8080/login/", {
                method: "POST",
                mode: "cors",
                body: formData,
            });

            if (!response.ok) {
                console.error(`API call failed with status: ${response.status}`)
            }
            try {
                const data = await response.json()
                const token = data.token
                console.log(token)
            } catch (error) {
                console.error(`Error parsing JSON: ${error}`)
            }
            
        } catch (error) {
            console.log(error)
        }

        
        
    }
}