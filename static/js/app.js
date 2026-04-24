function toggleMenu() {
    document.body.classList.toggle('menu-open');
}

async function loadPage(page) {
    const res = await fetch('/pages/' + page + '.html');
    const html = await res.text();
    document.getElementById('content').innerHTML = html;

    if (page === 'input') {
        loadTable();
    }
}

function toggleForm() {
    const form = document.getElementById('form');
    form.style.display = form.style.display === 'none' ? 'block' : 'none';
}

async function loadTable() {
    const res = await fetch('http://localhost:8080/api/commonTable');
    const data = await res.json();

    const tbody = document.querySelector('#table tbody');
    if (!tbody) return;

    tbody.innerHTML = '';

    data.forEach(row => {
        tbody.innerHTML += `
        <tr>
            <td>${row.position ?? ''}</td>
            <td>${row.client ?? ''}</td>
            <td>${row.operations ?? ''}</td>
            <td>${row.measure ?? ''}</td>
            <td>${row.min ?? ''}</td>
            <td>${row.max ?? ''}</td>
            <td>${row.period_type ?? ''}</td>
            <td>${row.period_count ?? ''}</td>
        </tr>`;
    });
}

async function addRecord() {
    const data = {
        position: document.getElementById('position').value,
        client: document.getElementById('client').value,
        operations: document.getElementById('operations').value,
        measure: document.getElementById('measure').value,
        min: document.getElementById('min').value,
        max: document.getElementById('max').value,
        period_type: document.getElementById('period_type').value,
        period_count: document.getElementById('period_count').value
    };

    await fetch('http://localhost:8080/api/addRecord', {
        method: 'POST',
        body: JSON.stringify(data)
    });

    loadTable();
}

loadPage('input');