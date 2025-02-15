document.addEventListener('DOMContentLoaded', () => {
    const token = localStorage.getItem('token');
    const userId = localStorage.getItem('user_id');

    if (!token) {
        alert('❌ Ошибка: Токен отсутствует. Войдите в систему.');
        window.location.href = 'index.html'; // Перенаправление на страницу входа
        return;
    }

    const issueForm = document.getElementById('issue-form');
    const issuesTableBody = document.querySelector('#issues-table tbody');

    // 🟢 Функция загрузки обращений пользователя
    async function loadIssues() {
        try {
            const response = await fetch('http://localhost:8081/api/issues', {
                method: 'GET',
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            });

            const issues = await response.json();
            if (response.ok) {
                issuesTableBody.innerHTML = ''; // Очищаем таблицу перед обновлением
                issues
                    .filter(issue => issue.user_id === Number(userId)) // Фильтруем по user_id
                    .forEach(issue => {
                        const row = document.createElement('tr');
                        row.innerHTML = `
                            <td>${issue.id}</td>
                            <td>${issue.title}</td>
                            <td>${issue.description}</td>
                            <td>${issue.status}</td>
                        `;
                        issuesTableBody.appendChild(row);
                    });
            } else {
                alert('❌ Ошибка загрузки обращений: ' + (issues.error || 'Неизвестная ошибка'));
            }
        } catch (error) {
            console.error('Ошибка загрузки обращений:', error);
            alert('❌ Ошибка соединения с сервером');
        }
    }

    // 🔹 Загружаем обращения при загрузке страницы
    loadIssues();

    // 🟢 Обработчик отправки формы
    issueForm.addEventListener('submit', async (e) => {
        e.preventDefault();

        const title = document.getElementById('title').value.trim();
        const description = document.getElementById('description').value.trim();

        if (!title || !description) {
            alert('❌ Заполните все поля');
            return;
        }

        try {
            const response = await fetch('http://localhost:8081/api/issues', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`
                },
                body: JSON.stringify({ title, description, status: 'Открыто', user_id: Number(userId) })
            });

            const data = await response.json();
            if (response.ok) {
                alert('✅ Проблема успешно отправлена!');
                issueForm.reset();
                loadIssues(); // 🔹 Обновляем таблицу
            } else {
                alert('❌ Ошибка: ' + (data.error || 'Неизвестная ошибка'));
            }
        } catch (error) {
            console.error('Ошибка запроса:', error);
            alert('❌ Ошибка соединения с сервером');
        }
    });
});
