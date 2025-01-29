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

window.onresize = resizeCanvas;
window.onload = setTimeout(() => {
  resizeCanvas();
  loop();
}, 1000);
let Balls;

let raf;
function loop() {
  G = GetUpdate();
  ctx.clearRect(0, 0, cvs.width, cvs.height);
  drawBackground(G.X, G.Y);
  drawFloor(G.Floor, G.X, G.Y);
  Boulder.render(G.Boulder, raf);
  Hades.render(G.Hades,raf)
  drawPlayer(G.Sisyphus, raf);
  updateUI(G);
  raf = requestAnimationFrame(loop);
}
