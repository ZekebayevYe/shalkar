document.addEventListener('DOMContentLoaded', async function () {
    const token = localStorage.getItem('token');
    const userId = localStorage.getItem('user_id');

   

    async function fetchHistory() {
        const response = await fetch(`http://localhost:8081/api/expenses/history?user_id=${userId}`, {
            headers: { 'Authorization': `Bearer ${token}` }
        });

        if (response.ok) {
            const data = await response.json();
            const historyList = document.getElementById('history');
            historyList.innerHTML = '';

            data.history.forEach(expense => {
                const li = document.createElement('li');
                li.textContent = `Date: ${expense.created_at}, Amount: ${expense.total} KZT`;
                historyList.appendChild(li);
            });
        }
    }

    document.getElementById('calculate').addEventListener('click', async function () {
        const electricity = parseFloat(document.getElementById('electricity').value) || 0;
        const water = parseFloat(document.getElementById('water').value) || 0;
        const gas = parseFloat(document.getElementById('gas').value) || 0;

        const response = await fetch('http://localhost:8081/api/expenses/calculate', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`
            },
            body: JSON.stringify({ user_id: userId, electricity, water, gas })
        });

        if (response.ok) {
            const data = await response.json();
            document.getElementById('total').textContent = `${data.total} KZT`;
            document.getElementById('elec-cost').textContent = `${data.electricity} KZT`;
            document.getElementById('water-cost').textContent = `${data.water} KZT`;
            document.getElementById('gas-cost').textContent = `${data.gas} KZT`;
            document.getElementById('result').classList.remove('hidden');
        }
    });

    fetchHistory();
});
