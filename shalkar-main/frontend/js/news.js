document.addEventListener('DOMContentLoaded', async () => {
    const token = localStorage.getItem('token');

    async function fetchNews() {
        try {
            const response = await fetch('http://localhost:8081/news', {
                headers: { 'Authorization': `Bearer ${token}` }
            });

            if (!response.ok) {
                throw new Error(`Ошибка HTTP: ${response.status}`);
            }

            const newsList = await response.json();
            console.log("Новости:", newsList);
            const newsContainer = document.getElementById('news-list');
            newsContainer.innerHTML = '';

            newsList.forEach(news => {
                const newsItem = document.createElement('div');
                newsItem.classList.add('swiper-slide', 'news-item');
                newsItem.innerHTML = `
                    <h3>${news.title}</h3>
                    <p>${news.description}</p>
                    <img src="${news.image_url}" alt="${news.title}" />
                `;
                newsContainer.appendChild(newsItem);
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
            console.error('Ошибка при загрузке новостей:', error);
        }
    }

    fetchNews();
});
