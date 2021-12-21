if (typeof window.ethereum !== 'undefined') {
    console.log('MetaMask is installed!');
}

var yghaccounts = [];
const ethereumButton = document.querySelector('.enableEthereumButton');
ethereumButton.addEventListener('click', () => {
    //Will Start the metamask extension
    var i = ethereum.request({ method: 'eth_requestAccounts' });
    getAccount();
    console.log(i)
});