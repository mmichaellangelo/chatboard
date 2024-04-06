

export interface LoginReturnData {
    success: boolean;
    errorMessage?: string;
    token?: string;

}

export async function login(username: string, password: string): Promise<LoginReturnData> {
    // convert username, password to FormData object
    const formData = buildFormData(username, password);
    let response = await fetch("http://api:8080/login/", {
                method: "POST",
                mode: "cors",
                body: formData,
            });

            console.log(`Response: ${response}`)

            if (!response.ok) {
                console.error(`API call failed with status: ${response.status}`)
                return {
                    success: false,
                    errorMessage: `Error, response: ${response}`,
                }
            }
            
            try {
                const data = await response.json();
                console.log(data)

                const token = data.session;

                return {
                    success: true,
                    token: token,
                }
                
            } catch (error: any) {
                return {
                    success: false,
                    errorMessage: error.message
                }
            }
        
}

function buildFormData(username: string, password: string): FormData {
    const formData = new FormData();
    formData.append("username", username);
    formData.append("password", password);
    return formData;
}