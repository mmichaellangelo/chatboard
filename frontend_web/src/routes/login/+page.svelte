<h1>Login</h1>
<div id="login_container">
    <form method="POST">
        <label for="username">Username</label>
            <input type="text" name="username">
            <br/>
        <label for="password">Password</label>
            <input type="password" name="password">
        <br/>
        <button type="submit">Log In</button>
        <span>Or <a href="/create-account">Create an Account</a></span>
    </form>
</div>

<style>
    #login_container {
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