const cvs = document.getElementById("canvas");
const ctx = cvs.getContext("2d");

var gameState = {
  View: {
    X: 0,
    Y: 0,
    W: 0,
    H: 0,
    D: 5,
  },
  PlayerA: {
    X: 0,
    Y: 0,
    A: 0,
    R: 0,
    D: 5,
  },
  PlayerB: {
    X: 0,
    Y: 0,
    A: 0,
    R: 0,
    D: 5,
  },
  Ball: {
    X: 0,
    Y: 0,
    A: 0,
    R: 0,
    D: 5,
  },
  Net: {
    X: 0,
    Y: 0,
    A: 0,
    R: 0,
    D: 5,
  },
  Floor: new Map(),
};

class Entity {
  spriteSheet = new Image();
  spriteWidth = 0;
  delayN = 10;
  maxN = 1;
  scale = 1;
  render(obj, N) {
    ctx.translate(obj.X, obj.Y);
    if (obj.A) {
      ctx.rotate(obj.A);
    }
    N = Math.round(N / this.delayN);
    let sx = this.spriteWidth * (N % this.maxN);
    let w = obj.R * 2;
    let h = (obj.R * 2 * this.spriteSheet.height) / this.spriteWidth;
    ctx.drawImage(
      this.spriteSheet,
      sx,
      0,
      this.spriteWidth,
      this.spriteSheet.height,
      -w / 2,
      -h / 2,
      w * this.scale,
      h * this.scale
    );
    ctx.resetTransform();
  }
}
class EntityCaped extends Entity {
  render(x, y, N) {
    N = Math.min(N, this.maxN - 2);
    super.render(x, y, N);
  }
}
const PlayerIdle = new Entity();
PlayerIdle.spriteWidth = 19;
PlayerIdle.maxN = 5;
PlayerIdle.spriteSheet.src = "../assets/PlayerIdle.png";

const PlayerLeft = new Entity();
PlayerLeft.spriteWidth = 20;
PlayerLeft.maxN = 6;
PlayerLeft.spriteSheet.src = "../assets/PlayerLeft.png";

const PlayerRight = new Entity();
PlayerRight.spriteWidth = 20;
PlayerRight.maxN = 6;
PlayerRight.spriteSheet.src = "../assets/PlayerRight.png";

const PlayerUp = new Entity();
PlayerUp.spriteWidth = 19;
PlayerUp.maxN = 3;
PlayerUp.spriteSheet.src = "../assets/PlayerUp.png";

const PlayerJump = new Entity();
PlayerJump.spriteWidth = 30;
PlayerJump.scale = 1.2;
PlayerJump.maxN = 5;
PlayerJump.spriteSheet.src = "../assets/PlayerJump.png";

const PlayerDown = new Entity();
PlayerDown.spriteWidth = 19;
PlayerDown.maxN = 3;
PlayerDown.spriteSheet.src = "../assets/PlayerDown.png";

const Ball = new Entity();
Ball.maxN = 7;
Ball.spriteWidth = 250;
Ball.spriteSheet.src = "../assets/Boulder.png";



function drawPlayer(obj, n) {
  switch (obj.D) {
    case "R":
      PlayerRight.render(obj, n);
      break;
    case "L":
      PlayerLeft.render(obj, n);
      break;
    case "U":
      PlayerUp.render(obj, n);
      break;
    case "D":
      PlayerDown.render(obj, n);
      break;
    case "J":
      PlayerJump.render(obj, n);
      break;
    default:
      PlayerIdle.render(obj, n);
  }
}



let backTexture = new Image();
backTexture.src = "../assets/BackTexture.jpg";

function drawBackground(dx, dy) {
  dx = (gameState.View.X / 1e2) % backTexture.width;
  dy = (gameState.View.Y / 1e2) % backTexture.height;
  ctx.save();
  ctx.fillStyle = "#f5f17f";
  ctx.fillRect(0, 0, cvs.width, cvs.height);
  let rect = ctx.rect(0, 0, cvs.width, backTexture.height);
  
  ctx.translate(-dx, 0);
  let backPattern = ctx.createPattern(backTexture, "repeat-x");
  ctx.fillStyle = backPattern;
  ctx.fill(rect);
  ctx.restore();
  ctx.resetTransform();
}



let floorTexture = new Image();
floorTexture.src = "../assets/GrassTexture.png";
function drawFloor(dx, dy) {
  dx = gameState.View.X % floorTexture.width;
  dy = gameState.View.Y % floorTexture.height;
  ctx.save();
  ctx.beginPath();
  ctx.moveTo(-gameState.View.D, gameState.Floor[-gameState.View.D]);
  for (let x = 0; x < cvs.width + gameState.View.D; x += gameState.View.D) {
    ctx.lineTo(x, gameState.Floor[x]);
  }
  ctx.lineTo(cvs.width + gameState.View.D, cvs.height + gameState.View.D);
  ctx.lineTo(-gameState.View.D, cvs.height + gameState.View.D);
  ctx.closePath();
  ctx.stroke();
  ctx.translate(-dx, -dy);
  let floorPattern = ctx.createPattern(floorTexture, "repeat");
  ctx.fillStyle = floorPattern;
  ctx.fill();
  ctx.restore();
  ctx.resetTransform();
}



let netTexture = new Image();
netTexture.src = "../assets/NetTexture.png";
const Net = {
  render(obj, n) {
    ctx.save();
    let netPattern = ctx.createPattern(netTexture, "repeat");
    ctx.fillStyle = netPattern;
    ctx.fillRect(
      obj.X - 0.02 * obj.R,
      obj.Y - obj.R,
      0.2* 2 * obj.R,
      obj.R * 2
    );
    ctx.restore();
  },
};
