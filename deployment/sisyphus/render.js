const cvs = document.getElementById("canvas");
const ctx = cvs.getContext("2d");

const boulderSpanX = document.getElementById("boulder-x")
const boulderSpanY = document.getElementById("boulder-y")
const boulderSpanR = document.getElementById("boulder-r")
const jumpSpan = document.getElementById("superjump")
const compassNeedle = document.getElementById("compass")

class Entity {
  spriteSheet = new Image();
  spriteWidth = 0;
  delayN = 10;
  maxN = 1;
  scale = 1;
  render(obj, N) {
    ctx.translate(obj.X, obj.Y);
    ctx.rotate(obj.A);
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

const Boulder = new Entity();
Boulder.maxN = 7;
Boulder.spriteWidth = 250;
Boulder.spriteSheet.src = "../assets/Boulder.png";

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

let floorTexture = new Image();
floorTexture.src = "../assets/GrassTexture.png";
function drawFloor(floorMap, dx, dy) {
  dx %= floorTexture.width;
  dy %= floorTexture.height;
  ctx.beginPath();
  ctx.moveTo(-playerView.D, floorMap[-playerView.D]);
  for (let x = 0; x < playerView.W + playerView.D; x += playerView.D) {
    ctx.lineTo(x, floorMap[x]);
  }
  ctx.lineTo(playerView.W + playerView.D, playerView.H + playerView.D);
  ctx.lineTo(-playerView.D, playerView.H + playerView.D);
  ctx.closePath();
  ctx.stroke();
  ctx.translate(-dx, -dy);
  let pattern = ctx.createPattern(floorTexture, "repeat");
  ctx.fillStyle = pattern;
  ctx.fill();
  ctx.resetTransform();
}

let backTexture = new Image();
backTexture.src = "../assets/BackTexture.jpg";

function drawBackground(dx, dy) {
  dx /= 1e2;
  dx %= backTexture.width;
  dy /= 1e2;
  dx %= backTexture.height;

  ctx.fillStyle = "#f5f17f";
  ctx.fillRect(0, 0, cvs.width, cvs.height);
  ctx.rect(0, 0, cvs.width, backTexture.height);

  ctx.translate(-dx, 0);
  let pattern = ctx.createPattern(backTexture, "repeat-x");
  ctx.fillStyle = pattern;
  ctx.fill();
  ctx.resetTransform();
}

function updateUI(obj){
  compassNeedle.style.transform = `rotate(${obj.Compass.A}rad)`
  boulderSpanX.innerText = Math.round(obj.Boulder.Meta.X/10)
  boulderSpanY.innerText = Math.round(obj.Boulder.Meta.Y/10)
  boulderSpanR.innerText = Math.round(obj.Boulder.R)/10
  jumpSpan.innerText = Math.round(obj.Sisyphus.Meta.Jump*10)/10

}