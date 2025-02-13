document.addEventListener('DOMContentLoaded', async () => {
    const newsCarousel = document.getElementById('news-carousel');
    const eventCarousel = document.getElementById('event-carousel');

    async function fetchData(url) {
        try {
            const response = await fetch(url);
            return await response.json();
        } catch (error) {
            console.error('Ошибка загрузки данных:', error);
        }
    }

    function createCard(item, type) {
        const div = document.createElement('div');
        div.classList.add(type === 'news' ? 'news-item' : 'event-item');
        div.innerHTML = `
            <h3>${item.title}</h3>
            <p>${item.content}</p>
            <button class="like-btn" data-id="${item.id}" data-type="${type}">❤️ ${item.likes}</button>
        `;
        return div;
    }

    async function loadNews() {
        const news = await fetchData('http://localhost:8081/news');
        news.forEach(item => newsCarousel.appendChild(createCard(item, 'news')));
    }

    async function loadEvents() {
        const events = await fetchData('http://localhost:8081/events');
        events.forEach(item => eventCarousel.appendChild(createCard(item, 'event')));
    }

    document.addEventListener('click', async (e) => {
        if (e.target.classList.contains('like-btn')) {
            const id = e.target.dataset.id;
            const type = e.target.dataset.type;
            const url = `http://localhost:8081/${type}/${id}/react`;
            
            const response = await fetch(url, { method: 'POST' });
            if (response.ok) {
                let likes = parseInt(e.target.textContent.replace('❤️', '').trim(), 10);
                e.target.textContent = `❤️ ${likes + 1}`;
            }
        }
    });

    loadNews();
    loadEvents();
});
