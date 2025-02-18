document.addEventListener('DOMContentLoaded', async function () {
    const token = localStorage.getItem('token');
    const role = localStorage.getItem('role');

    if (!token) {
        window.location.href = 'index.html';
        return;
    }

    const adminPanel = document.getElementById('admin-panel');

    if (role !== 'admin' && adminPanel) {
        adminPanel.style.display = 'none'; 
    }

    function getFileIcon(filename) {
        const ext = filename.split('.').pop().toLowerCase();
        const icons = {
            pdf: 'pdf.png',
            doc: 'word.png',
            docx: 'word.png',
            xls: 'excel.png',
            xlsx: 'excel.png',
            jpg: 'image.png',
            jpeg: 'image.png',
            png: 'image.png',
            txt: 'text.png',
            zip: 'zip.png',
            rar: 'zip.png',
            pptx: 'pptx.png',
        };
        return `../icons/${icons[ext] || 'file.png'}`;
    }

    async function fetchFiles() {
        try {
            const response = await fetch('http://localhost:8081/api/files', {
                headers: { 'Authorization': `Bearer ${token}` }
            });

            if (!response.ok) {
                throw new Error('Ошибка при загрузке списка файлов');
            }

            const files = await response.json();
            const fileList = document.getElementById('file-list');
            fileList.innerHTML = '';

            files.forEach(file => {
                const fileCard = document.createElement('div');
                fileCard.classList.add('file-card');

                fileCard.innerHTML = `
                    <img src="${getFileIcon(file.name)}" alt="file-icon" class="file-icon" data-id="${file.id}" data-name="${file.name}">
                    <p class="file-name" data-id="${file.id}" data-name="${file.name}">${file.name}</p>
                    ${role === 'admin' ? `<button class="delete-btn" data-id="${file.id}">Удалить</button>` : ''}
                `;

                fileList.appendChild(fileCard);
            });

            document.querySelectorAll('.file-icon, .file-name').forEach(element => {
                element.addEventListener('click', function () {
                    const fileId = this.dataset.id;
                    const fileName = this.dataset.name;
                    if (fileId && fileName) {
                        downloadFile(fileId, fileName);
                    }
                });
            });

            if (role === 'admin') {
                document.querySelectorAll('.delete-btn').forEach(btn => {
                    btn.addEventListener('click', async function () {
                        const fileId = this.dataset.id;
                        await deleteFile(fileId);
                        fetchFiles();
                    });
                });
            }
        } catch (error) {
            console.error('Ошибка:', error);
            alert('Не удалось загрузить файлы');
        }
    }

    async function downloadFile(fileId, fileName) {
        try {
            console.log(`📥 Скачивание файла: ID=${fileId}, Name=${fileName}`);

            const response = await fetch(`http://localhost:8081/api/download/${fileId}`, {
                headers: { 'Authorization': `Bearer ${token}` }
            });

            if (!response.ok) {
                throw new Error(`Ошибка при скачивании файла (код: ${response.status})`);
            }

            const blob = await response.blob();
            const url = window.URL.createObjectURL(blob);

            const a = document.createElement('a');
            a.href = url;
            a.download = fileName;
            document.body.appendChild(a);
            a.click();
            document.body.removeChild(a);
            window.URL.revokeObjectURL(url);

        } catch (error) {
            console.error('Ошибка при скачивании:', error);
            alert('Не удалось скачать файл');
        }
    }

    async function deleteFile(fileId) {
        try {
            const response = await fetch(`http://localhost:8081/api/admin/files/${fileId}`, {
                method: 'DELETE',
                headers: { 'Authorization': `Bearer ${token}` }
            });

            if (!response.ok) {
                throw new Error('Ошибка при удалении файла');
            }
        } catch (error) {
            console.error('Ошибка:', error);
            alert('Не удалось удалить файл');
        }
    }

    document.getElementById('upload-btn')?.addEventListener('click', async function () {
        const fileInput = document.getElementById('file-input');
        const file = fileInput.files[0];

        if (!file) return alert('Выберите файл');

        const formData = new FormData();
        formData.append('file', file);

        try {
            const response = await fetch('http://localhost:8081/api/admin/upload', {
                method: 'POST',
                headers: { 'Authorization': `Bearer ${token}` },
                body: formData
            });

            if (!response.ok) {
                throw new Error('Ошибка при загрузке файла');
            }

            fetchFiles();
        } catch (error) {
            console.error('Ошибка:', error);
            alert('Не удалось загрузить файл');
        }
    });

    fetchFiles();
});
