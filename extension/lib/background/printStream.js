async function printStream(stream) {
    const reader = stream.getReader();
    const decoder = new TextDecoder();

    try {
        while (true) {
            const { done, value } = await reader.read();
            if (done) {
                break;
            }
            const decodedString = decoder.decode(value);
            console.log("Decoded response from service:", decodedString);
        }
    } finally {
        reader.releaseLock();
    }
}

module.exports = printStream;
