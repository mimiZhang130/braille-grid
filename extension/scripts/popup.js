document.addEventListener('DOMContentLoaded', function() {
    var setApiKeyButton = document.getElementById('setApiKeyButton');
    var startButton = document.getElementById('startButton');
    var nextButton = document.getElementById('nextButton');
    var apiKeyInput = document.getElementById('apiKeyInput');

    setApiKeyButton.addEventListener('click', function() {
        var apiKey = apiKeyInput.value;
        chrome.runtime.sendMessage({ action: "apiKey", apiKey: apiKey }, function(response) {
            console.log(response.message);
        });
    });

    startButton.addEventListener('click', function() {
        chrome.runtime.sendMessage({ action: "start" }, function(response) {
            console.log(response.message);
        });
    });

    nextButton.addEventListener('click', function() {
        chrome.runtime.sendMessage({ action: "next" }, function(response) {
            console.log(response.message);
        });
    });

});

