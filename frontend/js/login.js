const userNameField = document.getElementById("username-login")
const passwordField = document.getElementById("password-login")

async function login() {
    console.log(userNameField.value)
    console.log(passwordField.value)
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
        return

    }
    let data = await res.json()
    window.localStorage.setItem('ssmc-jwt', data["data"]);
    window.location.href = "/";
}
