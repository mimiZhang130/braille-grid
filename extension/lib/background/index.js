const tokenize = require('./tokenize');
const chat = require('./chat');

var tokens = [];
var apiKey = '';
var curInd = 0;
chrome.runtime.onMessage.addListener(function(request, sender, sendResponse) {

    if (request.action == "dom") {
        const text = request.text;
        tokens = tokenize(text, 100);
        sendResponse({ message: 'DOM OK' });
        return true;
    } else if (request.action == "apiKey") {
        apiKey = request.apiKey;
    } else if (request.action == "start") {
        curInd = 0;
        console.log(tokens, apiKey);
    } else if (request.action == "next") {
        let tok = tokens[curInd];
        chat(apiKey, tok);
        curInd = curInd + 1;
    }

    sendResponse({ message: 'OK' });
    return true;
});

