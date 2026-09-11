
async function copyTextToClipboard(buttonId, source) {
    const content = document.getElementById(`${source}`);
    const button = document.getElementById(buttonId);
   
    try {
        const passContent = content.firstChild; 
        await navigator.clipboard.writeText(passContent.textContent);

        button.style.backgroundColor = '#00a63e';
        button.innerText = 'Copied!!';

        setTimeout(() => {
            button.style.backgroundColor = '#333';
            button.innerText = 'Copy';
        }, 2000);

    } catch(error) {
        const card = document.getElementById('error-card');
        card.classList.toggle('hidden')

        setTimeout(() => {
            card.classList.toggle('hidden'); 
        }, 3000);

        console.log(error); 
    }
}
