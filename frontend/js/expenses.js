document.getElementById("expense-form").addEventListener("submit", async function (event) {
    event.preventDefault(); 

    const coldWater = parseFloat(document.getElementById("cold_water").value) || 0;
    const hotWater = parseFloat(document.getElementById("hot_water").value) || 0;
    const heating = parseFloat(document.getElementById("heating").value) || 0;
    const gas = parseFloat(document.getElementById("gas").value) || 0;
    const electricity = parseFloat(document.getElementById("electricity").value) || 0;

    const requestBody = {
        cold_water: coldWater,
        hot_water: hotWater,
        heating: heating,
        gas: gas,
        electricity: electricity
    };

    console.log("📤 Отправляем JSON:", JSON.stringify(requestBody));

    try {
        const response = await fetch("http://localhost:8081/api/expenses", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${localStorage.getItem("token")}`
            },
            body: JSON.stringify(requestBody)
        });

        const data = await response.json();
        console.log("📥 Ответ сервера:", data);

        if (response.ok) {
            document.querySelector(".expense-result").classList.remove("hidden");

            document.getElementById("detail-total-cost").textContent =
                `Total: ${data.total.toFixed(2)} KZT`;

            document.getElementById("detail-cold-water").textContent =
                `Cold Water: ${data.breakdown.cold_water} m³ → ${(data.breakdown.cold_water * 100).toFixed(2)} KZT`;

            document.getElementById("detail-hot-water").textContent =
                `Hot Water: ${data.breakdown.hot_water} m³ → ${(data.breakdown.hot_water * 250).toFixed(2)} KZT`;

            document.getElementById("detail-heating").textContent =
                `Heating: ${data.breakdown.heating} units → ${(data.breakdown.heating * 200).toFixed(2)} KZT`;

            document.getElementById("detail-gas").textContent =
                `Gas: ${data.breakdown.gas} m³ → ${(data.breakdown.gas * 70).toFixed(2)} KZT`;

            document.getElementById("detail-electricity").textContent =
                `Electricity: ${data.breakdown.electricity} kWh → ${(data.breakdown.electricity * 20).toFixed(2)} KZT`;
        } else {
            console.error("❌ Ошибка сервера:", data);
            alert("Failed to calculate expenses. See console for details.");
        }
    } catch (error) {
        console.error("❌ Ошибка запроса:", error);
    }
});

async function viewExpenseDetails(expenseId) {
    console.log(`🔍 Fetching details for expense ID: ${expenseId}`);

    const response = await fetch(`http://localhost:8081/api/expenses/${expenseId}`, {
        headers: { "Authorization": `Bearer ${localStorage.getItem("token")}` }
    });

    const data = await response.json();
    console.log("📥 API Response:", data);

    if (!data.details) {
        console.error("🚨 API Response Missing 'details':", data);
        alert("Failed to load expense details. Check backend.");
        return;
    }

    let detailsBlock = document.getElementById(`details-${expenseId}`);
    if (!detailsBlock) {
        console.warn(`🚨 No matching details block found for expense ID: ${expenseId}, creating one.`);

        detailsBlock = document.createElement("div");
        detailsBlock.id = `details-${expenseId}`;
        detailsBlock.classList.add("expense-details");

        detailsBlock.innerHTML = `
            <h3>Expense Breakdown</h3>
            <p><strong>Total Cost:</strong> <span id="detail-total-cost-${expenseId}"></span></p>
            <p><strong>Cold Water:</strong> <span id="detail-cold-water-${expenseId}"></span></p>
            <p><strong>Hot Water:</strong> <span id="detail-hot-water-${expenseId}"></span></p>
            <p><strong>Heating:</strong> <span id="detail-heating-${expenseId}"></span></p>
            <p><strong>Gas:</strong> <span id="detail-gas-${expenseId}"></span></p>
            <p><strong>Electricity:</strong> <span id="detail-electricity-${expenseId}"></span></p>
            <p><strong>Date Recorded:</strong> <span id="detail-date-${expenseId}"></span></p>
        `;
        document.getElementById("expense-history").appendChild(detailsBlock);
    }

    document.getElementById(`detail-total-cost-${expenseId}`).textContent =
        `Total Cost: ${data.details.total_cost.toFixed(2)} KZT`;

    document.getElementById(`detail-cold-water-${expenseId}`).textContent =
        `${data.details.cold_water} m³ → ${(data.details.cold_water * 100).toFixed(2)} KZT`;

    document.getElementById(`detail-hot-water-${expenseId}`).textContent =
        `${data.details.hot_water} m³ → ${(data.details.hot_water * 250).toFixed(2)} KZT`;

    document.getElementById(`detail-heating-${expenseId}`).textContent =
        `${data.details.heating} units → ${(data.details.heating * 200).toFixed(2)} KZT`;

    document.getElementById(`detail-gas-${expenseId}`).textContent =
        `${data.details.gas} m³ → ${(data.details.gas * 70).toFixed(2)} KZT`;

    document.getElementById(`detail-electricity-${expenseId}`).textContent =
        `${data.details.electricity} kWh → ${(data.details.electricity * 20).toFixed(2)} KZT`;

    document.getElementById(`detail-date-${expenseId}`).textContent =
        `Recorded On: ${new Date(data.details.created_at).toLocaleDateString()} at ${new Date(data.details.created_at).toLocaleTimeString()}`;

    detailsBlock.classList.toggle("hidden"); // Переключение видимости
}


async function fetchHistory() {
    const token = localStorage.getItem("token");
    
    try {
        const response = await fetch("http://localhost:8081/api/expenses/history", {
            headers: { "Authorization": `Bearer ${token}` }
        });

        if (!response.ok) {
            console.error("❌ Ошибка сервера:", response.status);
            alert("Ошибка загрузки истории расходов!");
            return;
        }

        const data = await response.json();
        console.log("📥 История расходов:", data);

        const historyList = document.getElementById("expense-history");
        historyList.innerHTML = "";

        if (!data.history || data.history.length === 0) {
            historyList.innerHTML = "<p>История пуста.</p>";
            return;
        }

        data.history.forEach(expense => {
            const listItem = document.createElement("li");
            listItem.innerHTML = `
                <h3>Expense ID: ${expense.id}</h3>
                <p><strong>Total Cost:</strong> ${expense.total_cost} KZT</p>
                <p><strong>Cold Water:</strong> ${expense.cold_water} m³ → ${(expense.cold_water * 100).toFixed(2)} KZT</p>
                <p><strong>Hot Water:</strong> ${expense.hot_water} m³ → ${(expense.hot_water * 250).toFixed(2)} KZT</p>
                <p><strong>Heating:</strong> ${expense.heating} units → ${(expense.heating * 200).toFixed(2)} KZT</p>
                <p><strong>Gas:</strong> ${expense.gas} m³ → ${(expense.gas * 70).toFixed(2)} KZT</p>
                <p><strong>Electricity:</strong> ${expense.electricity} kWh → ${(expense.electricity * 20).toFixed(2)} KZT</p>
                <p><strong>Date Recorded:</strong> ${new Date(expense.created_at).toLocaleDateString()} at ${new Date(expense.created_at).toLocaleTimeString()}</p>
                <hr>
            `;
            historyList.appendChild(listItem);
        });

    } catch (error) {
        console.error("❌ Ошибка запроса:", error);
        alert("Ошибка загрузки истории расходов!");
    }
}

// Загружаем историю при загрузке страницы
document.addEventListener("DOMContentLoaded", fetchHistory);
