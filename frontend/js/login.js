async function login() {
    const userNameFieldLogin = document.getElementById("username-login")
    const passwordFieldLogin = document.getElementById("password-login")
    const res = await fetch("/login", {
        "method": "POST",
        "headers": {
            "Content-Type": "application/json"
        },
        "body": JSON.stringify({
            username: userNameFieldLogin.value,
            password: passwordFieldLogin.value,
        })
    })
    let data = await res.json()
    if (res.status === 401) {
        alert("Couldn't Login - Check your crendentials")
    } else if (res.status === 500) {
        alert("Error logging in")
    } else if (res.status === 200){
        window.localStorage.setItem('ssmc-jwt', data["data"]);
        window.location.href = "/";
    }
}

function logout() {
    window.localStorage.removeItem("ssmc-jwt")
    window.location.href = "/login.html";
}

async function register() {
    const userNameFieldRegister = document.getElementById("username-register")
    const passwordFieldRegister = document.getElementById("password-register")
    const confirmPasswordField = document.getElementById("confirm-password-register")
    const emailField = document.getElementById("email-register")
    const nameField = document.getElementById("name-register")
    const roleForm = document.getElementById("role-register")
    const secretField = document.getElementById("secret-register")
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
            role: Number(roleForm.value),
            secret: secretField.value,
        })
    })
    let data = await res.json()
    if (res.status === 401) {
        alert("Invalid secret key")
    } else if (res.status === 500) {
        alert("username or email alread exists")
    } else if (res.status === 400) {
        alert("Invalid input")
    } else if (res.status === 200) {
        window.localStorage.setItem('ssmc-jwt', data["data"]);
        window.location.href = "/";
    }
}


function openRestrationWindow() {
    var openelement=document.getElementsByClassName("register-window");
    openelement[0].style.display="inline-block";

}

function closeRestrationWindow() {
    var closeelement=document.getElementsByClassName("register-window");
    closeelement[0].style.display="none";
    closeelement[0].style.animation="fadeEffect .8s";
}
