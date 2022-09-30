const grahas = [
    'sun',
    'moon',
    'mars',
    'mercury',
    'jupiter',
    'venus',
    'saturn',
    'rahu',
    'ketu'
];

function showGraha(name) {
    for (const graha of grahas) {
        let id = document.getElementById(`${graha}-attr`);
        if (graha === name) {
            id.style.display = 'block';
        } else {
            id.style.display = 'none';
        }
    }
}