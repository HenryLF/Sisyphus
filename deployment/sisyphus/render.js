const cvs = document.getElementById("canvas");
const ctx = cvs.getContext("2d");

const boulderSpanX = document.getElementById("boulder-x");
const boulderSpanY = document.getElementById("boulder-y");
const boulderSpanR = document.getElementById("boulder-r");
const compassNeedle = document.getElementById("compass");

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
      (-w * this.scale) / 2,
      (-h * this.scale) / 2,
      w * this.scale,
      h * this.scale
    );
    ctx.resetTransform();
  }
}
class EntityMiror extends Entity {
  render(obj, N) {
    ctx.translate(obj.X, obj.Y);
    ctx.rotate(obj.A);
    ctx.scale(-1, 1);
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
      (w * this.scale) / 2,
      (-h * this.scale) / 2,
      -w * this.scale,
      h * this.scale
    );
    ctx.resetTransform();
  }
}
class EntityReverse extends Entity {
  render(obj, N) {
    N = this.maxN - (N % this.maxN);
    super.render(obj, N);
  }
}
class EntityCaped extends Entity {
  render(obj, N) {
    N = Math.min(N, (this.maxN - 1) * this.delayN);
    super.render(obj, N);
  }
}

class EntityCapedMiror extends EntityMiror {
  render(obj, N) {
    N = Math.min(N, (this.maxN - 1) * this.delayN);
    super.render(obj, N);
  }
}
class EntityCapedReverse extends EntityReverse {
  render(obj, N) {
    N = Math.min(N, (this.maxN - 1) * this.delayN);
    super.render(obj, N);
  }
}
class EntityCapedReverseMiror extends EntityMiror {
  render(obj, N) {
    N = Math.min(N, (this.maxN - 1) * this.delayN);
    N = this.maxN - (N % this.maxN);
    super.render(obj, N);
  }
}

const PlayerIdle_img = new Image();
PlayerIdle_img.src = "../assets/PlayerIdle.png";

const PlayerIdle = new Entity();
PlayerIdle.spriteWidth = 40;
PlayerIdle.maxN = 7;
PlayerIdle.spriteSheet = PlayerIdle_img;

const PlayerRight_img = new Image();
PlayerRight_img.src = "../assets/PlayerRight.png";

const PlayerRight = new Entity();
PlayerRight.spriteWidth = 50;
PlayerRight.scale = 1.25;
PlayerRight.maxN = 6;
PlayerRight.spriteSheet = PlayerRight_img;

const PlayerLeft = new EntityMiror();
PlayerLeft.spriteWidth = 50;
PlayerLeft.maxN = 6;
PlayerLeft.scale = 1.25;
PlayerLeft.spriteSheet = PlayerRight_img;

const PlayerHit_img = new Image();
PlayerHit_img.src = "../assets/PlayerHit.png";

const PlayerHit = new EntityCaped();
PlayerHit.spriteWidth = 62;
PlayerHit.scale = 1.5;
PlayerHit.maxN = 4;
PlayerHit.spriteSheet = PlayerHit_img;

const PlayerHitLeft = new EntityCapedMiror();
PlayerHitLeft.spriteWidth = 62;
PlayerHitLeft.scale = 1.5;
PlayerHitLeft.maxN = 4;
PlayerHitLeft.spriteSheet = PlayerHit_img;

const PlayerUp_img = new Image();
PlayerUp_img.src = "../assets/PlayerUp.png";

const PlayerUp = new EntityCaped();
PlayerUp.spriteWidth = 50;
PlayerUp.maxN = 7;
PlayerUp.delayN = 5;
PlayerUp.scale = 1.25;
PlayerUp.spriteSheet = PlayerUp_img;

const PlayerUpLeft = new EntityCapedMiror();
PlayerUpLeft.spriteWidth = 50;
PlayerUpLeft.maxN = 7;
PlayerUpLeft.scale = 1.25;
PlayerUpLeft.spriteSheet = PlayerUp_img;

const PlayerDown = new EntityCapedReverse();
PlayerDown.spriteWidth = 50;
PlayerDown.scale = 1.25;
PlayerDown.maxN = 7;
PlayerDown.spriteSheet = PlayerUp_img;

const PlayerDownLeft = new EntityCapedReverseMiror();
PlayerDownLeft.spriteWidth = 50;
PlayerDownLeft.scale = 1.25;
PlayerDownLeft.maxN = 7;
PlayerDownLeft.spriteSheet = PlayerUp_img;

const Boulder = new Entity();
Boulder.maxN = 7;
Boulder.spriteWidth = 250;
Boulder.spriteSheet.src = "../assets/Boulder.png";

const Hades = new Entity();
Hades.maxN = 6;
Hades.scale = 3;
Hades.spriteWidth = 50;
Hades.spriteSheet.src = "../assets/Hades.png";

function drawPlayer(obj, n) {
  switch (obj.D) {
    case "R":
      PlayerRight.render(obj, n);
      break;
    case "L":
      PlayerLeft.render(obj, n);
      break;
    case "UR":
      PlayerUp.render(obj, n);
      break;
    case "UL":
      PlayerLeft.render(obj, n);
      break;
    case "DR":
      PlayerDown.render(obj, n);
      break;
    case "DL":
      PlayerDownLeft.render(obj, n);
      break;
    case "HR":
      PlayerHit.render(obj, n);
      break;
    case "HL":
      PlayerHitLeft.render(obj, n);
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

function updateUI(obj) {
  compassNeedle.style.transform = `rotate(${obj.Compass.A}rad)`;
  boulderSpanX.innerText = Math.round(obj.Boulder.Meta.X / 10);
  boulderSpanY.innerText = -Math.round(obj.Boulder.Meta.Y / 10);
  boulderSpanR.innerText = Math.round(obj.Boulder.R) / 10;
}
