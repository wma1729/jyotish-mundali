const contents = [
    'charts',
    'grahas',
    'bhavas'
];

function showContent(category) {
    for (let c of contents) {
        let id = document.getElementById(`analysis-${c}`);
        if (c === category) {
            id.style.display = 'flex';
        } else {
            id.style.display = 'none';
        }
    }
}