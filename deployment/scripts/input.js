const userInput = {
  Up: false,
  Down: false,
  Left: false,
  Right: false,
  Hit: false,
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
      userInput.Hit = true;
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
      userInput.Hit = false;
      break;
  }
});

const hitButton = document.getElementById("hit-button");
hitButton.onmousedown = (e) => {
  e.preventDefault();
  userInput.Hit = true;
};
hitButton.onmouseup = (e) => {
  userInput.Hit = false;
};
hitButton.ontouchstart = (e) => {
  e.preventDefault();
  userInput.Hit = true;
};
hitButton.ontouchend = (e) => {
  e.preventDefault();
  userInput.Hit = false;
};
hitButton.ontouchcancel = (e) => {
  e.preventDefault();
  userInput.Hit = false;
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
