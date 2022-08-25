async function deleteProfileFromServer(id) {
    const response = await fetch(`/profiles/${id}`, {
        method: 'DELETE',
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin',
        headers: {
            'Content-Type': 'application/json'
        }
    });
    return response.status;
}

async function deleteProfile(element, id) {
    let status = await deleteProfileFromServer(id);
    if (status === 200) {
        let td = element.parentNode;
        let tr = td.parentNode;
        let tbody = tr.parentNode;
        tbody.removeChild(tr);
    }
}