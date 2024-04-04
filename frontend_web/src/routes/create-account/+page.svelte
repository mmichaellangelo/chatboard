<h1>Create Account</h1>
<div id="form_container">
    <form method="post">
        <label for="email">Email</label>
        <input type="email"><br/>
        <label for="username">Username</label>
        <input type="text" name="username"> <br/>
        <label for="password">Password</label>
        <input type="password"><br/>
        <button type="submit">Create Account</button>
    </form>
</div>

<style>
    #form_container {
        display: flex;
        max-width: 20rem;
        padding: 1rem;
    }

    form {
        text-align: right;
    }

    
    label {
        margin-right: auto;
    }

    input {
        margin-bottom: 1rem;
        
    }
</style>