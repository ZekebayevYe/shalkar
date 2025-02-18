document.addEventListener("DOMContentLoaded", function () {
    fetchUserFeedback();
});

document.getElementById("feedback-form").addEventListener("submit", async function (event) {
    event.preventDefault();

    const categoryElem = document.getElementById("feedback-category");
    const ratingElem = document.getElementById("feedback-rating");
    const commentElem = document.getElementById("feedback-comment");
    

    if (!categoryElem || !ratingElem || !commentElem) {
        alert("Form elements not found. Please check the form.");
        return;
    }

    const category = categoryElem.value;
    const rating = parseInt(ratingElem.value, 10);
    const comment = commentElem.value.trim();
    const token = localStorage.getItem("token");

    if (!token) {
        alert("User not logged in. Please login first.");
        return;
    }

    const requestBody = {
        category: category,
        rating: rating,
        comment: comment
    };

    try {
        const response = await fetch("http://localhost:8081/api/feedback", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`
            },
            body: JSON.stringify(requestBody)
        });

        const text = await response.text();
        if (!text) {
            throw new Error("Empty response from server");
        }

        const data = JSON.parse(text);
        console.log("Server Response:", data);

        if (!response.ok) {
            alert(`Error: ${data.error || "Something went wrong"}`);
            return;
        }

        alert("Feedback submitted successfully!");
        document.getElementById("feedback-form").reset();
    } catch (error) {
        console.error("Request Error:", error);
        alert("Failed to submit feedback. Check console for details.");
    }
});

async function fetchUserFeedback() {
    const token = localStorage.getItem("token");
    const response = await fetch("http://localhost:8081/api/feedback", {
        headers: { "Authorization": `Bearer ${token}` }
    });

    const data = await response.json();
    const historyList = document.getElementById("feedback-history");
    historyList.innerHTML = "";

    if (data.feedback && Array.isArray(data.feedback)) {
        data.feedback.forEach(item => {
            const listItem = document.createElement("li");
            listItem.innerHTML = `
                <p><strong>Category:</strong> ${item.category} | <strong>Rating:</strong> ${item.rating}/5</p>
                <p>${item.comment}</p>
                <p><em>Date: ${new Date(item.created_at).toLocaleDateString()}</em></p>
            `;
            historyList.appendChild(listItem);
        });
    } else {
        console.error("Unexpected response structure:", data);
    }
}
