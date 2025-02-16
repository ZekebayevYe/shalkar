document.addEventListener('DOMContentLoaded', async function () {
    const token = localStorage.getItem('token');
    const userId = localStorage.getItem('user_id');


    const chatHistory = document.getElementById("chat-history");
    const messageInput = document.getElementById("message");
    const sendButton = document.getElementById("send");

    async function fetchChatHistory() {
        const response = await fetch(`http://localhost:8081/api/chat/history?user_id=${userId}`, {
            headers: { 'Authorization': `Bearer ${token}` }
        });

        if (response.ok) {
            const data = await response.json();
            chatHistory.innerHTML = "";

            data.messages.forEach(msg => {
                const messageDiv = document.createElement("div");
                messageDiv.classList.add("message");
                messageDiv.classList.add(msg.user_id === userId ? "user-message" : "support-message");
                messageDiv.textContent = msg.message;
                chatHistory.appendChild(messageDiv);
            });

            chatHistory.scrollTop = chatHistory.scrollHeight; 
        }
    }

    async function sendMessage() {
        const messageText = messageInput.value.trim();
        if (messageText === "") return;

        const response = await fetch("http://localhost:8081/api/chat/send", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`
            },
            body: JSON.stringify({ user_id: userId, message: messageText })
        });

        if (response.ok) {
            messageInput.value = "";
            fetchChatHistory();
        }
    }

    sendButton.addEventListener("click", sendMessage);
    messageInput.addEventListener("keypress", (event) => {
        if (event.key === "Enter") sendMessage();
    });

    fetchChatHistory();
    setInterval(fetchChatHistory, 5000); 
});
