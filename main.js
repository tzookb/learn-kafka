const http = require("http");

const host = '0.0.0.0';
const port = 80;

const requestListener = function (req, res) {
    console.log('xxxxxx');
    res.writeHead(200);
    res.end("My first server!");
};

const server = http.createServer(requestListener);
server.listen(port, host, () => {
    console.log(`Server is running on http://${host}:${port}`);
});