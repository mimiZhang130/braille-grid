const text = document.body.innerText;

if (text) {
    chrome.runtime.sendMessage({action:'dom', text: text }, function(response) {
        console.log('innerText sent to background script');
    });
}

