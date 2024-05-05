function tokenize(text, groupSize) {
    const words = text.split(/\s+/);
    
    const paragraphs = [];

    let start = 0;
    let end = 0;

    while (start < words.length) {
        end = Math.min(start + groupSize, words.length);

        const paragraph = words.slice(start, end).join(" ");

        paragraphs.push(paragraph);

        start = end;
    }

    return paragraphs;
}

module.exports = tokenize;
