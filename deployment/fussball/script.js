var A = 5;

function resizeCanvas() {
  cvs.width = window.innerWidth;
  cvs.height = window.innerHeight;
  InitView(cvs.width, cvs.height);
}
window.onresize = resizeCanvas;

window.onload = () => {
  setTimeout(() => {
    resizeCanvas();
    loop();
  }, 500);
};

let raf
function loop() {
  GetUpdate()
  ctx.clearRect(0,0,cvs.width,cvs.height);
  drawBackground()
  drawFloor()
  Net.render(gameState.Net)
  Ball.render(gameState.Ball,raf)
  drawPlayer(gameState.PlayerA,raf)
  drawPlayer(gameState.PlayerB,raf)
  raf = requestAnimationFrame(loop)
}
