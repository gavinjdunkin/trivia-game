// app.js
const socket = new WebSocket('ws://localhost:3000/ws');

// Get the username from the user
const username = prompt('Enter your username:');

// Join the game by sending the username to the server
socket.send(JSON.stringify(username));

// Reference to HTML elements
const optionsElement = document.getElementById('options');
const scoreElement = document.getElementById('score');
const resultElement = document.getElementById('result');

socket.onmessage = (event) => {

    console.log(event.data);

    const gameState = JSON.parse(event.data);

    // Process the game state and update the UI
    switch(gameState.type) {
        case "question":
            updateUI(gameState);
            break;
        case "result":
            while( optionsElement.firstChild ){
                optionsElement.removeChild( optionsElement.firstChild );
            }
            if (gameState.correct == true)
            {
                resultElement.textContent = `CORRECT`;
            }
            else
            {
                resultElement.textContent = `INCORRECT`;
            }
            scoreElement.textContent = `Score: ${gameState.score}`;

            

    }
    
};

// Example: Function to update the UI with the received game state
function updateUI(gameState) {
    scoreElement.textContent = ``;
    resultElement.textContent = ``;

    // Update options
    gameState.options.forEach((option, index) => {
        const button = document.createElement('button');
        button.textContent = `${index + 1}. ${option}`;
        button.addEventListener('click', () => sendAnswer(index + 1));
        optionsElement.appendChild(button);
    });
}

function sendAnswer(index) {
    answer = ""
    switch(index) {
        case 1:
            answer = "a";
            break;
        case 2:
            answer = "b";
            break;
        case 3:
            answer = "c";
            break;
        case 4:
            answer = "d";
            break;
        default:
            answer = " ";
            break;
    }
    socket.send(answer);
}
