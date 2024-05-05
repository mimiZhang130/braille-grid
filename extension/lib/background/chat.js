const OpenAI = require('openai');
const br = require('braille');
const printStream = require('./printStream');
const url = 'http://localhost:3000';

async function chat(key, text) {
    const openai = new OpenAI({ apiKey: key });
    const completion = await openai.chat.completions.create({
        messages: [{
            role: "assistant", content: `
You are assisting blind individuals by summarizing text on a webpage that will later be converted into braille format. 
They will interface with the website using this summary.
Concisely summarize the following:
${text}
` }],
        model: "gpt-3.5-turbo",
    });
    const summary = completion.choices[0].message.content;

    const braille = br.toBraille(summary);

    fetch(url, {
        method: 'POST',
        body: braille
    })
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response;
        })
        .then(data => {
            printStream(data.body);
        })
        .catch(error => {
            console.error('Error:', error.message);
        });
}

module.exports = chat;

