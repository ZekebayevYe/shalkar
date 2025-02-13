async function apiRequest(url, method = 'GET', body = null) {
    const token = localStorage.getItem('token');

    const response = await fetch(url, {
        method,
        headers: { 
            'Content-Type': 'application/json', 
            'Authorization': `Bearer ${token}` 
        },
        body: body ? JSON.stringify(body) : null
    });

    return response.json();
}
