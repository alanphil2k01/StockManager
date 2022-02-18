const userNameField = document.getElementById("username-login")
const passwordField = document.getElementById("password-login")

async function login() {
    const res = await fetch("/login", {
        "method": "POST",
        "headers": {
            "Content-Type": "application/json"
        },
        "body": JSON.stringify({
            username: userNameField.value,
            password: passwordField.value,
        })
    })
    if (res.status === 401) {
        alert("Couldn't Login - Check your crendentials")
    } else if (res.status === 500) {
        alert("Error logging in")
    } else if (res.status === 200){
        let data = await res.json()
        window.localStorage.setItem('ssmc-jwt', data["data"]);
        window.location.href = "/";
    }
}

function logout() {
    window.localStorage.removeItem("ssmc-jwt")
    window.location.href = "/login.html";
}

const userNameFieldRegister = document.getElementById("username-register")
const passwordFieldRegister = document.getElementById("password-register")
const confirmPasswordField = document.getElementById("confirm-password-register")
const emailField = document.getElementById("email-register")
const nameField = document.getElementById("name-register")
const roleForm = document.getElementById("role-register")
const secretField = document.getElementById("secret-register")

async function register() {
    if(passwordFieldRegister.value != confirmPasswordField.value) {
        alert("passwords do not match")
        return
    }
    const res = await fetch("/register", {
        "method": "POST",
        "headers": {
            "Content-Type": "application/json"
        },
        "body": JSON.stringify({
            username: userNameFieldRegister.value,
            password: passwordFieldRegister.value,
            email: emailField.value,
            name: nameField.value,
            role: roleForm.value,
            secret: secretField.value,
        })
    })
    let data = await res.json()
    console.log(data["data"])
    if (res.status === 401) {
        alert("Invalid Secret Key")
        return

    } else if (res.status === 500) {
        alert("Couldn't Register")
        return
    } else if (res.status === 400) {
        alert("Invalid input")
        return
    } else if (res.status === 200) {
        window.localStorage.setItem('ssmc-jwt', data["data"]);
        window.location.href = "/";
    }
}
