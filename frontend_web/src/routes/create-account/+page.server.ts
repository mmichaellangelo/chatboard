import type { Actions } from "@sveltejs/kit";

export const actions: Actions = {
    default: async ({request, fetch}) => {
        const formData = await request.formData();
        
    }
}