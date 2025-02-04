function login() {
    fetch("http://localhost:8080/auth/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
            username: document.getElementById("login-username").value,
            password: document.getElementById("login-password").value
        })
    })
    .then(res => res.json())
    .then(data => {
        if (data.token) {
            document.getElementById("login-result").innerText = "Успешный вход!";
            localStorage.setItem("token", data.token);
        } else {
            document.getElementById("login-result").innerText = data.error || "Ошибка!";
        }
    });
}

function register() {
    fetch("http://localhost:8081/auth/register", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
            username: document.getElementById("register-username").value,
            password: document.getElementById("register-password").value,
            role: document.getElementById("register-role").value
        })
    })
    .then(res => res.json())
    .then(data => {
        document.getElementById("register-result").innerText = data.message || data.error;
    });
}
