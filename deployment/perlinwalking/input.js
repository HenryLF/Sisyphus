const userInput = {
  Up: false,
  Down: false,
  Left: false,
  Right: false,
};

window.addEventListener("keydown", (e) => {
  switch (e.key) {
    case "ArrowUp":
      e.preventDefault()
      userInput.Up = true;
      break;
    case "ArrowDown":
      e.preventDefault()
      userInput.Down = true;
      break;
    case "ArrowLeft":
      e.preventDefault()
      userInput.Left = true;
      break;
    case "ArrowRight":
      e.preventDefault()
      userInput.Right = true;
      break;
  }
});
window.addEventListener("keyup", (e) => {
    switch (e.key) {
      case "ArrowUp":
        userInput.Up = false;
        break;
      case "ArrowDown":
        userInput.Down = false;
        break;
      case "ArrowLeft":
        userInput.Left = false;
        break;
      case "ArrowRight":
        userInput.Right = false;
        break;
    }
  });

