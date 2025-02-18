document.addEventListener('DOMContentLoaded', async () => {
    const token = localStorage.getItem('token');

    async function fetchEvents() { // Исправлено имя функции
        try {
            const response = await fetch('http://localhost:8081/events', {
                headers: { 'Authorization': `Bearer ${token}` }
            });

            if (!response.ok) {
                throw new Error(`Ошибка HTTP: ${response.status}`);
            }

            const eventsList = await response.json();
            console.log("Events:", eventsList);
            const eventsContainer = document.getElementById('events-list');
            eventsContainer.innerHTML = '';

            eventsList.forEach(event => { // Исправлена переменная итерации
                const eventItem = document.createElement('div'); // Исправлено имя переменной
                eventItem.classList.add('swiper-slide', 'events-item');
                eventItem.innerHTML = `
                    <h3>${event.title}</h3>
                    <p>${event.description}</p>
                    <img src="${event.image_url}" alt="${event.title}" />
                `;
                eventsContainer.appendChild(eventItem); // Исправлено имя контейнера
            });

            // Инициализация Swiper после загрузки данных
            new Swiper('.news-carousel', {
                loop: true,
                navigation: {
                    nextEl: '.swiper-button-next',
                    prevEl: '.swiper-button-prev'
                },
                pagination: {
                    el: '.swiper-pagination',
                    clickable: true
                }
            });

        } catch (error) {
            console.error('Ошибка при загрузке событий:', error);
        }
    }

    fetchEvents(); // Исправлено имя вызова функции
});
