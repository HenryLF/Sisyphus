const userInput = {
  Up: false,
  Down: false,
  Left: false,
  Right: false,
  Jump: false,
};
window.userInput = userInput;

window.addEventListener("keydown", (e) => {
  switch (e.key) {
    case "ArrowUp":
      e.preventDefault();
      userInput.Up = true;
      break;
    case "ArrowDown":
      e.preventDefault();
      userInput.Down = true;
      break;
    case "ArrowLeft":
      e.preventDefault();
      userInput.Left = true;
      break;
    case "ArrowRight":
      e.preventDefault();
      userInput.Right = true;
      break;
    case " ":
      e.preventDefault();
      userInput.Jump = true;
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
    case " ":
      e.preventDefault();
      userInput.Jump = false;
      break;
  }
});

const jumpButton = document.getElementById("jump-button");
jumpButton.ontouchstart = (e) => {
  e.preventDefault();
  userInput.Jump = true;
};
jumpButton.ontouchend = (e) => {
  e.preventDefault();
  userInput.Jump = false;
};
jumpButton.ontouchcancel = (e) => {
  e.preventDefault();
  userInput.Jump = false;
};

var joy = new JoyStick("joyDiv", {}, (dt) => {
  x = parseInt(dt.x);
  y = parseInt(dt.y);
  if (x > 50) {
    userInput.Right = true;
  } else {
    userInput.Right = false;
  }
  if (x < -50) {
    userInput.Left = true;
  } else {
    userInput.Left = false;
  }
  if (y > 50) {
    userInput.Up = true;
  } else {
    userInput.Up = false;
  }
  if (y < -50) {
    userInput.Down = true;
  } else {
    userInput.Down = false;
  }
});
