document.addEventListener("DOMContentLoaded", function () {
    const token = localStorage.getItem("token");

    if (!token) {
        alert("You are not authorized!");
        window.location.href = "../pages/index.html"; // Redirect to login
        return;
    }

    document.getElementById("logout").addEventListener("click", function () {
        localStorage.removeItem("token");
        localStorage.removeItem("user_id");
        alert("Logged out successfully!");
        window.location.href = "../pages/index.html";
    });
});
