let websocket = null;

window.addEventListener("load", () => {
    document.onkeydown = function (e) {
        let keyCode = e.key;

        if (keyCode === "Enter") {
            sendMessage();
        }
    };
});

function sendMessage() {
    let username = document.getElementById("username");
    let content = document.getElementById("content");

    websocket.send(
        JSON.stringify({
            username: username.value,
            content: content.value,
        })
    );

    document.getElementById("content").value = "";
}

window.addEventListener('DOMContentLoaded', (_) => {
    websocket = new WebSocket("ws://" + window.location.host + "/ws");

    let messages = document.getElementById("messages");

    websocket.addEventListener("message", function (payload) {
        let data = JSON.parse(payload.data);
        let message = document.createElement('div');

        message.setAttribute('class', 'message')
        message.appendChild(document.createTextNode(`${data.username}: ${data.content}`))

        messages.appendChild(message);
        messages.appendChild(document.createElement('br'));

        message.scrollIntoView();
    });
});
