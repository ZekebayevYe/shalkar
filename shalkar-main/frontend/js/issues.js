document.addEventListener('DOMContentLoaded', async () => {
    const issueForm = document.getElementById('issue-form');
    const titleInput = document.getElementById('title');
    const descriptionInput = document.getElementById('description');
    
    // Проверяем, есть ли уже блок уведомлений, если нет - создаем
    let notification = document.getElementById('notification');
    if (!notification) {
        notification = document.createElement('div');
        notification.id = 'notification';
        notification.className = 'notification';
        document.body.appendChild(notification);
    }

    issueForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        const title = titleInput.value.trim();
        const description = descriptionInput.value.trim();

        if (!title || !description) {
            showNotification('⚠️ Заполните все поля!', 'error');
            return;
        }

        try {
            console.log('Отправляем запрос:', { title, description, status: 'Открыто' });

            const response = await fetch('http://localhost:8081/api/issues', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ title, description, status: 'Открыто' })
            });

            console.log('Ответ от сервера:', response);
            
            if (response.ok) {
                showNotification('✅ Проблема успешно отправлена!', 'success');
                titleInput.value = '';
                descriptionInput.value = '';
            } else {
                const errorMessage = await response.text();
                showNotification(`❌ Ошибка: ${errorMessage}`, 'error');
                console.error('Ошибка сервера:', errorMessage);
            }
        } catch (error) {
            console.error('Ошибка запроса:', error);
            showNotification('❌ Ошибка сети. Попробуйте снова!', 'error');
        }
    });

    function showNotification(message, type) {
        notification.textContent = message;
        notification.className = `notification ${type}`;
        notification.style.display = 'block';

        setTimeout(() => {
            notification.style.display = 'none';
        }, 3000);
    }
});
