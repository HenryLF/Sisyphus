const cvs = document.getElementById("canvas");
const ctx = cvs.getContext("2d");

class Entity {
  spriteSheet = new Image();
  spriteWidth = 0;
  delayN = 10;
  delayI = 0;
  maxN = 1;
  render(x, y,A, N) {
    ctx.translate(x,y)
    ctx.rotate(A)
    ctx.translate(-x,-y)
    N = Math.round(N / this.delayN);
    let sx = this.spriteWidth * (N % this.maxN);
    ctx.drawImage(
      this.spriteSheet,
      sx,
      0,
      this.spriteWidth,
      this.spriteSheet.height,
      x - this.spriteSheet.height / 2,
      y - this.spriteWidth / 2,
      this.spriteWidth,
      this.spriteSheet.height
    );
    ctx.resetTransform()
  }
}
class EntityCaped extends Entity {
  render(x, y, N) {
    N = Math.min(N, this.maxN - 1);
    super.render(x, y, N);
  }
}

let PlayerIdle_img = new Image();
PlayerIdle_img.src = "../assets/PlayerIdle.png";
class PlayerIdle extends Entity {
  spriteWidth = 19;
  maxN = 5;
  delayI = 50;
  spriteSheet = PlayerIdle_img;
}

let PlayerLeft_img = new Image();
PlayerLeft_img.src = "../assets/PlayerLeft.png";
class PlayerLeft extends Entity {
  spriteWidth = 20;
  maxN = 6;
  spriteSheet = PlayerLeft_img;
}

let PlayerRight_img = new Image();
PlayerRight_img.src = "../assets/PlayerRight.png";
class PlayerRight extends PlayerLeft {
  spriteSheet = PlayerRight_img;
}

let PlayerJump_img = new Image();
PlayerJump_img.src = "../assets/PlayerJump.png";
class PlayerJump extends EntityCaped {
  maxN = 3;
  spriteWidth = 19;
  spriteSheet = PlayerJump_img;
}

let PlayerDown_img = new Image();
PlayerDown_img.src = "../assets/PlayerDown.png";
class PlayerDown extends EntityCaped {
  maxN = 3;
  spriteWidth = 19;
  spriteSheet = PlayerDown_img;
}

const Ball = {
  color: "#0000FF",
  render(obj) {
    ctx.save();
    ctx.fillStyle = this.color;
    let ball = new Path2D();
    ball.arc(obj.X, obj.Y, 10, 0, Math.PI * 2);
    ball.closePath();
    ctx.fill(ball);
    ctx.restore();
  },
};

function drawPlayer(obj, n) {
  let p;
  switch (obj.D) {
    case "R":
      p = new PlayerRight();
      break;
    case "L":
      p = new PlayerLeft();
      break;
    case "U":
      p = new PlayerJump();
      break;
    case "D":
      p = new PlayerDown();
      break;
    default:
      p = new PlayerIdle();
  }
  p.render(obj.X, obj.Y,obj.A, n);
}

function drawFloor(floorMap) {
  ctx.fillStyle = "green";
  ctx.strokeStyle = "black";

  ctx.beginPath();
  ctx.moveTo(-playerView.D, floorMap[-playerView.D]);
  for (let x = 0; x < playerView.W + playerView.D; x += playerView.D) {
    ctx.lineTo(x, floorMap[x]);
  }
  ctx.lineTo(playerView.W + playerView.D, playerView.H + playerView.D);
  ctx.lineTo(-playerView.D, playerView.H + playerView.D);
  ctx.closePath();
  ctx.stroke();
  ctx.fill();
}
