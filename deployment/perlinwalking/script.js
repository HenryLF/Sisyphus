var playerView = {
  X: 0,
  Y: 0,
  W: 0,
  H: 0,
  D: 5,
};

function resizeCanvas() {
  cvs.width = window.innerWidth * 0.98;
  cvs.height = window.innerHeight * 0.98;
  playerView = InitView(cvs.width, cvs.height);
}

cvs.onclick = (e) => {
  let x = e.pageX - cvs.offsetLeft;
  let y = e.pageY - cvs.offsetTop;
  NewBall(x, y);
};

window.onresize = resizeCanvas;
window.onload = setTimeout(() => {
  resizeCanvas();
  loop();
}, 1000);
let Balls;

let raf;
function loop() {
  GetUpdate(userInput.Up, userInput.Left, userInput.Down, userInput.Right);
  playerView = SlowCenter();
  floorMap = GetMap();
  ctx.clearRect(0, 0, cvs.width, cvs.height);
  drawFloor();
  C = GetPlayer();
  drawPlayer(C, raf);
  Balls = GetBalls();
  for (let b of Object.values(Balls)) {
    Ball.render(b);
  }

  raf = requestAnimationFrame(loop);
}

const setMapButton = document.getElementById("setmap");
const mapParams = document.getElementById("map-params");

setMapButton.onclick = (e) => {
  e.preventDefault();
  let Dt = new FormData(mapParams);
  SetMapParams(
    parseFloat(Dt.get("alpha")),
    parseFloat(Dt.get("beta")),
    parseInt(Dt.get("nit")),
    parseFloat(Dt.get("xscale")),
    parseFloat(Dt.get("yscale"))
  );
};
