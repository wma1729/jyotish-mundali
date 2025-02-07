const contents = [
    'lagna-chart',
    'grahas-relations',
    'grahas-aspects',
    'bhavas'
];

function showContent(category) {
    console.log(category)
    for (let c of contents) {
        let id = document.getElementById(`analysis-${c}`);
        if (c === category) {
            id.style.display = 'flex';
        } else {
            id.style.display = 'none';
        }
    }
}